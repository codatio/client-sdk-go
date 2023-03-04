package operations

import (
	"net/http"
)

type CreateRuleWebhookNotifiers struct {
	Emails  []string `json:"emails,omitempty"`
	Webhook *string  `json:"webhook,omitempty"`
}

// CreateRuleWebhook
// Configuration to alert to a url or list of email addresses based on the given type / condition.
type CreateRuleWebhook struct {
	CompanyID *string                    `json:"companyId,omitempty"`
	ID        string                     `json:"id"`
	Notifiers CreateRuleWebhookNotifiers `json:"notifiers"`
	Type      string                     `json:"type"`
}

type CreateRuleRequest struct {
	Request *CreateRuleWebhook `request:"mediaType=application/json"`
}

type CreateRule401ApplicationJSON struct {
	CanBeRetried      *string `json:"canBeRetried,omitempty"`
	CorrelationID     *string `json:"correlationId,omitempty"`
	DetailedErrorCode *int64  `json:"detailedErrorCode,omitempty"`
	Error             *string `json:"error,omitempty"`
	Service           *string `json:"service,omitempty"`
	StatusCode        *int64  `json:"statusCode,omitempty"`
}

type CreateRuleResponse struct {
	ContentType                        string
	StatusCode                         int
	RawResponse                        *http.Response
	Webhook                            *CreateRuleWebhook
	CreateRule401ApplicationJSONObject *CreateRule401ApplicationJSON
}
