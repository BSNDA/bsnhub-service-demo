package main

import (
	"opb-contract-call-service-provider/mysql"
	txstore "opb-contract-call-service-provider/mysql/store"
	"github.com/spf13/cobra"

	"opb-contract-call-service-provider/app"
	"opb-contract-call-service-provider/common"
	contractservice "opb-contract-call-service-provider/contract-service"
	"opb-contract-call-service-provider/iservice"
	"opb-contract-call-service-provider/server"
	"opb-contract-call-service-provider/store"
)

const (
	_HttpPort = "base.http_port"
)

func StartCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "start",
		Short:   "Start provider daemon",
		Example: `irita-opb-provider start [config-file]`,
		Args:    cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			configFileName := ""

			if len(args) == 0 {
				configFileName = common.DefaultConfigFileName
			} else {
				configFileName = args[0]
			}

			config, err := common.LoadYAMLConfig(configFileName)
			if err != nil {
				return err
			}

			logger := common.Logger

			store, err := store.NewStore(config.GetString(common.ConfigKeyStorePath))
			if err != nil {
				return err
			}
			chainManager := server.NewChainManager(store)

			iserviceClient := iservice.MakeServiceClientWrapper(iservice.NewConfig(config))

			contractService, err := contractservice.BuildContractService(config, chainManager)
			if err != nil {
				return err
			}

			contractService.Logger = logger
			appInstance := app.NewApp(iserviceClient, contractService, logger)

			//set service name
			appInstance.SetServiceName(config.GetString(common.ConfigKeyServiceName))

			mysqlConfig := mysql.NewConfig(config)
			txstore.InitMysql(mysqlConfig.DSN())

			//mysql.NewDB(mysqlConfig)
			//defer mysql.Close()

			// set test api handle
			server.SetTestCallBack(contractService.Callback)

			httpPort := config.GetInt(_HttpPort)
			if httpPort == 0 {
				httpPort = 8083
			}

			go server.StartWebServer(chainManager, httpPort)
			appInstance.Start()

			return nil
		},
	}

	return cmd
}
