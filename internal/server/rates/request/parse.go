package request

import "strings"

func parseCurrencies(input string) []string {
	return filterEmpty(strings.FieldsFunc(input, func(r rune) bool {
		return r == ','
	}))
}

func filterEmpty(slice []string) []string {
	var out []string
	for _, item := range slice {
		item = strings.TrimSpace(item)
		if item != "" {
			out = append(out, item)
		}
	}
	return out
}
