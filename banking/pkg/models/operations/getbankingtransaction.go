package operations

import (
	"github.com/codatio/client-sdk-go/banking/pkg/models/shared"
)

type GetBankingTransactionPathParams struct {
	CompanyID     string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID  string `pathParam:"style=simple,explode=false,name=connectionId"`
	TransactionID string `pathParam:"style=simple,explode=false,name=transactionId"`
}

type GetBankingTransactionSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type GetBankingTransactionRequest struct {
	PathParams GetBankingTransactionPathParams
	Security   GetBankingTransactionSecurity
}

type GetBankingTransactionResponse struct {
	ContentType string
	StatusCode  int
	Transaction *shared.Transaction
}
