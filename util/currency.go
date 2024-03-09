package util

const (
	USD = "USD"
	EUR = "EUR"
	CAD = "CAD"
	AKZ = "AKZ"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, CAD, AKZ:
		return true
	}
	return false
}
