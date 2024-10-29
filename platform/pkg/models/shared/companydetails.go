// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type CompanyDetails struct {
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
	Created *string `json:"created,omitempty"`
	// Name of user that created the company in Codat.
	CreatedByUserName *string `json:"createdByUserName,omitempty"`
	// Additional information about the company. This can be used to store foreign IDs, references, etc.
	Description *string `json:"description,omitempty"`
	// Unique identifier for your SMB in Codat.
	ID string `json:"id"`
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
	LastSync *string `json:"lastSync,omitempty"`
	// The name of the company
	Name string `json:"name"`
	// An array of products that are currently enabled for the company.
	Products []string `json:"products,omitempty"`
	// The `redirect` [Link URL](https://docs.codat.io/auth-flow/authorize-hosted-link) enabling the customer to start their auth flow journey for the company.
	Redirect string `json:"redirect"`
	// A collection of user-defined key-value pairs that store custom metadata against the company.
	Tags map[string]string `json:"tags,omitempty"`
}

func (o *CompanyDetails) GetCreated() *string {
	if o == nil {
		return nil
	}
	return o.Created
}

func (o *CompanyDetails) GetCreatedByUserName() *string {
	if o == nil {
		return nil
	}
	return o.CreatedByUserName
}

func (o *CompanyDetails) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *CompanyDetails) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *CompanyDetails) GetLastSync() *string {
	if o == nil {
		return nil
	}
	return o.LastSync
}

func (o *CompanyDetails) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *CompanyDetails) GetProducts() []string {
	if o == nil {
		return nil
	}
	return o.Products
}

func (o *CompanyDetails) GetRedirect() string {
	if o == nil {
		return ""
	}
	return o.Redirect
}

func (o *CompanyDetails) GetTags() map[string]string {
	if o == nil {
		return nil
	}
	return o.Tags
}
