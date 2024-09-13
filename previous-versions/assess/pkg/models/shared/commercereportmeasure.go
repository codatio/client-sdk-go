// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type CommerceReportMeasure struct {
	// The measure's display name.
	DisplayName *string `json:"displayName,omitempty"`
	// The measure's index.
	Index *int64 `json:"index,omitempty"`
	// The measure's type.
	Type *string `json:"type,omitempty"`
	// The measure's units e.g. percentage (%).
	Units *string `json:"units,omitempty"`
}

func (o *CommerceReportMeasure) GetDisplayName() *string {
	if o == nil {
		return nil
	}
	return o.DisplayName
}

func (o *CommerceReportMeasure) GetIndex() *int64 {
	if o == nil {
		return nil
	}
	return o.Index
}

func (o *CommerceReportMeasure) GetType() *string {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *CommerceReportMeasure) GetUnits() *string {
	if o == nil {
		return nil
	}
	return o.Units
}
