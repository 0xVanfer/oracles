package oracles

import (
	"github.com/0xVanfer/abigen/chainlink/chainlinkOracle"
	"github.com/0xVanfer/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/shopspring/decimal"
)

// Get the price by chainlink oracle.
//
// Return value already divided by decimals.
func GetPriceByChainlink(oracleAddr string, client bind.ContractBackend) (decimal.Decimal, error) {
	oracle, _ := chainlinkOracle.NewChainlinkOracle(types.ToAddress(oracleAddr), client)
	res, err := oracle.LatestAnswer(nil)
	if err != nil {
		return decimal.Zero, err
	}
	decimals, err := oracle.Decimals(nil)
	if err != nil {
		return decimal.Zero, err
	}
	return types.ToDecimal(res).Div(decimal.New(1, int32(decimals))), nil
}
