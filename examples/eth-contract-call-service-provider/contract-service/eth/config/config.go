package config

import (
	"encoding/json"
	"github.com/spf13/viper"
	"math/rand"

	"eth-contract-call-service-provider/common"
)

const (
	Prefix = "eth"

	ChainID         = "chain_id"
	GasLimit        = "gas_limit"
	GasPrice        = "gas_price"
	Key             = "key"
	Passphrase      = "passphrase"
	Nodes           = "nodes"

	TargetEventName  = "iservice_event_name"
	TargetEventSig   = "iservice_event_sig"
)

// BaseConfig defines the base config
type BaseConfig struct {
	ChainID         string            `yaml:"chain_id"`
	GasLimit        uint64            `yaml:"gas_limit"`
	GasPrice        uint64            `yaml:"gas_price"`
	Key             string            `yaml:"key"`
	Passphrase      string            `yaml:"passphrase"`
	NodesMap        map[string]string `yaml:"nodes"`
	MonitorInterval uint64
	IServiceEventName  string `yaml:"iservice_event_name"`
	IServiceEventSig   string `yaml:"iservice_event_sig"`
}

// ChainParams defines the params for the specific chain
type ChainParams struct {
	NodeURL        []string `json:"nodes"`
	ChainID        string    `json:"chainId"`
	TargetCoreAddr string   `json:"targetCoreAddr"`
}

// Config defines the specific chain config
type Config struct {
	BaseConfig
	ChainParams
}

// NewBaseConfig constructs a new BaseConfig instance from viper
func NewBaseConfig(v *viper.Viper) (*BaseConfig, error) {
	return &BaseConfig{
		ChainID:         v.GetString(common.GetConfigKey(Prefix, ChainID)),
		GasLimit:        v.GetUint64(common.GetConfigKey(Prefix, GasLimit)),
		GasPrice:        v.GetUint64(common.GetConfigKey(Prefix, GasPrice)),
		Key:             v.GetString(common.GetConfigKey(Prefix, Key)),
		Passphrase:      v.GetString(common.GetConfigKey(Prefix, Passphrase)),
		NodesMap:        v.GetStringMapString(common.GetConfigKey(Prefix, Nodes)),
		IServiceEventName:  v.GetString(common.GetConfigKey(Prefix, TargetEventName)),
		IServiceEventSig:   v.GetString(common.GetConfigKey(Prefix, TargetEventSig)),
	}, nil
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
