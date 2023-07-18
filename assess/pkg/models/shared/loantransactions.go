// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// LoanTransactions - OK
type LoanTransactions struct {
	// If there are no errors, an empty array is returned.
	Errors     []interface{}               `json:"errors,omitempty"`
	ReportInfo *LoanTransactionsReportInfo `json:"reportInfo,omitempty"`
	// Contains object of reporting properties. The loan ref will reference a different object depending on the integration type.
	ReportItems []ReportItems1 `json:"reportItems,omitempty"`
}

func (o *LoanTransactions) GetErrors() []interface{} {
	if o == nil {
		return nil
	}
	return o.Errors
}

func (o *LoanTransactions) GetReportInfo() *LoanTransactionsReportInfo {
	if o == nil {
		return nil
	}
	return o.ReportInfo
}

func (o *LoanTransactions) GetReportItems() []ReportItems1 {
	if o == nil {
		return nil
	}
	return o.ReportItems
}
