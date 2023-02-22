package shared

import (
	"time"
)

type EnhancedReportReportInfo struct {
	CompanyName   *string    `json:"companyName,omitempty"`
	Currency      *string    `json:"currency,omitempty"`
	GeneratedDate *time.Time `json:"generatedDate,omitempty"`
	ReportName    *string    `json:"reportName,omitempty"`
}

type EnhancedReportReportItemsAccountCategoryLevels struct {
	Confidence *float64 `json:"confidence,omitempty"`
	LevelName  *string  `json:"levelName,omitempty"`
}

type EnhancedReportReportItemsAccountCategory struct {
	Levels []EnhancedReportReportItemsAccountCategoryLevels `json:"levels,omitempty"`
	Status *string                                          `json:"status,omitempty"`
}

type EnhancedReportReportItems struct {
	AccountCategory *EnhancedReportReportItemsAccountCategory `json:"accountCategory,omitempty"`
	AccountID       *string                                   `json:"accountId,omitempty"`
	AccountName     *string                                   `json:"accountName,omitempty"`
	Balance         *string                                   `json:"balance,omitempty"`
	Date            *time.Time                                `json:"date,omitempty"`
}

type EnhancedReport struct {
	ReportInfo  *EnhancedReportReportInfo   `json:"reportInfo,omitempty"`
	ReportItems []EnhancedReportReportItems `json:"reportItems,omitempty"`
}
