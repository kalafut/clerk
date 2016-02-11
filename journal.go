package main

import "sort"

// Journal is the highest level container, containing transactions and all related
// accounts and commodities.
type Journal struct {
	rootAcct *Account
	txs      []*Tx
}

func NewJournal() Journal {
	return Journal{
		rootAcct: NewRootAccount(),
		txs:      []*Tx{},
	}
}

func (jrnl *Journal) All() []*Tx {
	txn := make([]*Tx, len(jrnl.txs))
	copy(txn, jrnl.txs)
	return txn
}

// Load populates the Journal transaction log with the values from the TxnReader.
// The existing transaction log is replaced, so this function is typically used
// to an intialize a Journal.
func (jrnl *Journal) Load(r TxReader) {
	txn := r.Read(jrnl.rootAcct)
	jrnl.txs = make([]*Tx, len(txn))
	copy(jrnl.txs, txn)
	jrnl.sort()
}

func (jrnl Journal) Store(w TxWriter) {
	w.Write(jrnl.txs)
}

func (jrnl *Journal) Add(t *Tx) {
	jrnl.txs = append(jrnl.txs, t)
	jrnl.sort()
}

func (jrnl *Journal) Del(t *Tx) bool {
	a := jrnl.txs
	for i, v := range jrnl.txs {
		if v == t {
			a, a[len(a)-1] = append(a[:i], a[i+1:]...), nil
			jrnl.txs = a
			return true
		}
	}

	return false
}

func (jrnl *Journal) Replace(dst, src *Tx) bool {
	if jrnl.Del(dst) {
		jrnl.Add(src)
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

func (jrnl *Journal) sort() {
	sort.Sort(byDate(jrnl.txs))
}
