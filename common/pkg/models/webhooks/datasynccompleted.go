package webhooks

import (
	"net/http"
)

type DataSyncCompletedResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}

type DataSyncCompletedDataSyncCompleteWebhookData struct {
	DataType  *string `json:"dataType,omitempty"`
	DatasetID *string `json:"datasetId,omitempty"`
}

// DataSyncCompletedDataSyncCompleteWebhook
// Webhook request body to notify the completion of a data sync.
type DataSyncCompletedDataSyncCompleteWebhook struct {
	AlertID          *string                                       `json:"alertId,omitempty"`
	ClientID         *string                                       `json:"clientId,omitempty"`
	ClientName       *string                                       `json:"clientName,omitempty"`
	CompanyID        *string                                       `json:"companyId,omitempty"`
	Data             *DataSyncCompletedDataSyncCompleteWebhookData `json:"data,omitempty"`
	DataConnectionID *string                                       `json:"dataConnectionId,omitempty"`
	Message          *string                                       `json:"message,omitempty"`
	RuleID           *string                                       `json:"ruleId,omitempty"`
	RuleType         *string                                       `json:"ruleType,omitempty"`
}

type DataSyncCompletedRequest struct {
	Request *DataSyncCompletedDataSyncCompleteWebhook `request:"mediaType=application/json"`
}
