package eth

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcmn "github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"strings"

	"eth-contract-call-service-provider/common"
	ethcfg "eth-contract-call-service-provider/contract-service/eth/config"
	"eth-contract-call-service-provider/server"
)

// EthChain defines the Eth chain
type EthChain struct {
	Client              *ethclient.Client
	ChainManager        *server.ChainManager
	BaseConfig          ethcfg.BaseConfig
	TargetCoreContract *TargetCoreEx // iService Core Extension contract session
	TargetCoreABI     abi.ABI                // parsed iService Core Extension ABI
}

// NewEthChain constructs a new EthChain instance
func NewEthChain(
	baseConfig ethcfg.BaseConfig,
	chainManager *server.ChainManager,
) *EthChain {
	return &EthChain{
		BaseConfig:   baseConfig,
		ChainManager: chainManager,
	}
}

// InstantiateClient instantiates the eth client according to the given chain params
func (ec *EthChain) InstantiateClient(
	chainParams ethcfg.ChainParams,
) error {
	config := ethcfg.Config{
		BaseConfig:  ec.BaseConfig,
		ChainParams: chainParams,
	}
	//将接口传递的节点名称通过配置转换为 节点地址，如果不在配置中，不转换
	//随机取一个传入的node
	nodeName := ethcfg.RandURL(config.NodeURL)
	//获取配置的nodeURL
	nodeUrl, ok := config.NodesMap[nodeName]
	if ok {
		nodeName = nodeUrl
	}

	client, err := ethclient.Dial(nodeUrl)
	if err != nil {
		common.Logger.Errorf("failed to connect to eth node: %s", err)
		return fmt.Errorf("failed to connect to eth node: %s", err)
	}

	targetCore, err := NewTargetCoreEx(ethcmn.HexToAddress(chainParams.TargetCoreAddr), client)
	if err != nil {
		common.Logger.Errorf("failed to instantiate the iservice core contract: %s", err)
	}

	targetCoreABI, err := abi.JSON(strings.NewReader(TargetCoreExABI))
	if err != nil {
		return fmt.Errorf("failed to parse iService Core Extension ABI: %s", err)
	}
	ec.Client = client
	ec.TargetCoreContract = targetCore
	ec.TargetCoreABI = targetCoreABI
	return nil
}

// CallContract calls the specified contract with the given contract address and data

// WaitForReceipt waits for the receipt of the given tx
func (ec *EthChain) WaitForReceipt(tx *ethtypes.Transaction, name string) (*ethtypes.Receipt, error) {
	common.Logger.Infof("%s: transaction sent, hash: %s", name, tx.Hash().Hex())

	receipt, err := bind.WaitMined(context.Background(), ec.Client, tx)
	if err != nil {
		return nil, fmt.Errorf("failed to mint the transaction %s: %s", tx.Hash().Hex(), err)
	}

	if receipt.Status != ethtypes.ReceiptStatusSuccessful {
		return nil, fmt.Errorf("transaction %s execution failed", tx.Hash().Hex())
	}

	common.Logger.Infof("%s: transaction %s execution succeeded", name, tx.Hash().Hex())

	return receipt, nil
}

// BuildAuthTransactor builds an authenticated transactor
func (ec *EthChain) BuildAuthTransactor() (*bind.TransactOpts, error) {
	privKey, err := crypto.HexToECDSA(ec.BaseConfig.Key)
	if err != nil {
		return nil, err
	}

	auth := bind.NewKeyedTransactor(privKey)

	nextNonce, err := ec.Client.PendingNonceAt(context.Background(), auth.From)
	if err != nil {
		return nil, err
	}

	auth.GasLimit = ec.BaseConfig.GasLimit
	auth.GasPrice = big.NewInt(int64(ec.BaseConfig.GasPrice))
	auth.Nonce = big.NewInt(int64(nextNonce))

	return auth, nil
}
