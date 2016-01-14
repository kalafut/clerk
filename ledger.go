package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"sort"
)

type Ledger struct {
	blocks       []Block
	transactions []Transaction
}

func NewLedger(data io.Reader) Ledger {
	transactions := ParseTransactions(data)
	return Ledger{
		transactions: transactions,
	}
}

func NewLedgerFromFile(filename string) Ledger {
	f, err := os.Open(filename)
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}
	return NewLedger(f)
}

func (l *Ledger) Sort() {
	sort.Sort(ByDate(l.blocks))
}

func (l Ledger) Export(w io.Writer) {
	for _, t := range l.transactions {
		fmt.Fprint(w, t.toCSV())
	}
}
