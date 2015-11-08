package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"

	"gopkg.in/alecthomas/kingpin.v2"
)

//type ynabRow struct {

//"Account","Flag","Check Number","Date","Payee","Category","Master Category","Sub Category","Memo","Outflow","Inflow","Cleared","Running Balance"

var (
	app       = kingpin.New("clerk", "Ledger Helper")
	filename  = app.Flag("filename", "Ledger filename").Short('f').Default("master.dat").String()
	inplace   = app.Flag("inplace", "Edit file in place").Short('i').Bool()
	outfile   = app.Flag("outfile", "Output file").Short('o').String()
	sortCmd   = app.Command("sort", "Sort the ledger by date.")
	dedupeCmd = app.Command("dedupe", "Deduplicate the ledger.")
	importCmd = app.Command("import", "Import from external sources.")
)

func main() {
	var f *os.File
	var output *bufio.Writer
	var tempBuffer bytes.Buffer

	cmd := kingpin.MustParse(app.Parse(os.Args[1:]))
	f, err := os.Open(*filename)
	if err != nil {
		log.Fatal(err)
	}
	ledger := NewLedger(f)
	f.Close()

	if *inplace {
		output = bufio.NewWriter(&tempBuffer)
	} else if *outfile != "" {
		f, err = os.Create(*outfile)
		defer f.Close()
		if err != nil {
			log.Fatal(err)
		}
		output = bufio.NewWriter(f)
	} else {
		output = bufio.NewWriter(os.Stdout)
	}

	defer output.Flush()

	switch cmd {
	case sortCmd.FullCommand():
		ledger.Sort()
		ledger.Export(output)
	case dedupeCmd.FullCommand():
		findDupes(ledger.blocks)
	}

	if *inplace {
		f, err = os.Create(*filename)
		if err != nil {
			log.Fatal(err)
		}
		output.Flush()
		f.Write(tempBuffer.Bytes())
		f.Close()
	}
}

// findDupes returns a list of likely duplicate blocks. Duplicates
// are block with the same date and transaction structure. The same
// accounts and amounts must be present in both for it to be dupe.
func findDupes(blocks []Block) {
	for i := range blocks {
		for j := i + 1; j < len(blocks); j++ {
			if blocks[i].IsDupe(blocks[j], 0) {
				fmt.Printf("%v,%v:%v\n", i, j, blocks[i].lines[0])
			}
		}
	}
}
