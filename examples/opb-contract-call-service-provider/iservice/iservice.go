package iservice

import (
	"encoding/json"
	"github.com/go-errors/errors"
	log "github.com/sirupsen/logrus"

	"github.com/bianjieai/iritamod-sdk-go/service"
	"github.com/bianjieai/iritamod-sdk-go/wasm"

	"github.com/irisnet/core-sdk-go/client"
	"github.com/irisnet/core-sdk-go/codec"
	cdctypes "github.com/irisnet/core-sdk-go/codec/types"
	cryptocodec "github.com/irisnet/core-sdk-go/crypto/codec"
	"github.com/irisnet/core-sdk-go/modules/bank"
	sdk "github.com/irisnet/core-sdk-go/types"
	storetypes "github.com/irisnet/core-sdk-go/types/store"
	txtypes "github.com/irisnet/core-sdk-go/types/tx"

	txStore "opb-contract-call-service-provider/mysql/store"
	"opb-contract-call-service-provider/types"
)

const (
	eventTypeNewBatchRequestProvider = "new_batch_request_provider"
	attributeKeyServiceName          = "service_name"
	attributeKeyProvider             = "provider"
	attributeKeyRequests             = "requests"
	attributeKeyRequestID            = "request_id"
)

// ServiceClientWrapper defines a wrapper for service client
type ServiceClientWrapper struct {
	ChainID     string
	NodeRPCAddr string

	KeyPath    string
	KeyName    string
	Passphrase string

	IritaClient *ServiceClient
}

type ServiceClient struct {
	encodingConfig sdk.EncodingConfig
	sdk.BaseClient
	Bank    bank.Client
	Service service.Client
	WASM    wasm.Client
}

func NewServiceClient(config sdk.ClientConfig) *ServiceClient {
	encodingConfig := makeEncodingConfig()
	baseClient := client.NewBaseClient(config, encodingConfig, nil)
	bankClient := bank.NewClient(baseClient, encodingConfig.Codec)
	serviceClient := service.NewClient(baseClient, encodingConfig.Codec)
	wasmClient := wasm.NewClient(baseClient)
	sc := &ServiceClient{
		encodingConfig: encodingConfig,
		BaseClient:     baseClient,
		Bank:           bankClient,
		Service:        serviceClient,
		WASM:           wasmClient,
	}

	sc.RegisterModule(
		bankClient,
		serviceClient,
		wasmClient,
	)

	return sc
}

func (sc *ServiceClient) RegisterModule(ms ...sdk.Module) {
	for _, m := range ms {
		m.RegisterInterfaceTypes(sc.encodingConfig.InterfaceRegistry)
	}
}

//client init
func makeEncodingConfig() sdk.EncodingConfig {
	amino := codec.NewLegacyAmino()
	interfaceRegistry := cdctypes.NewInterfaceRegistry()
	marshaler := codec.NewProtoCodec(interfaceRegistry)
	txCfg := txtypes.NewTxConfig(marshaler, txtypes.DefaultSignModes)

	encodingConfig := sdk.EncodingConfig{
		InterfaceRegistry: interfaceRegistry,
		TxConfig:          txCfg,
		Amino:             amino,
		Codec:             marshaler,
	}
	registerLegacyAminoCodec(encodingConfig.Amino)
	registerInterfaces(encodingConfig.InterfaceRegistry)
	return encodingConfig
}

// RegisterLegacyAminoCodec registers the sdk message type.
func registerLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterInterface((*sdk.Msg)(nil), nil)
	cdc.RegisterInterface((*sdk.Tx)(nil), nil)
	cryptocodec.RegisterCrypto(cdc)
}

// RegisterInterfaces registers the sdk message type.
func registerInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterInterface("cosmos.v1beta1.Msg", (*sdk.Msg)(nil))
	txtypes.RegisterInterfaces(registry)
	cryptocodec.RegisterInterfaces(registry)
}

// NewServiceClientWrapper constructs a new ServiceClientWrapper
func NewServiceClientWrapper(
	chainID string,
	nodeRPCAddr string,
	nodeGRPCAddr string,
	keyName string,
	passphrase string,
	keyArmor string,
	feeString string,
) ServiceClientWrapper {
	if len(chainID) == 0 {
		chainID = defaultChainID
	}

	if len(nodeRPCAddr) == 0 {
		nodeRPCAddr = defaultNodeRPCAddr
	}

	if len(nodeGRPCAddr) == 0 {
		nodeGRPCAddr = defaultNodeGRPCAddr
	}

	if keyName == "" || passphrase == "" || keyArmor == "" {
		panic("account miss: key_name or passphrase or KeyArmor is missing")
	}

	if len(feeString) == 0 {
		feeString = defaultFee
	}
	fee, err := sdk.ParseDecCoins(feeString)
	if err != nil {
		panic(err)
	}

	config, err := sdk.NewClientConfig(
		nodeRPCAddr,
		nodeGRPCAddr,
		chainID,
		sdk.FeeOption(fee),
		sdk.GasOption(defaultGas),
		sdk.ModeOption(defaultBroadcastMode),
		sdk.AlgoOption(defaultKeyAlgorithm),
		sdk.KeyDAOOption(storetypes.NewMemory(nil)),
		sdk.TimeoutOption(5),
	)
	if err != nil {
		panic(err)
	}

	wrapper := ServiceClientWrapper{
		ChainID:     chainID,
		NodeRPCAddr: nodeRPCAddr,
		KeyName:     keyName,
		Passphrase:  passphrase,
		IritaClient: NewServiceClient(config),
	}
	addr, err := wrapper.ImportKey(keyName, passphrase, keyArmor)
	if err != nil {
		panic("import key missing: " + err.Error())
	}
	log.WithField("addr", addr).Info("import key success")

	return wrapper
}

// MakeServiceClientWrapper builds a ServiceClientWrapper from the given config
func MakeServiceClientWrapper(config Config) ServiceClientWrapper {
	return NewServiceClientWrapper(
		config.ChainID,
		config.NodeRPCAddr,
		config.NodeGRPCAddr,
		config.Account.KeyName,
		config.Account.Passphrase,
		config.Account.KeyArmor,
		config.Fee,
	)
}

// SubscribeServiceRequest wraps service.SubscribeServiceRequest
func (s ServiceClientWrapper) SubscribeServiceRequest(serviceName string, cb service.RespondCallback) error {
	baseTx := s.BuildBaseTx()
	provider, e := s.IritaClient.QueryAddress(baseTx.From, baseTx.Password)
	if e != nil {
		return errors.New(e)
	}

	builder := sdk.NewEventQueryBuilder().AddCondition(
		sdk.NewCond(eventTypeNewBatchRequestProvider, attributeKeyServiceName).EQ(sdk.EventValue(serviceName)),
	).AddCondition(
		sdk.NewCond(eventTypeNewBatchRequestProvider, attributeKeyProvider).EQ(sdk.EventValue(provider.String())),
	)

	_, err := s.IritaClient.SubscribeNewBlock(builder, func(block sdk.EventDataNewBlock) {
		msgs := s.GenServiceResponseMsgs(block.ResultEndBlock.Events, serviceName, provider, cb)
		if msgs == nil || len(msgs) == 0 {
			s.IritaClient.Logger().Error("no message created",
				"serviceName", serviceName,
				"provider", provider,
			)
		}
		for _, msg := range msgs {
			msg, ok := msg.(*service.MsgRespondService)

			data := &txStore.ProviderResInfo{
				TxStatus: txStore.TxStatus_Success,
				ErrMsg:   "",
			}

			if ok {
				data.IcRequestId = msg.RequestId
			}

			resTx, err := s.IritaClient.BuildAndSend([]sdk.Msg{msg}, baseTx)
			if err != nil {
				data.TxStatus = txStore.TxStatus_Error
				data.ErrMsg = err.Error()
				s.IritaClient.Logger().Error("provider respond failed", "errMsg", err.Error())
				//mysql.TxErrCollection(msg.RequestId, err.Error())
			} else {
				data.HUBResTxId = resTx.Hash.String()
				//mysql.OnInterchainResponseSent(msg.RequestId, resTx.Hash)
			}

			txStore.ProviderCallBackTransRecord(data)
		}
	})
	return err
}

func (s ServiceClientWrapper) GenServiceResponseMsgs(events sdk.StringEvents, serviceName string,
	provider sdk.AccAddress,
	handler service.RespondCallback) (msgs []sdk.Msg) {

	var ids []string
	for _, e := range events {
		if e.Type != eventTypeNewBatchRequestProvider {
			continue
		}
		attributes := sdk.Attributes(e.Attributes)
		svcName := attributes.GetValue(attributeKeyServiceName)
		prov := attributes.GetValue(attributeKeyProvider)
		if svcName == serviceName && prov == provider.String() {
			reqIDsStr := attributes.GetValue(attributeKeyRequests)
			var idsTemp []string
			if err := json.Unmarshal([]byte(reqIDsStr), &idsTemp); err != nil {
				s.IritaClient.Logger().Error(
					"service request don't exist",
					attributeKeyRequestID, reqIDsStr,
					attributeKeyServiceName, serviceName,
					attributeKeyProvider, provider.String(),
					"errMsg", err.Error(),
				)
				return
			}
			ids = append(ids, idsTemp...)
		}
	}

	for _, reqID := range ids {
		request, err := s.IritaClient.Service.QueryServiceRequest(reqID)
		if err != nil {
			s.IritaClient.Logger().Error(
				"service request don't exist",
				attributeKeyRequestID, reqID,
				attributeKeyServiceName, serviceName,
				attributeKeyProvider, provider.String(),
				"errMsg", err.Error(),
			)
			continue
		}
		//check again
		providerStr := provider.String()
		if providerStr == request.Provider && request.ServiceName == serviceName {
			output, result := handler(request.RequestContextID, reqID, request.Input)
			var resultObj types.Result
			json.Unmarshal([]byte(result), &resultObj)
			if resultObj.Code != 204 {
				msgs = append(msgs, &service.MsgRespondService{
					RequestId: reqID,
					Provider:  providerStr,
					Output:    output,
					Result:    result,
				})
			}
		}
	}
	return msgs
}

// DefineService wraps iservice.DefineService
func (s ServiceClientWrapper) DefineService(
	serviceName string,
	description string,
	authorDescription string,
	tags []string,
	schemas string,
) error {
	request := service.DefineServiceRequest{
		ServiceName:       serviceName,
		Description:       description,
		AuthorDescription: authorDescription,
		Tags:              tags,
		Schemas:           schemas,
	}

	_, err := s.IritaClient.Service.DefineService(request, s.BuildBaseTx())

	return err
}

// BindService wraps iservice.BindService
func (s ServiceClientWrapper) BindService(
	serviceName string,
	deposit string,
	pricing string,
	options string,
	qos uint64,
) error {
	depositCoins, err := sdk.ParseDecCoins(deposit)
	if err != nil {
		return err
	}

	provider, err := s.ShowKey(s.KeyName, s.Passphrase)
	if err != nil {
		return err
	}

	request := service.BindServiceRequest{
		ServiceName: serviceName,
		Deposit:     depositCoins,
		Pricing:     pricing,
		Options:     options,
		QoS:         qos,
		Provider:    provider,
	}

	_, err = s.IritaClient.Service.BindService(request, s.BuildBaseTx())

	return err
}

// BuildBaseTx builds a base tx
func (s ServiceClientWrapper) BuildBaseTx() sdk.BaseTx {
	return sdk.BaseTx{
		From:     s.KeyName,
		Password: s.Passphrase,
	}
}
