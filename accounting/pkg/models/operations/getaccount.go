package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
)

type GetAccountPathParams struct {
	AccountID string `pathParam:"style=simple,explode=false,name=accountId"`
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type GetAccountSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type GetAccountRequest struct {
	PathParams GetAccountPathParams
	Security   GetAccountSecurity
}

type GetAccountResponse struct {
	Account     *shared.Account
	ContentType string
	StatusCode  int
}
