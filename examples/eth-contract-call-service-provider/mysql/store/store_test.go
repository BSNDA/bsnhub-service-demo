package store

import (
	"fmt"
	"testing"
)


func TestRelayerResponeRecord(t *testing.T) {
	data :=&RelayerResInfo{
		RequestId: "requestID",
		TxStatus: TxStatus_Success,
		ErrMsg: "",
	}
	defer func(d *RelayerResInfo) {

		fmt.Println(d.RequestId)
	}(data)


	data.RequestId = "123456"

}