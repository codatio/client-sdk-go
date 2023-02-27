package operations

type GetCompanyConfigurationPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type GetCompanyConfigurationRequest struct {
	PathParams GetCompanyConfigurationPathParams
}

type GetCompanyConfiguration200TextJSONBankAccount struct {
	ID *string `json:"id,omitempty"`
}

type GetCompanyConfiguration200TextJSONCustomer struct {
	ID *string `json:"id,omitempty"`
}

type GetCompanyConfiguration200TextJSONSupplier struct {
	ID *string `json:"id,omitempty"`
}

type GetCompanyConfiguration200TextJSON struct {
	BankAccount *GetCompanyConfiguration200TextJSONBankAccount `json:"bankAccount,omitempty"`
	Customer    *GetCompanyConfiguration200TextJSONCustomer    `json:"customer,omitempty"`
	Supplier    *GetCompanyConfiguration200TextJSONSupplier    `json:"supplier,omitempty"`
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
	GetCompanyConfiguration200ApplicationJSONObject *GetCompanyConfiguration200ApplicationJSON
	GetCompanyConfiguration200TextJSONObject        *GetCompanyConfiguration200TextJSON
	GetCompanyConfiguration200TextPlainObject       *string
}
