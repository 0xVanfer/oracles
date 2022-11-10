package oracles

import (
	"fmt"

	"github.com/0xVanfer/chainId"
)

func ExampleGetChainlinkOracle() {
	fmt.Println(GetChainlinkOracle(chainId.BinanceSmartChainName, "CRV", "usd"))
	fmt.Println(GetChainlinkOracle(chainId.AvalancheChainName, "link", "avax"))
	fmt.Println(GetChainlinkOracle(chainId.EthereumChainName, "1inch", "eth"))
	// Output:
	// 0x2e1C3b6Fcae47b20Dd343D9354F7B1140a1E6B27
	// 0x1b8a25F73c9420dD507406C3A3816A276b62f56a
	// 0x72AFAECF99C9d9C8215fF44C77B94B99C28741e8
}
