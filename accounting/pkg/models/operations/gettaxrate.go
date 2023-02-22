package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
)

type GetTaxRatePathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	TaxRateID string `pathParam:"style=simple,explode=false,name=taxRateId"`
}

type GetTaxRateSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type GetTaxRateRequest struct {
	PathParams GetTaxRatePathParams
	Security   GetTaxRateSecurity
}

type GetTaxRateResponse struct {
	ContentType string
	StatusCode  int
	TaxRate     *shared.TaxRate
}
