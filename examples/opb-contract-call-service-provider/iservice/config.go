package iservice

import (
	"fmt"
	sdk "github.com/irisnet/core-sdk-go/types"
	"os"
	"path/filepath"

	"github.com/spf13/viper"

	"opb-contract-call-service-provider/common"
)

// default config variables
var (
	defaultChainID       = "irita-hub"
	defaultNodeRPCAddr   = "http://127.0.0.1:26657"
	defaultNodeGRPCAddr  = "127.0.0.1:9090"
	defaultKeyPath       = os.ExpandEnv(filepath.Join("$HOME", ".iritacli"))
	defaultGas           = uint64(200000)
	defaultFee           = "4point"
	defaultBroadcastMode = sdk.Commit
	defaultKeyAlgorithm  = "sm2"
)

const (
	Prefix       = "iservice"
	ChainID      = "chain_id"
	NodeRPCAddr  = "node_rpc_addr"
	NodeGRPCAddr = "node_grpc_addr"
	AccountKey   = "account"
	Fee          = "fee"
	KeyPath      = "key_path"
	KeyName      = "key_name"
	Passphrase   = "passphrase"
)

// Config is a config struct for iservice
type Config struct {
	ChainID      string `yaml:"chain_id"`
	NodeRPCAddr  string `yaml:"node_rpc_addr"`
	NodeGRPCAddr string `yaml:"node_grpc_addr"`

	NodesMap map[string]string `yaml:"nodes"`
	Account  Account           `yaml:"account"`

	Fee string `yaml:"fee"`
}

type Account struct {
	KeyName    string `yaml:"key_name" mapstructure:"key_name"`
	Passphrase string `yaml:"passphrase"`
	KeyArmor   string `yaml:"key_armor" mapstructure:"key_armor"`
}

// NewConfig constructs a new Config from viper
func NewConfig(v *viper.Viper) Config {
	account := Account{}
	err := v.UnmarshalKey(common.GetConfigKey(Prefix, AccountKey), &account)
	if err != nil {
		fmt.Println(err)
	}
	return Config{
		ChainID:      v.GetString(common.GetConfigKey(Prefix, ChainID)),
		NodeRPCAddr:  v.GetString(common.GetConfigKey(Prefix, NodeRPCAddr)),
		NodeGRPCAddr: v.GetString(common.GetConfigKey(Prefix, NodeGRPCAddr)),
		Account:      account,
		Fee:          v.GetString(common.GetConfigKey(Prefix, Fee)),
	}
}
