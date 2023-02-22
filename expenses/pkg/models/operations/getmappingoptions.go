package operations

import (
	"github.com/codatio/client-sdk-go/expenses/pkg/models/shared"
)

type GetMappingOptionsPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type GetMappingOptionsRequest struct {
	PathParams GetMappingOptionsPathParams
}

type GetMappingOptionsResponse struct {
	ContentType    string
	MappingOptions *shared.MappingOptions
	StatusCode     int
}
