package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
)

type GetTransferPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
	TransferID   string `pathParam:"style=simple,explode=false,name=transferId"`
}

type GetTransferSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type GetTransferRequest struct {
	PathParams GetTransferPathParams
	Security   GetTransferSecurity
}

type GetTransferResponse struct {
	ContentType string
	StatusCode  int
	Transfer    *shared.Transfer
}
