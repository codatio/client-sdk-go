package operations

import (
	"github.com/codatio/client-sdk-go/synccommerce/pkg/models/shared"
)

type AddDataConnectionPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type AddDataConnectionRequest struct {
	PathParams AddDataConnectionPathParams
	Request    *string `request:"mediaType=application/json"`
}

type AddDataConnectionResponse struct {
	CodatPublicAPIModelsCompanyDataConnection *shared.CodatPublicAPIModelsCompanyDataConnection
	ContentType                               string
	StatusCode                                int
}
