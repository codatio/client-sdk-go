// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/sync-for-payables/v3/pkg/models/shared"
	"net/http"
)

type GetPullOperationRequest struct {
	// Unique identifier for a company.
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	// Unique identifier for the dataset that completed its sync.
	DatasetID string `pathParam:"style=simple,explode=false,name=datasetId"`
}

func (o *GetPullOperationRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *GetPullOperationRequest) GetDatasetID() string {
	if o == nil {
		return ""
	}
	return o.DatasetID
}

type GetPullOperationResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// OK
	PullOperation *shared.PullOperation
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetPullOperationResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetPullOperationResponse) GetPullOperation() *shared.PullOperation {
	if o == nil {
		return nil
	}
	return o.PullOperation
}

func (o *GetPullOperationResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetPullOperationResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}