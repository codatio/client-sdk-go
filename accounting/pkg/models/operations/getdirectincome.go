package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
)

type GetDirectIncomePathParams struct {
	CompanyID      string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID   string `pathParam:"style=simple,explode=false,name=connectionId"`
	DirectIncomeID string `pathParam:"style=simple,explode=false,name=directIncomeId"`
}

type GetDirectIncomeSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type GetDirectIncomeRequest struct {
	PathParams GetDirectIncomePathParams
	Security   GetDirectIncomeSecurity
}

type GetDirectIncomeResponse struct {
	ContentType  string
	DirectIncome *shared.DirectIncome
	StatusCode   int
}
