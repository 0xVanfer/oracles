package oracles

import (
	"strings"
	"time"

	"github.com/0xVanfer/chainId"
)

// Mapping pair to oracle address.
func getChainlinkCompleteMap() *ChainlinkOracleMap {
	completeMap := make(map[string]map[string]string)
	// Request the complete map.
	allInfo, err := reqOracleAddresses()
	if err != nil {
		return nil
	}
	// For every chain, create a smaller chainMap.
	for chain, info := range allInfo {
		chainMap := make(map[string]string)
		for _, networks := range info.Networks {
			// Only choose mainnet.
			if !strings.Contains(networks.Name, "Mainnet") || networks.DataType != "default" {
				continue
			}
			for _, proxy := range networks.Proxies {
				chainMap[proxy.Pair] = proxy.Proxy
			}
		}
		// Adapt to the chain names.
		for shouldbe, actually := range chainMapping {
			if strings.EqualFold(chain, actually) {
				chain = shouldbe
			}
		}
		completeMap[chain] = chainMap
	}
	return &ChainlinkOracleMap{
		AddressMap: completeMap,
		UpdateAt:   time.Now(),
	}
}

// Some chain naming is diffent from chainlist. Use a map to pair these chains.
var chainMapping = map[string]string{
	chainId.BinanceSmartChainName: "bnb-chain",
	chainId.HecoChainName:         "heco-chain",
}
