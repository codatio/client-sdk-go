package operations

import (
	"github.com/codatio/client-sdk-go/assess/pkg/models/shared"
)

type GetDataCompaniesCompanyIDConnectionsConnectionIDAssessSubscriptionsProcessPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type GetDataCompaniesCompanyIDConnectionsConnectionIDAssessSubscriptionsProcessRequest struct {
	PathParams GetDataCompaniesCompanyIDConnectionsConnectionIDAssessSubscriptionsProcessPathParams
}

type GetDataCompaniesCompanyIDConnectionsConnectionIDAssessSubscriptionsProcessResponse struct {
	ContentType string
	Report      *shared.Report
	StatusCode  int
}
