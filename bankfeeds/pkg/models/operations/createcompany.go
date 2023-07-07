// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/bankfeeds/pkg/models/shared"
	"net/http"
)

type CreateCompanyResponse struct {
	// OK
	Company     *shared.Company
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// The request made is not valid.
	Schema *shared.Schema
}
