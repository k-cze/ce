package crypto

type CryptoRate struct {
	CryptoCurrency string  `csv:"CryptoCurrency"`
	DecimalPlaces  int     `csv:"Decimal places"`
	RateUSD        float64 `csv:"Rate (to USD)"`
}
