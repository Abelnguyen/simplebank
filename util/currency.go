package util

// constants for all supported currencies
const (
	USD = "USD"
	EUR = "EUR"
	VND = "VND"
)

// ISSupportedCurrency returns true if the currency is supported
func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, VND:
		return true
	}
	return false
}
