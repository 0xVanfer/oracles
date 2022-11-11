package oracles

import (
	"fmt"
	"strings"
	"time"

	"github.com/0xVanfer/chainId"
)

// Complete map of chainlink oracles.
type ChainlinkOracleMap struct {
	AddressMap map[string]map[string]string
	UpdateAt   time.Time
}

// Used to choose oracle from the complete map.
type Selectors struct {
	Network  string // Default: "ethereum".
	Currency string // Default: "USD". Can be "ETH", "AVAX", etc.
	Symbol   string // Must not be empty.
}

// Set the map to up to date status.
func (m *ChainlinkOracleMap) UpdateMap() {
	newMap := getChainlinkCompleteMap()
	if newMap == nil {
		fmt.Println("update chainlink oracles map failed")
		return
	}
	*m = *newMap
}

// Get the specific oracle base on selectors. Symbol must not be empty.
//
// If oracle not found, will return "".
func (m *ChainlinkOracleMap) ChooseOracle(selectors Selectors) string {
	// Network is defaulted as ethereum.
	network := chainId.EthereumChainName
	// Currency is defaulted as "USD".
	currency := "USD"
	symbol := strings.ToUpper(selectors.Symbol)
	// Symbol not defined, return empty string.
	if symbol == "" {
		return ""
	}
	if selectors.Network != "" {
		network = selectors.Network
	}
	if selectors.Currency != "" {
		network = strings.ToUpper(selectors.Currency)
	}
	chainMap := (*m).AddressMap[network]
	return chainMap[symbol+" / "+currency]
}
