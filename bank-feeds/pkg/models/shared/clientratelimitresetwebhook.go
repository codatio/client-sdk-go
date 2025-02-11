// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// ClientRateLimitResetWebhook - Webhook request body for a client that has had their rate limit reset.
type ClientRateLimitResetWebhook struct {
	// Unique identifier of the webhook event.
	AlertID *string `json:"AlertId,omitempty"`
	// Unique identifier for your client in Codat.
	ClientID *string `json:"ClientId,omitempty"`
	// Name of your client in Codat.
	ClientName *string                          `json:"ClientName,omitempty"`
	Data       *ClientRateLimitResetWebhookData `json:"Data,omitempty"`
	// A human-readable message about the webhook.
	Message *string `json:"Message,omitempty"`
	// Unique identifier for the rule.
	//
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	RuleID *string `json:"RuleId,omitempty"`
	// The type of rule.
	RuleType *string `json:"RuleType,omitempty"`
}

func (o *ClientRateLimitResetWebhook) GetAlertID() *string {
	if o == nil {
		return nil
	}
	return o.AlertID
}

func (o *ClientRateLimitResetWebhook) GetClientID() *string {
	if o == nil {
		return nil
	}
	return o.ClientID
}

func (o *ClientRateLimitResetWebhook) GetClientName() *string {
	if o == nil {
		return nil
	}
	return o.ClientName
}

func (o *ClientRateLimitResetWebhook) GetData() *ClientRateLimitResetWebhookData {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *ClientRateLimitResetWebhook) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *ClientRateLimitResetWebhook) GetRuleID() *string {
	if o == nil {
		return nil
	}
	return o.RuleID
}

func (o *ClientRateLimitResetWebhook) GetRuleType() *string {
	if o == nil {
		return nil
	}
	return o.RuleType
}
