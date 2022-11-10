package oracles

import (
	"errors"

	"github.com/0xVanfer/abigen/chainlink/chainlinkOracle"
	"github.com/0xVanfer/chainId"
	"github.com/0xVanfer/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/shopspring/decimal"
)

func GetPrice(network, symbol, currency string, client bind.ContractBackend) (decimal.Decimal, error) {
	oracleAddr := GetChainlinkOracle(chainId.AvalancheChainName, "eth", "usd")
	if oracleAddr == "" {
		return decimal.Zero, errors.New("no such oracle")
	}
	return GetPriceByOracle(oracleAddr, client)
}

func GetPriceByOracle(oracleAddr string, client bind.ContractBackend) (decimal.Decimal, error) {
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
