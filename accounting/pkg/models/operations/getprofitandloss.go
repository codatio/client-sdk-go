package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"time"
)

type GetProfitAndLossPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type GetProfitAndLossQueryParams struct {
	PeriodLength     int        `queryParam:"style=form,explode=true,name=periodLength"`
	PeriodsToCompare int        `queryParam:"style=form,explode=true,name=periodsToCompare"`
	StartMonth       *time.Time `queryParam:"style=form,explode=true,name=startMonth"`
}

type GetProfitAndLossSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type GetProfitAndLossRequest struct {
	PathParams  GetProfitAndLossPathParams
	QueryParams GetProfitAndLossQueryParams
	Security    GetProfitAndLossSecurity
}

type GetProfitAndLossResponse struct {
	ContentType           string
	ProfitAndLossResponse *shared.ProfitAndLossResponse
	StatusCode            int
}
