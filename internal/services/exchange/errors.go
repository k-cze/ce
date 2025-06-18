package exchange

import "errors"

var (
	errCryptoCurrencyNotSupported = errors.New("requested cryptocurrency is not supported")
)
