package operations

import (
	"github.com/codatio/client-sdk-go/synccommerce/pkg/models/shared"
)

type UpdateDataConnectionPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type UpdateDataConnectionRequest struct {
	PathParams UpdateDataConnectionPathParams
	Request    *shared.CodatPublicAPIModelsCompanyPatchDataConnectionModel `request:"mediaType=application/json"`
}

type UpdateDataConnectionResponse struct {
	CodatPublicAPIModelsCompanyDataConnection *shared.CodatPublicAPIModelsCompanyDataConnection
	ContentType                               string
	StatusCode                                int
}
