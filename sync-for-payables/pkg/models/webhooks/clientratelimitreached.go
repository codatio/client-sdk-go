// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package webhooks

import (
	"github.com/codatio/client-sdk-go/sync-for-payables/v4/pkg/models/shared"
)

type ClientRateLimitReachedResponse1 struct {
	HTTPMeta shared.HTTPMetadata `json:"-"`
}

func (o *ClientRateLimitReachedResponse1) GetHTTPMeta() shared.HTTPMetadata {
	if o == nil {
		return shared.HTTPMetadata{}
	}
	return o.HTTPMeta
}
