// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type DatasetStatusChangedErrorWebhookData struct {
	// Data type the sync completed for.
	DataType *string `json:"dataType,omitempty"`
	// Unique identifier for the dataset that completed its sync.
	DatasetID *string `json:"datasetId,omitempty"`
	// The current status of the dataset's sync.
	DatasetStatus *string `json:"datasetStatus,omitempty"`
}

// DatasetStatusChangedErrorWebhook - Webhook request body to notify that a data synchronization has completed.
type DatasetStatusChangedErrorWebhook struct {
	// Unique identifier of the webhook event.
	AlertID *string `json:"alertId,omitempty"`
	// Unique identifier for your SMB in Codat.
	CompanyID *string                               `json:"companyId,omitempty"`
	Data      *DatasetStatusChangedErrorWebhookData `json:"data,omitempty"`
	// A human readable message about the webhook.
	Message *string `json:"message,omitempty"`
	// Unique identifier for the rule.
	RuleID *string `json:"ruleId,omitempty"`
	// The type of rule.
	Type *string `json:"type,omitempty"`
}
