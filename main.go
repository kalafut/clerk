package main

import (
	"bufio"
	"bytes"
	"log"
	"os"

	"gopkg.in/alecthomas/kingpin.v2"
)

//type ynabRow struct {

//"Account","Flag","Check Number","Date","Payee","Category","Master Category","Sub Category","Memo","Outflow","Inflow","Cleared","Running Balance"

var (
	app        = kingpin.New("clerk", "clerk Helper")
	ledgerFile = app.Flag("filename", "clerk filename").Short('f').Default("master.dat").String()
	importFile = app.Flag("csv", "CSV filename").String()
	inplace    = app.Flag("inplace", "Edit file in place").Short('i').Bool()
	outfile    = app.Flag("outfile", "Output file").Short('o').String()
	sortCmd    = app.Command("sort", "Sort the ledger by date.")
	dedupeCmd  = app.Command("dedupe", "Deduplicate the ledger.")
	importCmd  = app.Command("import", "Import from external sources.")
)

func main() {
	var f *os.File
	var output *bufio.Writer
	var tempBuffer bytes.Buffer

	cmd := kingpin.MustParse(app.Parse(os.Args[1:]))
	f, err := os.Open(*ledgerFile)
	if err != nil {
		log.Fatal(err)
	}
	ledger := NewLedgerReader(f)
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
		//Sort()
		//Export(output)
		_ = ledger
	case dedupeCmd.FullCommand():
		//FindDupes(ledger)
	case importCmd.FullCommand():
		//Import(*ledgerFile, *importFile)
	}

	if *inplace {
		f, err = os.Create(*ledgerFile)
		if err != nil {
			log.Fatal(err)
		}
		output.Flush()
		f.Write(tempBuffer.Bytes())
		f.Close()
	}
}
