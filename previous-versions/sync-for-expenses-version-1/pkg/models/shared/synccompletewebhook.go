// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// SyncCompleteWebhook - Webhook request body used to notify that a sync has completed.
type SyncCompleteWebhook struct {
	// Unique identifier of the webhook event.
	AlertID *string `json:"AlertId,omitempty"`
	// Unique identifier for your client in Codat.
	ClientID *string `json:"ClientId,omitempty"`
	// Name of your client in Codat.
	ClientName *string `json:"ClientName,omitempty"`
	// Unique identifier for your SMB in Codat.
	CompanyID *string                  `json:"CompanyId,omitempty"`
	Data      *SyncCompleteWebhookData `json:"Data,omitempty"`
	// A human-readable message about the webhook.
	Message *string `json:"Message,omitempty"`
	// Unique identifier for the rule.
	//
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	RuleID *string `json:"RuleId,omitempty"`
	// The type of rule.
	RuleType *string `json:"RuleType,omitempty"`
}

func (o *SyncCompleteWebhook) GetAlertID() *string {
	if o == nil {
		return nil
	}
	return o.AlertID
}

func (o *SyncCompleteWebhook) GetClientID() *string {
	if o == nil {
		return nil
	}
	return o.ClientID
}

func (o *SyncCompleteWebhook) GetClientName() *string {
	if o == nil {
		return nil
	}
	return o.ClientName
}

func (o *SyncCompleteWebhook) GetCompanyID() *string {
	if o == nil {
		return nil
	}
	return o.CompanyID
}

func (o *SyncCompleteWebhook) GetData() *SyncCompleteWebhookData {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *SyncCompleteWebhook) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *SyncCompleteWebhook) GetRuleID() *string {
	if o == nil {
		return nil
	}
	return o.RuleID
}

func (o *SyncCompleteWebhook) GetRuleType() *string {
	if o == nil {
		return nil
	}
	return o.RuleType
}
