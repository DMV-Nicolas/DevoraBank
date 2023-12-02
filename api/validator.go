package api

import (
	"github.com/DMV-Nicolas/sakurabank/util"
	"github.com/go-playground/validator/v10"
)

var validCurrency validator.Func = func(fl validator.FieldLevel) bool {
	if currency, ok := fl.Field().Interface().(string); ok {
		// chech if currency is supported or not
		return util.IsSupportedCurrency(currency)
	}
	return false
}
