package constants

type Network string

func (n Network) String() string {
	return string(n)
}

const (
	NetworkBTC  Network = "Bitcoin"
	NetworkETH  Network = "Ethereum"
	NetworkTRON Network = "TRON"
	NetworkLTC  Network = "LITECOIN"
	NetworkIMX  Network = "Immutable zkEVM"
	NetworkARB  Network = "Arbitrum"
)
