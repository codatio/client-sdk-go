package operations

import (
	"github.com/codatio/client-sdk-go/expenses/pkg/models/shared"
)

type SaveCompanyConfigurationPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type SaveCompanyConfigurationRequest struct {
	PathParams SaveCompanyConfigurationPathParams
	Request    *shared.CompanyConfigResourceRepresentation `request:"mediaType=application/json"`
}

type SaveCompanyConfigurationResponse struct {
	CodatErrorMessage                   *shared.CodatErrorMessage
	CompanyConfigResourceRepresentation *shared.CompanyConfigResourceRepresentation
	ContentType                         string
	StatusCode                          int
}
