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
		"http://183.252.194.19:26657",
		"183.252.194.19:9090",
		"node0",
		"12345678",
		"-----BEGIN TENDERMINT PRIVATE KEY-----\nsalt: E9153622631E90E305B065101DD445A0\ntype: sm2\nkdf: bcrypt\n\nTUPYO8bYtJ6ZXAATk7G1+NvB99nrx1Gxj8jmwj1Stw8kiWP4jfnuaugYbJ/AQSAv\nAgfhA4RW/MbEUd/V64kMdpm1GEClZUS8HzXBlls=\n=TqK0\n-----END TENDERMINT PRIVATE KEY-----",
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
