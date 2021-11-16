package contract_service

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/bianjieai/iritamod-sdk-go/service"
	"github.com/bianjieai/iritamod-sdk-go/wasm"
	sdk "github.com/irisnet/core-sdk-go/types"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"opb-contract-call-service-provider/contract-service/opb"
	"opb-contract-call-service-provider/mysql/store"
	"strconv"
	"strings"
	"time"

	"opb-contract-call-service-provider/contract-service/opb/config"

	"opb-contract-call-service-provider/server"
	"opb-contract-call-service-provider/types"
)

// ContractService defines the contract service
type ContractService struct {
	opbClient *opb.OpbChain
	Logger    *log.Logger
}

// BuildContractService builds a ContractService instance from the given config
func BuildContractService(v *viper.Viper, chainManager *server.ChainManager) (ContractService, error) {
	baseConfig, err := config.NewBaseConfig(v)
	if err != nil {
		return ContractService{}, err
	}

	return ContractService{
		opbClient: opb.NewOpbChain(*baseConfig, chainManager),
	}, nil
}

// Callback implements the iservice.RespondCallback interface
func (cs ContractService) Callback(reqCtxID, reqID, input string) (output string, result string) {

	cs.Logger.Infof("service request received, request id: %s", reqID)
	res := &types.Result{
		Code: 200,
	}

	var txHash string
	var callResult string
	var reqHeader types.Header

	defer func() {
		resBz, _ := json.Marshal(res)
		result = string(resBz)

		if res.Code == 200 {
			var outputBz []byte
			var headerBz []byte
			headerBz, _ = json.Marshal(reqHeader)
			outputBz, _ = json.Marshal(types.Output{Result: callResult, TxHash: txHash})
			output = fmt.Sprintf(`{"header":%s,"body":%s}`, headerBz, outputBz)
		}

		cs.Logger.Infof("request processed, result: %s, output: %s", result, output)
	}()
	cs.Logger.Infof("Input is %s ", input)
	var request types.Input
	err := json.Unmarshal([]byte(input), &request)
	if err != nil {
		//参数不符合规则，直接不处理
		res.Code = 500
		res.Message = fmt.Sprintf("can not parse request [%s] input json string : %s", reqID, err.Error())
		return
	}

	reqHeader = request.Header

	id, _ := strconv.ParseInt(request.Dest.ChainID, 10, 64)
	chainParams, err := cs.opbClient.ChainManager.GetChainParams(types.GetChainID(id))
	if err != nil {
		res.Code = 204
		res.Message = "chain params not exist"
		cs.Logger.Error("chain params not exist: ", err)

		return
	}

	//mysql.OnServiceRequestReceived(reqID, request.Dest.ChainID)

	// instantiate the opb client with the specified group id and chain id
	err = cs.opbClient.InstantiateClient(chainParams)
	if err != nil {
		res.Code = 500
		res.Message = "failed to connect to the opb node"
		cs.Logger.Error("InstantiateClient err: ", err)
		return
	}

	switch request.Dest.EndpointType {
	case "service":
		var invokeServiceReq types.InvokeServiceReq
		err := json.Unmarshal(request.CallData, &invokeServiceReq)
		if err != nil {
			res.Code = 500
			res.Message = fmt.Sprintf("can not parse request [%s] input json string : %s", reqID, err.Error())
			return
		}
		if invokeServiceReq.Providers == nil {
			invokeServiceReq.Providers = []string{request.Dest.EndpointAddress}
		}
		serviceFeeCap, err := sdk.ParseDecCoins(invokeServiceReq.ServiceFeeCap)
		if err != nil {
			res.Code = 500
			res.Message = fmt.Sprintf("can not parse request [%s] input json string : %s", reqID, err.Error())
			return
		}

		targetReqCtxID, _, err := cs.opbClient.OpbClient.Service.InvokeService(service.InvokeServiceRequest{
			ServiceName:   invokeServiceReq.ServiceName,
			Providers:     invokeServiceReq.Providers,
			Input:         invokeServiceReq.Input,
			ServiceFeeCap: serviceFeeCap,
			Timeout:       invokeServiceReq.Timeout,
		}, cs.opbClient.BuildBaseTx())
		if err != nil {
			res.Code = 500
			res.Message = fmt.Sprintf("invoke service failed : %s", err.Error())
			return
		}

		requests, err := cs.opbClient.OpbClient.Service.QueryRequestsByReqCtx(targetReqCtxID, 1, nil)
		if err != nil {
			res.Code = 500
			res.Message = fmt.Sprintf("query requests by ReqCtx failed: %s", err.Error())
			return
		}

		if len(requests) == 0 {
			res.Code = 500
			res.Message = fmt.Sprintf("no service request initiated on %s", request.Dest.ChainID)
			return
		}
		ch := make(chan int)
		callbackWrapper := func(reqCtxID, reqID, responses string) {
			resp := types.ResponseAdaptor{
				StatusCode: 200,
				// todo result is missing in sdk
				//Result:     result,
				Output: responses,
			}
			callResultbyte, _ := json.Marshal(resp)
			callResult = string(callResultbyte)
			ch <- 1
		}

		cs.Logger.Infof("waiting for the service response on %s", request.Dest.ChainID)
		subscription, err := cs.opbClient.OpbClient.Service.SubscribeServiceResponse(targetReqCtxID, callbackWrapper)
		if err != nil {
			res.Code = 500
			res.Message = fmt.Sprintf("no service request initiated on %s", request.Dest.ChainID)
			return
		}
		go func() {
			for {
				status, err2 := cs.opbClient.OpbClient.Status(context.Background())
				req, err3 := cs.opbClient.OpbClient.Service.QueryServiceRequest(requests[0].ID)
				if err != nil || err2 != nil || err3 != nil || status.SyncInfo.LatestBlockHeight > req.ExpirationHeight {
					cs.Logger.Infof("HUB Unsubscribe RequestID is %s", requests[0].ID)
					res.Code = 500
					res.Message = fmt.Sprintf("call service timeout")
					_ = cs.opbClient.OpbClient.Unsubscribe(subscription)
					ch <- 1
					break
				}
				time.Sleep(time.Second)
			}
		}()
		<-ch

	default:
		execAbi := wasm.NewContractABI().
			WithMethod("call_service").
			WithArgs("request_id", request.ReqSequence).
			WithArgs("endpoint_address", request.Dest.EndpointAddress).
			WithArgs("call_data", base64.StdEncoding.EncodeToString(request.CallData))

		resultTx, err := cs.opbClient.OpbClient.WASM.Execute(chainParams.TargetCoreAddr, execAbi, nil, cs.opbClient.BuildBaseTx())
		if err != nil {
			//mysql.TxErrCollection(reqID, err.Error())
			cs.Logger.Errorf("Opb ChainId %s Chaincode %s CallService has error %v", request.Dest.ChainID, request.Dest.EndpointAddress, err)
			//不包含重复交易，再记录
			if strings.Contains(err.Error(), "Unauthorized") {
				cs.Logger.Infof("the request has been received or received invalid transaction ,not record trans")
				res.Code = 204
				res.Message = err.Error()

				return
			} else {
				cs.Logger.Infof("call fabric error don't has 'the request has been received',record trans")
				store.InitProviderTransRecord(request.Header.ReqSequence, request.Dest.ChainID, reqID, "", err.Error(), store.TxStatus_Error)
				//store.TargetChainInfo(&InsectCrossInfo)
				res.Code = 500
				res.Message = err.Error()
				return
			}
		}
		txHash = resultTx.Hash.String()

		err = cs.opbClient.WaitForSuccess(resultTx.Hash.String(), "callService")
		if err != nil {
			//mysql.TxErrCollection(reqID, err.Error())
			cs.Logger.Errorf("Opb ChainId %s Chaincode %s WaitForReceipt has error %v", request.Dest.ChainID, request.Dest.EndpointAddress, err)
			//不包含重复交易，再记录
			if strings.Contains(err.Error(), "Unauthorized") {
				cs.Logger.Infof("the request has been received or received invalid transaction ,not record trans")
				res.Code = 204
				res.Message = err.Error()

				return
			} else {
				cs.Logger.Infof("call fabric error don't has 'the request has been received',record trans")
				store.InitProviderTransRecord(request.Header.ReqSequence, request.Dest.ChainID, reqID, "", err.Error(), store.TxStatus_Error)
				//store.TargetChainInfo(&InsectCrossInfo)
				res.Code = 500
				res.Message = err.Error()
				return
			}
		}

		for _, e := range resultTx.TxResult.Events {
			if e.Type == "wasm" && len(e.Attributes) > 1 {
				for _, attr := range e.Attributes {
					if string(attr.Key) == "_contract_address" && string(attr.Value) == request.Dest.EndpointAddress {
						for _, attr = range e.Attributes {
							if string(attr.Key) == "result" {
								callResult = string(attr.Value)
							}
						}
					}
				}
			}
		}
	}

	//mysql.OnContractTxSend(reqID, txHash)
	store.InitProviderTransRecord(request.Header.ReqSequence, request.Dest.ChainID, reqID, txHash, "", store.TxStatus_Unknow)
	return output, result
}
