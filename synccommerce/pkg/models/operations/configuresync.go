package operations

import (
	"net/http"
	"time"
)

type ConfigureSyncPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type ConfigureSyncRequestBodyConfigurationFeesAccountsAccountOptions struct {
	Classification *string `json:"classification,omitempty"`
	ID             *string `json:"id,omitempty"`
	Name           *string `json:"name,omitempty"`
	NominalCode    *string `json:"nominalCode,omitempty"`
}

type ConfigureSyncRequestBodyConfigurationFeesAccounts struct {
	AccountOptions    []ConfigureSyncRequestBodyConfigurationFeesAccountsAccountOptions `json:"accountOptions,omitempty"`
	DescriptionText   *string                                                           `json:"descriptionText,omitempty"`
	LabelText         *string                                                           `json:"labelText,omitempty"`
	Required          *bool                                                             `json:"required,omitempty"`
	SelectedAccountID *string                                                           `json:"selectedAccountId,omitempty"`
}

type ConfigureSyncRequestBodyConfigurationFeesFeesSupplierSupplierOptions struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type ConfigureSyncRequestBodyConfigurationFeesFeesSupplier struct {
	SelectedSupplierID *string                                                                `json:"selectedSupplierId,omitempty"`
	SupplierOptions    []ConfigureSyncRequestBodyConfigurationFeesFeesSupplierSupplierOptions `json:"supplierOptions,omitempty"`
}

type ConfigureSyncRequestBodyConfigurationFees struct {
	Accounts     map[string]ConfigureSyncRequestBodyConfigurationFeesAccounts `json:"accounts,omitempty"`
	FeesSupplier *ConfigureSyncRequestBodyConfigurationFeesFeesSupplier       `json:"feesSupplier,omitempty"`
	SyncFees     *bool                                                        `json:"syncFees,omitempty"`
}

type ConfigureSyncRequestBodyConfigurationNewPaymentsAccountsAccountOptions struct {
	Classification *string `json:"classification,omitempty"`
	ID             *string `json:"id,omitempty"`
	Name           *string `json:"name,omitempty"`
	NominalCode    *string `json:"nominalCode,omitempty"`
}

type ConfigureSyncRequestBodyConfigurationNewPaymentsAccounts struct {
	AccountOptions    []ConfigureSyncRequestBodyConfigurationNewPaymentsAccountsAccountOptions `json:"accountOptions,omitempty"`
	DescriptionText   *string                                                                  `json:"descriptionText,omitempty"`
	LabelText         *string                                                                  `json:"labelText,omitempty"`
	Required          *bool                                                                    `json:"required,omitempty"`
	SelectedAccountID *string                                                                  `json:"selectedAccountId,omitempty"`
}

type ConfigureSyncRequestBodyConfigurationNewPayments struct {
	Accounts     map[string]ConfigureSyncRequestBodyConfigurationNewPaymentsAccounts `json:"accounts,omitempty"`
	SyncPayments *bool                                                               `json:"syncPayments,omitempty"`
}

type ConfigureSyncRequestBodyConfigurationPaymentsAccountsAccountOptions struct {
	Classification *string `json:"classification,omitempty"`
	ID             *string `json:"id,omitempty"`
	Name           *string `json:"name,omitempty"`
	NominalCode    *string `json:"nominalCode,omitempty"`
}

type ConfigureSyncRequestBodyConfigurationPaymentsAccounts struct {
	AccountOptions    []ConfigureSyncRequestBodyConfigurationPaymentsAccountsAccountOptions `json:"accountOptions,omitempty"`
	DescriptionText   *string                                                               `json:"descriptionText,omitempty"`
	LabelText         *string                                                               `json:"labelText,omitempty"`
	Required          *bool                                                                 `json:"required,omitempty"`
	SelectedAccountID *string                                                               `json:"selectedAccountId,omitempty"`
}

type ConfigureSyncRequestBodyConfigurationPayments struct {
	Accounts     map[string]ConfigureSyncRequestBodyConfigurationPaymentsAccounts `json:"accounts,omitempty"`
	SyncPayments *bool                                                            `json:"syncPayments,omitempty"`
}

type ConfigureSyncRequestBodyConfigurationSalesAccountsAccountOptions struct {
	Classification *string `json:"classification,omitempty"`
	ID             *string `json:"id,omitempty"`
	Name           *string `json:"name,omitempty"`
	NominalCode    *string `json:"nominalCode,omitempty"`
}

type ConfigureSyncRequestBodyConfigurationSalesAccounts struct {
	AccountOptions    []ConfigureSyncRequestBodyConfigurationSalesAccountsAccountOptions `json:"accountOptions,omitempty"`
	DescriptionText   *string                                                            `json:"descriptionText,omitempty"`
	LabelText         *string                                                            `json:"labelText,omitempty"`
	Required          *bool                                                              `json:"required,omitempty"`
	SelectedAccountID *string                                                            `json:"selectedAccountId,omitempty"`
}

type ConfigureSyncRequestBodyConfigurationSalesGroupingGroupingLevelsInvoiceLevel struct {
	GroupByOptions         []string `json:"groupByOptions,omitempty"`
	SelectedGroupByOptions []string `json:"selectedGroupByOptions,omitempty"`
}

type ConfigureSyncRequestBodyConfigurationSalesGroupingGroupingLevelsInvoiceLineLevel struct {
	GroupByOptions         []string `json:"groupByOptions,omitempty"`
	SelectedGroupByOptions []string `json:"selectedGroupByOptions,omitempty"`
}

type ConfigureSyncRequestBodyConfigurationSalesGroupingGroupingLevels struct {
	InvoiceLevel     *ConfigureSyncRequestBodyConfigurationSalesGroupingGroupingLevelsInvoiceLevel     `json:"invoiceLevel,omitempty"`
	InvoiceLineLevel *ConfigureSyncRequestBodyConfigurationSalesGroupingGroupingLevelsInvoiceLineLevel `json:"invoiceLineLevel,omitempty"`
}

type ConfigureSyncRequestBodyConfigurationSalesGroupingGroupingPeriod struct {
	GroupingPeriodOptions  []string `json:"groupingPeriodOptions,omitempty"`
	SelectedGroupingPeriod *string  `json:"selectedGroupingPeriod,omitempty"`
}

type ConfigureSyncRequestBodyConfigurationSalesGrouping struct {
	GroupingLevels *ConfigureSyncRequestBodyConfigurationSalesGroupingGroupingLevels `json:"groupingLevels,omitempty"`
	GroupingPeriod *ConfigureSyncRequestBodyConfigurationSalesGroupingGroupingPeriod `json:"groupingPeriod,omitempty"`
}

type ConfigureSyncRequestBodyConfigurationSalesInvoiceStatus struct {
	InvoiceStatusOptions  []string `json:"invoiceStatusOptions,omitempty"`
	SelectedInvoiceStatus *string  `json:"selectedInvoiceStatus,omitempty"`
}

type ConfigureSyncRequestBodyConfigurationSalesNewTaxRatesAccountingTaxRateOptions struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type ConfigureSyncRequestBodyConfigurationSalesNewTaxRatesCommerceTaxRateOptions struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type ConfigureSyncRequestBodyConfigurationSalesNewTaxRatesDefaultZeroTaxRateOptions struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type ConfigureSyncRequestBodyConfigurationSalesNewTaxRatesTaxRateMappings struct {
	SelectedAccountingTaxRateID *string  `json:"selectedAccountingTaxRateId,omitempty"`
	SelectedCommerceTaxRateIds  []string `json:"selectedCommerceTaxRateIds,omitempty"`
}

type ConfigureSyncRequestBodyConfigurationSalesNewTaxRates struct {
	AccountingTaxRateOptions     []ConfigureSyncRequestBodyConfigurationSalesNewTaxRatesAccountingTaxRateOptions  `json:"accountingTaxRateOptions,omitempty"`
	CommerceTaxRateOptions       []ConfigureSyncRequestBodyConfigurationSalesNewTaxRatesCommerceTaxRateOptions    `json:"commerceTaxRateOptions,omitempty"`
	DefaultZeroTaxRateOptions    []ConfigureSyncRequestBodyConfigurationSalesNewTaxRatesDefaultZeroTaxRateOptions `json:"defaultZeroTaxRateOptions,omitempty"`
	SelectedDefaultZeroTaxRateID *string                                                                          `json:"selectedDefaultZeroTaxRateId,omitempty"`
	TaxRateMappings              []ConfigureSyncRequestBodyConfigurationSalesNewTaxRatesTaxRateMappings           `json:"taxRateMappings,omitempty"`
}

type ConfigureSyncRequestBodyConfigurationSalesSalesCustomerCustomerOptions struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type ConfigureSyncRequestBodyConfigurationSalesSalesCustomer struct {
	CustomerOptions    []ConfigureSyncRequestBodyConfigurationSalesSalesCustomerCustomerOptions `json:"customerOptions,omitempty"`
	SelectedCustomerID *string                                                                  `json:"selectedCustomerId,omitempty"`
}

type ConfigureSyncRequestBodyConfigurationSalesTaxRatesTaxRateOptions struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type ConfigureSyncRequestBodyConfigurationSalesTaxRates struct {
	SelectedTaxRateID *string                                                            `json:"selectedTaxRateId,omitempty"`
	TaxRateOptions    []ConfigureSyncRequestBodyConfigurationSalesTaxRatesTaxRateOptions `json:"taxRateOptions,omitempty"`
}

type ConfigureSyncRequestBodyConfigurationSales struct {
	Accounts      map[string]ConfigureSyncRequestBodyConfigurationSalesAccounts `json:"accounts,omitempty"`
	Grouping      *ConfigureSyncRequestBodyConfigurationSalesGrouping           `json:"grouping,omitempty"`
	InvoiceStatus *ConfigureSyncRequestBodyConfigurationSalesInvoiceStatus      `json:"invoiceStatus,omitempty"`
	NewTaxRates   *ConfigureSyncRequestBodyConfigurationSalesNewTaxRates        `json:"newTaxRates,omitempty"`
	SalesCustomer *ConfigureSyncRequestBodyConfigurationSalesSalesCustomer      `json:"salesCustomer,omitempty"`
	SyncSales     *bool                                                         `json:"syncSales,omitempty"`
	TaxRates      map[string]ConfigureSyncRequestBodyConfigurationSalesTaxRates `json:"taxRates,omitempty"`
}

type ConfigureSyncRequestBodyConfiguration struct {
	Fees        *ConfigureSyncRequestBodyConfigurationFees        `json:"fees,omitempty"`
	NewPayments *ConfigureSyncRequestBodyConfigurationNewPayments `json:"newPayments,omitempty"`
	Payments    *ConfigureSyncRequestBodyConfigurationPayments    `json:"payments,omitempty"`
	Sales       *ConfigureSyncRequestBodyConfigurationSales       `json:"sales,omitempty"`
}

type ConfigureSyncRequestBodySchedule struct {
	FrequencyOptions  []string   `json:"frequencyOptions,omitempty"`
	SelectedFrequency *string    `json:"selectedFrequency,omitempty"`
	StartDate         *time.Time `json:"startDate,omitempty"`
	SyncHourUtc       *int       `json:"syncHourUtc,omitempty"`
	TimeZone          *string    `json:"timeZone,omitempty"`
}

type ConfigureSyncRequestBody struct {
	AccountingSoftwareCompanyName *string                                `json:"accountingSoftwareCompanyName,omitempty"`
	CompanyID                     *string                                `json:"companyId,omitempty"`
	Configuration                 *ConfigureSyncRequestBodyConfiguration `json:"configuration,omitempty"`
	Configured                    *bool                                  `json:"configured,omitempty"`
	Enabled                       *bool                                  `json:"enabled,omitempty"`
	Schedule                      *ConfigureSyncRequestBodySchedule      `json:"schedule,omitempty"`
}

type ConfigureSyncRequest struct {
	PathParams ConfigureSyncPathParams
	Request    *ConfigureSyncRequestBody `request:"mediaType=application/json"`
}

type ConfigureSync200ApplicationJSONConfigurationFeesAccountsAccountOptions struct {
	Classification *string `json:"classification,omitempty"`
	ID             *string `json:"id,omitempty"`
	Name           *string `json:"name,omitempty"`
	NominalCode    *string `json:"nominalCode,omitempty"`
}

type ConfigureSync200ApplicationJSONConfigurationFeesAccounts struct {
	AccountOptions    []ConfigureSync200ApplicationJSONConfigurationFeesAccountsAccountOptions `json:"accountOptions,omitempty"`
	DescriptionText   *string                                                                  `json:"descriptionText,omitempty"`
	LabelText         *string                                                                  `json:"labelText,omitempty"`
	Required          *bool                                                                    `json:"required,omitempty"`
	SelectedAccountID *string                                                                  `json:"selectedAccountId,omitempty"`
}

type ConfigureSync200ApplicationJSONConfigurationFeesFeesSupplierSupplierOptions struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type ConfigureSync200ApplicationJSONConfigurationFeesFeesSupplier struct {
	SelectedSupplierID *string                                                                       `json:"selectedSupplierId,omitempty"`
	SupplierOptions    []ConfigureSync200ApplicationJSONConfigurationFeesFeesSupplierSupplierOptions `json:"supplierOptions,omitempty"`
}

type ConfigureSync200ApplicationJSONConfigurationFees struct {
	Accounts     map[string]ConfigureSync200ApplicationJSONConfigurationFeesAccounts `json:"accounts,omitempty"`
	FeesSupplier *ConfigureSync200ApplicationJSONConfigurationFeesFeesSupplier       `json:"feesSupplier,omitempty"`
	SyncFees     *bool                                                               `json:"syncFees,omitempty"`
}

type ConfigureSync200ApplicationJSONConfigurationNewPaymentsAccountsAccountOptions struct {
	Classification *string `json:"classification,omitempty"`
	ID             *string `json:"id,omitempty"`
	Name           *string `json:"name,omitempty"`
	NominalCode    *string `json:"nominalCode,omitempty"`
}

type ConfigureSync200ApplicationJSONConfigurationNewPaymentsAccounts struct {
	AccountOptions    []ConfigureSync200ApplicationJSONConfigurationNewPaymentsAccountsAccountOptions `json:"accountOptions,omitempty"`
	DescriptionText   *string                                                                         `json:"descriptionText,omitempty"`
	LabelText         *string                                                                         `json:"labelText,omitempty"`
	Required          *bool                                                                           `json:"required,omitempty"`
	SelectedAccountID *string                                                                         `json:"selectedAccountId,omitempty"`
}

type ConfigureSync200ApplicationJSONConfigurationNewPayments struct {
	Accounts     map[string]ConfigureSync200ApplicationJSONConfigurationNewPaymentsAccounts `json:"accounts,omitempty"`
	SyncPayments *bool                                                                      `json:"syncPayments,omitempty"`
}

type ConfigureSync200ApplicationJSONConfigurationPaymentsAccountsAccountOptions struct {
	Classification *string `json:"classification,omitempty"`
	ID             *string `json:"id,omitempty"`
	Name           *string `json:"name,omitempty"`
	NominalCode    *string `json:"nominalCode,omitempty"`
}

type ConfigureSync200ApplicationJSONConfigurationPaymentsAccounts struct {
	AccountOptions    []ConfigureSync200ApplicationJSONConfigurationPaymentsAccountsAccountOptions `json:"accountOptions,omitempty"`
	DescriptionText   *string                                                                      `json:"descriptionText,omitempty"`
	LabelText         *string                                                                      `json:"labelText,omitempty"`
	Required          *bool                                                                        `json:"required,omitempty"`
	SelectedAccountID *string                                                                      `json:"selectedAccountId,omitempty"`
}

type ConfigureSync200ApplicationJSONConfigurationPayments struct {
	Accounts     map[string]ConfigureSync200ApplicationJSONConfigurationPaymentsAccounts `json:"accounts,omitempty"`
	SyncPayments *bool                                                                   `json:"syncPayments,omitempty"`
}

type ConfigureSync200ApplicationJSONConfigurationSalesAccountsAccountOptions struct {
	Classification *string `json:"classification,omitempty"`
	ID             *string `json:"id,omitempty"`
	Name           *string `json:"name,omitempty"`
	NominalCode    *string `json:"nominalCode,omitempty"`
}

type ConfigureSync200ApplicationJSONConfigurationSalesAccounts struct {
	AccountOptions    []ConfigureSync200ApplicationJSONConfigurationSalesAccountsAccountOptions `json:"accountOptions,omitempty"`
	DescriptionText   *string                                                                   `json:"descriptionText,omitempty"`
	LabelText         *string                                                                   `json:"labelText,omitempty"`
	Required          *bool                                                                     `json:"required,omitempty"`
	SelectedAccountID *string                                                                   `json:"selectedAccountId,omitempty"`
}

type ConfigureSync200ApplicationJSONConfigurationSalesGroupingGroupingLevelsInvoiceLevel struct {
	GroupByOptions         []string `json:"groupByOptions,omitempty"`
	SelectedGroupByOptions []string `json:"selectedGroupByOptions,omitempty"`
}

type ConfigureSync200ApplicationJSONConfigurationSalesGroupingGroupingLevelsInvoiceLineLevel struct {
	GroupByOptions         []string `json:"groupByOptions,omitempty"`
	SelectedGroupByOptions []string `json:"selectedGroupByOptions,omitempty"`
}

type ConfigureSync200ApplicationJSONConfigurationSalesGroupingGroupingLevels struct {
	InvoiceLevel     *ConfigureSync200ApplicationJSONConfigurationSalesGroupingGroupingLevelsInvoiceLevel     `json:"invoiceLevel,omitempty"`
	InvoiceLineLevel *ConfigureSync200ApplicationJSONConfigurationSalesGroupingGroupingLevelsInvoiceLineLevel `json:"invoiceLineLevel,omitempty"`
}

type ConfigureSync200ApplicationJSONConfigurationSalesGroupingGroupingPeriod struct {
	GroupingPeriodOptions  []string `json:"groupingPeriodOptions,omitempty"`
	SelectedGroupingPeriod *string  `json:"selectedGroupingPeriod,omitempty"`
}

type ConfigureSync200ApplicationJSONConfigurationSalesGrouping struct {
	GroupingLevels *ConfigureSync200ApplicationJSONConfigurationSalesGroupingGroupingLevels `json:"groupingLevels,omitempty"`
	GroupingPeriod *ConfigureSync200ApplicationJSONConfigurationSalesGroupingGroupingPeriod `json:"groupingPeriod,omitempty"`
}

type ConfigureSync200ApplicationJSONConfigurationSalesInvoiceStatus struct {
	InvoiceStatusOptions  []string `json:"invoiceStatusOptions,omitempty"`
	SelectedInvoiceStatus *string  `json:"selectedInvoiceStatus,omitempty"`
}

type ConfigureSync200ApplicationJSONConfigurationSalesNewTaxRatesAccountingTaxRateOptions struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type ConfigureSync200ApplicationJSONConfigurationSalesNewTaxRatesCommerceTaxRateOptions struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type ConfigureSync200ApplicationJSONConfigurationSalesNewTaxRatesDefaultZeroTaxRateOptions struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type ConfigureSync200ApplicationJSONConfigurationSalesNewTaxRatesTaxRateMappings struct {
	SelectedAccountingTaxRateID *string  `json:"selectedAccountingTaxRateId,omitempty"`
	SelectedCommerceTaxRateIds  []string `json:"selectedCommerceTaxRateIds,omitempty"`
}

type ConfigureSync200ApplicationJSONConfigurationSalesNewTaxRates struct {
	AccountingTaxRateOptions     []ConfigureSync200ApplicationJSONConfigurationSalesNewTaxRatesAccountingTaxRateOptions  `json:"accountingTaxRateOptions,omitempty"`
	CommerceTaxRateOptions       []ConfigureSync200ApplicationJSONConfigurationSalesNewTaxRatesCommerceTaxRateOptions    `json:"commerceTaxRateOptions,omitempty"`
	DefaultZeroTaxRateOptions    []ConfigureSync200ApplicationJSONConfigurationSalesNewTaxRatesDefaultZeroTaxRateOptions `json:"defaultZeroTaxRateOptions,omitempty"`
	SelectedDefaultZeroTaxRateID *string                                                                                 `json:"selectedDefaultZeroTaxRateId,omitempty"`
	TaxRateMappings              []ConfigureSync200ApplicationJSONConfigurationSalesNewTaxRatesTaxRateMappings           `json:"taxRateMappings,omitempty"`
}

type ConfigureSync200ApplicationJSONConfigurationSalesSalesCustomerCustomerOptions struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type ConfigureSync200ApplicationJSONConfigurationSalesSalesCustomer struct {
	CustomerOptions    []ConfigureSync200ApplicationJSONConfigurationSalesSalesCustomerCustomerOptions `json:"customerOptions,omitempty"`
	SelectedCustomerID *string                                                                         `json:"selectedCustomerId,omitempty"`
}

type ConfigureSync200ApplicationJSONConfigurationSalesTaxRatesTaxRateOptions struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type ConfigureSync200ApplicationJSONConfigurationSalesTaxRates struct {
	SelectedTaxRateID *string                                                                   `json:"selectedTaxRateId,omitempty"`
	TaxRateOptions    []ConfigureSync200ApplicationJSONConfigurationSalesTaxRatesTaxRateOptions `json:"taxRateOptions,omitempty"`
}

type ConfigureSync200ApplicationJSONConfigurationSales struct {
	Accounts      map[string]ConfigureSync200ApplicationJSONConfigurationSalesAccounts `json:"accounts,omitempty"`
	Grouping      *ConfigureSync200ApplicationJSONConfigurationSalesGrouping           `json:"grouping,omitempty"`
	InvoiceStatus *ConfigureSync200ApplicationJSONConfigurationSalesInvoiceStatus      `json:"invoiceStatus,omitempty"`
	NewTaxRates   *ConfigureSync200ApplicationJSONConfigurationSalesNewTaxRates        `json:"newTaxRates,omitempty"`
	SalesCustomer *ConfigureSync200ApplicationJSONConfigurationSalesSalesCustomer      `json:"salesCustomer,omitempty"`
	SyncSales     *bool                                                                `json:"syncSales,omitempty"`
	TaxRates      map[string]ConfigureSync200ApplicationJSONConfigurationSalesTaxRates `json:"taxRates,omitempty"`
}

type ConfigureSync200ApplicationJSONConfiguration struct {
	Fees        *ConfigureSync200ApplicationJSONConfigurationFees        `json:"fees,omitempty"`
	NewPayments *ConfigureSync200ApplicationJSONConfigurationNewPayments `json:"newPayments,omitempty"`
	Payments    *ConfigureSync200ApplicationJSONConfigurationPayments    `json:"payments,omitempty"`
	Sales       *ConfigureSync200ApplicationJSONConfigurationSales       `json:"sales,omitempty"`
}

type ConfigureSync200ApplicationJSONSchedule struct {
	FrequencyOptions  []string   `json:"frequencyOptions,omitempty"`
	SelectedFrequency *string    `json:"selectedFrequency,omitempty"`
	StartDate         *time.Time `json:"startDate,omitempty"`
	SyncHourUtc       *int       `json:"syncHourUtc,omitempty"`
	TimeZone          *string    `json:"timeZone,omitempty"`
}

type ConfigureSync200ApplicationJSON struct {
	AccountingSoftwareCompanyName *string                                       `json:"accountingSoftwareCompanyName,omitempty"`
	CompanyID                     *string                                       `json:"companyId,omitempty"`
	Configuration                 *ConfigureSync200ApplicationJSONConfiguration `json:"configuration,omitempty"`
	Configured                    *bool                                         `json:"configured,omitempty"`
	Enabled                       *bool                                         `json:"enabled,omitempty"`
	Schedule                      *ConfigureSync200ApplicationJSONSchedule      `json:"schedule,omitempty"`
}

type ConfigureSyncResponse struct {
	ContentType                           string
	StatusCode                            int
	RawResponse                           *http.Response
	ConfigureSync200ApplicationJSONObject *ConfigureSync200ApplicationJSON
}
