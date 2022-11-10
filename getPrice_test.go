package oracles

import (
	"fmt"
	"testing"

	"github.com/0xVanfer/chainId"
	"github.com/ethereum/go-ethereum/ethclient"
)

func TestGetPrice(t *testing.T) {
	client, err := ethclient.Dial("http://192.168.1.104:9655/ext/bc/C/rpc")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(GetPrice(chainId.AvalancheChainName, "eth", "usd", client))
}
