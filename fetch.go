package oracles

import (
	"github.com/imroc/req"
)

type chainlinkReq struct {
	Title            string        `json:"title"`
	NetworkStatusURL string        `json:"networkStatusUrl"`
	Networks         []networkInfo `json:"networks"`
}

// Example:
//
//	{
//		"name": "Ethereum Mainnet",
//		"url": "https://etherscan.io/address/%s",
//		"dataType": "default",
//		"networkType": "mainnet",
//		"proxies": [
//		{
//		"pair": "1INCH / ETH",
//		"assetName": "1inch",
//		"deviationThreshold": 2,
//		"heartbeat": "24h",
//		"decimals": 18,
//		"proxy": "0x72AFAECF99C9d9C8215fF44C77B94B99C28741e8",
//		"feedCategory": "verified",
//		"feedType": "Crypto"
//		},
//		...
//		]
//	}
type networkInfo struct {
	Name        string       `json:"name"`
	URL         string       `json:"url"`
	DataType    string       `json:"dataType"`
	NetworkType string       `json:"networkType"`
	Proxies     []singlePair `json:"proxies"`
}

type singlePair struct {
	Pair               string  `json:"pair"`
	AssetName          string  `json:"assetName"`
	DeviationThreshold float64 `json:"deviationThreshold"`
	Heartbeat          string  `json:"heartbeat"`
	Decimals           int     `json:"decimals"`
	Proxy              string  `json:"proxy"`
	FeedCategory       string  `json:"feedCategory"`
	FeedType           string  `json:"feedType"`
	ShutdownDate       string  `json:"shutdownDate,omitempty"`
}

// Request result of chainlink oracle addresses.
type chainlinkInfo map[string]chainlinkReq

// Request all the chainlink oracle addresses.
func reqOracleAddresses() (allInfo chainlinkInfo, err error) {
	url := "https://cl-docs-addresses.web.app/addresses.json"
	r, _ := req.Get(url)
	err = r.ToJSON(&allInfo)
	return
}
