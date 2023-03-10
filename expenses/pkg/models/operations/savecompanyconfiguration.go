package operations

import (
	"net/http"
)

type SaveCompanyConfigurationPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type SaveCompanyConfigurationRequestBodyBankAccount struct {
	ID *string `json:"id,omitempty"`
}

type SaveCompanyConfigurationRequestBodyCustomer struct {
	ID *string `json:"id,omitempty"`
}

type SaveCompanyConfigurationRequestBodySupplier struct {
	ID *string `json:"id,omitempty"`
}

type SaveCompanyConfigurationRequestBody struct {
	BankAccount *SaveCompanyConfigurationRequestBodyBankAccount `json:"bankAccount,omitempty"`
	Customer    *SaveCompanyConfigurationRequestBodyCustomer    `json:"customer,omitempty"`
	Supplier    *SaveCompanyConfigurationRequestBodySupplier    `json:"supplier,omitempty"`
}

type SaveCompanyConfigurationRequest struct {
	PathParams SaveCompanyConfigurationPathParams
	Request    *SaveCompanyConfigurationRequestBody `request:"mediaType=application/json"`
}

type SaveCompanyConfiguration400ApplicationJSONValidationErrors struct {
	ItemID        *string `json:"itemId,omitempty"`
	Message       *string `json:"message,omitempty"`
	RuleID        *string `json:"ruleId,omitempty"`
	ValidatorName *string `json:"validatorName,omitempty"`
}

type SaveCompanyConfiguration400ApplicationJSONValidationInternals struct {
	ItemID        *string `json:"itemId,omitempty"`
	Message       *string `json:"message,omitempty"`
	RuleID        *string `json:"ruleId,omitempty"`
	ValidatorName *string `json:"validatorName,omitempty"`
}

type SaveCompanyConfiguration400ApplicationJSONValidationWarnings struct {
	ItemID        *string `json:"itemId,omitempty"`
	Message       *string `json:"message,omitempty"`
	RuleID        *string `json:"ruleId,omitempty"`
	ValidatorName *string `json:"validatorName,omitempty"`
}

type SaveCompanyConfiguration400ApplicationJSONValidation struct {
	Errors       []SaveCompanyConfiguration400ApplicationJSONValidationErrors    `json:"errors,omitempty"`
	HasErrors    *bool                                                           `json:"hasErrors,omitempty"`
	HasInternals *bool                                                           `json:"hasInternals,omitempty"`
	HasWarnings  *bool                                                           `json:"hasWarnings,omitempty"`
	Internals    []SaveCompanyConfiguration400ApplicationJSONValidationInternals `json:"internals,omitempty"`
	Warnings     []SaveCompanyConfiguration400ApplicationJSONValidationWarnings  `json:"warnings,omitempty"`
}

type SaveCompanyConfiguration400ApplicationJSON struct {
	CanBeRetried      *string                                               `json:"canBeRetried,omitempty"`
	CorrelationID     *string                                               `json:"correlationId,omitempty"`
	DetailedErrorCode *int64                                                `json:"detailedErrorCode,omitempty"`
	Error             *string                                               `json:"error,omitempty"`
	Inner             *string                                               `json:"inner,omitempty"`
	Service           *string                                               `json:"service,omitempty"`
	StatusCode        *int64                                                `json:"statusCode,omitempty"`
	Validation        *SaveCompanyConfiguration400ApplicationJSONValidation `json:"validation,omitempty"`
}

type SaveCompanyConfiguration200ApplicationJSONBankAccount struct {
	ID *string `json:"id,omitempty"`
}

type SaveCompanyConfiguration200ApplicationJSONCustomer struct {
	ID *string `json:"id,omitempty"`
}

type SaveCompanyConfiguration200ApplicationJSONSupplier struct {
	ID *string `json:"id,omitempty"`
}

type SaveCompanyConfiguration200ApplicationJSON struct {
	BankAccount *SaveCompanyConfiguration200ApplicationJSONBankAccount `json:"bankAccount,omitempty"`
	Customer    *SaveCompanyConfiguration200ApplicationJSONCustomer    `json:"customer,omitempty"`
	Supplier    *SaveCompanyConfiguration200ApplicationJSONSupplier    `json:"supplier,omitempty"`
}

type SaveCompanyConfigurationResponse struct {
	ContentType                                      string
	StatusCode                                       int
	RawResponse                                      *http.Response
	SaveCompanyConfiguration200ApplicationJSONObject *SaveCompanyConfiguration200ApplicationJSON
	SaveCompanyConfiguration400ApplicationJSONObject *SaveCompanyConfiguration400ApplicationJSON
}
