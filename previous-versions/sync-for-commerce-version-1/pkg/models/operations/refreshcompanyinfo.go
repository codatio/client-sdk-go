// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1/pkg/models/shared"
	"net/http"
)

type RefreshCompanyInfoRequest struct {
	// Unique identifier for a company.
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

func (o *RefreshCompanyInfoRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

type RefreshCompanyInfoResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Success
	Dataset *shared.Dataset
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *RefreshCompanyInfoResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *RefreshCompanyInfoResponse) GetDataset() *shared.Dataset {
	if o == nil {
		return nil
	}
	return o.Dataset
}

func (o *RefreshCompanyInfoResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *RefreshCompanyInfoResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
