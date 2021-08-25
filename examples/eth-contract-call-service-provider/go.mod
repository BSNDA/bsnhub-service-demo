module eth-contract-call-service-provider

require (
	github.com/cockroachdb/pebble v0.0.0-20210406003833-3d4c32f510a8
	github.com/ethereum/go-ethereum v1.10.9-0.20210824190246-fe2f153b556a
	github.com/gin-gonic/gin v1.7.1
	github.com/go-sql-driver/mysql v1.4.1
	github.com/irisnet/service-sdk-go v1.0.0
	github.com/sirupsen/logrus v1.6.0
	github.com/spf13/cobra v1.1.1
	github.com/spf13/viper v1.7.1
)

replace (
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4
	github.com/tendermint/tendermint => github.com/bianjieai/tendermint v0.34.1-irita-210113
)

go 1.14
