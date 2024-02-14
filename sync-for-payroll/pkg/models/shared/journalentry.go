// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// JournalEntry - > **Language tip:** For the top-level record of a company's financial transactions, refer to the [Journals](https://docs.codat.io/sync-for-payroll-api#/schemas/Journal) data type
//
// > View the coverage for journal entries in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=journalEntries" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// A journal entry report shows the entries made in a company's general ledger, or [accounts](https://docs.codat.io/sync-for-payroll-api#/schemas/Account), when transactions are approved. The journal line items for each journal entry should balance.
//
// A journal entry line item is a single transaction line on the journal entry. For example:
//
// - When a journal entry is recording a receipt of cash, the credit to accounts receivable and the debit to cash are separate line items.
// - When a company needs to recognise revenue from an annual contract on a monthly basis, on receipt of cash for month one, they make a debit to deferred income and a credit to revenue.
//
// In Codat a journal entry contains details of:
//
// - The date on which the entry was created and posted.
// - Itemised lines, including amounts and currency.
// - A reference to the associated accounts.
// - A reference to the underlying record. For example, the invoice, bill, or other data type that triggered the posting of the journal entry to the general ledger.
//
// > **Pushing journal entries**
// > Codat only supports journal entries in the base currency of the company that are pushed into accounts denominated in the same base currency.
type JournalEntry struct {
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
	CreatedOn *string `json:"createdOn,omitempty"`
	// Optional description of the journal entry.
	Description *string `json:"description,omitempty"`
	// Unique identifier of the journal entry for the company in the accounting platform.
	ID *string `json:"id,omitempty"`
	// An array of journal lines.
	JournalLines []JournalLine `json:"journalLines,omitempty"`
	// Links journal entries to the relevant journal in accounting integrations that use multi-book accounting (multiple journals).
	JournalRef   *JournalRef `json:"journalRef,omitempty"`
	Metadata     *Metadata   `json:"metadata,omitempty"`
	ModifiedDate *string     `json:"modifiedDate,omitempty"`
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
	PostedOn *string `json:"postedOn,omitempty"`
	// Links a journal entry to the underlying record that created it.
	RecordRef          *JournalEntryRecordRef `json:"recordRef,omitempty"`
	SourceModifiedDate *string                `json:"sourceModifiedDate,omitempty"`
	// Supplemental data is additional data you can include in our standard data types.
	//
	// It is referenced as a configured dynamic key value pair that is unique to the accounting platform. [Learn more](https://docs.codat.io/using-the-api/supplemental-data/overview) about supplemental data.
	SupplementalData *SupplementalData `json:"supplementalData,omitempty"`
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
	UpdatedOn *string `json:"updatedOn,omitempty"`
}

func (o *JournalEntry) GetCreatedOn() *string {
	if o == nil {
		return nil
	}
	return o.CreatedOn
}

func (o *JournalEntry) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *JournalEntry) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *JournalEntry) GetJournalLines() []JournalLine {
	if o == nil {
		return nil
	}
	return o.JournalLines
}

func (o *JournalEntry) GetJournalRef() *JournalRef {
	if o == nil {
		return nil
	}
	return o.JournalRef
}

func (o *JournalEntry) GetMetadata() *Metadata {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *JournalEntry) GetModifiedDate() *string {
	if o == nil {
		return nil
	}
	return o.ModifiedDate
}

func (o *JournalEntry) GetPostedOn() *string {
	if o == nil {
		return nil
	}
	return o.PostedOn
}

func (o *JournalEntry) GetRecordRef() *JournalEntryRecordRef {
	if o == nil {
		return nil
	}
	return o.RecordRef
}

func (o *JournalEntry) GetSourceModifiedDate() *string {
	if o == nil {
		return nil
	}
	return o.SourceModifiedDate
}

func (o *JournalEntry) GetSupplementalData() *SupplementalData {
	if o == nil {
		return nil
	}
	return o.SupplementalData
}

func (o *JournalEntry) GetUpdatedOn() *string {
	if o == nil {
		return nil
	}
	return o.UpdatedOn
}
