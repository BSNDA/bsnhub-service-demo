package test

import (
	"github.com/bianjieai/iritamod-sdk-go/wasm"
	sdk "github.com/irisnet/core-sdk-go/types"
	"opb-contract-call-service-provider/iservice"
	"testing"
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
	)
}

func TestExecute(t *testing.T) {

	fee := sdk.NewDecCoins(
		sdk.NewDecCoin(
			"upoint",
			sdk.NewInt(400)))

	resultTx, err := svcClient.IritaClient.WASM.Execute(
		"iaa1nc5tatafv6eyq7llkr2gv50ff9e22mnfgrs38c",
		wasm.NewContractABI().WithMethod("hello").WithArgs("words", "ori"),
		sdk.NewCoins(sdk.NewCoin("upoint", sdk.NewInt(1000000))),
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
}
