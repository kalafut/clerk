package clerk

type TransactionReader interface {
	Read(root *Account) []Transaction
}
