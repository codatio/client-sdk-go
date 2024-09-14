// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/codatio/client-sdk-go/sync-for-expenses/v4/pkg/utils"
	"github.com/ericlagergren/decimal"
)

// AccountReference - Reference of the account you are transferring money from.
type AccountReference struct {
	// 'id' from the Accounts data type.
	ID string `json:"id"`
}

func (o *AccountReference) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type From struct {
	// Reference of the account you are transferring money from.
	AccountRef AccountReference `json:"accountRef"`
	// Amount that has been transferred from the account in the native currency of the account.
	Amount *decimal.Big `decimal:"number" json:"amount"`
}

func (f From) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(f, "", false)
}

func (f *From) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &f, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *From) GetAccountRef() AccountReference {
	if o == nil {
		return AccountReference{}
	}
	return o.AccountRef
}

func (o *From) GetAmount() *decimal.Big {
	if o == nil {
		return new(decimal.Big).SetFloat64(0.0)
	}
	return o.Amount
}

// TransferTransactionRequestAccountReference - Reference of the account you are transferring money to.
type TransferTransactionRequestAccountReference struct {
	// 'id' from the Accounts data type.
	ID string `json:"id"`
}

func (o *TransferTransactionRequestAccountReference) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type To struct {
	// Reference of the account you are transferring money to.
	AccountRef TransferTransactionRequestAccountReference `json:"accountRef"`
	// Amount that has been transferred to the account in the native currency of the account.
	Amount *decimal.Big `decimal:"number" json:"amount"`
}

func (t To) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *To) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *To) GetAccountRef() TransferTransactionRequestAccountReference {
	if o == nil {
		return TransferTransactionRequestAccountReference{}
	}
	return o.AccountRef
}

func (o *To) GetAmount() *decimal.Big {
	if o == nil {
		return new(decimal.Big).SetFloat64(0.0)
	}
	return o.Amount
}

type TransferTransactionRequest struct {
	Date string `json:"date"`
	// Any private, company notes about the transaction.
	Description *string `json:"description,omitempty"`
	From        From    `json:"from"`
	To          To      `json:"to"`
}

func (o *TransferTransactionRequest) GetDate() string {
	if o == nil {
		return ""
	}
	return o.Date
}

func (o *TransferTransactionRequest) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *TransferTransactionRequest) GetFrom() From {
	if o == nil {
		return From{}
	}
	return o.From
}

func (o *TransferTransactionRequest) GetTo() To {
	if o == nil {
		return To{}
	}
	return o.To
}
