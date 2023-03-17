package operations

import (
	"net/http"
)

type CheckDataStatusRequest struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	DatasetID string `pathParam:"style=simple,explode=false,name=datasetId"`
}

type CheckDataStatusResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
