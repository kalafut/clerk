package main

import (
	"fmt"
	"io"
	"sort"
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
	ldg.sort()
}

func (ldg Ledger) Export(w io.Writer) {
	for _, t := range ldg.transactions {
		fmt.Fprint(w, t.toCSV())
	}
}

func (ldg *Ledger) Add(t *Transaction) {
	ldg.transactions = append(ldg.transactions, t)
	ldg.sort()
}

func (ldg *Ledger) Del(t *Transaction) bool {
	a := ldg.transactions
	for i, v := range ldg.transactions {
		if v == t {
			a, a[len(a)-1] = append(a[:i], a[i+1:]...), nil
			ldg.transactions = a
			return true
		}
	}

	return false
}

func (ldg *Ledger) Replace(dst, src *Transaction) bool {
	if ldg.Del(dst) {
		ldg.Add(src)
		return true
	}

	return false
}

// Sorting support

type byDate []*Transaction

func (a byDate) Len() int      { return len(a) }
func (a byDate) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a byDate) Less(i, j int) bool {
	if a[i].Date().Before(a[j].Date()) {
		return true
	}
	if a[i].Date().Equal(a[j].Date()) && a[i].Summary() < a[j].Summary() {
		return true
	}
	return false
}

func (ldg *Ledger) sort() {
	sort.Sort(byDate(ldg.transactions))
}
