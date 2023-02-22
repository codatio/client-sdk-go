package shared

type CodatErrorMessageValidationErrors struct {
	ItemID        *string `json:"itemId,omitempty"`
	Message       *string `json:"message,omitempty"`
	RuleID        *string `json:"ruleId,omitempty"`
	ValidatorName *string `json:"validatorName,omitempty"`
}

type CodatErrorMessageValidationInternals struct {
	ItemID        *string `json:"itemId,omitempty"`
	Message       *string `json:"message,omitempty"`
	RuleID        *string `json:"ruleId,omitempty"`
	ValidatorName *string `json:"validatorName,omitempty"`
}

type CodatErrorMessageValidationWarnings struct {
	ItemID        *string `json:"itemId,omitempty"`
	Message       *string `json:"message,omitempty"`
	RuleID        *string `json:"ruleId,omitempty"`
	ValidatorName *string `json:"validatorName,omitempty"`
}

type CodatErrorMessageValidation struct {
	Errors       []CodatErrorMessageValidationErrors    `json:"errors,omitempty"`
	HasErrors    *bool                                  `json:"hasErrors,omitempty"`
	HasInternals *bool                                  `json:"hasInternals,omitempty"`
	HasWarnings  *bool                                  `json:"hasWarnings,omitempty"`
	Internals    []CodatErrorMessageValidationInternals `json:"internals,omitempty"`
	Warnings     []CodatErrorMessageValidationWarnings  `json:"warnings,omitempty"`
}

type CodatErrorMessage struct {
	CanBeRetried      *string                      `json:"canBeRetried,omitempty"`
	CorrelationID     *string                      `json:"correlationId,omitempty"`
	DetailedErrorCode *int64                       `json:"detailedErrorCode,omitempty"`
	Error             *string                      `json:"error,omitempty"`
	Inner             *string                      `json:"inner,omitempty"`
	Service           *string                      `json:"service,omitempty"`
	StatusCode        *int64                       `json:"statusCode,omitempty"`
	Validation        *CodatErrorMessageValidation `json:"validation,omitempty"`
}
