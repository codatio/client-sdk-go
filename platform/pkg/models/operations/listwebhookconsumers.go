// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/platform/v5/pkg/models/shared"
	"net/http"
)

type ListWebhookConsumersResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	WebhookConsumers *shared.WebhookConsumers
}

func (o *ListWebhookConsumersResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListWebhookConsumersResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListWebhookConsumersResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListWebhookConsumersResponse) GetWebhookConsumers() *shared.WebhookConsumers {
	if o == nil {
		return nil
	}
	return o.WebhookConsumers
}
