package shared

// BankTransactions
// > **Accessing Bank Accounts through Banking API**
// >
// > This datatype was originally used for accessing bank account data both in accounting integrations and open banking aggregators.
// >
// > To view bank account data through the Banking API, please refer to the new datatype [here](https://docs.codat.io/docs/data-model-banking-banking-transactions)
//
// > View the coverage for bank transactions in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=bankTransactions" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// Transactional banking data for a specific company and account.
//
// Bank transactions include the:
// * Amount of the transaction.
// * Current account balance.
// * Transaction type, for example, credit, debit, or transfer.
type BankTransactions struct {
	AccountID       *string             `json:"accountId,omitempty"`
	ContractVersion *string             `json:"contractVersion,omitempty"`
	Transactions    []BankStatementLine `json:"transactions,omitempty"`
}
