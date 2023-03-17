package webhooks

import (
	"net/http"
)

type PushOperationStatusHasChangedResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}

type PushOperationStatusHasChangedPushOperationStatusChangedWebhookData struct {
	DataType         *string `json:"dataType,omitempty"`
	PushOperationKey *string `json:"pushOperationKey,omitempty"`
	Status           *string `json:"status,omitempty"`
}

// PushOperationStatusHasChangedPushOperationStatusChangedWebhook
// Webhook request body for a push operation status change.
type PushOperationStatusHasChangedPushOperationStatusChangedWebhook struct {
	AlertID   *string                                                             `json:"alertId,omitempty"`
	CompanyID *string                                                             `json:"companyId,omitempty"`
	Data      *PushOperationStatusHasChangedPushOperationStatusChangedWebhookData `json:"data,omitempty"`
	Message   *string                                                             `json:"message,omitempty"`
	RuleID    *string                                                             `json:"ruleId,omitempty"`
	RuleType  *string                                                             `json:"ruleType,omitempty"`
}
