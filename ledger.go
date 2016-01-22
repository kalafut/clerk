package main

import (
	"fmt"
	"io"
	"sort"
)

// Ledger is the highest level container, containing transactions and all related
// accounts and commodities.
type Ledger struct {
	rootAcct *Account
	txs      []*Tx
}

func NewLedger() Ledger {
	return Ledger{
		rootAcct: NewRootAccount(),
		txs:      []*Tx{},
	}
}

func (ldg *Ledger) All() []*Tx {
	txn := make([]*Tx, len(ldg.txs))
	copy(txn, ldg.txs)
	return txn
}

// Load populates the Ledger transaction log with the values from the TxnReader.
// The existing transaction log is replaced, so this function is typically used
// to an intialize a Ledger.
func (ldg *Ledger) Load(r TxReader) {
	txn := r.Read(ldg.rootAcct)
	ldg.txs = make([]*Tx, len(txn))
	copy(ldg.txs, txn)
	ldg.sort()
}

func (ldg Ledger) Export(w io.Writer) {
	for _, t := range ldg.txs {
		fmt.Fprint(w, t.toCSV())
	}
}

func (ldg *Ledger) Add(t *Tx) {
	ldg.txs = append(ldg.txs, t)
	ldg.sort()
}

func (ldg *Ledger) Del(t *Tx) bool {
	a := ldg.txs
	for i, v := range ldg.txs {
		if v == t {
			a, a[len(a)-1] = append(a[:i], a[i+1:]...), nil
			ldg.txs = a
			return true
		}
	}

	return false
}

func (ldg *Ledger) Replace(dst, src *Tx) bool {
	if ldg.Del(dst) {
		ldg.Add(src)
		return true
	}

	return false
}

// Sorting support

type byDate []*Tx

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
	sort.Sort(byDate(ldg.txs))
}
