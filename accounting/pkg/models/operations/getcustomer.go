package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
)

type GetCustomerPathParams struct {
	CompanyID  string `pathParam:"style=simple,explode=false,name=companyId"`
	CustomerID string `pathParam:"style=simple,explode=false,name=customerId"`
}

type GetCustomerSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type GetCustomerRequest struct {
	PathParams GetCustomerPathParams
	Security   GetCustomerSecurity
}

type GetCustomerResponse struct {
	ContentType string
	Customer    *shared.Customer
	StatusCode  int
}
