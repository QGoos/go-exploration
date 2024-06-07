package iteration

type Transaction struct {
	From string
	To   string
	Sum  float64
}

type Account struct {
	Name    string
	Balance float64
}

func NewTransaction(from, to Account, sum float64) Transaction {
	return Transaction{From: from.Name, To: to.Name, Sum: sum}
}

func NewBalanceFor(account Account, transactions []Transaction) Account {
	return Reduce(
		transactions,
		applyTransaction,
		account,
	)
}

func applyTransaction(account Account, transaction Transaction) Account {
	if transaction.From == account.Name {
		account.Balance -= transaction.Sum
	}
	if transaction.To == account.Name {
		account.Balance += transaction.Sum
	}
	return account
}

// func NewBalanceFor(transactions []Transaction, name string) float64 {
// 	adjustBalance := func(currentBalance float64, t Transaction) float64 {
// 		if t.From == name {
// 			return currentBalance - t.Sum
// 		}
// 		if t.To == name {
// 			return currentBalance + t.Sum
// 		}
// 		return currentBalance
// 	}
// 	return Reduce(transactions, adjustBalance, 0.0)
// }
