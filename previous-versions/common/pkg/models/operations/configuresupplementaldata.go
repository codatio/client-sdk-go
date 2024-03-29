// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/codatio/client-sdk-go/previous-versions/common/pkg/models/shared"
	"net/http"
)

// ConfigureSupplementalDataDataType - Data types that support supplemental data
type ConfigureSupplementalDataDataType string

const (
	ConfigureSupplementalDataDataTypeChartOfAccounts           ConfigureSupplementalDataDataType = "chartOfAccounts"
	ConfigureSupplementalDataDataTypeBills                     ConfigureSupplementalDataDataType = "bills"
	ConfigureSupplementalDataDataTypeCompany                   ConfigureSupplementalDataDataType = "company"
	ConfigureSupplementalDataDataTypeCreditNotes               ConfigureSupplementalDataDataType = "creditNotes"
	ConfigureSupplementalDataDataTypeCustomers                 ConfigureSupplementalDataDataType = "customers"
	ConfigureSupplementalDataDataTypeInvoices                  ConfigureSupplementalDataDataType = "invoices"
	ConfigureSupplementalDataDataTypeItems                     ConfigureSupplementalDataDataType = "items"
	ConfigureSupplementalDataDataTypeJournalEntries            ConfigureSupplementalDataDataType = "journalEntries"
	ConfigureSupplementalDataDataTypeSuppliers                 ConfigureSupplementalDataDataType = "suppliers"
	ConfigureSupplementalDataDataTypeTaxRates                  ConfigureSupplementalDataDataType = "taxRates"
	ConfigureSupplementalDataDataTypeCommerceCompanyInfo       ConfigureSupplementalDataDataType = "commerce-companyInfo"
	ConfigureSupplementalDataDataTypeCommerceCustomers         ConfigureSupplementalDataDataType = "commerce-customers"
	ConfigureSupplementalDataDataTypeCommerceDisputes          ConfigureSupplementalDataDataType = "commerce-disputes"
	ConfigureSupplementalDataDataTypeCommerceLocations         ConfigureSupplementalDataDataType = "commerce-locations"
	ConfigureSupplementalDataDataTypeCommerceOrders            ConfigureSupplementalDataDataType = "commerce-orders"
	ConfigureSupplementalDataDataTypeCommercePayments          ConfigureSupplementalDataDataType = "commerce-payments"
	ConfigureSupplementalDataDataTypeCommercePaymentMethods    ConfigureSupplementalDataDataType = "commerce-paymentMethods"
	ConfigureSupplementalDataDataTypeCommerceProducts          ConfigureSupplementalDataDataType = "commerce-products"
	ConfigureSupplementalDataDataTypeCommerceProductCategories ConfigureSupplementalDataDataType = "commerce-productCategories"
	ConfigureSupplementalDataDataTypeCommerceTaxComponents     ConfigureSupplementalDataDataType = "commerce-taxComponents"
	ConfigureSupplementalDataDataTypeCommerceTransactions      ConfigureSupplementalDataDataType = "commerce-transactions"
)

func (e ConfigureSupplementalDataDataType) ToPointer() *ConfigureSupplementalDataDataType {
	return &e
}

func (e *ConfigureSupplementalDataDataType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "chartOfAccounts":
		fallthrough
	case "bills":
		fallthrough
	case "company":
		fallthrough
	case "creditNotes":
		fallthrough
	case "customers":
		fallthrough
	case "invoices":
		fallthrough
	case "items":
		fallthrough
	case "journalEntries":
		fallthrough
	case "suppliers":
		fallthrough
	case "taxRates":
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
	case "commerce-payments":
		fallthrough
	case "commerce-paymentMethods":
		fallthrough
	case "commerce-products":
		fallthrough
	case "commerce-productCategories":
		fallthrough
	case "commerce-taxComponents":
		fallthrough
	case "commerce-transactions":
		*e = ConfigureSupplementalDataDataType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConfigureSupplementalDataDataType: %v", v)
	}
}

type ConfigureSupplementalDataRequest struct {
	// The configuration for the specified platform and data type.
	SupplementalDataConfiguration *shared.SupplementalDataConfiguration `request:"mediaType=application/json"`
	// Supported supplemental data data type.
	DataType ConfigureSupplementalDataDataType `pathParam:"style=simple,explode=false,name=dataType"`
	// A unique 4-letter key to represent a platform in each integration. View [accounting](https://docs.codat.io/integrations/accounting/overview#platform-keys), [banking](https://docs.codat.io/integrations/banking/overview#platform-keys), and [commerce](https://docs.codat.io/integrations/commerce/overview#platform-keys) platform keys.
	PlatformKey string `pathParam:"style=simple,explode=false,name=platformKey"`
}

func (o *ConfigureSupplementalDataRequest) GetSupplementalDataConfiguration() *shared.SupplementalDataConfiguration {
	if o == nil {
		return nil
	}
	return o.SupplementalDataConfiguration
}

func (o *ConfigureSupplementalDataRequest) GetDataType() ConfigureSupplementalDataDataType {
	if o == nil {
		return ConfigureSupplementalDataDataType("")
	}
	return o.DataType
}

func (o *ConfigureSupplementalDataRequest) GetPlatformKey() string {
	if o == nil {
		return ""
	}
	return o.PlatformKey
}

type ConfigureSupplementalDataResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Your API request was not properly authorized.
	ErrorMessage *shared.ErrorMessage
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ConfigureSupplementalDataResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ConfigureSupplementalDataResponse) GetErrorMessage() *shared.ErrorMessage {
	if o == nil {
		return nil
	}
	return o.ErrorMessage
}

func (o *ConfigureSupplementalDataResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ConfigureSupplementalDataResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
