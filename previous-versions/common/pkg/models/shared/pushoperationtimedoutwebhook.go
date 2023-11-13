// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// PushOperationTimedOutWebhook - Webhook request body notifying that a push push operation has timed out.
type PushOperationTimedOutWebhook struct {
	// Unique identifier of the webhook event.
	AlertID *string `json:"AlertId,omitempty"`
	// Unique identifier for your client in Codat.
	ClientID *string `json:"ClientId,omitempty"`
	// Name of your client in Codat.
	ClientName *string `json:"ClientName,omitempty"`
	// Unique identifier for your SMB in Codat.
	CompanyID *string                           `json:"CompanyId,omitempty"`
	Data      *PushOperationTimedOutWebhookData `json:"Data,omitempty"`
	// Unique identifier for a company's data connection.
	DataConnectionID *string `json:"DataConnectionId,omitempty"`
	// A human readable message about the webhook.
	Message *string `json:"Message,omitempty"`
	// Unique identifier for the rule.
	RuleID *string `json:"RuleId,omitempty"`
	// The type of rule.
	RuleType *string `json:"RuleType,omitempty"`
}

func (o *PushOperationTimedOutWebhook) GetAlertID() *string {
	if o == nil {
		return nil
	}
	return o.AlertID
}

func (o *PushOperationTimedOutWebhook) GetClientID() *string {
	if o == nil {
		return nil
	}
	return o.ClientID
}

func (o *PushOperationTimedOutWebhook) GetClientName() *string {
	if o == nil {
		return nil
	}
	return o.ClientName
}

func (o *PushOperationTimedOutWebhook) GetCompanyID() *string {
	if o == nil {
		return nil
	}
	return o.CompanyID
}

func (o *PushOperationTimedOutWebhook) GetData() *PushOperationTimedOutWebhookData {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *PushOperationTimedOutWebhook) GetDataConnectionID() *string {
	if o == nil {
		return nil
	}
	return o.DataConnectionID
}

func (o *PushOperationTimedOutWebhook) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *PushOperationTimedOutWebhook) GetRuleID() *string {
	if o == nil {
		return nil
	}
	return o.RuleID
}

func (o *PushOperationTimedOutWebhook) GetRuleType() *string {
	if o == nil {
		return nil
	}
	return o.RuleType
}
