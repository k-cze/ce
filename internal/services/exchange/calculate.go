package exchange

func calculateAmount(rateFrom float64, rateTo float64, dp int) string {
	return formatFloatWithDecimals(rateFrom/rateTo, dp)
}
