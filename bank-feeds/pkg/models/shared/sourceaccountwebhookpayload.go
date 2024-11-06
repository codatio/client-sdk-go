// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type SourceAccountWebhookPayload struct {
	// Unique identifier for your SMB in Codat.
	CompanyID *string `json:"companyId,omitempty"`
	// Unique identifier for a company's data connection.
	ConnectionID     *string           `json:"connectionId,omitempty"`
	ReferenceCompany *CompanyReference `json:"referenceCompany,omitempty"`
	// The target bank account in a supported accounting software for ingestion into a bank feed.
	SourceAccount *SourceAccount `json:"sourceAccount,omitempty"`
}

func (o *SourceAccountWebhookPayload) GetCompanyID() *string {
	if o == nil {
		return nil
	}
	return o.CompanyID
}

func (o *SourceAccountWebhookPayload) GetConnectionID() *string {
	if o == nil {
		return nil
	}
	return o.ConnectionID
}

func (o *SourceAccountWebhookPayload) GetReferenceCompany() *CompanyReference {
	if o == nil {
		return nil
	}
	return o.ReferenceCompany
}

func (o *SourceAccountWebhookPayload) GetSourceAccount() *SourceAccount {
	if o == nil {
		return nil
	}
	return o.SourceAccount
}
