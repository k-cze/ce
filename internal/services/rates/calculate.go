package rates

type ExchangeMatrix map[string]map[string]float64

func calculateExchangeMatrix(rates map[string]float64, currencies []string) ExchangeMatrix {
	matrix := make(ExchangeMatrix)

	for _, from := range currencies {
		fromRate, ok := getRate(rates, from)
		if !ok {
			continue
		}

		matrix[from] = make(map[string]float64)

		for _, to := range currencies {
			if from == to {
				continue
			}

			toRate, ok := getRate(rates, to)
			if !ok {
				continue
			}

			matrix[from][to] = toRate / fromRate
		}
	}

	return matrix
}

func getRate(rates map[string]float64, currency string) (float64, bool) {
	rate, ok := rates[currency]
	return rate, ok
}
