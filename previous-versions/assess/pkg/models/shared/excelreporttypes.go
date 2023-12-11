// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// ExcelReportTypes - The type of the report requested in the query string.
type ExcelReportTypes string

const (
	ExcelReportTypesAudit              ExcelReportTypes = "audit"
	ExcelReportTypesEnhancedFinancials ExcelReportTypes = "enhancedFinancials"
	ExcelReportTypesEnhancedInvoices   ExcelReportTypes = "enhancedInvoices"
	ExcelReportTypesEnhancedCashFlow   ExcelReportTypes = "enhancedCashFlow"
)

func (e ExcelReportTypes) ToPointer() *ExcelReportTypes {
	return &e
}

func (e *ExcelReportTypes) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "audit":
		fallthrough
	case "enhancedFinancials":
		fallthrough
	case "enhancedInvoices":
		fallthrough
	case "enhancedCashFlow":
		*e = ExcelReportTypes(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ExcelReportTypes: %v", v)
	}
}
