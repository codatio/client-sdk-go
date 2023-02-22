package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
)

type GetPaymentPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	PaymentID string `pathParam:"style=simple,explode=false,name=paymentId"`
}

type GetPaymentSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type GetPaymentRequest struct {
	PathParams GetPaymentPathParams
	Security   GetPaymentSecurity
}

type GetPaymentResponse struct {
	ContentType string
	Payment     *shared.Payment
	StatusCode  int
}
