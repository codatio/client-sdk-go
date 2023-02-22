package operations

import (
	"github.com/codatio/client-sdk-go/commerce/pkg/models/shared"
)

type GetCommerceInfoPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type GetCommerceInfoRequest struct {
	PathParams GetCommerceInfoPathParams
}

type GetCommerceInfoResponse struct {
	CompanyInfo *shared.CompanyInfo
	ContentType string
	StatusCode  int
}
