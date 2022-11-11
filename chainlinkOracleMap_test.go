package oracles

import (
	"fmt"
	"testing"

	"github.com/0xVanfer/chainId"
)

func TestMap(t *testing.T) {
	mapp := NewChainlinkMap()
	fmt.Println(mapp.UpdateAt)
	mapp.UpdateMap()
	fmt.Println(mapp.UpdateAt)
	address := mapp.ChooseOracle(Selectors{Network: chainId.AvalancheChainName, Symbol: "avax"})
	fmt.Println(address)
}
