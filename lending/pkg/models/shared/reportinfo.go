// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// ReportInfo - Report additional information, which is specific to Lending API reports.
type ReportInfo struct {
	// The name of the company being queried.
	CompanyName *string `json:"companyName,omitempty"`
	// Date the report was generated.
	GeneratedDate *string `json:"generatedDate,omitempty"`
	// The number of the page queried.
	PageNumber *int64 `json:"pageNumber,omitempty"`
	// The number of transactions returned per page.
	PageSize *int64 `json:"pageSize,omitempty"`
	// Name of the report.
	ReportName *string `json:"reportName,omitempty"`
	// The total number of transactions available for a company for the period specified in the query string.
	TotalResults *int64 `json:"totalResults,omitempty"`
}

func (o *ReportInfo) GetCompanyName() *string {
	if o == nil {
		return nil
	}
	return o.CompanyName
}

func (o *ReportInfo) GetGeneratedDate() *string {
	if o == nil {
		return nil
	}
	return o.GeneratedDate
}

func (o *ReportInfo) GetPageNumber() *int64 {
	if o == nil {
		return nil
	}
	return o.PageNumber
}

func (o *ReportInfo) GetPageSize() *int64 {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *ReportInfo) GetReportName() *string {
	if o == nil {
		return nil
	}
	return o.ReportName
}

func (o *ReportInfo) GetTotalResults() *int64 {
	if o == nil {
		return nil
	}
	return o.TotalResults
}
