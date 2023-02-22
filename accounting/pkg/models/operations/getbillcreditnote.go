package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
)

type GetBillCreditNotePathParams struct {
	BillCreditNoteID string `pathParam:"style=simple,explode=false,name=billCreditNoteId"`
	CompanyID        string `pathParam:"style=simple,explode=false,name=companyId"`
}

type GetBillCreditNoteSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type GetBillCreditNoteRequest struct {
	PathParams GetBillCreditNotePathParams
	Security   GetBillCreditNoteSecurity
}

type GetBillCreditNoteResponse struct {
	BillCreditNote *shared.BillCreditNote
	ContentType    string
	StatusCode     int
}
