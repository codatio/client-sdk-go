// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type CompanyDataConnectionStatusChangedWebhookData struct {
	// Unique identifier for a company's data connection.
	DataConnectionID *string `json:"dataConnectionId,omitempty"`
	// The current authorization status of the data connection.
	NewStatus *ConnectionStatusEnum `json:"newStatus,omitempty"`
	// The current authorization status of the data connection.
	OldStatus *ConnectionStatusEnum `json:"oldStatus,omitempty"`
	// A unique 4-letter key to represent a platform in each integration. View [accounting](https://docs.codat.io/integrations/accounting/accounting-platform-keys), [banking](https://docs.codat.io/integrations/banking/banking-platform-keys), and [commerce](https://docs.codat.io/integrations/commerce/commerce-platform-keys) platform keys.
	PlatformKey *string `json:"platformKey,omitempty"`
}

// CompanyDataConnectionStatusChangedWebhook - Webhook request body for a company's data connection status changed.
type CompanyDataConnectionStatusChangedWebhook struct {
	// Unique identifier of the alert.
	AlertID *string `json:"alertId,omitempty"`
	// Unique identifier for your SMB in Codat.
	CompanyID *string                                        `json:"companyId,omitempty"`
	Data      *CompanyDataConnectionStatusChangedWebhookData `json:"data,omitempty"`
	// A human readable message about the webhook.
	Message *string `json:"message,omitempty"`
	// Unique identifier for the rule.
	RuleID *string `json:"ruleId,omitempty"`
	// The type of rule.
	RuleType *string `json:"ruleType,omitempty"`
}
