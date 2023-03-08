package operations

import (
	"net/http"
	"time"
)

type GetAccountsForEnhancedBalanceSheetPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type GetAccountsForEnhancedBalanceSheetQueryParams struct {
	NumberOfPeriods int64  `queryParam:"style=form,explode=true,name=numberOfPeriods"`
	ReportDate      string `queryParam:"style=form,explode=true,name=reportDate"`
}

type GetAccountsForEnhancedBalanceSheetRequest struct {
	PathParams  GetAccountsForEnhancedBalanceSheetPathParams
	QueryParams GetAccountsForEnhancedBalanceSheetQueryParams
}

type GetAccountsForEnhancedBalanceSheetEnhancedReportReportInfo struct {
	CompanyName   *string    `json:"companyName,omitempty"`
	Currency      *string    `json:"currency,omitempty"`
	GeneratedDate *time.Time `json:"generatedDate,omitempty"`
	ReportName    *string    `json:"reportName,omitempty"`
}

type GetAccountsForEnhancedBalanceSheetEnhancedReportReportItemsAccountCategoryLevels struct {
	Confidence *float64 `json:"confidence,omitempty"`
	LevelName  *string  `json:"levelName,omitempty"`
}

type GetAccountsForEnhancedBalanceSheetEnhancedReportReportItemsAccountCategory struct {
	Levels []GetAccountsForEnhancedBalanceSheetEnhancedReportReportItemsAccountCategoryLevels `json:"levels,omitempty"`
	Status *string                                                                            `json:"status,omitempty"`
}

type GetAccountsForEnhancedBalanceSheetEnhancedReportReportItems struct {
	AccountCategory *GetAccountsForEnhancedBalanceSheetEnhancedReportReportItemsAccountCategory `json:"accountCategory,omitempty"`
	AccountID       *string                                                                     `json:"accountId,omitempty"`
	AccountName     *string                                                                     `json:"accountName,omitempty"`
	Balance         *string                                                                     `json:"balance,omitempty"`
	Date            *time.Time                                                                  `json:"date,omitempty"`
}

type GetAccountsForEnhancedBalanceSheetEnhancedReport struct {
	ReportInfo  *GetAccountsForEnhancedBalanceSheetEnhancedReportReportInfo   `json:"reportInfo,omitempty"`
	ReportItems []GetAccountsForEnhancedBalanceSheetEnhancedReportReportItems `json:"reportItems,omitempty"`
}

type GetAccountsForEnhancedBalanceSheetResponse struct {
	ContentType    string
	EnhancedReport *GetAccountsForEnhancedBalanceSheetEnhancedReport
	StatusCode     int
	RawResponse    *http.Response
}
