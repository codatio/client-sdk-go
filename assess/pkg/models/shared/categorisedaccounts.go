// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// CategorisedAccounts - OK
type CategorisedAccounts struct {
	Links      Links `json:"_links"`
	PageNumber int64 `json:"pageNumber"`
	PageSize   int64 `json:"pageSize"`
	// A list confirmed and suggested account categories.
	Results      []CategorisedAccount `json:"results,omitempty"`
	TotalResults int64                `json:"totalResults"`
}