package operations

import (
	"github.com/codatio/client-sdk-go/banking/pkg/models/shared"
)

type ListBankingTransactionsPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type ListBankingTransactionsQueryParams struct {
	OrderBy  *string  `queryParam:"style=form,explode=true,name=orderBy"`
	Page     float64  `queryParam:"style=form,explode=true,name=page"`
	PageSize *float64 `queryParam:"style=form,explode=true,name=pageSize"`
	Query    *string  `queryParam:"style=form,explode=true,name=query"`
}

type ListBankingTransactionsSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type ListBankingTransactionsRequest struct {
	PathParams  ListBankingTransactionsPathParams
	QueryParams ListBankingTransactionsQueryParams
	Security    ListBankingTransactionsSecurity
}

type ListBankingTransactions200ApplicationJSON struct {
	Results *shared.Transaction `json:"results,omitempty"`
}

type ListBankingTransactionsResponse struct {
	ContentType                                     string
	StatusCode                                      int
	ListBankingTransactions200ApplicationJSONObject *ListBankingTransactions200ApplicationJSON
}
