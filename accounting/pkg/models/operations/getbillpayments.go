package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
)

type GetBillPaymentsPathParams struct {
	BillPaymentID string `pathParam:"style=simple,explode=false,name=billPaymentId"`
	CompanyID     string `pathParam:"style=simple,explode=false,name=companyId"`
}

type GetBillPaymentsSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type GetBillPaymentsRequest struct {
	PathParams GetBillPaymentsPathParams
	Security   GetBillPaymentsSecurity
}

type GetBillPaymentsResponse struct {
	BillPayment *shared.BillPayment
	ContentType string
	StatusCode  int
}
