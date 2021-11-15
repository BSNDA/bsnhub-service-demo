package opb

import (
	"fmt"
	sdk "github.com/irisnet/core-sdk-go/types"
	storetypes "github.com/irisnet/core-sdk-go/types/store"
	log "github.com/sirupsen/logrus"
	"opb-contract-call-service-provider/common"
	opbcfg "opb-contract-call-service-provider/contract-service/opb/config"
	"opb-contract-call-service-provider/iservice"
	"opb-contract-call-service-provider/server"
)

const (
	defaultKeyAlgorithm = "sm2"
)

// OpbChain defines the Opb chain
type OpbChain struct {
	OpbClient    *iservice.ServiceClient
	ChainManager *server.ChainManager
	BaseConfig   opbcfg.BaseConfig
}

// NewOpbChain constructs a new OpbChain instance
func NewOpbChain(
	baseConfig opbcfg.BaseConfig,
	chainManager *server.ChainManager,
) *OpbChain {
	return &OpbChain{
		BaseConfig:   baseConfig,
		ChainManager: chainManager,
	}
}

// BuildBaseTx builds a base tx
func (opb *OpbChain) BuildBaseTx() sdk.BaseTx {
	return sdk.BaseTx{
		From:     opb.BaseConfig.Account.KeyName,
		Password: opb.BaseConfig.Account.Passphrase,
		Mode:     sdk.Commit,
	}
}

// InstantiateClient instantiates the opb client according to the given chain params
func (f *OpbChain) InstantiateClient(
	chainParams opbcfg.ChainParams,
) error {
	config := opbcfg.Config{
		BaseConfig:  f.BaseConfig,
		ChainParams: chainParams,
	}

	//将接口传递的节点名称通过配置转换为 节点地址，如果不在配置中，不转换
	//随机取一个传入的node
	nodeName := opbcfg.RandURL(config.ChainParams.NodeURLs)
	var rpcAddr string
	var grpcAddr string
	//获取配置的nodeURL
	rpcAddrstr, ok := config.RpcAddrsMap[nodeName]
	if ok {
		rpcAddr = rpcAddrstr
	}
	grpcAddrstr, ok := config.GrpcAddrsMap[nodeName]
	if ok {
		grpcAddr = grpcAddrstr
	}
	fees, err := sdk.ParseDecCoins(config.DefaultFee)
	if err != nil {
		return err
	}
	options := []sdk.Option{
		sdk.CachedOption(true),
		sdk.KeyDAOOption(storetypes.NewMemory(nil)),
		sdk.FeeOption(fees),
		sdk.GasOption(config.DefaultGas),
		sdk.TimeoutOption(config.Timeout),
		sdk.AlgoOption(defaultKeyAlgorithm),
	}

	clientConfig, err := sdk.NewClientConfig(rpcAddr, grpcAddr, config.BaseConfig.ChainId, options...)

	if err != nil {
		common.Logger.Errorf("failed to get the sdk clientConfig: %s", err)
		return fmt.Errorf("failed to get the sdk clientConfig: %s", err)
	}

	opbClient := iservice.NewServiceClient(clientConfig)
	f.OpbClient = opbClient

	// import
	addr, err := f.OpbClient.Import(config.Account.KeyName, config.Account.Passphrase, config.Account.KeyArmor)
	if err != nil {
		return fmt.Errorf("opb chain import key failed: %s", err)
	}
	log.WithField("addr", addr).Info("import key success")
	return nil
}

// waitForSuccess waits for the receipt of the given tx
func (opb *OpbChain) WaitForSuccess(txHash string, name string) error {
	common.Logger.Infof("%s: transaction sent to %s, hash: %s", name, opb.BaseConfig.ChainId, txHash)

	tx, _ := opb.OpbClient.QueryTx(txHash)
	if tx.TxResult.Code != 0 {
		return fmt.Errorf("transaction %s execution failed: %s", txHash, tx.TxResult.Log)
	}

	common.Logger.Infof("%s: transaction %s execution succeeded", name, txHash)

	return nil
}
