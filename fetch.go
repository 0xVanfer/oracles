package oracles

import (
	"github.com/imroc/req"
)

type networkInfo struct {
	Title            string `json:"title"`
	NetworkStatusURL string `json:"networkStatusUrl"`
	Networks         []struct {
		Name        string `json:"name"`
		URL         string `json:"url"`
		DataType    string `json:"dataType"`
		NetworkType string `json:"networkType"`
		Proxies     []struct {
			Pair               string  `json:"pair"`
			AssetName          string  `json:"assetName"`
			DeviationThreshold float64 `json:"deviationThreshold"`
			Heartbeat          string  `json:"heartbeat"`
			Decimals           int     `json:"decimals"`
			Proxy              string  `json:"proxy"`
			FeedCategory       string  `json:"feedCategory"`
			FeedType           string  `json:"feedType"`
			ShutdownDate       string  `json:"shutdownDate,omitempty"`
		} `json:"proxies"`
	} `json:"networks"`
}

type chainlinkInfo map[string]networkInfo

func reqOracleAddresses() (allInfo chainlinkInfo, err error) {
	url := "https://cl-docs-addresses.web.app/addresses.json"
	r, _ := req.Get(url)
	err = r.ToJSON(&allInfo)
	return
}
