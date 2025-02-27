package config

import (
	"encoding/json"
	"math/rand"
	"strconv"
	"strings"

	"github.com/spf13/viper"

	"opb-contract-call-service-provider/common"
)

const (
	Prefix = "opb"

	ChainId      = "chain_id"
	KeyPath      = "key_path"
	KeyName      = "key_name"
	Passphrase   = "passphrase"
	RpcAddrsMap  = "rpc_addrs"
	GrpcAddrsMap = "grpc_addrs"
	DefaultFee   = "default_fee"
	Timeout      = "timeout"
	DefaultGas   = "default_gas"
	AccountKey   = "account"
)

// BaseConfig defines the base config
type BaseConfig struct {
	Account      Account `yaml:"account"`
	RpcAddrsMap  map[string]string
	GrpcAddrsMap map[string]string
	ChainId      string
	DefaultFee   string
	Timeout      uint
	DefaultGas   uint64
}

type Account struct {
	KeyName    string `yaml:"key_name" mapstructure:"key_name"`
	Passphrase string `yaml:"passphrase"`
	KeyArmor   string `yaml:"key_armor" mapstructure:"key_armor"`
}

// ChainParams defines the params for the specific chain
type ChainParams struct {
	NodeURLs       []string `json:"nodes"`
	ChainID        int64    `json:"chainId"`
	TargetCoreAddr string   `json:"targetCoreAddr"`
}

// Config defines the specific chain config
type Config struct {
	BaseConfig
	ChainParams
}

// NewBaseConfig constructs a new BaseConfig instance from viper
func NewBaseConfig(v *viper.Viper) (*BaseConfig, error) {
	account := Account{}
	err := v.UnmarshalKey(common.GetConfigKey(Prefix, AccountKey), &account)
	if err != nil {
		return nil, err
	}
	config := new(BaseConfig)
	config.ChainId = v.GetString(common.GetConfigKey(Prefix, ChainId))
	config.RpcAddrsMap = v.GetStringMapString(common.GetConfigKey(Prefix, RpcAddrsMap))
	config.GrpcAddrsMap = v.GetStringMapString(common.GetConfigKey(Prefix, GrpcAddrsMap))
	config.DefaultFee = v.GetString(common.GetConfigKey(Prefix, DefaultFee))
	config.Timeout = v.GetUint(common.GetConfigKey(Prefix, Timeout))
	config.DefaultGas = v.GetUint64(common.GetConfigKey(Prefix, DefaultGas))
	config.Account = account

	return config, nil
}

// ValidateBaseConfig validates if the given bytes is valid BaseConfig
func ValidateBaseConfig(baseCfg []byte) error {
	var baseConfig BaseConfig
	return json.Unmarshal(baseCfg, &baseConfig)
}

func RandURL(m []string) string {
	if len(m) == 0 {
		return ""
	}
	for _, index := range rand.Perm(len(m)) {
		return m[index]
	}
	return ""
}

// GetChainID returns the unique opb chain id from the ChainID
func GetOpbChainID(ChainID string) int64 {
	chainInfos := strings.Split(ChainID, "-")
	opbChainID, _ := strconv.ParseInt(chainInfos[2], 10, 64)
	return opbChainID
}

// GetGroupID returns the unique opb group id from the ChainID
func GetOpbGroupID(ChainID string) int {
	chainInfos := strings.Split(ChainID, "-")
	opbGroupID, _ := strconv.Atoi(chainInfos[1])
	return opbGroupID
}
