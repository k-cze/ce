package request

import "errors"

var (
	errMissingAmount                      = errors.New("amount parameter is required")
	errInvalidAmountMustBeNumber          = errors.New("invalid amount parameter: must be a number")
	errInvalidAmountMustBeGreaterThanZero = errors.New("amount must be greater than zero")
)
