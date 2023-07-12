// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SyncStartedWebhookData struct {
	// In Codat's data model, dates and times are represented using the <a class="external" href="https://en.wikipedia.org/wiki/ISO_8601" target="_blank">ISO 8601 standard</a>. Date and time fields are formatted as strings; for example:
	//
	// ```
	// 2020-10-08T22:40:50Z
	// 2021-01-01T00:00:00
	// ```
	//
	//
	//
	// When syncing data that contains `DateTime` fields from Codat, make sure you support the following cases when reading time information:
	//
	// - Coordinated Universal Time (UTC): `2021-11-15T06:00:00Z`
	// - Unqualified local time: `2021-11-15T01:00:00`
	// - UTC time offsets: `2021-11-15T01:00:00-05:00`
	//
	// > Time zones
	// >
	// > Not all dates from Codat will contain information about time zones.
	// > Where it is not available from the underlying platform, Codat will return these as times local to the business whose data has been synced.
	SyncDateRangeFinishUtc *string `json:"SyncDateRangeFinishUtc,omitempty"`
	// In Codat's data model, dates and times are represented using the <a class="external" href="https://en.wikipedia.org/wiki/ISO_8601" target="_blank">ISO 8601 standard</a>. Date and time fields are formatted as strings; for example:
	//
	// ```
	// 2020-10-08T22:40:50Z
	// 2021-01-01T00:00:00
	// ```
	//
	//
	//
	// When syncing data that contains `DateTime` fields from Codat, make sure you support the following cases when reading time information:
	//
	// - Coordinated Universal Time (UTC): `2021-11-15T06:00:00Z`
	// - Unqualified local time: `2021-11-15T01:00:00`
	// - UTC time offsets: `2021-11-15T01:00:00-05:00`
	//
	// > Time zones
	// >
	// > Not all dates from Codat will contain information about time zones.
	// > Where it is not available from the underlying platform, Codat will return these as times local to the business whose data has been synced.
	SyncDateRangeStartUtc *string `json:"SyncDateRangeStartUtc,omitempty"`
	SyncID                *string `json:"syncId,omitempty"`
	// The type of sync being performed.
	SyncType *string `json:"syncType,omitempty"`
}

// SyncStartedWebhook - Webhook request body used to notify that a sync has started.
type SyncStartedWebhook struct {
	// Unique identifier of the webhook event.
	AlertID *string `json:"AlertId,omitempty"`
	// Unique identifier for your client in Codat.
	ClientID *string `json:"ClientId,omitempty"`
	// Name of your client in Codat.
	ClientName *string `json:"ClientName,omitempty"`
	// Unique identifier for your SMB in Codat.
	CompanyID *string                 `json:"CompanyId,omitempty"`
	Data      *SyncStartedWebhookData `json:"Data,omitempty"`
	// A human readable message about the webhook.
	Message *string `json:"Message,omitempty"`
	// Unique identifier for the rule.
	RuleID *string `json:"RuleId,omitempty"`
	// The type of rule.
	RuleType *string `json:"RuleType,omitempty"`
}