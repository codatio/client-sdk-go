package operations

import (
	"github.com/codatio/client-sdk-go/bankfeeds/pkg/models/shared"
)

type UpdateBankFeedPathParams struct {
	BankAccountID string `pathParam:"style=simple,explode=false,name=bankAccountId"`
	CompanyID     string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID  string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type UpdateBankFeedSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type UpdateBankFeedRequest struct {
	PathParams UpdateBankFeedPathParams
	Request    *shared.BankFeedBankAccount `request:"mediaType=application/json"`
	Security   UpdateBankFeedSecurity
}

type UpdateBankFeedResponse struct {
	BankFeedBankAccount *shared.BankFeedBankAccount
	ContentType         string
	StatusCode          int
}
