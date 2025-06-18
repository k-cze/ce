package request

import (
	"strconv"
	"strings"
)

func amountFromString(input string) (amount, error) {
	input = strings.TrimSpace(input)
	if input == "" {
		return 0, errMissingAmount
	}

	val, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, errInvalidAmountMustBeNumber
	}

	a := amount(val)
	if err := a.validate(); err != nil {
		return 0, err
	}

	return a, nil
}
