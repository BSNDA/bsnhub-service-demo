package test

import (
	"github.com/bianjieai/iritamod-sdk-go/wasm"
	sdk "github.com/irisnet/core-sdk-go/types"
	"opb-contract-call-service-provider/iservice"
	"testing"
	"time"
)

var svcClient iservice.ServiceClientWrapper

func init() {
	svcClient = iservice.NewServiceClientWrapper(
		"wenchangchain",
		"http://10.1.4.149:36657",
		"10.1.4.149:39090",
		"node0",
		"12345678",
		"-----BEGIN TENDERMINT PRIVATE KEY-----\nkdf: bcrypt\nsalt: E82064503E284EE753B13E9424B08B4C\ntype: sm2\n\nqLgix+DPFfNY+TpWWlNmquy3jUDR314/dJmIxw8JCWGiSn4deFtp8IWGH/mnVe6S\nNdGt6OJ2SbwO098fk16Gw6RO+MgVjShVMXbkggc=\n=h7AT\n-----END TENDERMINT PRIVATE KEY-----",
		"",
	)
}

func TestExecute(t *testing.T) {

	fee := sdk.NewDecCoins(
		sdk.NewDecCoin(
			"uirita",
			sdk.NewInt(400)))
	callData, err2 := wasm.NewContractABI().WithMethod("hello").WithArgs("words", "ori2").Build()
	if err2 != nil {
		t.Error(err2)
		return
	}
	resultTx, err := svcClient.IritaClient.WASM.Execute(
		"iaa14hj2tavq8fpesdwxxcu44rty3hh90vhudk32rt",
		wasm.NewContractABI().WithMethod("call_service").
			WithArgs("call_data", callData).
			WithArgs("endpoint_address", "iaa1nc5tatafv6eyq7llkr2gv50ff9e22mnfgrs38c").
			WithArgs("request_id", "aaaaaaa"+time.Now().String()),
		sdk.NewCoins(sdk.NewCoin("uirita", sdk.NewInt(1000000))),
		sdk.BaseTx{
			From:     "node0",
			Password: "12345678",
			Fee:      fee,
		},
	)

	if err != nil {
		t.Error(err)
		return
	}
	t.Log(resultTx)
	for _, e := range resultTx.TxResult.Events {
		if e.Type == "wasm" && len(e.Attributes) > 1 {
			for _, attr := range e.Attributes {
				if string(attr.Key) == "_contract_address" && string(attr.Value) == "iaa1nc5tatafv6eyq7llkr2gv50ff9e22mnfgrs38c" {
					for _, attr = range e.Attributes {
						if string(attr.Key) == "result" {
							t.Log(string(attr.Value))
						}
					}
				}
			}
		}
	}
}
