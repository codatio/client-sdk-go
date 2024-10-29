// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/codatio/client-sdk-go/platform/v4/pkg/models/shared"
	"net/http"
)

// DataType - Data types that support supplemental data
type DataType string

const (
	DataTypeChartOfAccounts           DataType = "chartOfAccounts"
	DataTypeBills                     DataType = "bills"
	DataTypeCompany                   DataType = "company"
	DataTypeCreditNotes               DataType = "creditNotes"
	DataTypeCustomers                 DataType = "customers"
	DataTypeInvoices                  DataType = "invoices"
	DataTypeItems                     DataType = "items"
	DataTypeJournalEntries            DataType = "journalEntries"
	DataTypeSuppliers                 DataType = "suppliers"
	DataTypeTaxRates                  DataType = "taxRates"
	DataTypeCommerceCompanyInfo       DataType = "commerce-companyInfo"
	DataTypeCommerceCustomers         DataType = "commerce-customers"
	DataTypeCommerceDisputes          DataType = "commerce-disputes"
	DataTypeCommerceLocations         DataType = "commerce-locations"
	DataTypeCommerceOrders            DataType = "commerce-orders"
	DataTypeCommercePayments          DataType = "commerce-payments"
	DataTypeCommercePaymentMethods    DataType = "commerce-paymentMethods"
	DataTypeCommerceProducts          DataType = "commerce-products"
	DataTypeCommerceProductCategories DataType = "commerce-productCategories"
	DataTypeCommerceTaxComponents     DataType = "commerce-taxComponents"
	DataTypeCommerceTransactions      DataType = "commerce-transactions"
)

func (e DataType) ToPointer() *DataType {
	return &e
}
func (e *DataType) UnmarshalJSON(data []byte) error {
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
		*e = DataType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DataType: %v", v)
	}
}

type ConfigureSupplementalDataRequest struct {
	// The configuration for the specified platform and data type.
	SupplementalDataConfiguration *shared.SupplementalDataConfiguration `request:"mediaType=application/json"`
	// Supported supplemental data data type.
	DataType DataType `pathParam:"style=simple,explode=false,name=dataType"`
	// A unique 4-letter key to represent a platform in each integration.
	PlatformKey string `pathParam:"style=simple,explode=false,name=platformKey"`
}

func (o *ConfigureSupplementalDataRequest) GetSupplementalDataConfiguration() *shared.SupplementalDataConfiguration {
	if o == nil {
		return nil
	}
	return o.SupplementalDataConfiguration
}

func (o *ConfigureSupplementalDataRequest) GetDataType() DataType {
	if o == nil {
		return DataType("")
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
