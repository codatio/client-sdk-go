package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
)

type GetBankAccountPathParams struct {
	AccountID    string `pathParam:"style=simple,explode=false,name=accountId"`
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type GetBankAccountSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type GetBankAccountRequest struct {
	PathParams GetBankAccountPathParams
	Security   GetBankAccountSecurity
}

type GetBankAccountResponse struct {
	BankAccount *shared.BankAccount
	ContentType string
	StatusCode  int
}
