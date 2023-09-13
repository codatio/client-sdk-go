// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type Attachment struct {
	// Unique ID of company in Codat
	CompanyID *string `json:"companyId,omitempty"`
	// Unique identifier of attachment
	ID *string `json:"id,omitempty"`
	// Unique identifier of transaction
	TransactionID *string `json:"transactionId,omitempty"`
}

func (o *Attachment) GetCompanyID() *string {
	if o == nil {
		return nil
	}
	return o.CompanyID
}

func (o *Attachment) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Attachment) GetTransactionID() *string {
	if o == nil {
		return nil
	}
	return o.TransactionID
}
