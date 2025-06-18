package response

import (
	ratesvc "ce/internal/services/rates"
	"encoding/json"
)

type RateEntry struct {
	From string  `json:"from"`
	To   string  `json:"to"`
	Rate float64 `json:"rate"`
}

type RatesResponse struct {
	Entries []*RateEntry
}

func (r *RatesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.Entries)
}

type RatesResponseOption func(r *RatesResponse)

func RatesResponseWithMatrix(m ratesvc.ExchangeMatrix) RatesResponseOption {
	return func(r *RatesResponse) {
		r.Entries = flattenMatrix(m)
	}
}

func New(opts ...RatesResponseOption) *RatesResponse {
	resp := new(RatesResponse)

	for _, opt := range opts {
		if opt == nil {
			continue
		}

		opt(resp)
	}

	return resp
}
