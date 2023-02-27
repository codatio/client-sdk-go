package operations

type DeleteCompaniesCompanyIDPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type DeleteCompaniesCompanyIDRequest struct {
	PathParams DeleteCompaniesCompanyIDPathParams
}

type DeleteCompaniesCompanyID401ApplicationJSON struct {
	CanBeRetried      *string `json:"canBeRetried,omitempty"`
	CorrelationID     *string `json:"correlationId,omitempty"`
	DetailedErrorCode *int64  `json:"detailedErrorCode,omitempty"`
	Error             *string `json:"error,omitempty"`
	Service           *string `json:"service,omitempty"`
	StatusCode        *int64  `json:"statusCode,omitempty"`
}

type DeleteCompaniesCompanyIDResponse struct {
	ContentType                                      string
	StatusCode                                       int
	DeleteCompaniesCompanyID401ApplicationJSONObject *DeleteCompaniesCompanyID401ApplicationJSON
}
