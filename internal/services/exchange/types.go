package exchange

type Amount struct {
	value            float64
	decimalPrecision int
}

func NewAmount(value float64, precision int) Amount {
	return Amount{
		value:            value,
		decimalPrecision: precision,
	}
}

func (a Amount) Value() float64 {
	return a.value
}

func (a Amount) Precision() int {
	return a.decimalPrecision
}
