package api

import (
	"github.com/go-playground/validator/v10"
	"github/adedaryorh/pooler_Remmitance_Application/utils"
)

var currencyValidator validator.Func = func(fl validator.FieldLevel) bool {
	if currency, ok := fl.Field().Interface().(string); ok {
		return utils.IsValidCurrency(currency)
	}
	return false

}
