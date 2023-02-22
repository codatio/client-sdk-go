package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
)

type GetBillPathParams struct {
	BillID    string `pathParam:"style=simple,explode=false,name=billId"`
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type GetBillSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type GetBillRequest struct {
	PathParams GetBillPathParams
	Security   GetBillSecurity
}

type GetBillResponse struct {
	Bill        *shared.Bill
	ContentType string
	StatusCode  int
}
