package crypto

type Interface interface {
	GetDecimalPlaces(cryptoName string) int
	GetRateUSD(cryptoName string) float64
	Exists(cryptoName string) bool
}

type Driver struct {
	data                   []*CryptoRate
	lookupByCryptoCurrency map[string]*CryptoRate
}

func (d *Driver) Append(r *CryptoRate) {
	d.data = append(d.data, r)
	if d.lookupByCryptoCurrency == nil {
		d.lookupByCryptoCurrency = make(map[string]*CryptoRate)
	}
	d.lookupByCryptoCurrency[r.CryptoCurrency] = r
}

func (d *Driver) All() []*CryptoRate {
	return d.data
}

func (d *Driver) GetDecimalPlaces(cryptoName string) int {
	if r, ok := d.lookupByCryptoCurrency[cryptoName]; ok {
		return r.DecimalPlaces
	}
	return 0
}

func (d *Driver) GetRateUSD(cryptoName string) float64 {
	if r, ok := d.lookupByCryptoCurrency[cryptoName]; ok {
		return r.RateUSD
	}
	return 0.0
}

func (d *Driver) Exists(cryptoName string) bool {
	if _, ok := d.lookupByCryptoCurrency[cryptoName]; ok {
		return true
	}
	return false
}
