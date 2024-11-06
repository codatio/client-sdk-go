// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/lending/v6/pkg/models/shared"
	"net/http"
)

type GetAccountingProfitAndLossRequest struct {
	// Unique identifier for a company.
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	// Number of months defining the period of interest.
	PeriodLength int `queryParam:"style=form,explode=true,name=periodLength"`
	// Number of periods with `periodLength` to compare.
	PeriodsToCompare int `queryParam:"style=form,explode=true,name=periodsToCompare"`
	// The month the report starts from.
	StartMonth *string `queryParam:"style=form,explode=true,name=startMonth"`
}

func (o *GetAccountingProfitAndLossRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *GetAccountingProfitAndLossRequest) GetPeriodLength() int {
	if o == nil {
		return 0
	}
	return o.PeriodLength
}

func (o *GetAccountingProfitAndLossRequest) GetPeriodsToCompare() int {
	if o == nil {
		return 0
	}
	return o.PeriodsToCompare
}

func (o *GetAccountingProfitAndLossRequest) GetStartMonth() *string {
	if o == nil {
		return nil
	}
	return o.StartMonth
}

type GetAccountingProfitAndLossResponse struct {
	// Success
	AccountingProfitAndLossReport *shared.AccountingProfitAndLossReport
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetAccountingProfitAndLossResponse) GetAccountingProfitAndLossReport() *shared.AccountingProfitAndLossReport {
	if o == nil {
		return nil
	}
	return o.AccountingProfitAndLossReport
}

func (o *GetAccountingProfitAndLossResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetAccountingProfitAndLossResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetAccountingProfitAndLossResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
