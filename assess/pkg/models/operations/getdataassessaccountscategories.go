package operations

type GetDataAssessAccountsCategoriesChartOfAccountCategory struct {
	DetailType            *string `json:"detailType,omitempty"`
	DetailTypeDescription *string `json:"detailTypeDescription,omitempty"`
	DetailTypeDisplayName *string `json:"detailTypeDisplayName,omitempty"`
	Subtype               *string `json:"subtype,omitempty"`
	SubtypeDisplayName    *string `json:"subtypeDisplayName,omitempty"`
	Type                  *string `json:"type,omitempty"`
}

type GetDataAssessAccountsCategoriesResponse struct {
	ContentType                                                 string
	StatusCode                                                  int
	GetDataAssessAccountsCategoriesChartOfAccountCategoryAllOfs []GetDataAssessAccountsCategoriesChartOfAccountCategory
}
