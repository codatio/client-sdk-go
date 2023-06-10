# ItemStatus

Current state of the item, either:

- `Active`: Available for use
- `Archived`: Unavailable
- `Unknown`

Due to a [limitation in Xero's API](https://docs.codat.io/integrations/accounting/xero/xero-faq#why-do-all-of-my-items-from-xero-have-their-status-as-unknown), all items from Xero are mapped as `Unknown`.


## Values

| Name                 | Value                |
| -------------------- | -------------------- |
| `ItemStatusUnknown`  | Unknown              |
| `ItemStatusActive`   | Active               |
| `ItemStatusArchived` | Archived             |