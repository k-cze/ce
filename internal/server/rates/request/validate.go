package request

type currencies []string

func (c currencies) validate() error {
	if len(c) == 0 {
		return errMissingCurrencies
	}
	if len(c) < 2 {
		return errInvalidCount
	}
	return nil
}
