// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package lending

type Banking struct {
	AccountBalances       *AccountBalances
	Accounts              *CodatLendingAccounts
	TransactionCategories *TransactionCategories
	Transactions          *CodatLendingBankingTransactions
	CategorizedStatement  *CategorizedStatement

	sdkConfiguration sdkConfiguration
}

func newBanking(sdkConfig sdkConfiguration) *Banking {
	return &Banking{
		sdkConfiguration:      sdkConfig,
		AccountBalances:       newAccountBalances(sdkConfig),
		Accounts:              newCodatLendingAccounts(sdkConfig),
		TransactionCategories: newTransactionCategories(sdkConfig),
		Transactions:          newCodatLendingBankingTransactions(sdkConfig),
		CategorizedStatement:  newCategorizedStatement(sdkConfig),
	}
}
