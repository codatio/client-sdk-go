# AccountingJournalEntryRecordReference

Links the current record to the underlying record or data type that created it. 

For example, if a journal entry is generated based on an invoice, this property allows you to connect the journal entry to the underlying invoice in our data model. 


## Fields

| Field                                       | Type                                        | Required                                    | Description                                 | Example                                     |
| ------------------------------------------- | ------------------------------------------- | ------------------------------------------- | ------------------------------------------- | ------------------------------------------- |
| `DataType`                                  | **string*                                   | :heavy_minus_sign:                          | Allowed name of the 'dataType'.             | journalEntry                                |
| `ID`                                        | **string*                                   | :heavy_minus_sign:                          | 'id' of the underlying record or data type. |                                             |