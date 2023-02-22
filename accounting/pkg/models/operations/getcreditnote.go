package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
)

type GetCreditNotePathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	CreditNoteID string `pathParam:"style=simple,explode=false,name=creditNoteId"`
}

type GetCreditNoteSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type GetCreditNoteRequest struct {
	PathParams GetCreditNotePathParams
	Security   GetCreditNoteSecurity
}

type GetCreditNoteResponse struct {
	ContentType string
	CreditNote  *shared.CreditNote
	StatusCode  int
}
