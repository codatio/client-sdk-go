package operations

import (
	"time"
)

type GetCompaniesCompanyIDReportsEnhancedBalanaceSheetAccountsPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type GetCompaniesCompanyIDReportsEnhancedBalanaceSheetAccountsQueryParams struct {
	NumberOfPeriods int64  `queryParam:"style=form,explode=true,name=numberOfPeriods"`
	ReportDate      string `queryParam:"style=form,explode=true,name=reportDate"`
}

type GetCompaniesCompanyIDReportsEnhancedBalanaceSheetAccountsRequest struct {
	PathParams  GetCompaniesCompanyIDReportsEnhancedBalanaceSheetAccountsPathParams
	QueryParams GetCompaniesCompanyIDReportsEnhancedBalanaceSheetAccountsQueryParams
}

type GetCompaniesCompanyIDReportsEnhancedBalanaceSheetAccountsEnhancedReportReportInfo struct {
	CompanyName   *string    `json:"companyName,omitempty"`
	Currency      *string    `json:"currency,omitempty"`
	GeneratedDate *time.Time `json:"generatedDate,omitempty"`
	ReportName    *string    `json:"reportName,omitempty"`
}

type GetCompaniesCompanyIDReportsEnhancedBalanaceSheetAccountsEnhancedReportReportItemsAccountCategoryLevels struct {
	Confidence *float64 `json:"confidence,omitempty"`
	LevelName  *string  `json:"levelName,omitempty"`
}

type GetCompaniesCompanyIDReportsEnhancedBalanaceSheetAccountsEnhancedReportReportItemsAccountCategory struct {
	Levels []GetCompaniesCompanyIDReportsEnhancedBalanaceSheetAccountsEnhancedReportReportItemsAccountCategoryLevels `json:"levels,omitempty"`
	Status *string                                                                                                   `json:"status,omitempty"`
}

type GetCompaniesCompanyIDReportsEnhancedBalanaceSheetAccountsEnhancedReportReportItems struct {
	AccountCategory *GetCompaniesCompanyIDReportsEnhancedBalanaceSheetAccountsEnhancedReportReportItemsAccountCategory `json:"accountCategory,omitempty"`
	AccountID       *string                                                                                            `json:"accountId,omitempty"`
	AccountName     *string                                                                                            `json:"accountName,omitempty"`
	Balance         *string                                                                                            `json:"balance,omitempty"`
	Date            *time.Time                                                                                         `json:"date,omitempty"`
}

type GetCompaniesCompanyIDReportsEnhancedBalanaceSheetAccountsEnhancedReport struct {
	ReportInfo  *GetCompaniesCompanyIDReportsEnhancedBalanaceSheetAccountsEnhancedReportReportInfo   `json:"reportInfo,omitempty"`
	ReportItems []GetCompaniesCompanyIDReportsEnhancedBalanaceSheetAccountsEnhancedReportReportItems `json:"reportItems,omitempty"`
}

type GetCompaniesCompanyIDReportsEnhancedBalanaceSheetAccountsResponse struct {
	ContentType    string
	EnhancedReport *GetCompaniesCompanyIDReportsEnhancedBalanaceSheetAccountsEnhancedReport
	StatusCode     int
}
