package mysql

import (
	"fmt"
	"github.com/spf13/viper"

	cfg "github.com/bianjieai/bsnhub-service-demo/examples/opb-contract-call-service-provider/common"
)

const (
	Prefix           = "mysql"
	DBName           = "db_name"
	DBUserName       = "db_user_name"
	DBUserPassphrase = "db_user_passphrase"
	Host             = "host"
	Port             = "port"
)

// Config represents the Fabric chain config
type Config struct {
	DBName           string   `yaml:"db_name"`
	DBUserName       string   `yaml:"db_user_name"`
	DBUserPassphrase string `yaml:"db_user_passphrase"`
	Host             string   `yaml:"host"`
	Port             string   `yaml:"port"`
}

func (mysqlConfig *Config)DSN() string {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true",
		mysqlConfig.DBUserName, mysqlConfig.DBUserPassphrase, mysqlConfig.Host, mysqlConfig.Port, mysqlConfig.DBName)

	return dsn
}

// NewConfig constructs a new Config from viper
func NewConfig(v *viper.Viper) Config {
	return Config{
		DBName:           v.GetString(cfg.GetConfigKey(Prefix, DBName)),
		DBUserName:       v.GetString(cfg.GetConfigKey(Prefix, DBUserName)),
		DBUserPassphrase: v.GetString(cfg.GetConfigKey(Prefix, DBUserPassphrase)),
		Host:             v.GetString(cfg.GetConfigKey(Prefix, Host)),
		Port:             v.GetString(cfg.GetConfigKey(Prefix, Port)),
	}
}