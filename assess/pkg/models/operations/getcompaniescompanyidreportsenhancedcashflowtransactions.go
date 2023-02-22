package operations

import (
	"github.com/codatio/client-sdk-go/assess/pkg/models/shared"
)

type GetCompaniesCompanyIDReportsEnhancedCashFlowTransactionsPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type GetCompaniesCompanyIDReportsEnhancedCashFlowTransactionsQueryParams struct {
	Page     float64  `queryParam:"style=form,explode=true,name=page"`
	PageSize *float64 `queryParam:"style=form,explode=true,name=pageSize"`
	Query    *string  `queryParam:"style=form,explode=true,name=query"`
}

type GetCompaniesCompanyIDReportsEnhancedCashFlowTransactionsRequest struct {
	PathParams  GetCompaniesCompanyIDReportsEnhancedCashFlowTransactionsPathParams
	QueryParams GetCompaniesCompanyIDReportsEnhancedCashFlowTransactionsQueryParams
}

type GetCompaniesCompanyIDReportsEnhancedCashFlowTransactionsResponse struct {
	ContentType                  string
	EnhancedCashFlowTransactions *shared.EnhancedCashFlowTransactions
	StatusCode                   int
}
