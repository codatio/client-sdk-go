package operations

import (
	"net/http"
	"time"
)

type UpdateAccountsCategoriesPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type UpdateAccountsCategoriesRequestBodyCategoriesAccountRef struct {
	ID string `json:"id"`
}

type UpdateAccountsCategoriesRequestBodyCategories struct {
	AccountRef *UpdateAccountsCategoriesRequestBodyCategoriesAccountRef `json:"accountRef,omitempty"`
	Confirmed  interface{}                                              `json:"confirmed,omitempty"`
}

type UpdateAccountsCategoriesRequestBody struct {
	Categories []UpdateAccountsCategoriesRequestBodyCategories `json:"categories,omitempty"`
}

type UpdateAccountsCategoriesRequest struct {
	PathParams UpdateAccountsCategoriesPathParams
	Request    *UpdateAccountsCategoriesRequestBody `request:"mediaType=application/json"`
}

// UpdateAccountsCategoriesCategorisedAccountAccountRef
// An object containing account reference data.
type UpdateAccountsCategoriesCategorisedAccountAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type UpdateAccountsCategoriesCategorisedAccountModifiedDate struct {
	DetailType   *string    `json:"detailType,omitempty"`
	ModifiedDate *time.Time `json:"modifiedDate,omitempty"`
	Subtype      *string    `json:"subtype,omitempty"`
	Type         *string    `json:"type,omitempty"`
}

type UpdateAccountsCategoriesCategorisedAccount struct {
	AccountRef *UpdateAccountsCategoriesCategorisedAccountAccountRef   `json:"accountRef,omitempty"`
	Confirmed  *UpdateAccountsCategoriesCategorisedAccountModifiedDate `json:"confirmed,omitempty"`
	Suggested  *UpdateAccountsCategoriesCategorisedAccountModifiedDate `json:"suggested,omitempty"`
}

type UpdateAccountsCategoriesResponse struct {
	CategorisedAccounts []UpdateAccountsCategoriesCategorisedAccount
	ContentType         string
	StatusCode          int
	RawResponse         *http.Response
}
