package constants

type Currency string

func (c Currency) String() string {
	return string(c)
}

const (
	CurrencyUSDC  Currency = "USDC"
	CurrencyUSDT  Currency = "USDT"
	CurrencyBTC   Currency = "BTC"
	CurrencyETH   Currency = "ETH"
	CurrencyLTC   Currency = "LTC"
	CurrencyTRX   Currency = "TRX"
	CurrencyIMX   Currency = "IMX"
	CurrencyBNB   Currency = "BNB"
	CurrencyOIK   Currency = "OIK"
	CurrencySUPER Currency = "SUPER"
	CurrencyULTI  Currency = "ULTI"
	CurrencyUSA   Currency = "USA"
	CurrencyCAD   Currency = "CAD"
	CurrencyHKD   Currency = "HKD"
)
