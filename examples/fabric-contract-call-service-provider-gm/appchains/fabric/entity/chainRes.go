package entity

import (
	"bsn-irita-fabric-provider-gm/appchains/fabric/metadata"
	"bsn-irita-fabric-provider-gm/common"
	"encoding/hex"
	"encoding/json"
	"github.com/BSNDA/fabric-sdk-go-gm/pkg/client/channel"
)

func NewFabricRes(fabRes channel.Response) *FabricRespone {
	res := &FabricRespone{
		TxValidationCode: int32(fabRes.TxValidationCode),
		TxId:             string(fabRes.TransactionID),
		Payload:          "0x" + hex.EncodeToString(fabRes.Payload),
		ChaincodeStatus:  fabRes.ChaincodeStatus,
	}

	if res.TxValidationCode != 0 {
		common.Logger.Errorf("call fabric res  TxValidationCode is %d", res.TxValidationCode)
	}
	return res
}

type FabricRespone struct {
	TxValidationCode int32  `json:"txValidationCode"`
	ChaincodeStatus  int32  `json:"chaincodeStatus"`
	TxId             string `json:"txId"`
	Payload          string `json:"payload"`
}

type OutPut struct {
	Header metadata.Header  `json:"header"`
	Body   FabricRespone `json:"body"`
}

func GetErrOutPut(header metadata.Header) string {

	outPut := &OutPut{
		Header: header,
		Body:   FabricRespone{},
	}

	jsonBytes, _ := json.Marshal(outPut)

	return string(jsonBytes)
}

func GetSuccessOutPut(header metadata.Header, res FabricRespone) string {
	outPut := &OutPut{
		Header: header,
		Body:   res,
	}
	jsonBytes, _ := json.Marshal(outPut)

	return string(jsonBytes)
}
