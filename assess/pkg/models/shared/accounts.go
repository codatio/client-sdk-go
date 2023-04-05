// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type Accounts struct {
	AccountName     *string    `json:"accountName,omitempty"`
	AccountProvider *string    `json:"accountProvider,omitempty"`
	AccountType     *string    `json:"accountType,omitempty"`
	Currency        *string    `json:"currency,omitempty"`
	CurrentBalance  *float64   `json:"currentBalance,omitempty"`
	PlatformName    *string    `json:"platformName,omitempty"`
	SourceRef       *SourceRef `json:"sourceRef,omitempty"`
}