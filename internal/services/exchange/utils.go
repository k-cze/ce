package exchange

import "fmt"

func formatFloatWithDecimals(value float64, decimalPlaces int) string {
	format := fmt.Sprintf("%%.%df", decimalPlaces)
	return fmt.Sprintf(format, value)
}
