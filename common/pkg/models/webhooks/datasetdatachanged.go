package webhooks

import (
	"net/http"
)

type DatasetDataChangedResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}

type DatasetDataChangedDatasetDataChangedWebhookData struct {
	DataType  *string `json:"dataType,omitempty"`
	DatasetID *string `json:"datasetId,omitempty"`
}

// DatasetDataChangedDatasetDataChangedWebhook
// Webhook request body to notify that a data synchronization has completed.
type DatasetDataChangedDatasetDataChangedWebhook struct {
	AlertID   *string                                          `json:"alertId,omitempty"`
	CompanyID *string                                          `json:"companyId,omitempty"`
	Data      *DatasetDataChangedDatasetDataChangedWebhookData `json:"data,omitempty"`
	Message   *string                                          `json:"message,omitempty"`
	RuleID    *string                                          `json:"ruleId,omitempty"`
	RuleType  *string                                          `json:"ruleType,omitempty"`
}

type DatasetDataChangedRequest struct {
	Request *DatasetDataChangedDatasetDataChangedWebhook `request:"mediaType=application/json"`
}
