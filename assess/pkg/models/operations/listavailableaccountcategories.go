package operations

import (
	"net/http"
)

type ListAvailableAccountCategoriesChartOfAccountCategory struct {
	DetailType            *string `json:"detailType,omitempty"`
	DetailTypeDescription *string `json:"detailTypeDescription,omitempty"`
	DetailTypeDisplayName *string `json:"detailTypeDisplayName,omitempty"`
	Subtype               *string `json:"subtype,omitempty"`
	SubtypeDisplayName    *string `json:"subtypeDisplayName,omitempty"`
	Type                  *string `json:"type,omitempty"`
}

type ListAvailableAccountCategoriesResponse struct {
	ContentType                                               string
	StatusCode                                                int
	RawResponse                                               *http.Response
	ListAvailableAccountCategoriesChartOfAccountCategoryAnies []ListAvailableAccountCategoriesChartOfAccountCategory
}
