package kucoin

import (
	"net/http"
)

// A CurrencyModel represents a model of known currency.
type CurrencyModel struct {
	Name              string `json:"name"`
	Currency          string `json:"currency"`
	FullName          string `json:"fullName"`
	Precision         uint8  `json:"precision"`
	WithdrawalMinSize string `json:"withdrawalMinSize"`
	WithdrawalMinFee  string `json:"withdrawalMinFee"`
	IsWithdrawEnabled bool   `json:"isWithdrawEnabled"`
	IsDepositEnabled  bool   `json:"isDepositEnabled"`
}

// A CurrenciesModel is the set of *CurrencyModel.
type CurrenciesModel []*CurrencyModel

// Currencies returns a list of known currencies.
func (as *ApiService) Currencies() (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/currencies", nil)
	return as.Call(req)
}

// Currency returns the details of the currency.
func (as *ApiService) Currency(currency string) (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/currencies/"+currency, nil)
	return as.Call(req)
}

// FiatPrice returns the fiat price for currency.
func (as *ApiService) FiatPrice() (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/prices", nil)
	return as.Call(req)
}
