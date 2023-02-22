package operations

import (
	"github.com/codatio/client-sdk-go/assess/pkg/models/shared"
)

type GetDataCompaniesCompanyIDConnectionsConnectionIDAssessSubscriptionsMrrPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type GetDataCompaniesCompanyIDConnectionsConnectionIDAssessSubscriptionsMrrRequest struct {
	PathParams GetDataCompaniesCompanyIDConnectionsConnectionIDAssessSubscriptionsMrrPathParams
}

type GetDataCompaniesCompanyIDConnectionsConnectionIDAssessSubscriptionsMrrResponse struct {
	ContentType string
	Report      *shared.Report
	StatusCode  int
}
