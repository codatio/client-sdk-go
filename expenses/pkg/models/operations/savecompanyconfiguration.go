package operations

type SaveCompanyConfigurationPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type SaveCompanyConfigurationApplicationWildcardPlusJSONBankAccount struct {
	ID *string `json:"id,omitempty"`
}

type SaveCompanyConfigurationApplicationWildcardPlusJSONCustomer struct {
	ID *string `json:"id,omitempty"`
}

type SaveCompanyConfigurationApplicationWildcardPlusJSONSupplier struct {
	ID *string `json:"id,omitempty"`
}

type SaveCompanyConfigurationApplicationWildcardPlusJSON struct {
	BankAccount *SaveCompanyConfigurationApplicationWildcardPlusJSONBankAccount `json:"bankAccount,omitempty"`
	Customer    *SaveCompanyConfigurationApplicationWildcardPlusJSONCustomer    `json:"customer,omitempty"`
	Supplier    *SaveCompanyConfigurationApplicationWildcardPlusJSONSupplier    `json:"supplier,omitempty"`
}

type SaveCompanyConfigurationApplicationJSONBankAccount struct {
	ID *string `json:"id,omitempty"`
}

type SaveCompanyConfigurationApplicationJSONCustomer struct {
	ID *string `json:"id,omitempty"`
}

type SaveCompanyConfigurationApplicationJSONSupplier struct {
	ID *string `json:"id,omitempty"`
}

type SaveCompanyConfigurationApplicationJSON struct {
	BankAccount *SaveCompanyConfigurationApplicationJSONBankAccount `json:"bankAccount,omitempty"`
	Customer    *SaveCompanyConfigurationApplicationJSONCustomer    `json:"customer,omitempty"`
	Supplier    *SaveCompanyConfigurationApplicationJSONSupplier    `json:"supplier,omitempty"`
}

type SaveCompanyConfigurationApplicationJSONPatchPlusJSONBankAccount struct {
	ID *string `json:"id,omitempty"`
}

type SaveCompanyConfigurationApplicationJSONPatchPlusJSONCustomer struct {
	ID *string `json:"id,omitempty"`
}

type SaveCompanyConfigurationApplicationJSONPatchPlusJSONSupplier struct {
	ID *string `json:"id,omitempty"`
}

type SaveCompanyConfigurationApplicationJSONPatchPlusJSON struct {
	BankAccount *SaveCompanyConfigurationApplicationJSONPatchPlusJSONBankAccount `json:"bankAccount,omitempty"`
	Customer    *SaveCompanyConfigurationApplicationJSONPatchPlusJSONCustomer    `json:"customer,omitempty"`
	Supplier    *SaveCompanyConfigurationApplicationJSONPatchPlusJSONSupplier    `json:"supplier,omitempty"`
}

type SaveCompanyConfigurationTextJSONBankAccount struct {
	ID *string `json:"id,omitempty"`
}

type SaveCompanyConfigurationTextJSONCustomer struct {
	ID *string `json:"id,omitempty"`
}

type SaveCompanyConfigurationTextJSONSupplier struct {
	ID *string `json:"id,omitempty"`
}

type SaveCompanyConfigurationTextJSON struct {
	BankAccount *SaveCompanyConfigurationTextJSONBankAccount `json:"bankAccount,omitempty"`
	Customer    *SaveCompanyConfigurationTextJSONCustomer    `json:"customer,omitempty"`
	Supplier    *SaveCompanyConfigurationTextJSONSupplier    `json:"supplier,omitempty"`
}

type SaveCompanyConfigurationRequests struct {
	Object  *SaveCompanyConfigurationApplicationWildcardPlusJSON  `request:"mediaType=application/*+json"`
	Object1 *SaveCompanyConfigurationApplicationJSON              `request:"mediaType=application/json"`
	Object2 *SaveCompanyConfigurationApplicationJSONPatchPlusJSON `request:"mediaType=application/json-patch+json"`
	Object3 *SaveCompanyConfigurationTextJSON                     `request:"mediaType=text/json"`
}

type SaveCompanyConfigurationRequest struct {
	PathParams SaveCompanyConfigurationPathParams
	Request    *SaveCompanyConfigurationRequests
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

type SaveCompanyConfiguration200TextJSONBankAccount struct {
	ID *string `json:"id,omitempty"`
}

type SaveCompanyConfiguration200TextJSONCustomer struct {
	ID *string `json:"id,omitempty"`
}

type SaveCompanyConfiguration200TextJSONSupplier struct {
	ID *string `json:"id,omitempty"`
}

type SaveCompanyConfiguration200TextJSON struct {
	BankAccount *SaveCompanyConfiguration200TextJSONBankAccount `json:"bankAccount,omitempty"`
	Customer    *SaveCompanyConfiguration200TextJSONCustomer    `json:"customer,omitempty"`
	Supplier    *SaveCompanyConfiguration200TextJSONSupplier    `json:"supplier,omitempty"`
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
	SaveCompanyConfiguration200ApplicationJSONObject *SaveCompanyConfiguration200ApplicationJSON
	SaveCompanyConfiguration200TextJSONObject        *SaveCompanyConfiguration200TextJSON
	SaveCompanyConfiguration200TextPlainObject       *string
	SaveCompanyConfiguration400ApplicationJSONObject *SaveCompanyConfiguration400ApplicationJSON
}
