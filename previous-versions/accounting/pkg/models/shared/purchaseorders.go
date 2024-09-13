// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type PurchaseOrders struct {
	Links Links `json:"_links"`
	// Current page number.
	PageNumber int64 `json:"pageNumber"`
	// Number of items to return in results array.
	PageSize int64           `json:"pageSize"`
	Results  []PurchaseOrder `json:"results,omitempty"`
	// Total number of items.
	TotalResults int64 `json:"totalResults"`
}

func (o *PurchaseOrders) GetLinks() Links {
	if o == nil {
		return Links{}
	}
	return o.Links
}

func (o *PurchaseOrders) GetPageNumber() int64 {
	if o == nil {
		return 0
	}
	return o.PageNumber
}

func (o *PurchaseOrders) GetPageSize() int64 {
	if o == nil {
		return 0
	}
	return o.PageSize
}

func (o *PurchaseOrders) GetResults() []PurchaseOrder {
	if o == nil {
		return nil
	}
	return o.Results
}

func (o *PurchaseOrders) GetTotalResults() int64 {
	if o == nil {
		return 0
	}
	return o.TotalResults
}
