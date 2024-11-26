// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/platform/v5/pkg/models/shared"
	"github.com/codatio/client-sdk-go/platform/v5/pkg/utils"
	"net/http"
)

type ListCustomDataTypeRecordsRequest struct {
	// Unique identifier for a company.
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	// Unique identifier for a connection.
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
	// Unique identifier for a custom data type.
	CustomDataIdentifier string `pathParam:"style=simple,explode=false,name=customDataIdentifier"`
	// Page number. [Read more](https://docs.codat.io/using-the-api/paging).
	Page *int `default:"1" queryParam:"style=form,explode=true,name=page"`
	// Number of records to return in a page. [Read more](https://docs.codat.io/using-the-api/paging).
	PageSize *int `default:"100" queryParam:"style=form,explode=true,name=pageSize"`
}

func (l ListCustomDataTypeRecordsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListCustomDataTypeRecordsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListCustomDataTypeRecordsRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *ListCustomDataTypeRecordsRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *ListCustomDataTypeRecordsRequest) GetCustomDataIdentifier() string {
	if o == nil {
		return ""
	}
	return o.CustomDataIdentifier
}

func (o *ListCustomDataTypeRecordsRequest) GetPage() *int {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *ListCustomDataTypeRecordsRequest) GetPageSize() *int {
	if o == nil {
		return nil
	}
	return o.PageSize
}

type ListCustomDataTypeRecordsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// OK
	CustomDataTypeRecords *shared.CustomDataTypeRecords
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ListCustomDataTypeRecordsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListCustomDataTypeRecordsResponse) GetCustomDataTypeRecords() *shared.CustomDataTypeRecords {
	if o == nil {
		return nil
	}
	return o.CustomDataTypeRecords
}

func (o *ListCustomDataTypeRecordsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListCustomDataTypeRecordsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
