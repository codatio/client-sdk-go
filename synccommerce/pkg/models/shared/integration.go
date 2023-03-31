// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// IntegrationDatatypeFeatureDataTypeEnum - Available Data types
type IntegrationDatatypeFeatureDataTypeEnum string

const (
	IntegrationDatatypeFeatureDataTypeEnumAccountTransactions          IntegrationDatatypeFeatureDataTypeEnum = "accountTransactions"
	IntegrationDatatypeFeatureDataTypeEnumBalanceSheet                 IntegrationDatatypeFeatureDataTypeEnum = "balanceSheet"
	IntegrationDatatypeFeatureDataTypeEnumBankAccounts                 IntegrationDatatypeFeatureDataTypeEnum = "bankAccounts"
	IntegrationDatatypeFeatureDataTypeEnumBankTransactions             IntegrationDatatypeFeatureDataTypeEnum = "bankTransactions"
	IntegrationDatatypeFeatureDataTypeEnumBillCreditNotes              IntegrationDatatypeFeatureDataTypeEnum = "billCreditNotes"
	IntegrationDatatypeFeatureDataTypeEnumBillPayments                 IntegrationDatatypeFeatureDataTypeEnum = "billPayments"
	IntegrationDatatypeFeatureDataTypeEnumBills                        IntegrationDatatypeFeatureDataTypeEnum = "bills"
	IntegrationDatatypeFeatureDataTypeEnumCashFlowStatement            IntegrationDatatypeFeatureDataTypeEnum = "cashFlowStatement"
	IntegrationDatatypeFeatureDataTypeEnumChartOfAccounts              IntegrationDatatypeFeatureDataTypeEnum = "chartOfAccounts"
	IntegrationDatatypeFeatureDataTypeEnumCompany                      IntegrationDatatypeFeatureDataTypeEnum = "company"
	IntegrationDatatypeFeatureDataTypeEnumCreditNotes                  IntegrationDatatypeFeatureDataTypeEnum = "creditNotes"
	IntegrationDatatypeFeatureDataTypeEnumCustomers                    IntegrationDatatypeFeatureDataTypeEnum = "customers"
	IntegrationDatatypeFeatureDataTypeEnumDirectCosts                  IntegrationDatatypeFeatureDataTypeEnum = "directCosts"
	IntegrationDatatypeFeatureDataTypeEnumDirectIncomes                IntegrationDatatypeFeatureDataTypeEnum = "directIncomes"
	IntegrationDatatypeFeatureDataTypeEnumInvoices                     IntegrationDatatypeFeatureDataTypeEnum = "invoices"
	IntegrationDatatypeFeatureDataTypeEnumItems                        IntegrationDatatypeFeatureDataTypeEnum = "items"
	IntegrationDatatypeFeatureDataTypeEnumJournalEntries               IntegrationDatatypeFeatureDataTypeEnum = "journalEntries"
	IntegrationDatatypeFeatureDataTypeEnumJournals                     IntegrationDatatypeFeatureDataTypeEnum = "journals"
	IntegrationDatatypeFeatureDataTypeEnumPaymentMethods               IntegrationDatatypeFeatureDataTypeEnum = "paymentMethods"
	IntegrationDatatypeFeatureDataTypeEnumPayments                     IntegrationDatatypeFeatureDataTypeEnum = "payments"
	IntegrationDatatypeFeatureDataTypeEnumProfitAndLoss                IntegrationDatatypeFeatureDataTypeEnum = "profitAndLoss"
	IntegrationDatatypeFeatureDataTypeEnumPurchaseOrders               IntegrationDatatypeFeatureDataTypeEnum = "purchaseOrders"
	IntegrationDatatypeFeatureDataTypeEnumSalesOrders                  IntegrationDatatypeFeatureDataTypeEnum = "salesOrders"
	IntegrationDatatypeFeatureDataTypeEnumSuppliers                    IntegrationDatatypeFeatureDataTypeEnum = "suppliers"
	IntegrationDatatypeFeatureDataTypeEnumTaxRates                     IntegrationDatatypeFeatureDataTypeEnum = "taxRates"
	IntegrationDatatypeFeatureDataTypeEnumTrackingCategories           IntegrationDatatypeFeatureDataTypeEnum = "trackingCategories"
	IntegrationDatatypeFeatureDataTypeEnumTransfers                    IntegrationDatatypeFeatureDataTypeEnum = "transfers"
	IntegrationDatatypeFeatureDataTypeEnumBankingAccountBalances       IntegrationDatatypeFeatureDataTypeEnum = "banking-accountBalances"
	IntegrationDatatypeFeatureDataTypeEnumBankingAccounts              IntegrationDatatypeFeatureDataTypeEnum = "banking-accounts"
	IntegrationDatatypeFeatureDataTypeEnumBankingTransactionCategories IntegrationDatatypeFeatureDataTypeEnum = "banking-transactionCategories"
	IntegrationDatatypeFeatureDataTypeEnumBankingTransactions          IntegrationDatatypeFeatureDataTypeEnum = "banking-transactions"
	IntegrationDatatypeFeatureDataTypeEnumCommerceCompanyInfo          IntegrationDatatypeFeatureDataTypeEnum = "commerce-companyInfo"
	IntegrationDatatypeFeatureDataTypeEnumCommerceCustomers            IntegrationDatatypeFeatureDataTypeEnum = "commerce-customers"
	IntegrationDatatypeFeatureDataTypeEnumCommerceDisputes             IntegrationDatatypeFeatureDataTypeEnum = "commerce-disputes"
	IntegrationDatatypeFeatureDataTypeEnumCommerceLocations            IntegrationDatatypeFeatureDataTypeEnum = "commerce-locations"
	IntegrationDatatypeFeatureDataTypeEnumCommerceOrders               IntegrationDatatypeFeatureDataTypeEnum = "commerce-orders"
	IntegrationDatatypeFeatureDataTypeEnumCommercePaymentMethods       IntegrationDatatypeFeatureDataTypeEnum = "commerce-paymentMethods"
	IntegrationDatatypeFeatureDataTypeEnumCommercePayments             IntegrationDatatypeFeatureDataTypeEnum = "commerce-payments"
	IntegrationDatatypeFeatureDataTypeEnumCommerceProductCategories    IntegrationDatatypeFeatureDataTypeEnum = "commerce-productCategories"
	IntegrationDatatypeFeatureDataTypeEnumCommerceProducts             IntegrationDatatypeFeatureDataTypeEnum = "commerce-products"
	IntegrationDatatypeFeatureDataTypeEnumCommerceTaxComponents        IntegrationDatatypeFeatureDataTypeEnum = "commerce-taxComponents"
	IntegrationDatatypeFeatureDataTypeEnumCommerceTransactions         IntegrationDatatypeFeatureDataTypeEnum = "commerce-transactions"
)

func (e *IntegrationDatatypeFeatureDataTypeEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "accountTransactions":
		fallthrough
	case "balanceSheet":
		fallthrough
	case "bankAccounts":
		fallthrough
	case "bankTransactions":
		fallthrough
	case "billCreditNotes":
		fallthrough
	case "billPayments":
		fallthrough
	case "bills":
		fallthrough
	case "cashFlowStatement":
		fallthrough
	case "chartOfAccounts":
		fallthrough
	case "company":
		fallthrough
	case "creditNotes":
		fallthrough
	case "customers":
		fallthrough
	case "directCosts":
		fallthrough
	case "directIncomes":
		fallthrough
	case "invoices":
		fallthrough
	case "items":
		fallthrough
	case "journalEntries":
		fallthrough
	case "journals":
		fallthrough
	case "paymentMethods":
		fallthrough
	case "payments":
		fallthrough
	case "profitAndLoss":
		fallthrough
	case "purchaseOrders":
		fallthrough
	case "salesOrders":
		fallthrough
	case "suppliers":
		fallthrough
	case "taxRates":
		fallthrough
	case "trackingCategories":
		fallthrough
	case "transfers":
		fallthrough
	case "banking-accountBalances":
		fallthrough
	case "banking-accounts":
		fallthrough
	case "banking-transactionCategories":
		fallthrough
	case "banking-transactions":
		fallthrough
	case "commerce-companyInfo":
		fallthrough
	case "commerce-customers":
		fallthrough
	case "commerce-disputes":
		fallthrough
	case "commerce-locations":
		fallthrough
	case "commerce-orders":
		fallthrough
	case "commerce-paymentMethods":
		fallthrough
	case "commerce-payments":
		fallthrough
	case "commerce-productCategories":
		fallthrough
	case "commerce-products":
		fallthrough
	case "commerce-taxComponents":
		fallthrough
	case "commerce-transactions":
		*e = IntegrationDatatypeFeatureDataTypeEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for IntegrationDatatypeFeatureDataTypeEnum: %s", s)
	}
}

type IntegrationDatatypeFeatureSupportedFeaturesFeatureStateEnum string

const (
	IntegrationDatatypeFeatureSupportedFeaturesFeatureStateEnumRelease        IntegrationDatatypeFeatureSupportedFeaturesFeatureStateEnum = "Release"
	IntegrationDatatypeFeatureSupportedFeaturesFeatureStateEnumBeta           IntegrationDatatypeFeatureSupportedFeaturesFeatureStateEnum = "Beta"
	IntegrationDatatypeFeatureSupportedFeaturesFeatureStateEnumDeprecated     IntegrationDatatypeFeatureSupportedFeaturesFeatureStateEnum = "Deprecated"
	IntegrationDatatypeFeatureSupportedFeaturesFeatureStateEnumNotSupported   IntegrationDatatypeFeatureSupportedFeaturesFeatureStateEnum = "NotSupported"
	IntegrationDatatypeFeatureSupportedFeaturesFeatureStateEnumNotImplemented IntegrationDatatypeFeatureSupportedFeaturesFeatureStateEnum = "NotImplemented"
)

func (e *IntegrationDatatypeFeatureSupportedFeaturesFeatureStateEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "Release":
		fallthrough
	case "Beta":
		fallthrough
	case "Deprecated":
		fallthrough
	case "NotSupported":
		fallthrough
	case "NotImplemented":
		*e = IntegrationDatatypeFeatureSupportedFeaturesFeatureStateEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for IntegrationDatatypeFeatureSupportedFeaturesFeatureStateEnum: %s", s)
	}
}

type IntegrationDatatypeFeatureSupportedFeaturesFeatureTypeEnum string

const (
	IntegrationDatatypeFeatureSupportedFeaturesFeatureTypeEnumGet                IntegrationDatatypeFeatureSupportedFeaturesFeatureTypeEnum = "Get"
	IntegrationDatatypeFeatureSupportedFeaturesFeatureTypeEnumPost               IntegrationDatatypeFeatureSupportedFeaturesFeatureTypeEnum = "Post"
	IntegrationDatatypeFeatureSupportedFeaturesFeatureTypeEnumCategorization     IntegrationDatatypeFeatureSupportedFeaturesFeatureTypeEnum = "Categorization"
	IntegrationDatatypeFeatureSupportedFeaturesFeatureTypeEnumDelete             IntegrationDatatypeFeatureSupportedFeaturesFeatureTypeEnum = "Delete"
	IntegrationDatatypeFeatureSupportedFeaturesFeatureTypeEnumPut                IntegrationDatatypeFeatureSupportedFeaturesFeatureTypeEnum = "Put"
	IntegrationDatatypeFeatureSupportedFeaturesFeatureTypeEnumGetAsPdf           IntegrationDatatypeFeatureSupportedFeaturesFeatureTypeEnum = "GetAsPdf"
	IntegrationDatatypeFeatureSupportedFeaturesFeatureTypeEnumDownloadAttachment IntegrationDatatypeFeatureSupportedFeaturesFeatureTypeEnum = "DownloadAttachment"
	IntegrationDatatypeFeatureSupportedFeaturesFeatureTypeEnumGetAttachment      IntegrationDatatypeFeatureSupportedFeaturesFeatureTypeEnum = "GetAttachment"
	IntegrationDatatypeFeatureSupportedFeaturesFeatureTypeEnumGetAttachments     IntegrationDatatypeFeatureSupportedFeaturesFeatureTypeEnum = "GetAttachments"
	IntegrationDatatypeFeatureSupportedFeaturesFeatureTypeEnumUploadAttachment   IntegrationDatatypeFeatureSupportedFeaturesFeatureTypeEnum = "UploadAttachment"
)

func (e *IntegrationDatatypeFeatureSupportedFeaturesFeatureTypeEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "Get":
		fallthrough
	case "Post":
		fallthrough
	case "Categorization":
		fallthrough
	case "Delete":
		fallthrough
	case "Put":
		fallthrough
	case "GetAsPdf":
		fallthrough
	case "DownloadAttachment":
		fallthrough
	case "GetAttachment":
		fallthrough
	case "GetAttachments":
		fallthrough
	case "UploadAttachment":
		*e = IntegrationDatatypeFeatureSupportedFeaturesFeatureTypeEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for IntegrationDatatypeFeatureSupportedFeaturesFeatureTypeEnum: %s", s)
	}
}

type IntegrationDatatypeFeatureSupportedFeatures struct {
	FeatureState IntegrationDatatypeFeatureSupportedFeaturesFeatureStateEnum `json:"featureState"`
	FeatureType  IntegrationDatatypeFeatureSupportedFeaturesFeatureTypeEnum  `json:"featureType"`
}

// IntegrationDatatypeFeature - Describes support for a given datatype and associated operations
type IntegrationDatatypeFeature struct {
	// Available Data types
	DataType          *IntegrationDatatypeFeatureDataTypeEnum       `json:"dataType,omitempty"`
	SupportedFeatures []IntegrationDatatypeFeatureSupportedFeatures `json:"supportedFeatures"`
}

// Integration - An integration that Codat supports
type Integration struct {
	DataProvidedBy   *string                      `json:"dataProvidedBy,omitempty"`
	DatatypeFeatures []IntegrationDatatypeFeature `json:"datatypeFeatures,omitempty"`
	// Whether this integration is enabled for your customers to use
	Enabled bool `json:"enabled"`
	// A Codat ID representing the integration.
	IntegrationID      *string `json:"integrationId,omitempty"`
	IsBeta             *bool   `json:"isBeta,omitempty"`
	IsOfflineConnector *bool   `json:"isOfflineConnector,omitempty"`
	// A unique 4-letter key to represent a platform in each integration. View [accounting](https://docs.codat.io/integrations/accounting/accounting-platform-keys), [banking](https://docs.codat.io/integrations/banking/banking-platform-keys), and [commerce](https://docs.codat.io/integrations/commerce/commerce-platform-keys) platform keys.
	Key     string `json:"key"`
	LogoURL string `json:"logoUrl"`
	Name    string `json:"name"`
	// A source-specific ID used to distinguish between different sources originating from the same data connection. In general, a data connection is a single data source. However, for TrueLayer, `sourceId` is associated with a specific bank and has a many-to-one relationship with the `integrationId`.
	SourceID *string `json:"sourceId,omitempty"`
	// The type of platform of the connection.
	SourceType *SourceTypeEnum `json:"sourceType,omitempty"`
}
