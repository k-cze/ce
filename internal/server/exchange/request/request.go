package request

type exchangeRequest struct {
	From   string
	To     string
	Amount amount
}

type exchangeRequestOption func(*exchangeRequest) (err error)

func New(opts ...exchangeRequestOption) (*exchangeRequest, error) {
	req := new(exchangeRequest)

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

func ExchangeRequestWithFrom(fromParam string) exchangeRequestOption {
	return func(er *exchangeRequest) (err error) {
		er.From = fromParam
		return
	}
}

func ExchangeRequestWithTo(toParam string) exchangeRequestOption {
	return func(er *exchangeRequest) (err error) {
		er.To = toParam
		return
	}
}

func ExchangeRequestWithAmount(amountParam string) exchangeRequestOption {
	return func(er *exchangeRequest) error {
		a, err := amountFromString(amountParam)
		if err != nil {
			return err
		}

		er.Amount = a
		return nil
	}
}
