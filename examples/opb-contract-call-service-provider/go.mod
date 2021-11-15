module opb-contract-call-service-provider

require (
	github.com/bianjieai/iritamod-sdk-go v0.0.0-20211115021642-7519137c6c9f
	github.com/cockroachdb/pebble v0.0.0-20210406003833-3d4c32f510a8
	github.com/gin-gonic/gin v1.7.1
	github.com/go-errors/errors v1.0.1
	github.com/go-sql-driver/mysql v1.4.0
	github.com/irisnet/core-sdk-go v0.0.0-20211104064902-e26c6107c96e
	github.com/sirupsen/logrus v1.6.0
	github.com/spf13/cobra v1.1.1
	github.com/spf13/viper v1.7.1
)

replace (
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4
	github.com/irisnet/core-sdk-go => github.com/irisnet/core-sdk-go v0.0.0-20211104064902-e26c6107c96e
	github.com/tendermint/tendermint => github.com/bianjieai/tendermint v0.34.1-irita-210113
)

go 1.16
