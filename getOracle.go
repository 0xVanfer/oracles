package oracles

import (
	"strings"

	"github.com/0xVanfer/chainId"
)

func GetChainlinkOraclesMapAll(currency string) map[string]map[string]string {
	mapp := make(map[string]map[string]string)
	allInfo, err := reqOracleAddresses()
	if err != nil {
		return nil
	}
	for network, info := range allInfo {
		mappp := make(map[string]string)
		_ = network
		for _, networks := range info.Networks {
			// must be Mainnet, default
			if !strings.Contains(networks.Name, "Mainnet") || networks.DataType != "default" {
				continue
			}
			for _, proxy := range networks.Proxies {
				if !strings.Contains(proxy.Pair, " / "+strings.ToUpper(currency)) {
					continue
				}
				symbol := strings.ReplaceAll(proxy.Pair, " / "+strings.ToUpper(currency), "")
				oracle := proxy.Proxy
				mappp[symbol] = oracle
			}
		}
		mapp[network] = mappp
	}
	return mapp
}

var chainmapping = map[string]string{
	chainId.BinanceSmartChainName: "bnb-chain",
	chainId.HecoChainName:         "heco-chain",
}

func GetChainlinkOraclesMap(network string, currency string) map[string]string {
	res := GetChainlinkOraclesMapAll(currency)[network]
	if res == nil {
		res = GetChainlinkOraclesMapAll(currency)[chainmapping[network]]
	}
	return res
}

func GetChainlinkOracle(network string, symbol string, currency string) string {
	return GetChainlinkOraclesMap(network, currency)[strings.ToUpper(symbol)]
}
