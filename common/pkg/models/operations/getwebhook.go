package operations

type GetWebhookPathParams struct {
	RuleID string `pathParam:"style=simple,explode=false,name=ruleId"`
}

type GetWebhookRequest struct {
	PathParams GetWebhookPathParams
}

type GetWebhook404ApplicationJSON struct {
	CanBeRetried      *string `json:"canBeRetried,omitempty"`
	CorrelationID     *string `json:"correlationId,omitempty"`
	DetailedErrorCode *int64  `json:"detailedErrorCode,omitempty"`
	Error             *string `json:"error,omitempty"`
	Service           *string `json:"service,omitempty"`
	StatusCode        *int64  `json:"statusCode,omitempty"`
}

type GetWebhook401ApplicationJSON struct {
	CanBeRetried      *string `json:"canBeRetried,omitempty"`
	CorrelationID     *string `json:"correlationId,omitempty"`
	DetailedErrorCode *int64  `json:"detailedErrorCode,omitempty"`
	Error             *string `json:"error,omitempty"`
	Service           *string `json:"service,omitempty"`
	StatusCode        *int64  `json:"statusCode,omitempty"`
}

type GetWebhookWebhookNotifiers struct {
	Emails  []string `json:"emails,omitempty"`
	Webhook *string  `json:"webhook,omitempty"`
}

// GetWebhookWebhook
// Configuration to alert to a url or list of email addresses based on the given type / condition.
type GetWebhookWebhook struct {
	CompanyID *string                    `json:"companyId,omitempty"`
	ID        string                     `json:"id"`
	Notifiers GetWebhookWebhookNotifiers `json:"notifiers"`
	Type      string                     `json:"type"`
}

type GetWebhookResponse struct {
	ContentType                        string
	StatusCode                         int
	Webhook                            *GetWebhookWebhook
	GetWebhook401ApplicationJSONObject *GetWebhook401ApplicationJSON
	GetWebhook404ApplicationJSONObject *GetWebhook404ApplicationJSON
}
