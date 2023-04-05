// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ReportInfo - Report additional information, which is specific to Assess reports
type ReportInfo struct {
	// Company name
	CompanyName *string `json:"companyName,omitempty"`
	// Date the report was generated
	GeneratedDate *string `json:"generatedDate,omitempty"`
	PageNumber    *int64  `json:"pageNumber,omitempty"`
	PageSize      *int64  `json:"pageSize,omitempty"`
	// Name of the report
	ReportName   *string `json:"reportName,omitempty"`
	TotalResults *int64  `json:"totalResults,omitempty"`
}