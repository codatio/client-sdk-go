// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// ClientRateLimitReachedWebhook - Webhook request body for a client that has reached their rate limit.
type ClientRateLimitReachedWebhook struct {
	// Unique identifier of the webhook event.
	AlertID *string `json:"AlertId,omitempty"`
	// Unique identifier for your client in Codat.
	ClientID *string `json:"ClientId,omitempty"`
	// Name of your client in Codat.
	ClientName *string                            `json:"ClientName,omitempty"`
	Data       *ClientRateLimitReachedWebhookData `json:"Data,omitempty"`
	// A human-readable message about the webhook.
	Message *string `json:"Message,omitempty"`
	// Unique identifier for the rule.
	//
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	RuleID *string `json:"RuleId,omitempty"`
	// The type of rule.
	RuleType *string `json:"RuleType,omitempty"`
}

func (o *ClientRateLimitReachedWebhook) GetAlertID() *string {
	if o == nil {
		return nil
	}
	return o.AlertID
}

func (o *ClientRateLimitReachedWebhook) GetClientID() *string {
	if o == nil {
		return nil
	}
	return o.ClientID
}

func (o *ClientRateLimitReachedWebhook) GetClientName() *string {
	if o == nil {
		return nil
	}
	return o.ClientName
}

func (o *ClientRateLimitReachedWebhook) GetData() *ClientRateLimitReachedWebhookData {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *ClientRateLimitReachedWebhook) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *ClientRateLimitReachedWebhook) GetRuleID() *string {
	if o == nil {
		return nil
	}
	return o.RuleID
}

func (o *ClientRateLimitReachedWebhook) GetRuleType() *string {
	if o == nil {
		return nil
	}
	return o.RuleType
}
