package operations

import (
	"time"
)

type GetAccountCategoryPathParams struct {
	AccountID    string `pathParam:"style=simple,explode=false,name=accountId"`
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type GetAccountCategoryRequest struct {
	PathParams GetAccountCategoryPathParams
}

// GetAccountCategoryCategorisedAccountAccountRef
// An object containing account reference data.
type GetAccountCategoryCategorisedAccountAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type GetAccountCategoryCategorisedAccountModifiedDate struct {
	DetailType   *string    `json:"detailType,omitempty"`
	ModifiedDate *time.Time `json:"modifiedDate,omitempty"`
	Subtype      *string    `json:"subtype,omitempty"`
	Type         *string    `json:"type,omitempty"`
}

type GetAccountCategoryCategorisedAccount struct {
	AccountRef *GetAccountCategoryCategorisedAccountAccountRef   `json:"accountRef,omitempty"`
	Confirmed  *GetAccountCategoryCategorisedAccountModifiedDate `json:"confirmed,omitempty"`
	Suggested  *GetAccountCategoryCategorisedAccountModifiedDate `json:"suggested,omitempty"`
}

type GetAccountCategoryResponse struct {
	CategorisedAccount *GetAccountCategoryCategorisedAccount
	ContentType        string
	StatusCode         int
}
