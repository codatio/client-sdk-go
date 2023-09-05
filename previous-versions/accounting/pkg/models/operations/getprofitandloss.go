// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/shared"
	"net/http"
)

type GetProfitAndLossRequest struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	// Number of months defining the period of interest.
	PeriodLength int `queryParam:"style=form,explode=true,name=periodLength"`
	// Number of periods with `periodLength` to compare.
	PeriodsToCompare int     `queryParam:"style=form,explode=true,name=periodsToCompare"`
	StartMonth       *string `queryParam:"style=form,explode=true,name=startMonth"`
}

func (o *GetProfitAndLossRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *GetProfitAndLossRequest) GetPeriodLength() int {
	if o == nil {
		return 0
	}
	return o.PeriodLength
}

func (o *GetProfitAndLossRequest) GetPeriodsToCompare() int {
	if o == nil {
		return 0
	}
	return o.PeriodsToCompare
}

func (o *GetProfitAndLossRequest) GetStartMonth() *string {
	if o == nil {
		return nil
	}
	return o.StartMonth
}

type GetProfitAndLossResponse struct {
	ContentType string
	// Your API request was not properly authorized.
	ErrorMessage *shared.ErrorMessage
	// Success
	ProfitAndLossReport *shared.ProfitAndLossReport1
	StatusCode          int
	RawResponse         *http.Response
}

func (o *GetProfitAndLossResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetProfitAndLossResponse) GetErrorMessage() *shared.ErrorMessage {
	if o == nil {
		return nil
	}
	return o.ErrorMessage
}

func (o *GetProfitAndLossResponse) GetProfitAndLossReport() *shared.ProfitAndLossReport1 {
	if o == nil {
		return nil
	}
	return o.ProfitAndLossReport
}

func (o *GetProfitAndLossResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetProfitAndLossResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}