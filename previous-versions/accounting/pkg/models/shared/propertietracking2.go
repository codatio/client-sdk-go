// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type PropertieTracking2 struct {
	RecordRefs []InvoiceTo `json:"recordRefs,omitempty"`
}

func (o *PropertieTracking2) GetRecordRefs() []InvoiceTo {
	if o == nil {
		return nil
	}
	return o.RecordRefs
}
