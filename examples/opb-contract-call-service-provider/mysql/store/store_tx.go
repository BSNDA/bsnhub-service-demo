package store

import (
	"fmt"
	"github.com/bianjieai/bsnhub-service-demo/examples/opb-contract-call-service-provider/common/mysql"
	logging "github.com/bianjieai/bsnhub-service-demo/examples/opb-contract-call-service-provider/common"
	"time"
)

const (
	Source_Relayer = 0
	Source_Provider = 1

	TxStatus_Unknow  = 0
	TxStatus_Success =1
	TxStatus_Error   = 2
)

var (
	source_service = Source_Provider
)

func NowTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}


//HandleInterchainRequest 初次处理请求

// Request_id、From_chainid、From_tx、To_chainid
// 请求HUB  Hub_req_tx、Ic_request_id
//记录 tx_status 0


// InitTransRecord 第一次初始化relauer的交易记录
func InitRelayerTransRecord(
	requestId string,
	fromChainId string,
	fromTxId string,
	toChainId string,

	hubReqTxId string,
	icRequestId string,
	txStatus int,
	errMsg string,
) {
	//

	logging.Logger.Infof("Init Relayer trans record , requestId is %s,tx status is %d",requestId,txStatus)

	insertsql := fmt.Sprintf("INSERT INTO %s ( " +
		"request_id, " +
		"from_chainId, " +
		"from_tx," +
		"hub_req_tx, " +
		"ic_request_id ," +
		"to_chainid ," +
		"tx_createtime, " +
		"tx_status, " +
		"error, " +
		"source_service ) "+
		"VALUES ( ?, ?, ?,?, ?, ?, ?,?,?,?);", _TabName_cc_Tx)

	lastId, rows, err := mysql.Exec(insertsql,
		requestId,
		fromChainId,
		fromTxId,
		hubReqTxId,
		icRequestId,
		toChainId,
		NowTime(),
		txStatus,
		errMsg,
		source_service)

	if err != nil {
		logging.Logger.Errorf("Init Relayer trans record Failed :%s", err.Error())
	} else {
		logging.Logger.Infof("Init Relayer trans record  lastId:%d ;rows:%d ", lastId, rows)
	}


	return
}

type RelayerResInfo struct {
	RequestId string
	FromResTxId string
	TxStatus int
	ErrMsg string
}

// callback
// Request_id Ic_request_id From_res_tx
// 记录 tx_status 1 or 2
func RelayerResponeRecord(data *RelayerResInfo) {

	logging.Logger.Infof("set relayer callback record , requestId is %s,tx status is %d",data.RequestId,data.TxStatus)
	sql :=fmt.Sprintf("update %s set from_res_tx = ? ,error = ? ,tx_time =? ,tx_status = ? where request_id = ? and source_service = %d",_TabName_cc_Tx,source_service)

	lastId, rows, err := mysql.Exec(sql,
		data.FromResTxId,
		data.ErrMsg,
		NowTime(),
		data.TxStatus,
		data.RequestId)

	if err != nil {
		logging.Logger.Errorf("set relayer callback record Failed :%s", err.Error())
	} else {
		logging.Logger.Infof("set relayer callback record lastId:%d ;rows:%d ", lastId, rows)
	}


}


//requestId ,to_chainid,ic_request_id ,to_tx,hub_res_tx ,tx_status,error,source_service

//InitProviderTransRecord
func InitProviderTransRecord(requestId ,to_chainid,ic_request_id ,to_tx ,error string,tx_status int)  {

	logging.Logger.Infof("Init Provider trans record , requestId is %s,tx status is %d",requestId,tx_status)
	insertsql :=fmt.Sprintf("INSERT INTO %s ( " +
		"request_id, " +
		"ic_request_id ," +
		"to_chainid ," +
		"to_tx, " +
		"tx_createtime, " +
		"tx_status, " +
		"error, " +
		"source_service ) "+
		"VALUES ( ?, ?, ?,?, ?, ?,?,?);", _TabName_cc_Tx)



	lastId, rows, err := mysql.Exec(insertsql,
		requestId,
		ic_request_id,
		to_chainid,
		to_tx,
		NowTime(),
		tx_status,
		error,
		source_service)

	if err != nil {
		logging.Logger.Errorf("Init Provider trans record Failed :%s", err.Error())
	} else {
		logging.Logger.Infof("Init Provider trans record  lastId:%d ;rows:%d ", lastId, rows)
	}

}

type ProviderResInfo struct {
	IcRequestId string
	HUBResTxId string
	TxStatus int
	ErrMsg string
}


//ic_request_id ,hub_res_tx
func ProviderCallBackTransRecord(data *ProviderResInfo )  {
	logging.Logger.Infof("set provider callback record , ic_requestId is %s,res txid is %s,tx status is %d",data.IcRequestId,data.HUBResTxId,data.TxStatus)
	if data.IcRequestId == "" {
		return
	}
	sql :=fmt.Sprintf("update %s set hub_res_tx = ? ,error = ? ,tx_time =? ,tx_status = ? where ic_request_id = ? and source_service = %d",_TabName_cc_Tx,source_service)

	lastId, rows, err := mysql.Exec(sql,
		data.HUBResTxId,
		data.ErrMsg,
		NowTime(),
		data.TxStatus,
		data.IcRequestId)

	if err != nil {
		logging.Logger.Errorf("set relayer callback record Failed :%s", err.Error())
	} else {
		logging.Logger.Infof("set relayer callback record lastId:%d ;rows:%d ", lastId, rows)
	}
}



/*

request_id			relayer/provider
from_chainid		relayer
from_tx				relayer
hub_req_tx			relayer
ic_request_id		relayer/provider
to_chainid			relayer/provider
to_tx				provider
hub_res_tx			provider
from_res_tx			relayer
tx_status 			relayer/provider
tx_time				relayer/provider
tx_createtime		relayer/provider
error				relayer/provider
source_service		relayer/provider

*/
































