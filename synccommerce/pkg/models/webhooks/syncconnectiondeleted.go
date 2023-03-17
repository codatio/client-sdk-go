package webhooks

import (
	"net/http"
)

type SyncConnectionDeletedResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}

// SyncConnectionDeletedAccountCategoriesUpdatedWebhook
// Webhook request body for account categories updated.
type SyncConnectionDeletedAccountCategoriesUpdatedWebhook struct {
	AlertID          *string `json:"alertId,omitempty"`
	ClientID         *string `json:"clientId,omitempty"`
	ClientName       *string `json:"clientName,omitempty"`
	CompanyID        *string `json:"companyId,omitempty"`
	DataConnectionID *string `json:"dataConnectionId,omitempty"`
	Message          *string `json:"message,omitempty"`
	RuleID           *string `json:"ruleId,omitempty"`
	RuleType         *string `json:"ruleType,omitempty"`
}
