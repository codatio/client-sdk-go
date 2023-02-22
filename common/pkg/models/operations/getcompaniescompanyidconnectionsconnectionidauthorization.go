package operations

import (
	"github.com/codatio/client-sdk-go/common/pkg/models/shared"
)

type GetCompaniesCompanyIDConnectionsConnectionIDAuthorizationPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type GetCompaniesCompanyIDConnectionsConnectionIDAuthorizationRequestBody struct {
	Property1 *string `json:"property1,omitempty"`
	Property2 *string `json:"property2,omitempty"`
	Property3 *string `json:"property3,omitempty"`
}

type GetCompaniesCompanyIDConnectionsConnectionIDAuthorizationRequest struct {
	PathParams GetCompaniesCompanyIDConnectionsConnectionIDAuthorizationPathParams
	Request    *GetCompaniesCompanyIDConnectionsConnectionIDAuthorizationRequestBody `request:"mediaType=application/json"`
}

type GetCompaniesCompanyIDConnectionsConnectionIDAuthorizationResponse struct {
	Connection  *shared.Connection
	ContentType string
	StatusCode  int
}
