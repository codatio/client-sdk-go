package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"time"
)

type GetCashFlowStatementPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type GetCashFlowStatementQueryParams struct {
	PeriodLength     int        `queryParam:"style=form,explode=true,name=periodLength"`
	PeriodsToCompare int        `queryParam:"style=form,explode=true,name=periodsToCompare"`
	StartMonth       *time.Time `queryParam:"style=form,explode=true,name=startMonth"`
}

type GetCashFlowStatementSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type GetCashFlowStatementRequest struct {
	PathParams  GetCashFlowStatementPathParams
	QueryParams GetCashFlowStatementQueryParams
	Security    GetCashFlowStatementSecurity
}

type GetCashFlowStatementResponse struct {
	CashFlowStatementResponse *shared.CashFlowStatementResponse
	ContentType               string
	StatusCode                int
}
