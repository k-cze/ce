package rates

import (
	"ce/internal/external/openexchangerates"
)

type Service interface {
	GetRates(currencies []string) (ExchangeMatrix, error)
}

type RatesService struct {
	client openexchangerates.ClientInterface
}

func NewService(client openexchangerates.ClientInterface) *RatesService {
	return &RatesService{client: client}
}

func (s *RatesService) GetRates(currencies []string) (ExchangeMatrix, error) {
	rates, err := s.client.FetchRates(currencies)
	if err != nil {
		return nil, err
	}

	return calculateExchangeMatrix(rates, currencies), nil
}
