// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type WebhookNotifier struct {
	Emails []string `json:"emails,omitempty"`
	// The URI the webhook service will use to post events.
	Webhook *string `json:"webhook,omitempty"`
}

func (o *WebhookNotifier) GetEmails() []string {
	if o == nil {
		return nil
	}
	return o.Emails
}

func (o *WebhookNotifier) GetWebhook() *string {
	if o == nil {
		return nil
	}
	return o.Webhook
}
