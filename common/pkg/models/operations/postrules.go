package operations

type PostRulesWebhookNotifiers struct {
	Emails  []string `json:"emails,omitempty"`
	Webhook *string  `json:"webhook,omitempty"`
}

// PostRulesWebhook
// Configuration to alert to a url or list of email addresses based on the given type / condition.
type PostRulesWebhook struct {
	CompanyID *string                   `json:"companyId,omitempty"`
	ID        string                    `json:"id"`
	Notifiers PostRulesWebhookNotifiers `json:"notifiers"`
	Type      string                    `json:"type"`
}

type PostRulesRequest struct {
	Request *PostRulesWebhook `request:"mediaType=application/json"`
}

type PostRules401ApplicationJSON struct {
	CanBeRetried      *string `json:"canBeRetried,omitempty"`
	CorrelationID     *string `json:"correlationId,omitempty"`
	DetailedErrorCode *int64  `json:"detailedErrorCode,omitempty"`
	Error             *string `json:"error,omitempty"`
	Service           *string `json:"service,omitempty"`
	StatusCode        *int64  `json:"statusCode,omitempty"`
}

type PostRulesResponse struct {
	ContentType                       string
	StatusCode                        int
	Webhook                           *PostRulesWebhook
	PostRules401ApplicationJSONObject *PostRules401ApplicationJSON
}
