package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
)

type GetDirectCostPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
	DirectCostID string `pathParam:"style=simple,explode=false,name=directCostId"`
}

type GetDirectCostSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type GetDirectCostRequest struct {
	PathParams GetDirectCostPathParams
	Security   GetDirectCostSecurity
}

type GetDirectCostResponse struct {
	ContentType string
	DirectCost  *shared.DirectCost
	StatusCode  int
}
