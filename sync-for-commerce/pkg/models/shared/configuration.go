// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// Configuration - Success
type Configuration struct {
	Fees        *Fees        `json:"fees,omitempty"`
	NewPayments *NewPayments `json:"newPayments,omitempty"`
	Payments    *Payments    `json:"payments,omitempty"`
	Sales       *Sales       `json:"sales,omitempty"`
}

func (o *Configuration) GetFees() *Fees {
	if o == nil {
		return nil
	}
	return o.Fees
}

func (o *Configuration) GetNewPayments() *NewPayments {
	if o == nil {
		return nil
	}
	return o.NewPayments
}

func (o *Configuration) GetPayments() *Payments {
	if o == nil {
		return nil
	}
	return o.Payments
}

func (o *Configuration) GetSales() *Sales {
	if o == nil {
		return nil
	}
	return o.Sales
}