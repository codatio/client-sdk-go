package shared

import (
	"time"
)

type EnhancedCashFlowTransactionsDataSourcesAccountsReportSourceRef struct {
	SourceType *string `json:"sourceType,omitempty"`
}

type EnhancedCashFlowTransactionsDataSourcesAccounts struct {
	AccountName     *string                                                         `json:"accountName,omitempty"`
	AccountProvider *string                                                         `json:"accountProvider,omitempty"`
	AccountType     *string                                                         `json:"accountType,omitempty"`
	Currency        *string                                                         `json:"currency,omitempty"`
	CurrentBalance  *float64                                                        `json:"currentBalance,omitempty"`
	PlatformName    *string                                                         `json:"platformName,omitempty"`
	SourceRef       *EnhancedCashFlowTransactionsDataSourcesAccountsReportSourceRef `json:"sourceRef,omitempty"`
}

type EnhancedCashFlowTransactionsDataSources struct {
	Accounts []EnhancedCashFlowTransactionsDataSourcesAccounts `json:"accounts,omitempty"`
}

type EnhancedCashFlowTransactionsReportInfo struct {
	CompanyName   *string `json:"companyName,omitempty"`
	GeneratedDate *string `json:"generatedDate,omitempty"`
	PageNumber    *int64  `json:"pageNumber,omitempty"`
	PageSize      *int64  `json:"pageSize,omitempty"`
	ReportName    *string `json:"reportName,omitempty"`
	TotalResults  *int64  `json:"totalResults,omitempty"`
}

type EnhancedCashFlowTransactionsReportItemsTransactionsTransactionCategory struct {
	Confidence *float64 `json:"confidence,omitempty"`
	Levels     []string `json:"levels,omitempty"`
}

type EnhancedCashFlowTransactionsReportItemsTransactions struct {
	Amount              *float64                                                                `json:"amount,omitempty"`
	Currency            *string                                                                 `json:"currency,omitempty"`
	Date                *time.Time                                                              `json:"date,omitempty"`
	Description         *string                                                                 `json:"description,omitempty"`
	ID                  *string                                                                 `json:"id,omitempty"`
	SourceRef           *SourceRef                                                              `json:"sourceRef,omitempty"`
	TransactionCategory *EnhancedCashFlowTransactionsReportItemsTransactionsTransactionCategory `json:"transactionCategory,omitempty"`
}

type EnhancedCashFlowTransactionsReportItems struct {
	Transactions []EnhancedCashFlowTransactionsReportItemsTransactions `json:"transactions,omitempty"`
}

type EnhancedCashFlowTransactions struct {
	DataSources []EnhancedCashFlowTransactionsDataSources `json:"dataSources,omitempty"`
	ReportInfo  *EnhancedCashFlowTransactionsReportInfo   `json:"reportInfo,omitempty"`
	ReportItems []EnhancedCashFlowTransactionsReportItems `json:"reportItems,omitempty"`
}
