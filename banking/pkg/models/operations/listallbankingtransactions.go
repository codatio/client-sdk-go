package operations

import (
	"github.com/codatio/client-sdk-go/banking/pkg/models/shared"
)

type ListAllBankingTransactionsPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type ListAllBankingTransactionsQueryParams struct {
	OrderBy  *string  `queryParam:"style=form,explode=true,name=orderBy"`
	Page     float64  `queryParam:"style=form,explode=true,name=page"`
	PageSize *float64 `queryParam:"style=form,explode=true,name=pageSize"`
	Query    *string  `queryParam:"style=form,explode=true,name=query"`
}

type ListAllBankingTransactionsSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type ListAllBankingTransactionsRequest struct {
	PathParams  ListAllBankingTransactionsPathParams
	QueryParams ListAllBankingTransactionsQueryParams
	Security    ListAllBankingTransactionsSecurity
}

type ListAllBankingTransactions200ApplicationJSON struct {
	Results *shared.Transaction `json:"results,omitempty"`
}

type ListAllBankingTransactionsResponse struct {
	ContentType                                        string
	StatusCode                                         int
	ListAllBankingTransactions200ApplicationJSONObject *ListAllBankingTransactions200ApplicationJSON
}
