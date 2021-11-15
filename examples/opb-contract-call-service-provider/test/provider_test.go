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
		"-----BEGIN TENDERMINT PRIVATE KEY-----\nkdf: bcrypt\nsalt: 28E4503C0FA024135BD8D06A457EB235\ntype: sm2\n\nJxEV2LpJzkgHodehJNEun73zj9aMRRXJh3g/DGrTR7Mz5DOvEFSR7nd9aXP+i0in\nBAqhpYt9hb6H/MbNty2kfqTBTcSTHAOzES8R1KQ=\n=HZRz\n-----END TENDERMINT PRIVATE KEY-----",
	)
}

func TestExecute(t *testing.T) {

	fee := sdk.NewDecCoins(
		sdk.NewDecCoin(
			"upoint",
			sdk.NewInt(400)))

	resultTx, err := svcClient.IritaClient.WASM.Execute(
		"iaa142r3lvrefe0xl0h6yvyql9d0fmvawmp5wwgp2j",
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
