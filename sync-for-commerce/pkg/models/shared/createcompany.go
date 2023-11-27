// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type CreateCompany struct {
	// Additional information about the company. This can be used to store foreign IDs, references, etc.
	Description *string `json:"description,omitempty"`
	// Name of company being connected.
	Name string `json:"name"`
}

func (o *CreateCompany) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *CreateCompany) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}
