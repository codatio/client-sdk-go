// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// Propertiestracking - Categories, and a project and customer, against which the item is tracked.
type Propertiestracking struct {
	CategoryRefs []TrackingCategoryRef `json:"categoryRefs"`
	CustomerRef  *CustomerRef          `json:"customerRef,omitempty"`
	IsBilledTo   BilledToType          `json:"isBilledTo"`
	IsRebilledTo BilledToType          `json:"isRebilledTo"`
	ProjectRef   *ProjectRef           `json:"projectRef,omitempty"`
}

func (o *Propertiestracking) GetCategoryRefs() []TrackingCategoryRef {
	if o == nil {
		return []TrackingCategoryRef{}
	}
	return o.CategoryRefs
}

func (o *Propertiestracking) GetCustomerRef() *CustomerRef {
	if o == nil {
		return nil
	}
	return o.CustomerRef
}

func (o *Propertiestracking) GetIsBilledTo() BilledToType {
	if o == nil {
		return BilledToType("")
	}
	return o.IsBilledTo
}

func (o *Propertiestracking) GetIsRebilledTo() BilledToType {
	if o == nil {
		return BilledToType("")
	}
	return o.IsRebilledTo
}

func (o *Propertiestracking) GetProjectRef() *ProjectRef {
	if o == nil {
		return nil
	}
	return o.ProjectRef
}
