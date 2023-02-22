package operations

import (
	"github.com/codatio/client-sdk-go/common/pkg/models/shared"
)

type GetWebhookPathParams struct {
	RuleID string `pathParam:"style=simple,explode=false,name=ruleId"`
}

type GetWebhookRequest struct {
	PathParams GetWebhookPathParams
}

type GetWebhook404ApplicationJSON struct {
	CanBeRetried      *string `json:"canBeRetried,omitempty"`
	CorrelationID     *string `json:"correlationId,omitempty"`
	DetailedErrorCode *int64  `json:"detailedErrorCode,omitempty"`
	Error             *string `json:"error,omitempty"`
	Service           *string `json:"service,omitempty"`
	StatusCode        *int64  `json:"statusCode,omitempty"`
}

type GetWebhook401ApplicationJSON struct {
	CanBeRetried      *string `json:"canBeRetried,omitempty"`
	CorrelationID     *string `json:"correlationId,omitempty"`
	DetailedErrorCode *int64  `json:"detailedErrorCode,omitempty"`
	Error             *string `json:"error,omitempty"`
	Service           *string `json:"service,omitempty"`
	StatusCode        *int64  `json:"statusCode,omitempty"`
}

type GetWebhookResponse struct {
	ContentType                        string
	Rule                               *shared.Rule
	StatusCode                         int
	GetWebhook401ApplicationJSONObject *GetWebhook401ApplicationJSON
	GetWebhook404ApplicationJSONObject *GetWebhook404ApplicationJSON
}
