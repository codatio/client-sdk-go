// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type PushValidationInfo struct {
	Information []PushFieldValidation `json:"information,omitempty"`
	Warnings    []PushFieldValidation `json:"warnings,omitempty"`
}

func (o *PushValidationInfo) GetInformation() []PushFieldValidation {
	if o == nil {
		return nil
	}
	return o.Information
}

func (o *PushValidationInfo) GetWarnings() []PushFieldValidation {
	if o == nil {
		return nil
	}
	return o.Warnings
}
