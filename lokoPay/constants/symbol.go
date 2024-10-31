package constants

type Symbol string

func (s Symbol) String() string {
	return string(s)
}

const (
	SymbolBTC     Symbol = "BTC"
	SymbolETH     Symbol = "ETH"
	SymbolTRX     Symbol = "TRX"
	SymbolLTC     Symbol = "LTC"
	SymbolIMX     Symbol = "IMX"
	SymbolUSDC    Symbol = "USDC"
	SymbolARB     Symbol = "ARB"
	SymbolERCUSDC Symbol = "ERCUSDC"
	SymbolERCIMX  Symbol = "ERCIMX"
	SymbolERCUSDT Symbol = "ERCUSDT"
	SymbolTRCUSDT Symbol = "TRCUSDT"
	SymbolIMXUSDC Symbol = "IMXUSDC"
	SymbolIMXUSDT Symbol = "IMXUSDT"
	SymbolIMXETH  Symbol = "IMXETH"
	SymbolARBUSDC Symbol = "ARBUSDC"
	SymbolARBUSDT Symbol = "ARBUSDT"
)
