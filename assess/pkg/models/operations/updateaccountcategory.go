package operations

import (
	"net/http"
	"time"
)

type UpdateAccountCategoryRequestBodyChartOfAccountCategory struct {
	DetailType *string `json:"detailType,omitempty"`
	Subtype    *string `json:"subtype,omitempty"`
	Type       *string `json:"type,omitempty"`
}

type UpdateAccountCategoryRequestBody struct {
	Confirmed UpdateAccountCategoryRequestBodyChartOfAccountCategory `json:"confirmed"`
}

type UpdateAccountCategoryRequest struct {
	RequestBody  *UpdateAccountCategoryRequestBody `request:"mediaType=application/json"`
	AccountID    string                            `pathParam:"style=simple,explode=false,name=accountId"`
	CompanyID    string                            `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string                            `pathParam:"style=simple,explode=false,name=connectionId"`
}

// UpdateAccountCategoryCategorisedAccountAccountRef
// An object containing account reference data.
type UpdateAccountCategoryCategorisedAccountAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type UpdateAccountCategoryCategorisedAccountModifiedDate struct {
	DetailType   *string    `json:"detailType,omitempty"`
	ModifiedDate *time.Time `json:"modifiedDate,omitempty"`
	Subtype      *string    `json:"subtype,omitempty"`
	Type         *string    `json:"type,omitempty"`
}

type UpdateAccountCategoryCategorisedAccount struct {
	AccountRef *UpdateAccountCategoryCategorisedAccountAccountRef   `json:"accountRef,omitempty"`
	Confirmed  *UpdateAccountCategoryCategorisedAccountModifiedDate `json:"confirmed,omitempty"`
	Suggested  *UpdateAccountCategoryCategorisedAccountModifiedDate `json:"suggested,omitempty"`
}

type UpdateAccountCategoryResponse struct {
	CategorisedAccount *UpdateAccountCategoryCategorisedAccount
	ContentType        string
	StatusCode         int
	RawResponse        *http.Response
}
