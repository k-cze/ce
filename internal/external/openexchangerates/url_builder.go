package openexchangerates

import (
	"net/url"
	"strings"
)

func buildRatesURL(apiKey string, currencies []string) (string, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return "", err
	}

	q := u.Query()
	q.Set(paramAppID, apiKey)
	q.Set(paramSymbols, strings.Join(currencies, ","))
	u.RawQuery = q.Encode()

	return u.String(), nil
}
