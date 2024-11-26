// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/previous-versions/commerce/v3/pkg/models/shared"
	"net/http"
)

type GetProductCategoryRequest struct {
	// Unique identifier for a company.
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	// Unique identifier for a connection.
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
	// Unique identifier for a product.
	ProductID string `pathParam:"style=simple,explode=false,name=productId"`
}

func (o *GetProductCategoryRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *GetProductCategoryRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *GetProductCategoryRequest) GetProductID() string {
	if o == nil {
		return ""
	}
	return o.ProductID
}

type GetProductCategoryResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// OK
	ProductCategory *shared.ProductCategory
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetProductCategoryResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetProductCategoryResponse) GetProductCategory() *shared.ProductCategory {
	if o == nil {
		return nil
	}
	return o.ProductCategory
}

func (o *GetProductCategoryResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetProductCategoryResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
