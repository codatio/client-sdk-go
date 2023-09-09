// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// AccountingJournalEntry - > **Language tip:** For the top-level record of a company's financial transactions, refer to the [Journals](https://docs.codat.io/accounting-api#/schemas/Journal) data type
//
// > View the coverage for journal entries in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=journalEntries" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// A journal entry report shows the entries made in a company's general ledger, or [accounts](https://docs.codat.io/accounting-api#/schemas/Account), when transactions are approved. The journal line items for each journal entry should balance.
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
type AccountingJournalEntry struct {
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
	// Links the current record to the underlying record or data type that created it.
	//
	// For example, if a journal entry is generated based on an invoice, this property allows you to connect the journal entry to the underlying invoice in our data model.
	RecordRef          *RecordRef `json:"recordRef,omitempty"`
	SourceModifiedDate *string    `json:"sourceModifiedDate,omitempty"`
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

func (o *AccountingJournalEntry) GetCreatedOn() *string {
	if o == nil {
		return nil
	}
	return o.CreatedOn
}

func (o *AccountingJournalEntry) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *AccountingJournalEntry) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *AccountingJournalEntry) GetJournalLines() []JournalLine {
	if o == nil {
		return nil
	}
	return o.JournalLines
}

func (o *AccountingJournalEntry) GetJournalRef() *JournalRef {
	if o == nil {
		return nil
	}
	return o.JournalRef
}

func (o *AccountingJournalEntry) GetMetadata() *Metadata {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *AccountingJournalEntry) GetModifiedDate() *string {
	if o == nil {
		return nil
	}
	return o.ModifiedDate
}

func (o *AccountingJournalEntry) GetPostedOn() *string {
	if o == nil {
		return nil
	}
	return o.PostedOn
}

func (o *AccountingJournalEntry) GetRecordRef() *RecordRef {
	if o == nil {
		return nil
	}
	return o.RecordRef
}

func (o *AccountingJournalEntry) GetSourceModifiedDate() *string {
	if o == nil {
		return nil
	}
	return o.SourceModifiedDate
}

func (o *AccountingJournalEntry) GetSupplementalData() *SupplementalData {
	if o == nil {
		return nil
	}
	return o.SupplementalData
}

func (o *AccountingJournalEntry) GetUpdatedOn() *string {
	if o == nil {
		return nil
	}
	return o.UpdatedOn
}
