// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// AccountBalances - Success
type AccountBalances struct {
	Links        Links            `json:"_links"`
	PageNumber   int64            `json:"pageNumber"`
	PageSize     int64            `json:"pageSize"`
	Results      []AccountBalance `json:"results,omitempty"`
	TotalResults int64            `json:"totalResults"`
}