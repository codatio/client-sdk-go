package operations

import (
	"github.com/codatio/client-sdk-go/bankfeed/pkg/models/shared"
)

type GetBankFeedPathParams struct {
	BankAccountID string `pathParam:"style=simple,explode=false,name=bankAccountId"`
	CompanyID     string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID  string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type GetBankFeedSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type GetBankFeedRequest struct {
	PathParams GetBankFeedPathParams
	Request    *shared.BankFeedBankAccount `request:"mediaType=application/json"`
	Security   GetBankFeedSecurity
}

type GetBankFeedResponse struct {
	BankFeedBankAccount *shared.BankFeedBankAccount
	ContentType         string
	StatusCode          int
}
