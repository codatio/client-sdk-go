package webhooks

type PushOperationHasTimedOutResponse struct {
	ContentType string
	StatusCode  int
}

type PushOperationHasTimedOutPushOperationTimedOutWebhookData struct {
	DataType         *string `json:"dataType,omitempty"`
	PushOperationKey *string `json:"pushOperationKey,omitempty"`
}

// PushOperationHasTimedOutPushOperationTimedOutWebhook
// Webhook request body notifying that a push push operation has timed out.
type PushOperationHasTimedOutPushOperationTimedOutWebhook struct {
	AlertID   *string                                                   `json:"alertId,omitempty"`
	CompanyID *string                                                   `json:"companyId,omitempty"`
	Data      *PushOperationHasTimedOutPushOperationTimedOutWebhookData `json:"data,omitempty"`
	Message   *string                                                   `json:"message,omitempty"`
	RuleID    *string                                                   `json:"ruleId,omitempty"`
	RuleType  *string                                                   `json:"ruleType,omitempty"`
}

type PushOperationHasTimedOutRequest struct {
	Request *PushOperationHasTimedOutPushOperationTimedOutWebhook `request:"mediaType=application/json"`
}
