package webhooks

import (
	"net/http"
)

type DatasetStatusHasChangedToAnErrorStateResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}

type DatasetStatusHasChangedToAnErrorStateDatasetDataErrorWebhookData struct {
	DataType      *string `json:"dataType,omitempty"`
	DatasetID     *string `json:"datasetId,omitempty"`
	DatasetStatus *string `json:"datasetStatus,omitempty"`
}

// DatasetStatusHasChangedToAnErrorStateDatasetDataErrorWebhook
// Webhook request body to notify that a data synchronization has completed.
type DatasetStatusHasChangedToAnErrorStateDatasetDataErrorWebhook struct {
	AlertID   *string                                                           `json:"alertId,omitempty"`
	CompanyID *string                                                           `json:"companyId,omitempty"`
	Data      *DatasetStatusHasChangedToAnErrorStateDatasetDataErrorWebhookData `json:"data,omitempty"`
	Message   *string                                                           `json:"message,omitempty"`
	RuleID    *string                                                           `json:"ruleId,omitempty"`
	RuleType  *string                                                           `json:"ruleType,omitempty"`
}

type DatasetStatusHasChangedToAnErrorStateRequest struct {
	Request *DatasetStatusHasChangedToAnErrorStateDatasetDataErrorWebhook `request:"mediaType=application/json"`
}
