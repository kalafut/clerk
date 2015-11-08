package main

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
	"time"
)

type Ledger struct {
	blocks []Block
}

func NewLedger(data io.Reader) Ledger {
	blocks := parse(data)
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

// parseFile read a legder-formatted text file and returns a slice of blocks
func parse(data io.Reader) []Block {
	var blocks []Block

	scanner := bufio.NewScanner(data)

	block := Block{}
	for scanner.Scan() {
		line := scanner.Text()

		if len(strings.TrimSpace(line)) == 0 {
			if !block.Empty() {
				blocks = append(blocks, block)
				block = Block{}
			}
		} else {
			t, err := time.Parse("2006/01/02", line[0:10])
			if err == nil {
				// Start a new block
				if !block.Empty() {
					blocks = append(blocks, block)
					block = Block{}
				}
				block.date = t
			}
			block.lines = append(block.lines, line)
		}
	}

	if !block.Empty() {
		blocks = append(blocks, block)
	}

	return blocks
}
