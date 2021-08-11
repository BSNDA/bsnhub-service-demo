package contract_service

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/bianjieai/bsnhub-service-demo/examples/opb-contract-call-service-provider/contract-service/opb"
	"github.com/bianjieai/bsnhub-service-demo/examples/opb-contract-call-service-provider/mysql/store"
	"strings"

	"strconv"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/bianjieai/bsnhub-service-demo/examples/opb-contract-call-service-provider/contract-service/opb/config"

	"github.com/bianjieai/bsnhub-service-demo/examples/opb-contract-call-service-provider/server"
	"github.com/bianjieai/bsnhub-service-demo/examples/opb-contract-call-service-provider/types"
	"github.com/bianjieai/irita-sdk-go/modules/wasm"
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
			outputBz, _ = json.Marshal(types.Output{Result: callResult,TxHash: txHash})
			output = fmt.Sprintf(`{"header":%s,"body":%s}`, headerBz, outputBz)
		}

		cs.Logger.Infof("request processed, result: %s, output: %s", result, output)
	}()
	cs.Logger.Infof("Input is %s ", input)
	var request types.Input
	err := json.Unmarshal([]byte(input), &request)
	if err != nil {
		//参数不符合规则，直接不处理
		res.Code = 204
		res.Message = fmt.Sprintf("can not parse request [%s] input json string : %s", reqID, err.Error())
		return
	}

	reqHeader = request.Header

	id, _ := strconv.ParseInt(request.Dest.ChainID, 10, 64)
	chainParams, err := cs.opbClient.ChainManager.GetChainParams(types.GetChainID(id))
	if err != nil {
		res.Code = 204
		res.Message = "chain params not exist"
		cs.Logger.Error("chain params not exist")
		return
	}

	//mysql.OnServiceRequestReceived(reqID, request.Dest.ChainID)

	// instantiate the opb client with the specified group id and chain id
	err = cs.opbClient.InstantiateClient(chainParams)
	if err != nil {
		res.Code = 500
		res.Message = "failed to connect to the opb node"

		return
	}

	execAbi := wasm.NewContractABI().
		WithMethod("call_service").
		WithArgs("request_id", request.ReqSequence).
		WithArgs("endpoint_address", request.Dest.EndpointAddress).
		WithArgs("call_data", base64.StdEncoding.EncodeToString(request.CallData))
	resultTx, err := cs.opbClient.OpbClient.WASM.Execute(chainParams.TargetCoreAddr, execAbi, nil, cs.opbClient.BuildBaseTx())
	if err != nil {
		//mysql.TxErrCollection(reqID, err.Error())
		cs.Logger.Errorf("FISCO ChainId %s Chaincode %s CallService has error %v", request.Dest.ChainID, request.Dest.EndpointAddress, err)
		//不包含重复交易，再记录
		if strings.Contains(err.Error(), "Unauthorized") {
			cs.Logger.Infof("the request has been received or received invalid transaction ,not record trans")
			res.Code = 204
			res.Message = err.Error()

			return
		}else  {
			cs.Logger.Infof("call fabric error don't has 'the request has been received',record trans")
			store.InitProviderTransRecord(request.Header.ReqSequence,request.Dest.ChainID,reqID,"",err.Error(),store.TxStatus_Error)
			//store.TargetChainInfo(&InsectCrossInfo)
			res.Code = 500
			res.Message = err.Error()
			return
		}
	}
	txHash = resultTx.Hash

	err = cs.opbClient.WaitForSuccess(resultTx.Hash, "callService")
	if err != nil {
		//mysql.TxErrCollection(reqID, err.Error())
		cs.Logger.Errorf("FISCO ChainId %s Chaincode %s WaitForReceipt has error %v", request.Dest.ChainID, request.Dest.EndpointAddress, err)
		//不包含重复交易，再记录
		if strings.Contains(err.Error(), "Unauthorized") {
			cs.Logger.Infof("the request has been received or received invalid transaction ,not record trans")
			res.Code = 204
			res.Message = err.Error()

			return
		}else  {
			cs.Logger.Infof("call fabric error don't has 'the request has been received',record trans")
			store.InitProviderTransRecord(request.Header.ReqSequence,request.Dest.ChainID,reqID,"",err.Error(),store.TxStatus_Error)
			//store.TargetChainInfo(&InsectCrossInfo)
			res.Code = 500
			res.Message = err.Error()
			return
		}
	}

	for _, e := range resultTx.Events {
		if e.Type == "wasm" && len(e.Attributes) > 1 {
			for _, attr := range e.Attributes {
				if attr.Key == "contract_address" && attr.Value == request.Dest.EndpointAddress {
					for _, attr = range e.Attributes {
						if attr.Key == "result" {
							callResult = attr.Value
						}
					}
				}
			}
		}
	}

	//mysql.OnContractTxSend(reqID, txHash)
	store.InitProviderTransRecord(request.Header.ReqSequence,request.Dest.ChainID,reqID,txHash,"",store.TxStatus_Unknow)
	return output, result
}
