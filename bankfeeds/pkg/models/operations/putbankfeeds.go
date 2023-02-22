package operations

import (
	"github.com/codatio/client-sdk-go/bankfeeds/pkg/models/shared"
)

type PutBankFeedsPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type PutBankFeedsSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type PutBankFeedsRequest struct {
	PathParams PutBankFeedsPathParams
	Request    []shared.BankFeedBankAccount `request:"mediaType=application/json"`
	Security   PutBankFeedsSecurity
}

type PutBankFeedsResponse struct {
	BankFeedBankAccounts []shared.BankFeedBankAccount
	ContentType          string
	StatusCode           int
}
