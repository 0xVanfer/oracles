# Oracles

## Chainlink Oracles

```go
m := NewChainlinkMap()
address := m.ChooseOracle(Selectors{Network: chainId.AvalancheChainName, Symbol: "avax"})
fmt.Println(address)
// Output:
// 0x0A77230d17318075983913bC2145DB16C7366156
```
