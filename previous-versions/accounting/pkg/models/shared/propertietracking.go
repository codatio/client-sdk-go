// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// PropertieTracking - Categories, and a project and customer, against which the item is tracked.
type PropertieTracking struct {
	CategoryRefs []TrackingCategoryRef  `json:"categoryRefs"`
	CustomerRef  *AccountingCustomerRef `json:"customerRef,omitempty"`
	// Defines if the invoice or credit note is billed/rebilled to a project or customer.
	IsBilledTo BilledToType `json:"isBilledTo"`
	// Defines if the invoice or credit note is billed/rebilled to a project or customer.
	IsRebilledTo BilledToType `json:"isRebilledTo"`
	ProjectRef   *ProjectRef  `json:"projectRef,omitempty"`
}

func (o *PropertieTracking) GetCategoryRefs() []TrackingCategoryRef {
	if o == nil {
		return []TrackingCategoryRef{}
	}
	return o.CategoryRefs
}

func (o *PropertieTracking) GetCustomerRef() *AccountingCustomerRef {
	if o == nil {
		return nil
	}
	return o.CustomerRef
}

func (o *PropertieTracking) GetIsBilledTo() BilledToType {
	if o == nil {
		return BilledToType("")
	}
	return o.IsBilledTo
}

func (o *PropertieTracking) GetIsRebilledTo() BilledToType {
	if o == nil {
		return BilledToType("")
	}
	return o.IsRebilledTo
}

func (o *PropertieTracking) GetProjectRef() *ProjectRef {
	if o == nil {
		return nil
	}
	return o.ProjectRef
}
