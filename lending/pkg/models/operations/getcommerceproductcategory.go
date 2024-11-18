// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/lending/v7/pkg/models/shared"
	"net/http"
)

type GetCommerceProductCategoryRequest struct {
	// Unique identifier for a company.
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	// Unique identifier for a connection.
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
	// Unique identifier for a product.
	ProductID string `pathParam:"style=simple,explode=false,name=productId"`
}

func (o *GetCommerceProductCategoryRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *GetCommerceProductCategoryRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *GetCommerceProductCategoryRequest) GetProductID() string {
	if o == nil {
		return ""
	}
	return o.ProductID
}

type GetCommerceProductCategoryResponse struct {
	// OK
	CommerceProductCategory *shared.CommerceProductCategory
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetCommerceProductCategoryResponse) GetCommerceProductCategory() *shared.CommerceProductCategory {
	if o == nil {
		return nil
	}
	return o.CommerceProductCategory
}

func (o *GetCommerceProductCategoryResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetCommerceProductCategoryResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetCommerceProductCategoryResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
