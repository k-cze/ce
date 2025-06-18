package exchange

import (
	"ce/internal/config"
	"ce/internal/csvloader"
)

type Service interface {
	GetExchange(from string, to string, amount float64) (*Amount, error)
}

type ExchangeService struct {
	drivers csvloader.Drivers
}

func NewService(cfg config.Config) *ExchangeService {
	return &ExchangeService{
		drivers: cfg.GetCSVDrivers(),
	}
}

func (s *ExchangeService) GetExchange(from string, to string, amount float64) (*Amount, error) {
	for _, cryptoName := range []string{from, to} {
		if !s.drivers.Crypto().Exists(cryptoName) {
			return nil, errCryptoCurrencyNotSupported
		}
	}

	var (
		dp                       = s.drivers.Crypto().GetDecimalPlaces(to)
		rateFrom                 = s.drivers.Crypto().GetRateUSD(from)
		rateTo                   = s.drivers.Crypto().GetRateUSD(to)
		getAfterConversionAmount = func() float64 {
			return (rateFrom * amount) / rateTo
		}
	)

	resultAmount := NewAmount(getAfterConversionAmount(), dp)
	return &resultAmount, nil
}
