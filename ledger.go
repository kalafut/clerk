package clerk

import (
	"fmt"
	"io"
	"log"
	"os"
	"sort"
)

type Ledger struct {
	blocks []Block
}

func NewLedger(data io.Reader) Ledger {
	blocks := ParseLines(data)
	return Ledger{
		blocks: blocks,
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
	for _, b := range l.blocks {
		for _, line := range b.lines {
			fmt.Fprintln(w, line)
		}
		fmt.Fprintln(w)
	}
}
