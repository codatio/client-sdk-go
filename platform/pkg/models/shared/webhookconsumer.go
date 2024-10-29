// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/codatio/client-sdk-go/platform/v4/pkg/utils"
)

// WebhookConsumer - A webhook consumer is an HTTP endpoint that developers can configure to subscribe to Codat's supported event types.
//
// See our documentation for more details on [Codat's webhook service](https://docs.codat.io/using-the-api/webhooks/overview).
type WebhookConsumer struct {
	// Unique identifier of the company to indicate company-specific events. The associated webhook consumer will receive events only for the specified ID.
	//
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	CompanyID *string `json:"companyId,omitempty"`
	// Company tags provide an additional way to filter messages, independent of event types. Company tags are case-sensitive, and only messages from companies with matching tags will be sent to this endpoint. Use the format `tagKey:tagValue`.
	CompanyTags []string `json:"companyTags,omitempty"`
	// Flag that enables or disables the endpoint from receiving events. Disabled when set to `true`.
	Disabled *bool `default:"false" json:"disabled"`
	// An array of event types the webhook consumer subscribes to.
	EventTypes []string `json:"eventTypes,omitempty"`
	// Unique identifier for the webhook consumer.
	ID *string `json:"id,omitempty"`
	// The URL that will consume webhook events dispatched by Codat.
	URL *string `json:"url,omitempty"`
}

func (w WebhookConsumer) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(w, "", false)
}

func (w *WebhookConsumer) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &w, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *WebhookConsumer) GetCompanyID() *string {
	if o == nil {
		return nil
	}
	return o.CompanyID
}

func (o *WebhookConsumer) GetCompanyTags() []string {
	if o == nil {
		return nil
	}
	return o.CompanyTags
}

func (o *WebhookConsumer) GetDisabled() *bool {
	if o == nil {
		return nil
	}
	return o.Disabled
}

func (o *WebhookConsumer) GetEventTypes() []string {
	if o == nil {
		return nil
	}
	return o.EventTypes
}

func (o *WebhookConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *WebhookConsumer) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}
