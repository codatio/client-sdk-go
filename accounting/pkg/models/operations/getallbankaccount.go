package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
)

type GetAllBankAccountPathParams struct {
	AccountID string `pathParam:"style=simple,explode=false,name=accountId"`
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type GetAllBankAccountQueryParams struct {
	Query *string `queryParam:"style=form,explode=true,name=query"`
}

type GetAllBankAccountSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type GetAllBankAccountRequest struct {
	PathParams  GetAllBankAccountPathParams
	QueryParams GetAllBankAccountQueryParams
	Security    GetAllBankAccountSecurity
}

type GetAllBankAccountResponse struct {
	BankStatementAccount *shared.BankStatementAccount
	ContentType          string
	StatusCode           int
}
