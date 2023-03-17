package webhooks

import (
	"net/http"
	"time"
)

type AccountCategoriesUpdatedResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}

type AccountCategoriesUpdatedAccountCategoriesUpdatedWebhookData struct {
	ModifiedDate *time.Time `json:"modifiedDate,omitempty"`
}

// AccountCategoriesUpdatedAccountCategoriesUpdatedWebhook
// Webhook request body for account categories updated.
type AccountCategoriesUpdatedAccountCategoriesUpdatedWebhook struct {
	AlertID          *string                                                      `json:"alertId,omitempty"`
	ClientID         *string                                                      `json:"clientId,omitempty"`
	ClientName       *string                                                      `json:"clientName,omitempty"`
	CompanyID        *string                                                      `json:"companyId,omitempty"`
	Data             *AccountCategoriesUpdatedAccountCategoriesUpdatedWebhookData `json:"data,omitempty"`
	DataConnectionID *string                                                      `json:"dataConnectionId,omitempty"`
	Message          *string                                                      `json:"message,omitempty"`
	RuleID           *string                                                      `json:"ruleId,omitempty"`
	RuleType         *string                                                      `json:"ruleType,omitempty"`
}
