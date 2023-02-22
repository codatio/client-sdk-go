package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
)

type GetPaymentMethodPathParams struct {
	CompanyID       string `pathParam:"style=simple,explode=false,name=companyId"`
	PaymentMethodID string `pathParam:"style=simple,explode=false,name=paymentMethodId"`
}

type GetPaymentMethodSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type GetPaymentMethodRequest struct {
	PathParams GetPaymentMethodPathParams
	Security   GetPaymentMethodSecurity
}

type GetPaymentMethodResponse struct {
	ContentType   string
	PaymentMethod *shared.PaymentMethod
	StatusCode    int
}
