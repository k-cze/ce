package response

import (
	"ce/internal/services/exchange"
	"encoding/json"
	"fmt"
	"strconv"
)

type ExchangeResponse struct {
	From   string          `json:"from"`
	To     string          `json:"to"`
	Amount formattedAmount `json:"amount"`
}

type ExchangeResponseOption func(r *ExchangeResponse)

func ExchangeResponseWithAmount(a exchange.Amount) ExchangeResponseOption {
	return func(r *ExchangeResponse) {
		r.Amount = formattedAmount{amount: a}
	}
}

func ExchangeResponseWithFrom(from string) ExchangeResponseOption {
	return func(r *ExchangeResponse) {
		r.From = from
	}
}

func ExchangeResponseWithTo(to string) ExchangeResponseOption {
	return func(r *ExchangeResponse) {
		r.To = to
	}
}

func New(opts ...ExchangeResponseOption) *ExchangeResponse {
	resp := new(ExchangeResponse)

	for _, opt := range opts {
		if opt == nil {
			continue
		}

		opt(resp)
	}

	return resp
}

type formattedAmount struct {
	amount exchange.Amount
}

func (fa formattedAmount) MarshalJSON() ([]byte, error) {
	str := fmt.Sprintf("%.*f", fa.amount.Precision(), fa.amount.Value())
	val, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return nil, err
	}
	return json.Marshal(val)
}
