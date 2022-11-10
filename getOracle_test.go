package oracles

import (
	"fmt"

	"github.com/0xVanfer/chainId"
)

func ExampleGetChainlinkOracle() {
	fmt.Println(GetChainlinkOracle(chainId.AvalancheChainName, "aave", "usd"))
	fmt.Println(GetChainlinkOracle(chainId.AvalancheChainName, "link", "avax"))
	fmt.Println(GetChainlinkOracle(chainId.EthereumChainName, "1inch", "eth"))
	// Output:
	// 0x3CA13391E9fb38a75330fb28f8cc2eB3D9ceceED
	// 0x1b8a25F73c9420dD507406C3A3816A276b62f56a
	// 0x72AFAECF99C9d9C8215fF44C77B94B99C28741e8
}
