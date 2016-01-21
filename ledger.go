package main

import (
	"fmt"
	"io"
)

// Ledger is the highest level container, containing transactions and all related
// accounts and commodities.
type Ledger struct {
	rootAccount  *Account
	transactions []*Transaction
}

type TxnReader interface {
	Read(root *Account) []*Transaction
}

func NewLedger() Ledger {
	return Ledger{
		rootAccount:  NewRootAccount(),
		transactions: []*Transaction{},
	}
}

func (ldg *Ledger) All() []*Transaction {
	txn := make([]*Transaction, len(ldg.transactions))
	copy(txn, ldg.transactions)
	return txn
}

// Load populates the Ledger transaction log with the values from the TxnReader.
// The existing transaction log is replaced, so this function is typically used
// to an intialize a Ledger.
func (ldg *Ledger) Load(r TxnReader) {
	txn := r.Read(ldg.rootAccount)
	ldg.transactions = make([]*Transaction, len(txn))
	copy(ldg.transactions, txn)
	//sort
}

func (l *Ledger) Sort() {
	//sort.Sort(ByDate(l.blocks))
}

func (l Ledger) Export(w io.Writer) {
	for _, t := range l.transactions {
		fmt.Fprint(w, t.toCSV())
	}
}

func (ldg *Ledger) Add(t *Transaction) {
	ldg.transactions = append(ldg.transactions, t)
}

func (ldg *Ledger) Del(t *Transaction) {
	for _, v := range ldg.transactions {
		if v.id == t.id {

		}
	}
}
