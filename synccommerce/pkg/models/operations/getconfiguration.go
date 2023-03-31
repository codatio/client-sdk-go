// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/synccommerce/pkg/models/shared"
	"net/http"
)

type GetConfigurationRequest struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type GetConfigurationResponse struct {
	// Success
	Configuration *shared.Configuration
	ContentType   string
	StatusCode    int
	RawResponse   *http.Response
}
