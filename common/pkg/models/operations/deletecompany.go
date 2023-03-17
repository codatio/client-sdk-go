package operations

import (
	"net/http"
)

type DeleteCompanyRequest struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type DeleteCompany401ApplicationJSON struct {
	CanBeRetried      *string `json:"canBeRetried,omitempty"`
	CorrelationID     *string `json:"correlationId,omitempty"`
	DetailedErrorCode *int64  `json:"detailedErrorCode,omitempty"`
	Error             *string `json:"error,omitempty"`
	Service           *string `json:"service,omitempty"`
	StatusCode        *int64  `json:"statusCode,omitempty"`
}

type DeleteCompanyResponse struct {
	ContentType                           string
	StatusCode                            int
	RawResponse                           *http.Response
	DeleteCompany401ApplicationJSONObject *DeleteCompany401ApplicationJSON
}
