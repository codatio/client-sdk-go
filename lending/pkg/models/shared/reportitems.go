// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/codatio/client-sdk-go/lending/v8/pkg/utils"
	"github.com/ericlagergren/decimal"
)

// LoanTransactionType - The type of loan transaction.
type LoanTransactionType string

const (
	LoanTransactionTypeInvestment      LoanTransactionType = "Investment"
	LoanTransactionTypeRepayment       LoanTransactionType = "Repayment"
	LoanTransactionTypeInterest        LoanTransactionType = "Interest"
	LoanTransactionTypeAccuredInterest LoanTransactionType = "AccuredInterest"
)

func (e LoanTransactionType) ToPointer() *LoanTransactionType {
	return &e
}
func (e *LoanTransactionType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Investment":
		fallthrough
	case "Repayment":
		fallthrough
	case "Interest":
		fallthrough
	case "AccuredInterest":
		*e = LoanTransactionType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LoanTransactionType: %v", v)
	}
}

type ReportItems struct {
	// The loan transaction amount.
	Amount *decimal.Big `decimal:"number" json:"amount,omitempty"`
	// In Codat's data model, dates and times are represented using the <a class="external" href="https://en.wikipedia.org/wiki/ISO_8601" target="_blank">ISO 8601 standard</a>. Date and time fields are formatted as strings; for example:
	//
	// ```
	// 2020-10-08T22:40:50Z
	// 2021-01-01T00:00:00
	// ```
	//
	//
	//
	// When syncing data that contains `DateTime` fields from Codat, make sure you support the following cases when reading time information:
	//
	// - Coordinated Universal Time (UTC): `2021-11-15T06:00:00Z`
	// - Unqualified local time: `2021-11-15T01:00:00`
	// - UTC time offsets: `2021-11-15T01:00:00-05:00`
	//
	// > Time zones
	// >
	// > Not all dates from Codat will contain information about time zones.
	// > Where it is not available from the underlying platform, Codat will return these as times local to the business whose data has been synced.
	Date    *string  `json:"date,omitempty"`
	ItemRef *ItemRef `json:"itemRef,omitempty"`
	// The name of lender providing the loan.
	Lender  *string  `json:"lender,omitempty"`
	LoanRef *LoanRef `json:"loanRef,omitempty"`
	// The type of loan transaction.
	LoanTransactionType *LoanTransactionType `json:"loanTransactionType,omitempty"`
}

func (r ReportItems) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *ReportItems) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ReportItems) GetAmount() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.Amount
}

func (o *ReportItems) GetDate() *string {
	if o == nil {
		return nil
	}
	return o.Date
}

func (o *ReportItems) GetItemRef() *ItemRef {
	if o == nil {
		return nil
	}
	return o.ItemRef
}

func (o *ReportItems) GetLender() *string {
	if o == nil {
		return nil
	}
	return o.Lender
}

func (o *ReportItems) GetLoanRef() *LoanRef {
	if o == nil {
		return nil
	}
	return o.LoanRef
}

func (o *ReportItems) GetLoanTransactionType() *LoanTransactionType {
	if o == nil {
		return nil
	}
	return o.LoanTransactionType
}
