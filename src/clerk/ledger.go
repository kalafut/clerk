package main

import (
	"fmt"
	"io"
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
