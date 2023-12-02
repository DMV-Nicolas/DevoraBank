package util

// Constants for all supported currencies
const (
	USD = "USD"
	EUR = "EUR"
	COP = "COP"
	ARS = "ARS"
	MXN = "MXN"
	UYU = "UYU"
	CLP = "CLP"
	PEN = "PEN"
	BRL = "BRL"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, COP, ARS, MXN, UYU, CLP, PEN, BRL:
		return true
	}
	return false
}
