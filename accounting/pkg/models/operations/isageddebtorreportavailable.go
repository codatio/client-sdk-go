package operations

type IsAgedDebtorReportAvailablePathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type IsAgedDebtorReportAvailableRequest struct {
	PathParams IsAgedDebtorReportAvailablePathParams
}

type IsAgedDebtorReportAvailableResponse struct {
	ContentType                                          string
	StatusCode                                           int
	IsAgedDebtorReportAvailable200ApplicationJSONBoolean *bool
}
