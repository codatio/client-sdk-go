package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
)

type GetCompanyInfoPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type GetCompanyInfoSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type GetCompanyInfoRequest struct {
	PathParams GetCompanyInfoPathParams
	Security   GetCompanyInfoSecurity
}

type GetCompanyInfoResponse struct {
	CompanyDataset *shared.CompanyDataset
	ContentType    string
	StatusCode     int
}
