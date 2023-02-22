package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"time"
)

type GetBalanceSheetPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type GetBalanceSheetQueryParams struct {
	PeriodLength     int        `queryParam:"style=form,explode=true,name=periodLength"`
	PeriodsToCompare int        `queryParam:"style=form,explode=true,name=periodsToCompare"`
	StartMonth       *time.Time `queryParam:"style=form,explode=true,name=startMonth"`
}

type GetBalanceSheetSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type GetBalanceSheetRequest struct {
	PathParams  GetBalanceSheetPathParams
	QueryParams GetBalanceSheetQueryParams
	Security    GetBalanceSheetSecurity
}

type GetBalanceSheetResponse struct {
	BalanceSheetResponse *shared.BalanceSheetResponse
	ContentType          string
	StatusCode           int
}
