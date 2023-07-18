// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/assess/pkg/models/shared"
	"net/http"
)

type GetDataIntegritySummariesRequest struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	// A key for a Codat data type.
	DataType shared.DataIntegrityDataType `pathParam:"style=simple,explode=false,name=dataType"`
	// Codat query string. [Read more](https://docs.codat.io/using-the-api/querying).
	Query *string `queryParam:"style=form,explode=true,name=query"`
}

func (o *GetDataIntegritySummariesRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *GetDataIntegritySummariesRequest) GetDataType() shared.DataIntegrityDataType {
	if o == nil {
		return DataIntegrityDataType("")
	}
	return o.DataType
}

func (o *GetDataIntegritySummariesRequest) GetQuery() *string {
	if o == nil {
		return nil
	}
	return o.Query
}

type GetDataIntegritySummariesResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	Summaries *shared.Summaries
	// Your API request was not properly authorized.
	Schema *shared.Schema
}

func (o *GetDataIntegritySummariesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetDataIntegritySummariesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetDataIntegritySummariesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetDataIntegritySummariesResponse) GetSummaries() *shared.Summaries {
	if o == nil {
		return nil
	}
	return o.Summaries
}

func (o *GetDataIntegritySummariesResponse) GetSchema() *shared.Schema {
	if o == nil {
		return nil
	}
	return o.Schema
}
