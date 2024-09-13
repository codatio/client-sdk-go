// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type CommerceConfiguration struct {
	// The country code outlining where the company is based.
	CountryCode *string                   `json:"countryCode,omitempty"`
	Fees        *FeesConfiguration        `json:"fees,omitempty"`
	MapSettings *ConfigurationMapSettings `json:"mapSettings,omitempty"`
	NewPayments *NewPaymentsConfiguration `json:"newPayments,omitempty"`
	Payments    *PaymentsConfiguration    `json:"payments,omitempty"`
	Sales       *SalesConfiguration       `json:"sales,omitempty"`
}

func (o *CommerceConfiguration) GetCountryCode() *string {
	if o == nil {
		return nil
	}
	return o.CountryCode
}

func (o *CommerceConfiguration) GetFees() *FeesConfiguration {
	if o == nil {
		return nil
	}
	return o.Fees
}

func (o *CommerceConfiguration) GetMapSettings() *ConfigurationMapSettings {
	if o == nil {
		return nil
	}
	return o.MapSettings
}

func (o *CommerceConfiguration) GetNewPayments() *NewPaymentsConfiguration {
	if o == nil {
		return nil
	}
	return o.NewPayments
}

func (o *CommerceConfiguration) GetPayments() *PaymentsConfiguration {
	if o == nil {
		return nil
	}
	return o.Payments
}

func (o *CommerceConfiguration) GetSales() *SalesConfiguration {
	if o == nil {
		return nil
	}
	return o.Sales
}
