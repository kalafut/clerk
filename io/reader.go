package io

import "github.com/kalafut/clerk/ledger"

type TransactionReader interface {
	Read(root *ledger.Account) []ledger.Transaction
}
