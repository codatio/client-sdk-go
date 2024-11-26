// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/previous-versions/commerce/v3/pkg/models/shared"
	"net/http"
)

type GetTaxComponentRequest struct {
	// Unique identifier for a company.
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	// Unique identifier for a connection.
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
	// Unique identifier for a tax component.
	TaxID string `pathParam:"style=simple,explode=false,name=taxId"`
}

func (o *GetTaxComponentRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *GetTaxComponentRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *GetTaxComponentRequest) GetTaxID() string {
	if o == nil {
		return ""
	}
	return o.TaxID
}

type GetTaxComponentResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	TaxComponent *shared.TaxComponent
}

func (o *GetTaxComponentResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetTaxComponentResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetTaxComponentResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetTaxComponentResponse) GetTaxComponent() *shared.TaxComponent {
	if o == nil {
		return nil
	}
	return o.TaxComponent
}
