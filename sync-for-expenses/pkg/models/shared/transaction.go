// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/codatio/client-sdk-go/sync-for-expenses/v3/pkg/utils"
)

type Transaction struct {
	AdditionalProperties map[string]interface{} `additionalProperties:"true" json:"-"`
	// Type of transaction that has been processed e.g. Expense or Bank Feed.
	IntegrationType *IntegrationType `default:"expenses" json:"integrationType"`
	// Metadata such as validation errors or the resulting record created in the accounting software.
	Message *string `json:"message,omitempty"`
	// Status of the transaction.
	Status *TransactionStatus `json:"status,omitempty"`
	// Your unique idenfier of the transaction.
	TransactionID *string `json:"transactionId,omitempty"`
}

func (t Transaction) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *Transaction) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Transaction) GetAdditionalProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

func (o *Transaction) GetIntegrationType() *IntegrationType {
	if o == nil {
		return nil
	}
	return o.IntegrationType
}

func (o *Transaction) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *Transaction) GetStatus() *TransactionStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *Transaction) GetTransactionID() *string {
	if o == nil {
		return nil
	}
	return o.TransactionID
}
