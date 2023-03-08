package operations

import (
	"net/http"
)

type GetCompanyConfigurationPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type GetCompanyConfigurationRequest struct {
	PathParams GetCompanyConfigurationPathParams
}

type GetCompanyConfiguration200ApplicationJSONBankAccount struct {
	ID *string `json:"id,omitempty"`
}

type GetCompanyConfiguration200ApplicationJSONCustomer struct {
	ID *string `json:"id,omitempty"`
}

type GetCompanyConfiguration200ApplicationJSONSupplier struct {
	ID *string `json:"id,omitempty"`
}

type GetCompanyConfiguration200ApplicationJSON struct {
	BankAccount *GetCompanyConfiguration200ApplicationJSONBankAccount `json:"bankAccount,omitempty"`
	Customer    *GetCompanyConfiguration200ApplicationJSONCustomer    `json:"customer,omitempty"`
	Supplier    *GetCompanyConfiguration200ApplicationJSONSupplier    `json:"supplier,omitempty"`
}

type GetCompanyConfigurationResponse struct {
	ContentType                                     string
	StatusCode                                      int
	RawResponse                                     *http.Response
	GetCompanyConfiguration200ApplicationJSONObject *GetCompanyConfiguration200ApplicationJSON
}
