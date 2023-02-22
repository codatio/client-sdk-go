package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
)

type GetAccountTransactionPathParams struct {
	AccountTransactionID string `pathParam:"style=simple,explode=false,name=accountTransactionId"`
	CompanyID            string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID         string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type GetAccountTransactionSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type GetAccountTransactionRequest struct {
	PathParams GetAccountTransactionPathParams
	Security   GetAccountTransactionSecurity
}

type GetAccountTransactionResponse struct {
	AccountTransaction *shared.AccountTransaction
	ContentType        string
	StatusCode         int
}
