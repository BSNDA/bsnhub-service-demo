package server

import (
	"encoding/json"
	"opb-contract-call-service-provider/contract-service/opb/config"
	"opb-contract-call-service-provider/store"
	"opb-contract-call-service-provider/types"
)

// ChainManager defines a service for app chains management
type ChainManager struct {
	store *store.Store
}

// NewChainManager constructs a new ChainManager instance
func NewChainManager(s *store.Store) *ChainManager {
	return &ChainManager{
		store: s,
	}
}

// AddChain adds a new app chain for the relayer
func (cm *ChainManager) AddChain(params []byte) (chainID string, err error) {
	chainID, err = types.GetChainIDFromBytes(params)
	if err != nil {
		return "", err
	}
	chainIDsbz, _ := cm.store.Get([]byte("chainIDs"))
	if chainIDsbz == nil {
		chainIDsbz, err = json.Marshal(map[string]string{})
		if err != nil {
			return "", err
		}
		err = cm.store.Set([]byte("chainIDs"), chainIDsbz)
		if err != nil {
			return "", err
		}
	} else {
		chainIDMap := map[string]bool{}
		json.Unmarshal(chainIDsbz, &chainIDMap)
		chainIDMap[chainID] = true
		bz, err := json.Marshal(chainIDMap)
		if err != nil {
			return "", err
		}
		err = cm.store.Set([]byte("chainIDs"), bz)
		if err != nil {
			return "", err
		}
	}
	err = cm.store.Set([]byte(chainID), params)
	if err != nil {
		return "", err
	}
	return chainID, nil
}

// GetChains gets all active app chains
func (cm *ChainManager) GetChains() ([]string, error) {
	chainIDMap := map[string]bool{}
	chainIDsbz, err := cm.store.Get([]byte("chainIDs"))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(chainIDsbz, &chainIDMap)
	if err != nil {
		return nil, err
	}
	chainIDs := []string{}
	for chainID, isexist := range chainIDMap {
		if isexist {
			chainIDs = append(chainIDs, chainID)
		}
	}
	return chainIDs, nil
}

// DeleteChain delete chain params by chain-id
func (cm *ChainManager) DeleteChain(chainID string) (err error) {
	chainIDMap := map[string]bool{}
	chainIDsbz, err := cm.store.Get([]byte("chainIDs"))
	if err != nil {
		return err
	}
	err = json.Unmarshal(chainIDsbz, &chainIDMap)
	if err != nil {
		return err
	}
	chainIDMap[chainID] = false
	bz, err := json.Marshal(chainIDMap)
	if err != nil {
		return err
	}
	err = cm.store.Set([]byte("chainIDs"), bz)
	if err != nil {
		return err
	}
	return cm.store.Delete([]byte(chainID))
}

// GetChainParams gets all chain params by chain-id
func (cm *ChainManager) GetChainParams(chainID string) (config.ChainParams, error) {
	var chainParams config.ChainParams
	chainParamsBz, err := cm.store.Get([]byte(chainID))
	if err != nil {
		return config.ChainParams{}, err
	}
	err = json.Unmarshal(chainParamsBz, &chainParams)
	if err != nil {
		return config.ChainParams{}, err
	}
	return chainParams, nil
}
