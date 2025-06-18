package request

type amount float64

func (a amount) validate() error {
	if a <= 0 {
		return errInvalidAmountMustBeGreaterThanZero
	}
	return nil
}
