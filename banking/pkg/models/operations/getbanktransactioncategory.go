package operations

import (
	"github.com/codatio/client-sdk-go/banking/pkg/models/shared"
)

type GetBankTransactionCategoryPathParams struct {
	CompanyID             string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID          string `pathParam:"style=simple,explode=false,name=connectionId"`
	TransactionCategoryID string `pathParam:"style=simple,explode=false,name=transactionCategoryId"`
}

type GetBankTransactionCategorySecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type GetBankTransactionCategoryRequest struct {
	PathParams GetBankTransactionCategoryPathParams
	Security   GetBankTransactionCategorySecurity
}

type GetBankTransactionCategoryResponse struct {
	ContentType         string
	StatusCode          int
	TransactionCategory *shared.TransactionCategory
}
