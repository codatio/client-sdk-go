package operations

import (
	"net/http"
	"time"
)

type GetAllBankAccountPathParams struct {
	AccountID string `pathParam:"style=simple,explode=false,name=accountId"`
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type GetAllBankAccountQueryParams struct {
	Query *string `queryParam:"style=form,explode=true,name=query"`
}

type GetAllBankAccountRequest struct {
	PathParams  GetAllBankAccountPathParams
	QueryParams GetAllBankAccountQueryParams
}

type GetAllBankAccount200ApplicationJSON struct {
	AccountName        *string    `json:"accountName,omitempty"`
	AccountNumber      *string    `json:"accountNumber,omitempty"`
	AvailableBalance   *float64   `json:"availableBalance,omitempty"`
	Balance            *float64   `json:"balance,omitempty"`
	Currency           *string    `json:"currency,omitempty"`
	FromDate           *time.Time `json:"fromDate,omitempty"`
	Iban               *string    `json:"iban,omitempty"`
	ID                 *string    `json:"id,omitempty"`
	Institution        *string    `json:"institution,omitempty"`
	ModifiedDate       *time.Time `json:"modifiedDate,omitempty"`
	NominalCode        *string    `json:"nominalCode,omitempty"`
	OverdraftLimit     *float64   `json:"overdraftLimit,omitempty"`
	SortCode           *string    `json:"sortCode,omitempty"`
	SourceModifiedDate *time.Time `json:"sourceModifiedDate,omitempty"`
	ToDate             *time.Time `json:"toDate,omitempty"`
}

type GetAllBankAccountResponse struct {
	ContentType                               string
	StatusCode                                int
	RawResponse                               *http.Response
	GetAllBankAccount200ApplicationJSONObject *GetAllBankAccount200ApplicationJSON
}
