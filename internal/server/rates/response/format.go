package response

import "ce/internal/services/rates"

func flattenMatrix(matrix rates.ExchangeMatrix) []*RateEntry {
	var result []*RateEntry

	for from, targets := range matrix {
		for to, rate := range targets {
			result = append(result, &RateEntry{
				From: from,
				To:   to,
				Rate: rate,
			})
		}
	}

	return result
}
