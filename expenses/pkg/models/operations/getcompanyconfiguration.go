package operations

import (
	"github.com/codatio/client-sdk-go/expenses/pkg/models/shared"
)

type GetCompanyConfigurationPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type GetCompanyConfigurationRequest struct {
	PathParams GetCompanyConfigurationPathParams
}

type GetCompanyConfigurationResponse struct {
	CompanyConfigResourceRepresentation *shared.CompanyConfigResourceRepresentation
	ContentType                         string
	StatusCode                          int
}
