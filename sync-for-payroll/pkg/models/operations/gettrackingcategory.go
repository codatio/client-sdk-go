// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/sync-for-payroll/v2/pkg/models/shared"
	"net/http"
)

type GetTrackingCategoryRequest struct {
	// Unique identifier for a company.
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	// Unique identifier for a tracking category.
	TrackingCategoryID string `pathParam:"style=simple,explode=false,name=trackingCategoryId"`
}

func (o *GetTrackingCategoryRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *GetTrackingCategoryRequest) GetTrackingCategoryID() string {
	if o == nil {
		return ""
	}
	return o.TrackingCategoryID
}

type GetTrackingCategoryResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Success
	TrackingCategoryTree *shared.TrackingCategoryTree
}

func (o *GetTrackingCategoryResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetTrackingCategoryResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetTrackingCategoryResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetTrackingCategoryResponse) GetTrackingCategoryTree() *shared.TrackingCategoryTree {
	if o == nil {
		return nil
	}
	return o.TrackingCategoryTree
}
