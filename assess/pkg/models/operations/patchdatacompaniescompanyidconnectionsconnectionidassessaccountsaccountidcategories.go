package operations

import (
	"time"
)

type PatchDataCompaniesCompanyIDConnectionsConnectionIDAssessAccountsAccountIDCategoriesPathParams struct {
	AccountID    string `pathParam:"style=simple,explode=false,name=accountId"`
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type PatchDataCompaniesCompanyIDConnectionsConnectionIDAssessAccountsAccountIDCategoriesRequestBodyChartOfAccountCategory struct {
	DetailType *string `json:"detailType,omitempty"`
	Subtype    *string `json:"subtype,omitempty"`
	Type       *string `json:"type,omitempty"`
}

type PatchDataCompaniesCompanyIDConnectionsConnectionIDAssessAccountsAccountIDCategoriesRequestBody struct {
	Confirmed PatchDataCompaniesCompanyIDConnectionsConnectionIDAssessAccountsAccountIDCategoriesRequestBodyChartOfAccountCategory `json:"confirmed"`
}

type PatchDataCompaniesCompanyIDConnectionsConnectionIDAssessAccountsAccountIDCategoriesRequest struct {
	PathParams PatchDataCompaniesCompanyIDConnectionsConnectionIDAssessAccountsAccountIDCategoriesPathParams
	Request    *PatchDataCompaniesCompanyIDConnectionsConnectionIDAssessAccountsAccountIDCategoriesRequestBody `request:"mediaType=application/json"`
}

// PatchDataCompaniesCompanyIDConnectionsConnectionIDAssessAccountsAccountIDCategoriesCategorisedAccountAccountRef
// An object containing account reference data.
type PatchDataCompaniesCompanyIDConnectionsConnectionIDAssessAccountsAccountIDCategoriesCategorisedAccountAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type PatchDataCompaniesCompanyIDConnectionsConnectionIDAssessAccountsAccountIDCategoriesCategorisedAccountModifiedDate struct {
	DetailType   *string    `json:"detailType,omitempty"`
	ModifiedDate *time.Time `json:"modifiedDate,omitempty"`
	Subtype      *string    `json:"subtype,omitempty"`
	Type         *string    `json:"type,omitempty"`
}

type PatchDataCompaniesCompanyIDConnectionsConnectionIDAssessAccountsAccountIDCategoriesCategorisedAccount struct {
	AccountRef *PatchDataCompaniesCompanyIDConnectionsConnectionIDAssessAccountsAccountIDCategoriesCategorisedAccountAccountRef   `json:"accountRef,omitempty"`
	Confirmed  *PatchDataCompaniesCompanyIDConnectionsConnectionIDAssessAccountsAccountIDCategoriesCategorisedAccountModifiedDate `json:"confirmed,omitempty"`
	Suggested  *PatchDataCompaniesCompanyIDConnectionsConnectionIDAssessAccountsAccountIDCategoriesCategorisedAccountModifiedDate `json:"suggested,omitempty"`
}

type PatchDataCompaniesCompanyIDConnectionsConnectionIDAssessAccountsAccountIDCategoriesResponse struct {
	CategorisedAccount *PatchDataCompaniesCompanyIDConnectionsConnectionIDAssessAccountsAccountIDCategoriesCategorisedAccount
	ContentType        string
	StatusCode         int
}
