// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package webhooks

import (
	"github.com/codatio/client-sdk-go/sync-for-payables/v3/pkg/models/shared"
)

type ClientRateLimitResetResponse1 struct {
	HTTPMeta shared.HTTPMetadata `json:"-"`
}

func (o *ClientRateLimitResetResponse1) GetHTTPMeta() shared.HTTPMetadata {
	if o == nil {
		return shared.HTTPMetadata{}
	}
	return o.HTTPMeta
}
