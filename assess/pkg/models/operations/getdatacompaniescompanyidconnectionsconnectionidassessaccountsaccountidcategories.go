package operations

import (
	"github.com/codatio/client-sdk-go/assess/pkg/models/shared"
)

type GetDataCompaniesCompanyIDConnectionsConnectionIDAssessAccountsAccountIDCategoriesPathParams struct {
	AccountID    string `pathParam:"style=simple,explode=false,name=accountId"`
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type GetDataCompaniesCompanyIDConnectionsConnectionIDAssessAccountsAccountIDCategoriesRequest struct {
	PathParams GetDataCompaniesCompanyIDConnectionsConnectionIDAssessAccountsAccountIDCategoriesPathParams
}

type GetDataCompaniesCompanyIDConnectionsConnectionIDAssessAccountsAccountIDCategoriesResponse struct {
	CategorisedAccount *shared.CategorisedAccount
	ContentType        string
	StatusCode         int
}
