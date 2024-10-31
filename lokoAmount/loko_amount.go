package lokoAmount

import (
	"errors"
	"math"
	"math/big"
)

type LokoCurrency struct {
	Symbol      string
	Unit        string
	MinDecimals int
}

var currencyMap = map[string]LokoCurrency{
	"USDC":     {"USDC", "USDC", 2},
	"BTC":      {"BTC", "BTC", 8},
	"ETH":      {"ETH", "ETH", 18},
	"IMX":      {"IMX", "IMX", 18},
	"TRX":      {"TRX", "TRX", 6},
	"LTC":      {"LTC", "LTC", 8},
	"ARB":      {"ARB", "ETH", 18},
	"ERCUSDC":  {"ERCUSDC", "USDC", 6},
	"ERCIMX":   {"ERCIMX", "IMX", 18},
	"ERCUSDT":  {"ERCUSDT", "USDT", 6},
	"TRCUSDT":  {"TRCUSDT", "USDT", 6},
	"IMXUSDC":  {"IMXUSDC", "USDC", 6},
	"IMXUSDT":  {"IMXUSDT", "USDT", 6},
	"IMXETH":   {"IMXETH", "ETH", 18},
	"ARBUSDC":  {"ARBUSDC", "USDC", 6},
	"ARBUSDT:": {"ARBUSDT", "USDT", 6},
}

func NewLokoCurrency(symbol string) (LokoCurrency, error) {
	if currency, ok := currencyMap[symbol]; ok {
		return currency, nil
	}
	return LokoCurrency{}, errors.New("currency not found")
}

func (c *LokoCurrency) GetSymbol() string {
	if c == nil {
		return ""
	}
	return c.Symbol
}

func (c *LokoCurrency) GetUnit() string {
	if c == nil {
		return ""
	}
	return c.Unit
}

type LokoAmount struct {
	Currency LokoCurrency
	Amount   *big.Float
}

func NewLokoAmountFromMinAmount(amount interface{}, symbol string) *LokoAmount {
	currency, err := NewLokoCurrency(symbol)
	if err != nil {
		return nil
	}
	amountBigFloat := big.NewFloat(0)
	switch amount.(type) {
	case *big.Int:
		amountBigFloat.SetInt(amount.(*big.Int))
	case string:
		amountBigFloat.SetString(amount.(string))
	}
	amountBigFloat = new(big.Float).Quo(amountBigFloat, big.NewFloat(math.Pow10(currency.MinDecimals)))
	return &LokoAmount{
		Currency: currency,
		Amount:   amountBigFloat,
	}
}

func NewLokoAmountFromAmount(amount float64, symbol string) *LokoAmount {
	currency, err := NewLokoCurrency(symbol)
	if err != nil {
		return nil
	}
	amountBigFloat := big.NewFloat(0)
	amountBigFloat.SetFloat64(amount)
	return &LokoAmount{
		Currency: currency,
		Amount:   amountBigFloat,
	}
}

func (l *LokoAmount) ToMinAmount() *big.Int {
	if l == nil {
		return big.NewInt(0)
	}
	amountBigFloat := new(big.Float).Mul(l.Amount, big.NewFloat(math.Pow10(l.Currency.MinDecimals)))
	amountBigInt, _ := new(big.Int).SetString(amountBigFloat.Text('f', 0), 10)
	return amountBigInt
}

func (l *LokoAmount) ToAmount() float64 {
	if l == nil {
		return 0
	}
	amountFloat64, _ := l.Amount.Float64()
	return amountFloat64
}

func (l *LokoAmount) GetSymbol() string {
	if l == nil {
		return ""
	}
	return l.Currency.Symbol
}

func (l *LokoAmount) GetUnit() string {
	if l == nil {
		return ""
	}
	return l.Currency.Unit
}
