package crypto

import (
	"errors"
	"strconv"
	"strings"
)

func ParseCryptoRate(row []string) (*CryptoRate, error) {
	const (
		cryptoCurrencyIdx = 0
		decimalPlaces     = 1
		rateUSD           = 2
	)

	if len(row) != 3 {
		return nil, errors.New("invalid number of fields")
	}

	d, err := strconv.Atoi(strings.TrimSpace(row[decimalPlaces]))
	if err != nil {
		return nil, err
	}

	rate, err := strconv.ParseFloat(strings.TrimSpace(row[rateUSD]), 64)
	if err != nil {
		return nil, err
	}

	return &CryptoRate{
		CryptoCurrency: strings.TrimSpace(row[cryptoCurrencyIdx]),
		DecimalPlaces:  d,
		RateUSD:        rate,
	}, nil
}
