// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1/pkg/models/shared"
	"net/http"
)

type UpdateConfigTextSyncFlowResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Success
	LocalizationInfo map[string]shared.Localization
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *UpdateConfigTextSyncFlowResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateConfigTextSyncFlowResponse) GetLocalizationInfo() map[string]shared.Localization {
	if o == nil {
		return nil
	}
	return o.LocalizationInfo
}

func (o *UpdateConfigTextSyncFlowResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateConfigTextSyncFlowResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
