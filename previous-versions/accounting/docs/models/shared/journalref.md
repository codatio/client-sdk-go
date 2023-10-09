# JournalRef

Links journal entries to the relevant journal in accounting integrations that use multi-book accounting (multiple journals).


## Fields

| Field                           | Type                            | Required                        | Description                     |
| ------------------------------- | ------------------------------- | ------------------------------- | ------------------------------- |
| `AdditionalProperties`          | map[string]*interface{}*        | :heavy_minus_sign:              | N/A                             |
| `ID`                            | *string*                        | :heavy_check_mark:              | GUID of the underlying journal. |
| `Name`                          | **string*                       | :heavy_minus_sign:              | Name of journal                 |