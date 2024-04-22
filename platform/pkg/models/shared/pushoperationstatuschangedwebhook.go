// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// PushOperationStatusChangedWebhook - Webhook request body for a push operation status change.
type PushOperationStatusChangedWebhook struct {
	// Unique identifier of the webhook event.
	AlertID *string `json:"AlertId,omitempty"`
	// Unique identifier for your client in Codat.
	ClientID *string `json:"ClientId,omitempty"`
	// Name of your client in Codat.
	ClientName *string `json:"ClientName,omitempty"`
	// Unique identifier for your SMB in Codat.
	CompanyID *string                                `json:"CompanyId,omitempty"`
	Data      *PushOperationStatusChangedWebhookData `json:"Data,omitempty"`
	// Unique identifier for a company's data connection.
	DataConnectionID *string `json:"DataConnectionId,omitempty"`
	// A human-readable message about the webhook.
	Message *string `json:"Message,omitempty"`
	// Unique identifier for the rule.
	//
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	RuleID *string `json:"RuleId,omitempty"`
	// The type of rule.
	RuleType *string `json:"RuleType,omitempty"`
}

func (o *PushOperationStatusChangedWebhook) GetAlertID() *string {
	if o == nil {
		return nil
	}
	return o.AlertID
}

func (o *PushOperationStatusChangedWebhook) GetClientID() *string {
	if o == nil {
		return nil
	}
	return o.ClientID
}

func (o *PushOperationStatusChangedWebhook) GetClientName() *string {
	if o == nil {
		return nil
	}
	return o.ClientName
}

func (o *PushOperationStatusChangedWebhook) GetCompanyID() *string {
	if o == nil {
		return nil
	}
	return o.CompanyID
}

func (o *PushOperationStatusChangedWebhook) GetData() *PushOperationStatusChangedWebhookData {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *PushOperationStatusChangedWebhook) GetDataConnectionID() *string {
	if o == nil {
		return nil
	}
	return o.DataConnectionID
}

func (o *PushOperationStatusChangedWebhook) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *PushOperationStatusChangedWebhook) GetRuleID() *string {
	if o == nil {
		return nil
	}
	return o.RuleID
}

func (o *PushOperationStatusChangedWebhook) GetRuleType() *string {
	if o == nil {
		return nil
	}
	return o.RuleType
}
