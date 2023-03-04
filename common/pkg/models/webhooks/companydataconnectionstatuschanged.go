package webhooks

import (
	"net/http"
)

type CompanyDataConnectionStatusChangedResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}

type CompanyDataConnectionStatusChangedCompanyDataConnectionStatusChangedWebhookDataDataConnectionStatusEnum string

const (
	CompanyDataConnectionStatusChangedCompanyDataConnectionStatusChangedWebhookDataDataConnectionStatusEnumPendingAuth  CompanyDataConnectionStatusChangedCompanyDataConnectionStatusChangedWebhookDataDataConnectionStatusEnum = "PendingAuth"
	CompanyDataConnectionStatusChangedCompanyDataConnectionStatusChangedWebhookDataDataConnectionStatusEnumLinked       CompanyDataConnectionStatusChangedCompanyDataConnectionStatusChangedWebhookDataDataConnectionStatusEnum = "Linked"
	CompanyDataConnectionStatusChangedCompanyDataConnectionStatusChangedWebhookDataDataConnectionStatusEnumUnlinked     CompanyDataConnectionStatusChangedCompanyDataConnectionStatusChangedWebhookDataDataConnectionStatusEnum = "Unlinked"
	CompanyDataConnectionStatusChangedCompanyDataConnectionStatusChangedWebhookDataDataConnectionStatusEnumDeauthorized CompanyDataConnectionStatusChangedCompanyDataConnectionStatusChangedWebhookDataDataConnectionStatusEnum = "Deauthorized"
)

type CompanyDataConnectionStatusChangedCompanyDataConnectionStatusChangedWebhookData struct {
	DataConnectionID *string                                                                                                  `json:"dataConnectionId,omitempty"`
	NewStatus        *CompanyDataConnectionStatusChangedCompanyDataConnectionStatusChangedWebhookDataDataConnectionStatusEnum `json:"newStatus,omitempty"`
	OldStatus        *CompanyDataConnectionStatusChangedCompanyDataConnectionStatusChangedWebhookDataDataConnectionStatusEnum `json:"oldStatus,omitempty"`
	PlatformKey      *string                                                                                                  `json:"platformKey,omitempty"`
}

// CompanyDataConnectionStatusChangedCompanyDataConnectionStatusChangedWebhook
// Webhook request body for a company's data connection status changed.
type CompanyDataConnectionStatusChangedCompanyDataConnectionStatusChangedWebhook struct {
	AlertID   *string                                                                          `json:"alertId,omitempty"`
	CompanyID *string                                                                          `json:"companyId,omitempty"`
	Data      *CompanyDataConnectionStatusChangedCompanyDataConnectionStatusChangedWebhookData `json:"data,omitempty"`
	Message   *string                                                                          `json:"message,omitempty"`
	RuleID    *string                                                                          `json:"ruleId,omitempty"`
	RuleType  *string                                                                          `json:"ruleType,omitempty"`
}

type CompanyDataConnectionStatusChangedRequest struct {
	Request *CompanyDataConnectionStatusChangedCompanyDataConnectionStatusChangedWebhook `request:"mediaType=application/json"`
}
