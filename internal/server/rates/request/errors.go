package request

import "errors"

var (
	errMissingCurrencies = errors.New("missing 'currencies' query param")
	errInvalidCount      = errors.New("at least two currencies are required")
)
