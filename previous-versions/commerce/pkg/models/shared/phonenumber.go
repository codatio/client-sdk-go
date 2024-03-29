// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type PhoneNumber struct {
	// A phone number.
	Number *string `json:"number"`
	// The type of phone number
	Type PhoneNumberType `json:"type"`
}

func (o *PhoneNumber) GetNumber() *string {
	if o == nil {
		return nil
	}
	return o.Number
}

func (o *PhoneNumber) GetType() PhoneNumberType {
	if o == nil {
		return PhoneNumberType("")
	}
	return o.Type
}
