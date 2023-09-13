// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package codatlending

type transactions struct {
	AccountTransactions *transactionsAccountTransactions
	DirectCosts         *transactionsDirectCosts
	JournalEntries      *transactionsJournalEntries
	Journals            *transactionsJournals
	Transfers           *transactionsTransfers

	sdkConfiguration sdkConfiguration
}

func newTransactions(sdkConfig sdkConfiguration) *transactions {
	return &transactions{
		sdkConfiguration:    sdkConfig,
		AccountTransactions: newTransactionsAccountTransactions(sdkConfig),
		DirectCosts:         newTransactionsDirectCosts(sdkConfig),
		JournalEntries:      newTransactionsJournalEntries(sdkConfig),
		Journals:            newTransactionsJournals(sdkConfig),
		Transfers:           newTransactionsTransfers(sdkConfig),
	}
}
