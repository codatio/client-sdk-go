package webhooks

type NewCompanySynchronizedResponse struct {
	ContentType string
	StatusCode  int
}

// NewCompanySynchronizedNewCompanySynchronizedWebhook
// Webhook request body to notify that a new company has successfully synchronized at least one dataType for the first time.
type NewCompanySynchronizedNewCompanySynchronizedWebhook struct {
	AlertID   *string `json:"alertId,omitempty"`
	CompanyID *string `json:"companyId,omitempty"`
	Message   *string `json:"message,omitempty"`
	RuleID    *string `json:"ruleId,omitempty"`
	RuleType  *string `json:"ruleType,omitempty"`
}

type NewCompanySynchronizedRequest struct {
	Request *NewCompanySynchronizedNewCompanySynchronizedWebhook `request:"mediaType=application/json"`
}
