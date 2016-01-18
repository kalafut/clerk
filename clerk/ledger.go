package clerk

import (
	"fmt"
	"io"
	"log"
	"os"
)

// Ledger is the highest level container, containing transactions and all related
// accounts and commodities.
type Ledger struct {
	rootAccount  *Account
	transactions []Transaction
}

func NewLedger(data io.Reader) Ledger {
	return Ledger{
		rootAccount:  NewRootAccount(),
		transactions: []Transaction{},
	}
}

func NewLedgerReader(data io.Reader) Ledger {
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
	return NewLedgerReader(f)
}

func (l *Ledger) Sort() {
	//sort.Sort(ByDate(l.blocks))
}

func (l Ledger) Export(w io.Writer) {
	for _, t := range l.transactions {
		fmt.Fprint(w, t.toCSV())
	}
}
