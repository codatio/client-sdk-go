package operations

import (
	"github.com/codatio/client-sdk-go/synccommerce/pkg/models/shared"
)

type PostCompaniesRequest struct {
	Request *shared.CodatSyncDirectAPIModelsAddCompanyModel `request:"mediaType=application/json"`
}

type PostCompaniesResponse struct {
	CodatSyncDirectAPIModelsCompany *shared.CodatSyncDirectAPIModelsCompany
	ContentType                     string
	StatusCode                      int
}
