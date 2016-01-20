package main

type TransactionReader interface {
	Read(root *Account) []Transaction
}
