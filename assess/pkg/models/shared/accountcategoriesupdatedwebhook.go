// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type AccountCategoriesUpdatedWebhookData struct {
	// The date on which this account categories were last modified in Codat.
	ModifiedDate *string `json:"modifiedDate,omitempty"`
}

// AccountCategoriesUpdatedWebhook - Webhook request body for account categories updated.
type AccountCategoriesUpdatedWebhook struct {
	// Unique identifier of the webhook event.
	AlertID *string `json:"AlertId,omitempty"`
	// Unique identifier for your client in Codat.
	ClientID *string `json:"ClientId,omitempty"`
	// Name of your client in Codat.
	ClientName *string `json:"ClientName,omitempty"`
	// Unique identifier for your SMB in Codat.
	CompanyID *string                              `json:"CompanyId,omitempty"`
	Data      *AccountCategoriesUpdatedWebhookData `json:"Data,omitempty"`
	// Unique identifier for a company's data connection.
	DataConnectionID *string `json:"DataConnectionId,omitempty"`
	// A human readable message about the webhook.
	Message *string `json:"Message,omitempty"`
	// Unique identifier for the rule.
	RuleID *string `json:"RuleId,omitempty"`
	// The type of rule.
	Type *string `json:"Type,omitempty"`
}
