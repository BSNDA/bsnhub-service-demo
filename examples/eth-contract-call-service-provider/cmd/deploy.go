package main

import (
	"io/ioutil"

	"github.com/spf13/cobra"

	"eth-contract-call-service-provider/app"
	"eth-contract-call-service-provider/common"
	contractservice "eth-contract-call-service-provider/contract-service"
	"eth-contract-call-service-provider/iservice"
	"eth-contract-call-service-provider/server"
	"eth-contract-call-service-provider/store"
)

func DeployCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "deploy",
		Short:   "Deploy iservice according to the metadata",
		Example: `irita-eth-provider deploy [config-file]`,
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

			schemas, err := ioutil.ReadFile("./metadata/service-definition.json")
			if err != nil {
				return err
			}

			pricing, err := ioutil.ReadFile("./metadata/service-binding.json")
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
			appInstance.SetServiceName(config.GetString(common.ConfigKeyServiceName))
			err = appInstance.DeployIService(string(schemas), string(pricing))
			if err != nil {
				return err
			}

			return nil
		},
	}

	return cmd
}
