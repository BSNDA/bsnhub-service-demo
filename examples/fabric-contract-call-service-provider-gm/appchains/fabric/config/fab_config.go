package config

import (
	"bsn-irita-fabric-provider-gm/appchains/fabric/config/redconfig"
	"bsn-irita-fabric-provider-gm/appchains/fabric/config/redconfig/configbackend"
	"github.com/BSNDA/fabric-sdk-go-gm/pkg/common/providers/core"
)

type FabricConfig struct {
	SdkConfig string

	OrgName     string
	MspUserName string

	OrgCode string
}

func (f *FabricConfig) GetSdkConfig(channelId string, nodes []string) core.ConfigProvider {

	ch := configbackend.ChannelConfig{ChannelId: channelId, PeerName: nodes[0]}

	var s []redconfig.SetOption
	s = append(s, redconfig.SetChannel(&ch))

	configProvider := redconfig.FromFile(f.SdkConfig, s)
	return configProvider

}
