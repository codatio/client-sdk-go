// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ErrorMessage - Your `query` parameter was not correctly formed
type ErrorMessage struct {
	CanBeRetried      *string `json:"canBeRetried,omitempty"`
	CorrelationID     *string `json:"correlationId,omitempty"`
	DetailedErrorCode *int64  `json:"detailedErrorCode,omitempty"`
	Error             *string `json:"error,omitempty"`
	Service           *string `json:"service,omitempty"`
	StatusCode        *int64  `json:"statusCode,omitempty"`
}
