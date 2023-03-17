package operations

import (
	"net/http"
)

type IntiateSyncRequestBody struct {
	DatasetIds []string `json:"datasetIds,omitempty"`
}

type IntiateSyncRequest struct {
	RequestBody *IntiateSyncRequestBody `request:"mediaType=application/json"`
	CompanyID   string                  `pathParam:"style=simple,explode=false,name=companyId"`
}

type IntiateSync422ApplicationJSONValidationErrors struct {
	ItemID        *string `json:"itemId,omitempty"`
	Message       *string `json:"message,omitempty"`
	RuleID        *string `json:"ruleId,omitempty"`
	ValidatorName *string `json:"validatorName,omitempty"`
}

type IntiateSync422ApplicationJSONValidationInternals struct {
	ItemID        *string `json:"itemId,omitempty"`
	Message       *string `json:"message,omitempty"`
	RuleID        *string `json:"ruleId,omitempty"`
	ValidatorName *string `json:"validatorName,omitempty"`
}

type IntiateSync422ApplicationJSONValidationWarnings struct {
	ItemID        *string `json:"itemId,omitempty"`
	Message       *string `json:"message,omitempty"`
	RuleID        *string `json:"ruleId,omitempty"`
	ValidatorName *string `json:"validatorName,omitempty"`
}

type IntiateSync422ApplicationJSONValidation struct {
	Errors       []IntiateSync422ApplicationJSONValidationErrors    `json:"errors,omitempty"`
	HasErrors    *bool                                              `json:"hasErrors,omitempty"`
	HasInternals *bool                                              `json:"hasInternals,omitempty"`
	HasWarnings  *bool                                              `json:"hasWarnings,omitempty"`
	Internals    []IntiateSync422ApplicationJSONValidationInternals `json:"internals,omitempty"`
	Warnings     []IntiateSync422ApplicationJSONValidationWarnings  `json:"warnings,omitempty"`
}

type IntiateSync422ApplicationJSON struct {
	CanBeRetried      *string                                  `json:"canBeRetried,omitempty"`
	CorrelationID     *string                                  `json:"correlationId,omitempty"`
	DetailedErrorCode *int64                                   `json:"detailedErrorCode,omitempty"`
	Error             *string                                  `json:"error,omitempty"`
	Inner             *string                                  `json:"inner,omitempty"`
	Service           *string                                  `json:"service,omitempty"`
	StatusCode        *int64                                   `json:"statusCode,omitempty"`
	Validation        *IntiateSync422ApplicationJSONValidation `json:"validation,omitempty"`
}

type IntiateSync404ApplicationJSONValidationErrors struct {
	ItemID        *string `json:"itemId,omitempty"`
	Message       *string `json:"message,omitempty"`
	RuleID        *string `json:"ruleId,omitempty"`
	ValidatorName *string `json:"validatorName,omitempty"`
}

type IntiateSync404ApplicationJSONValidationInternals struct {
	ItemID        *string `json:"itemId,omitempty"`
	Message       *string `json:"message,omitempty"`
	RuleID        *string `json:"ruleId,omitempty"`
	ValidatorName *string `json:"validatorName,omitempty"`
}

type IntiateSync404ApplicationJSONValidationWarnings struct {
	ItemID        *string `json:"itemId,omitempty"`
	Message       *string `json:"message,omitempty"`
	RuleID        *string `json:"ruleId,omitempty"`
	ValidatorName *string `json:"validatorName,omitempty"`
}

type IntiateSync404ApplicationJSONValidation struct {
	Errors       []IntiateSync404ApplicationJSONValidationErrors    `json:"errors,omitempty"`
	HasErrors    *bool                                              `json:"hasErrors,omitempty"`
	HasInternals *bool                                              `json:"hasInternals,omitempty"`
	HasWarnings  *bool                                              `json:"hasWarnings,omitempty"`
	Internals    []IntiateSync404ApplicationJSONValidationInternals `json:"internals,omitempty"`
	Warnings     []IntiateSync404ApplicationJSONValidationWarnings  `json:"warnings,omitempty"`
}

type IntiateSync404ApplicationJSON struct {
	CanBeRetried      *string                                  `json:"canBeRetried,omitempty"`
	CorrelationID     *string                                  `json:"correlationId,omitempty"`
	DetailedErrorCode *int64                                   `json:"detailedErrorCode,omitempty"`
	Error             *string                                  `json:"error,omitempty"`
	Inner             *string                                  `json:"inner,omitempty"`
	Service           *string                                  `json:"service,omitempty"`
	StatusCode        *int64                                   `json:"statusCode,omitempty"`
	Validation        *IntiateSync404ApplicationJSONValidation `json:"validation,omitempty"`
}

type IntiateSync400ApplicationJSONValidationErrors struct {
	ItemID        *string `json:"itemId,omitempty"`
	Message       *string `json:"message,omitempty"`
	RuleID        *string `json:"ruleId,omitempty"`
	ValidatorName *string `json:"validatorName,omitempty"`
}

type IntiateSync400ApplicationJSONValidationInternals struct {
	ItemID        *string `json:"itemId,omitempty"`
	Message       *string `json:"message,omitempty"`
	RuleID        *string `json:"ruleId,omitempty"`
	ValidatorName *string `json:"validatorName,omitempty"`
}

type IntiateSync400ApplicationJSONValidationWarnings struct {
	ItemID        *string `json:"itemId,omitempty"`
	Message       *string `json:"message,omitempty"`
	RuleID        *string `json:"ruleId,omitempty"`
	ValidatorName *string `json:"validatorName,omitempty"`
}

type IntiateSync400ApplicationJSONValidation struct {
	Errors       []IntiateSync400ApplicationJSONValidationErrors    `json:"errors,omitempty"`
	HasErrors    *bool                                              `json:"hasErrors,omitempty"`
	HasInternals *bool                                              `json:"hasInternals,omitempty"`
	HasWarnings  *bool                                              `json:"hasWarnings,omitempty"`
	Internals    []IntiateSync400ApplicationJSONValidationInternals `json:"internals,omitempty"`
	Warnings     []IntiateSync400ApplicationJSONValidationWarnings  `json:"warnings,omitempty"`
}

type IntiateSync400ApplicationJSON struct {
	CanBeRetried      *string                                  `json:"canBeRetried,omitempty"`
	CorrelationID     *string                                  `json:"correlationId,omitempty"`
	DetailedErrorCode *int64                                   `json:"detailedErrorCode,omitempty"`
	Error             *string                                  `json:"error,omitempty"`
	Inner             *string                                  `json:"inner,omitempty"`
	Service           *string                                  `json:"service,omitempty"`
	StatusCode        *int64                                   `json:"statusCode,omitempty"`
	Validation        *IntiateSync400ApplicationJSONValidation `json:"validation,omitempty"`
}

type IntiateSync202ApplicationJSON struct {
	SyncID *string `json:"syncId,omitempty"`
}

type IntiateSyncResponse struct {
	ContentType                         string
	StatusCode                          int
	RawResponse                         *http.Response
	IntiateSync202ApplicationJSONObject *IntiateSync202ApplicationJSON
	IntiateSync400ApplicationJSONObject *IntiateSync400ApplicationJSON
	IntiateSync404ApplicationJSONObject *IntiateSync404ApplicationJSON
	IntiateSync422ApplicationJSONObject *IntiateSync422ApplicationJSON
}
