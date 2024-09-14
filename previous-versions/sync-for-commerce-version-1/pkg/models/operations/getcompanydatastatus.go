// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1/pkg/models/shared"
	"net/http"
)

type GetCompanyDataStatusRequest struct {
	// Unique identifier for a company.
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

func (o *GetCompanyDataStatusRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

// GetCompanyDataStatusDataStatuses - OK
type GetCompanyDataStatusDataStatuses struct {
	// Describes the state of data in the Codat cache for a company and data type
	AccountTransactions *shared.DataStatus `json:"accountTransactions,omitempty"`
	// Describes the state of data in the Codat cache for a company and data type
	BalanceSheet *shared.DataStatus `json:"balanceSheet,omitempty"`
	// Describes the state of data in the Codat cache for a company and data type
	BankAccounts *shared.DataStatus `json:"bankAccounts,omitempty"`
	// Describes the state of data in the Codat cache for a company and data type
	BankTransactions *shared.DataStatus `json:"bankTransactions,omitempty"`
	// Describes the state of data in the Codat cache for a company and data type
	BankingAccountBalances *shared.DataStatus `json:"banking-accountBalances,omitempty"`
	// Describes the state of data in the Codat cache for a company and data type
	BankingAccounts *shared.DataStatus `json:"banking-accounts,omitempty"`
	// Describes the state of data in the Codat cache for a company and data type
	BankingTransactionCategories *shared.DataStatus `json:"banking-transactionCategories,omitempty"`
	// Describes the state of data in the Codat cache for a company and data type
	BankingTransactions *shared.DataStatus `json:"banking-transactions,omitempty"`
	// Describes the state of data in the Codat cache for a company and data type
	BillCreditNotes *shared.DataStatus `json:"billCreditNotes,omitempty"`
	// Describes the state of data in the Codat cache for a company and data type
	BillPayments *shared.DataStatus `json:"billPayments,omitempty"`
	// Describes the state of data in the Codat cache for a company and data type
	Bills *shared.DataStatus `json:"bills,omitempty"`
	// Describes the state of data in the Codat cache for a company and data type
	CashFlowStatement *shared.DataStatus `json:"cashFlowStatement,omitempty"`
	// Describes the state of data in the Codat cache for a company and data type
	ChartOfAccounts *shared.DataStatus `json:"chartOfAccounts,omitempty"`
	// Describes the state of data in the Codat cache for a company and data type
	CommerceCompanyInfo *shared.DataStatus `json:"commerce-companyInfo,omitempty"`
	// Describes the state of data in the Codat cache for a company and data type
	CommerceCustomers *shared.DataStatus `json:"commerce-customers,omitempty"`
	// Describes the state of data in the Codat cache for a company and data type
	CommerceDisputes *shared.DataStatus `json:"commerce-disputes,omitempty"`
	// Describes the state of data in the Codat cache for a company and data type
	CommerceLocations *shared.DataStatus `json:"commerce-locations,omitempty"`
	// Describes the state of data in the Codat cache for a company and data type
	CommerceOrders *shared.DataStatus `json:"commerce-orders,omitempty"`
	// Describes the state of data in the Codat cache for a company and data type
	CommercePaymentMethods *shared.DataStatus `json:"commerce-paymentMethods,omitempty"`
	// Describes the state of data in the Codat cache for a company and data type
	CommercePayments *shared.DataStatus `json:"commerce-payments,omitempty"`
	// Describes the state of data in the Codat cache for a company and data type
	CommerceProductCategories *shared.DataStatus `json:"commerce-productCategories,omitempty"`
	// Describes the state of data in the Codat cache for a company and data type
	CommerceProducts *shared.DataStatus `json:"commerce-products,omitempty"`
	// Describes the state of data in the Codat cache for a company and data type
	CommerceTaxComponents *shared.DataStatus `json:"commerce-taxComponents,omitempty"`
	// Describes the state of data in the Codat cache for a company and data type
	CommerceTransactions *shared.DataStatus `json:"commerce-transactions,omitempty"`
	// Describes the state of data in the Codat cache for a company and data type
	Company *shared.DataStatus `json:"company,omitempty"`
	// Describes the state of data in the Codat cache for a company and data type
	CreditNotes *shared.DataStatus `json:"creditNotes,omitempty"`
	// Describes the state of data in the Codat cache for a company and data type
	Customers *shared.DataStatus `json:"customers,omitempty"`
	// Describes the state of data in the Codat cache for a company and data type
	DirectCosts *shared.DataStatus `json:"directCosts,omitempty"`
	// Describes the state of data in the Codat cache for a company and data type
	DirectIncomes *shared.DataStatus `json:"directIncomes,omitempty"`
	// Describes the state of data in the Codat cache for a company and data type
	Invoices *shared.DataStatus `json:"invoices,omitempty"`
	// Describes the state of data in the Codat cache for a company and data type
	ItemReceipts *shared.DataStatus `json:"itemReceipts,omitempty"`
	// Describes the state of data in the Codat cache for a company and data type
	Items *shared.DataStatus `json:"items,omitempty"`
	// Describes the state of data in the Codat cache for a company and data type
	JournalEntries *shared.DataStatus `json:"journalEntries,omitempty"`
	// Describes the state of data in the Codat cache for a company and data type
	Journals *shared.DataStatus `json:"journals,omitempty"`
	// Describes the state of data in the Codat cache for a company and data type
	PaymentMethods *shared.DataStatus `json:"paymentMethods,omitempty"`
	// Describes the state of data in the Codat cache for a company and data type
	Payments *shared.DataStatus `json:"payments,omitempty"`
	// Describes the state of data in the Codat cache for a company and data type
	ProfitAndLoss *shared.DataStatus `json:"profitAndLoss,omitempty"`
	// Describes the state of data in the Codat cache for a company and data type
	PurchaseOrders *shared.DataStatus `json:"purchaseOrders,omitempty"`
	// Describes the state of data in the Codat cache for a company and data type
	SalesOrders *shared.DataStatus `json:"salesOrders,omitempty"`
	// Describes the state of data in the Codat cache for a company and data type
	Suppliers *shared.DataStatus `json:"suppliers,omitempty"`
	// Describes the state of data in the Codat cache for a company and data type
	TaxRates *shared.DataStatus `json:"taxRates,omitempty"`
	// Describes the state of data in the Codat cache for a company and data type
	TrackingCategories *shared.DataStatus `json:"trackingCategories,omitempty"`
	// Describes the state of data in the Codat cache for a company and data type
	Transfers *shared.DataStatus `json:"transfers,omitempty"`
}

func (o *GetCompanyDataStatusDataStatuses) GetAccountTransactions() *shared.DataStatus {
	if o == nil {
		return nil
	}
	return o.AccountTransactions
}

func (o *GetCompanyDataStatusDataStatuses) GetBalanceSheet() *shared.DataStatus {
	if o == nil {
		return nil
	}
	return o.BalanceSheet
}

func (o *GetCompanyDataStatusDataStatuses) GetBankAccounts() *shared.DataStatus {
	if o == nil {
		return nil
	}
	return o.BankAccounts
}

func (o *GetCompanyDataStatusDataStatuses) GetBankTransactions() *shared.DataStatus {
	if o == nil {
		return nil
	}
	return o.BankTransactions
}

func (o *GetCompanyDataStatusDataStatuses) GetBankingAccountBalances() *shared.DataStatus {
	if o == nil {
		return nil
	}
	return o.BankingAccountBalances
}

func (o *GetCompanyDataStatusDataStatuses) GetBankingAccounts() *shared.DataStatus {
	if o == nil {
		return nil
	}
	return o.BankingAccounts
}

func (o *GetCompanyDataStatusDataStatuses) GetBankingTransactionCategories() *shared.DataStatus {
	if o == nil {
		return nil
	}
	return o.BankingTransactionCategories
}

func (o *GetCompanyDataStatusDataStatuses) GetBankingTransactions() *shared.DataStatus {
	if o == nil {
		return nil
	}
	return o.BankingTransactions
}

func (o *GetCompanyDataStatusDataStatuses) GetBillCreditNotes() *shared.DataStatus {
	if o == nil {
		return nil
	}
	return o.BillCreditNotes
}

func (o *GetCompanyDataStatusDataStatuses) GetBillPayments() *shared.DataStatus {
	if o == nil {
		return nil
	}
	return o.BillPayments
}

func (o *GetCompanyDataStatusDataStatuses) GetBills() *shared.DataStatus {
	if o == nil {
		return nil
	}
	return o.Bills
}

func (o *GetCompanyDataStatusDataStatuses) GetCashFlowStatement() *shared.DataStatus {
	if o == nil {
		return nil
	}
	return o.CashFlowStatement
}

func (o *GetCompanyDataStatusDataStatuses) GetChartOfAccounts() *shared.DataStatus {
	if o == nil {
		return nil
	}
	return o.ChartOfAccounts
}

func (o *GetCompanyDataStatusDataStatuses) GetCommerceCompanyInfo() *shared.DataStatus {
	if o == nil {
		return nil
	}
	return o.CommerceCompanyInfo
}

func (o *GetCompanyDataStatusDataStatuses) GetCommerceCustomers() *shared.DataStatus {
	if o == nil {
		return nil
	}
	return o.CommerceCustomers
}

func (o *GetCompanyDataStatusDataStatuses) GetCommerceDisputes() *shared.DataStatus {
	if o == nil {
		return nil
	}
	return o.CommerceDisputes
}

func (o *GetCompanyDataStatusDataStatuses) GetCommerceLocations() *shared.DataStatus {
	if o == nil {
		return nil
	}
	return o.CommerceLocations
}

func (o *GetCompanyDataStatusDataStatuses) GetCommerceOrders() *shared.DataStatus {
	if o == nil {
		return nil
	}
	return o.CommerceOrders
}

func (o *GetCompanyDataStatusDataStatuses) GetCommercePaymentMethods() *shared.DataStatus {
	if o == nil {
		return nil
	}
	return o.CommercePaymentMethods
}

func (o *GetCompanyDataStatusDataStatuses) GetCommercePayments() *shared.DataStatus {
	if o == nil {
		return nil
	}
	return o.CommercePayments
}

func (o *GetCompanyDataStatusDataStatuses) GetCommerceProductCategories() *shared.DataStatus {
	if o == nil {
		return nil
	}
	return o.CommerceProductCategories
}

func (o *GetCompanyDataStatusDataStatuses) GetCommerceProducts() *shared.DataStatus {
	if o == nil {
		return nil
	}
	return o.CommerceProducts
}

func (o *GetCompanyDataStatusDataStatuses) GetCommerceTaxComponents() *shared.DataStatus {
	if o == nil {
		return nil
	}
	return o.CommerceTaxComponents
}

func (o *GetCompanyDataStatusDataStatuses) GetCommerceTransactions() *shared.DataStatus {
	if o == nil {
		return nil
	}
	return o.CommerceTransactions
}

func (o *GetCompanyDataStatusDataStatuses) GetCompany() *shared.DataStatus {
	if o == nil {
		return nil
	}
	return o.Company
}

func (o *GetCompanyDataStatusDataStatuses) GetCreditNotes() *shared.DataStatus {
	if o == nil {
		return nil
	}
	return o.CreditNotes
}

func (o *GetCompanyDataStatusDataStatuses) GetCustomers() *shared.DataStatus {
	if o == nil {
		return nil
	}
	return o.Customers
}

func (o *GetCompanyDataStatusDataStatuses) GetDirectCosts() *shared.DataStatus {
	if o == nil {
		return nil
	}
	return o.DirectCosts
}

func (o *GetCompanyDataStatusDataStatuses) GetDirectIncomes() *shared.DataStatus {
	if o == nil {
		return nil
	}
	return o.DirectIncomes
}

func (o *GetCompanyDataStatusDataStatuses) GetInvoices() *shared.DataStatus {
	if o == nil {
		return nil
	}
	return o.Invoices
}

func (o *GetCompanyDataStatusDataStatuses) GetItemReceipts() *shared.DataStatus {
	if o == nil {
		return nil
	}
	return o.ItemReceipts
}

func (o *GetCompanyDataStatusDataStatuses) GetItems() *shared.DataStatus {
	if o == nil {
		return nil
	}
	return o.Items
}

func (o *GetCompanyDataStatusDataStatuses) GetJournalEntries() *shared.DataStatus {
	if o == nil {
		return nil
	}
	return o.JournalEntries
}

func (o *GetCompanyDataStatusDataStatuses) GetJournals() *shared.DataStatus {
	if o == nil {
		return nil
	}
	return o.Journals
}

func (o *GetCompanyDataStatusDataStatuses) GetPaymentMethods() *shared.DataStatus {
	if o == nil {
		return nil
	}
	return o.PaymentMethods
}

func (o *GetCompanyDataStatusDataStatuses) GetPayments() *shared.DataStatus {
	if o == nil {
		return nil
	}
	return o.Payments
}

func (o *GetCompanyDataStatusDataStatuses) GetProfitAndLoss() *shared.DataStatus {
	if o == nil {
		return nil
	}
	return o.ProfitAndLoss
}

func (o *GetCompanyDataStatusDataStatuses) GetPurchaseOrders() *shared.DataStatus {
	if o == nil {
		return nil
	}
	return o.PurchaseOrders
}

func (o *GetCompanyDataStatusDataStatuses) GetSalesOrders() *shared.DataStatus {
	if o == nil {
		return nil
	}
	return o.SalesOrders
}

func (o *GetCompanyDataStatusDataStatuses) GetSuppliers() *shared.DataStatus {
	if o == nil {
		return nil
	}
	return o.Suppliers
}

func (o *GetCompanyDataStatusDataStatuses) GetTaxRates() *shared.DataStatus {
	if o == nil {
		return nil
	}
	return o.TaxRates
}

func (o *GetCompanyDataStatusDataStatuses) GetTrackingCategories() *shared.DataStatus {
	if o == nil {
		return nil
	}
	return o.TrackingCategories
}

func (o *GetCompanyDataStatusDataStatuses) GetTransfers() *shared.DataStatus {
	if o == nil {
		return nil
	}
	return o.Transfers
}

type GetCompanyDataStatusResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// OK
	DataStatuses *GetCompanyDataStatusDataStatuses
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetCompanyDataStatusResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetCompanyDataStatusResponse) GetDataStatuses() *GetCompanyDataStatusDataStatuses {
	if o == nil {
		return nil
	}
	return o.DataStatuses
}

func (o *GetCompanyDataStatusResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetCompanyDataStatusResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
