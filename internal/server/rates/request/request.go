package request

type ratesRequest struct {
	Currencies currencies
}

type ratesRequestOption func(*ratesRequest) (err error)

func New(opts ...ratesRequestOption) (*ratesRequest, error) {
	req := new(ratesRequest)

	for _, opt := range opts {
		if opt == nil {
			continue
		}

		if err := opt(req); err != nil {
			return nil, err
		}
	}

	return req, nil
}

func RatesRequestWithCurrencies(currenciesParam string) ratesRequestOption {
	return func(rr *ratesRequest) (err error) {
		parsed := parseCurrencies(currenciesParam)
		if err = currencies(parsed).validate(); err != nil {
			return err
		}

		rr.Currencies = parsed
		return
	}
}
