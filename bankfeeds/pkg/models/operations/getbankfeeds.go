package operations

import (
	"github.com/codatio/client-sdk-go/bankfeeds/pkg/models/shared"
)

type GetBankFeedsPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type GetBankFeedsSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type GetBankFeedsRequest struct {
	PathParams GetBankFeedsPathParams
	Security   GetBankFeedsSecurity
}

type GetBankFeedsResponse struct {
	BankFeedBankAccounts []shared.BankFeedBankAccount
	ContentType          string
	StatusCode           int
}
