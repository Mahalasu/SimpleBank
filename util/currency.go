package util

const (
	USD = "USD"
	EUR = "EUR"
	CAD = "CAD"
	RMB = "RMB"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, CAD, RMB:
		return true
	}
	return false
}
