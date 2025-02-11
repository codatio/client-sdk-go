// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// CompanyReferenceLinks - A collection of links for the company.
type CompanyReferenceLinks struct {
	// Link to the company page in the portal.
	Portal *string `json:"portal,omitempty"`
}

func (o *CompanyReferenceLinks) GetPortal() *string {
	if o == nil {
		return nil
	}
	return o.Portal
}

type CompanyReference struct {
	// Unique identifier for your SMB in Codat.
	ID *string `json:"id,omitempty"`
	// The name of the company
	Name *string `json:"name,omitempty"`
	// Additional information about the company. This can be used to store foreign IDs, references, etc.
	Description *string `json:"description,omitempty"`
	// A collection of links for the company.
	Links *CompanyReferenceLinks `json:"links,omitempty"`
	// A collection of user-defined key-value pairs that store custom metadata against the company.
	Tags map[string]string `json:"tags,omitempty"`
}

func (o *CompanyReference) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *CompanyReference) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *CompanyReference) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *CompanyReference) GetLinks() *CompanyReferenceLinks {
	if o == nil {
		return nil
	}
	return o.Links
}

func (o *CompanyReference) GetTags() map[string]string {
	if o == nil {
		return nil
	}
	return o.Tags
}
