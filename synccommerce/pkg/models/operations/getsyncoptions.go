package operations

import (
	"net/http"
	"time"
)

type GetSyncOptionsRequest struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type GetSyncOptions200ApplicationJSONConfigurationFeesAccountsAccountOptions struct {
	Classification *string `json:"classification,omitempty"`
	ID             *string `json:"id,omitempty"`
	Name           *string `json:"name,omitempty"`
	NominalCode    *string `json:"nominalCode,omitempty"`
}

type GetSyncOptions200ApplicationJSONConfigurationFeesAccounts struct {
	AccountOptions    []GetSyncOptions200ApplicationJSONConfigurationFeesAccountsAccountOptions `json:"accountOptions,omitempty"`
	DescriptionText   *string                                                                   `json:"descriptionText,omitempty"`
	LabelText         *string                                                                   `json:"labelText,omitempty"`
	Required          *bool                                                                     `json:"required,omitempty"`
	SelectedAccountID *string                                                                   `json:"selectedAccountId,omitempty"`
}

type GetSyncOptions200ApplicationJSONConfigurationFeesFeesSupplierSupplierOptions struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type GetSyncOptions200ApplicationJSONConfigurationFeesFeesSupplier struct {
	SelectedSupplierID *string                                                                        `json:"selectedSupplierId,omitempty"`
	SupplierOptions    []GetSyncOptions200ApplicationJSONConfigurationFeesFeesSupplierSupplierOptions `json:"supplierOptions,omitempty"`
}

type GetSyncOptions200ApplicationJSONConfigurationFees struct {
	Accounts     map[string]GetSyncOptions200ApplicationJSONConfigurationFeesAccounts `json:"accounts,omitempty"`
	FeesSupplier *GetSyncOptions200ApplicationJSONConfigurationFeesFeesSupplier       `json:"feesSupplier,omitempty"`
	SyncFees     *bool                                                                `json:"syncFees,omitempty"`
}

type GetSyncOptions200ApplicationJSONConfigurationNewPaymentsAccountsAccountOptions struct {
	Classification *string `json:"classification,omitempty"`
	ID             *string `json:"id,omitempty"`
	Name           *string `json:"name,omitempty"`
	NominalCode    *string `json:"nominalCode,omitempty"`
}

type GetSyncOptions200ApplicationJSONConfigurationNewPaymentsAccounts struct {
	AccountOptions    []GetSyncOptions200ApplicationJSONConfigurationNewPaymentsAccountsAccountOptions `json:"accountOptions,omitempty"`
	DescriptionText   *string                                                                          `json:"descriptionText,omitempty"`
	LabelText         *string                                                                          `json:"labelText,omitempty"`
	Required          *bool                                                                            `json:"required,omitempty"`
	SelectedAccountID *string                                                                          `json:"selectedAccountId,omitempty"`
}

type GetSyncOptions200ApplicationJSONConfigurationNewPayments struct {
	Accounts     map[string]GetSyncOptions200ApplicationJSONConfigurationNewPaymentsAccounts `json:"accounts,omitempty"`
	SyncPayments *bool                                                                       `json:"syncPayments,omitempty"`
}

type GetSyncOptions200ApplicationJSONConfigurationPaymentsAccountsAccountOptions struct {
	Classification *string `json:"classification,omitempty"`
	ID             *string `json:"id,omitempty"`
	Name           *string `json:"name,omitempty"`
	NominalCode    *string `json:"nominalCode,omitempty"`
}

type GetSyncOptions200ApplicationJSONConfigurationPaymentsAccounts struct {
	AccountOptions    []GetSyncOptions200ApplicationJSONConfigurationPaymentsAccountsAccountOptions `json:"accountOptions,omitempty"`
	DescriptionText   *string                                                                       `json:"descriptionText,omitempty"`
	LabelText         *string                                                                       `json:"labelText,omitempty"`
	Required          *bool                                                                         `json:"required,omitempty"`
	SelectedAccountID *string                                                                       `json:"selectedAccountId,omitempty"`
}

type GetSyncOptions200ApplicationJSONConfigurationPayments struct {
	Accounts     map[string]GetSyncOptions200ApplicationJSONConfigurationPaymentsAccounts `json:"accounts,omitempty"`
	SyncPayments *bool                                                                    `json:"syncPayments,omitempty"`
}

type GetSyncOptions200ApplicationJSONConfigurationSalesAccountsAccountOptions struct {
	Classification *string `json:"classification,omitempty"`
	ID             *string `json:"id,omitempty"`
	Name           *string `json:"name,omitempty"`
	NominalCode    *string `json:"nominalCode,omitempty"`
}

type GetSyncOptions200ApplicationJSONConfigurationSalesAccounts struct {
	AccountOptions    []GetSyncOptions200ApplicationJSONConfigurationSalesAccountsAccountOptions `json:"accountOptions,omitempty"`
	DescriptionText   *string                                                                    `json:"descriptionText,omitempty"`
	LabelText         *string                                                                    `json:"labelText,omitempty"`
	Required          *bool                                                                      `json:"required,omitempty"`
	SelectedAccountID *string                                                                    `json:"selectedAccountId,omitempty"`
}

type GetSyncOptions200ApplicationJSONConfigurationSalesGroupingGroupingLevelsInvoiceLevel struct {
	GroupByOptions         []string `json:"groupByOptions,omitempty"`
	SelectedGroupByOptions []string `json:"selectedGroupByOptions,omitempty"`
}

type GetSyncOptions200ApplicationJSONConfigurationSalesGroupingGroupingLevelsInvoiceLineLevel struct {
	GroupByOptions         []string `json:"groupByOptions,omitempty"`
	SelectedGroupByOptions []string `json:"selectedGroupByOptions,omitempty"`
}

type GetSyncOptions200ApplicationJSONConfigurationSalesGroupingGroupingLevels struct {
	InvoiceLevel     *GetSyncOptions200ApplicationJSONConfigurationSalesGroupingGroupingLevelsInvoiceLevel     `json:"invoiceLevel,omitempty"`
	InvoiceLineLevel *GetSyncOptions200ApplicationJSONConfigurationSalesGroupingGroupingLevelsInvoiceLineLevel `json:"invoiceLineLevel,omitempty"`
}

type GetSyncOptions200ApplicationJSONConfigurationSalesGroupingGroupingPeriod struct {
	GroupingPeriodOptions  []string `json:"groupingPeriodOptions,omitempty"`
	SelectedGroupingPeriod *string  `json:"selectedGroupingPeriod,omitempty"`
}

type GetSyncOptions200ApplicationJSONConfigurationSalesGrouping struct {
	GroupingLevels *GetSyncOptions200ApplicationJSONConfigurationSalesGroupingGroupingLevels `json:"groupingLevels,omitempty"`
	GroupingPeriod *GetSyncOptions200ApplicationJSONConfigurationSalesGroupingGroupingPeriod `json:"groupingPeriod,omitempty"`
}

type GetSyncOptions200ApplicationJSONConfigurationSalesInvoiceStatus struct {
	InvoiceStatusOptions  []string `json:"invoiceStatusOptions,omitempty"`
	SelectedInvoiceStatus *string  `json:"selectedInvoiceStatus,omitempty"`
}

type GetSyncOptions200ApplicationJSONConfigurationSalesNewTaxRatesAccountingTaxRateOptions struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type GetSyncOptions200ApplicationJSONConfigurationSalesNewTaxRatesCommerceTaxRateOptions struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type GetSyncOptions200ApplicationJSONConfigurationSalesNewTaxRatesDefaultZeroTaxRateOptions struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type GetSyncOptions200ApplicationJSONConfigurationSalesNewTaxRatesTaxRateMappings struct {
	SelectedAccountingTaxRateID *string  `json:"selectedAccountingTaxRateId,omitempty"`
	SelectedCommerceTaxRateIds  []string `json:"selectedCommerceTaxRateIds,omitempty"`
}

type GetSyncOptions200ApplicationJSONConfigurationSalesNewTaxRates struct {
	AccountingTaxRateOptions     []GetSyncOptions200ApplicationJSONConfigurationSalesNewTaxRatesAccountingTaxRateOptions  `json:"accountingTaxRateOptions,omitempty"`
	CommerceTaxRateOptions       []GetSyncOptions200ApplicationJSONConfigurationSalesNewTaxRatesCommerceTaxRateOptions    `json:"commerceTaxRateOptions,omitempty"`
	DefaultZeroTaxRateOptions    []GetSyncOptions200ApplicationJSONConfigurationSalesNewTaxRatesDefaultZeroTaxRateOptions `json:"defaultZeroTaxRateOptions,omitempty"`
	SelectedDefaultZeroTaxRateID *string                                                                                  `json:"selectedDefaultZeroTaxRateId,omitempty"`
	TaxRateMappings              []GetSyncOptions200ApplicationJSONConfigurationSalesNewTaxRatesTaxRateMappings           `json:"taxRateMappings,omitempty"`
}

type GetSyncOptions200ApplicationJSONConfigurationSalesSalesCustomerCustomerOptions struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type GetSyncOptions200ApplicationJSONConfigurationSalesSalesCustomer struct {
	CustomerOptions    []GetSyncOptions200ApplicationJSONConfigurationSalesSalesCustomerCustomerOptions `json:"customerOptions,omitempty"`
	SelectedCustomerID *string                                                                          `json:"selectedCustomerId,omitempty"`
}

type GetSyncOptions200ApplicationJSONConfigurationSalesTaxRatesTaxRateOptions struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type GetSyncOptions200ApplicationJSONConfigurationSalesTaxRates struct {
	SelectedTaxRateID *string                                                                    `json:"selectedTaxRateId,omitempty"`
	TaxRateOptions    []GetSyncOptions200ApplicationJSONConfigurationSalesTaxRatesTaxRateOptions `json:"taxRateOptions,omitempty"`
}

type GetSyncOptions200ApplicationJSONConfigurationSales struct {
	Accounts      map[string]GetSyncOptions200ApplicationJSONConfigurationSalesAccounts `json:"accounts,omitempty"`
	Grouping      *GetSyncOptions200ApplicationJSONConfigurationSalesGrouping           `json:"grouping,omitempty"`
	InvoiceStatus *GetSyncOptions200ApplicationJSONConfigurationSalesInvoiceStatus      `json:"invoiceStatus,omitempty"`
	NewTaxRates   *GetSyncOptions200ApplicationJSONConfigurationSalesNewTaxRates        `json:"newTaxRates,omitempty"`
	SalesCustomer *GetSyncOptions200ApplicationJSONConfigurationSalesSalesCustomer      `json:"salesCustomer,omitempty"`
	SyncSales     *bool                                                                 `json:"syncSales,omitempty"`
	TaxRates      map[string]GetSyncOptions200ApplicationJSONConfigurationSalesTaxRates `json:"taxRates,omitempty"`
}

type GetSyncOptions200ApplicationJSONConfiguration struct {
	Fees        *GetSyncOptions200ApplicationJSONConfigurationFees        `json:"fees,omitempty"`
	NewPayments *GetSyncOptions200ApplicationJSONConfigurationNewPayments `json:"newPayments,omitempty"`
	Payments    *GetSyncOptions200ApplicationJSONConfigurationPayments    `json:"payments,omitempty"`
	Sales       *GetSyncOptions200ApplicationJSONConfigurationSales       `json:"sales,omitempty"`
}

type GetSyncOptions200ApplicationJSONSchedule struct {
	FrequencyOptions  []string   `json:"frequencyOptions,omitempty"`
	SelectedFrequency *string    `json:"selectedFrequency,omitempty"`
	StartDate         *time.Time `json:"startDate,omitempty"`
	SyncHourUtc       *int       `json:"syncHourUtc,omitempty"`
	TimeZone          *string    `json:"timeZone,omitempty"`
}

type GetSyncOptions200ApplicationJSON struct {
	AccountingSoftwareCompanyName *string                                        `json:"accountingSoftwareCompanyName,omitempty"`
	CompanyID                     *string                                        `json:"companyId,omitempty"`
	Configuration                 *GetSyncOptions200ApplicationJSONConfiguration `json:"configuration,omitempty"`
	Configured                    *bool                                          `json:"configured,omitempty"`
	Enabled                       *bool                                          `json:"enabled,omitempty"`
	Schedule                      *GetSyncOptions200ApplicationJSONSchedule      `json:"schedule,omitempty"`
}

type GetSyncOptionsResponse struct {
	ContentType                            string
	StatusCode                             int
	RawResponse                            *http.Response
	GetSyncOptions200ApplicationJSONObject *GetSyncOptions200ApplicationJSON
}
