package main

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"

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

func readCSV(filename string) [][]string {
	transfers := map[string]struct {
		src  string
		dest string
		amnt string
	}{}
	_ = transfers

	f, err := os.Open(filename)
	defer f.Close()

	if err != nil {
		log.Fatal(err)
	}

	//reader := bufio.NewReader(f)
	reader := csv.NewReader(f)
	rows, err := reader.ReadAll()

	if err != nil {
		log.Fatal(err)
	}

	return rows[1:]
}

func ynabRowConv(row []string) []string {
	var ledger []string
	var amt string

	srcAcct := row[0]
	destAcct := row[5]
	payee := row[4]
	outflow := row[9]
	inflow := row[10]

	// normalize amounts
	if outflow != "$0.00" && inflow != "$0.00" {
		log.Fatal("Both inflow and outflow?!?")
	}

	if inflow != "$0.00" {
		amt = "-" + inflow
	} else {
		amt = outflow
	}

	// Handle transfers
	if strings.HasPrefix(payee, "Transfer : ") {
		destAcct = payee[11:]
		payee = ""
	}

	if destAcct == "" {
		return ledger
	}

	date := row[3]
	dateConverted := date[6:10] + "/" + date[0:2] + "/" + date[3:5]

	ledger = append(ledger, fmt.Sprintf("%s  %s", dateConverted, payee))
	ledger = append(ledger, fmt.Sprintf("    %s    %s", destAcct, amt))
	ledger = append(ledger, fmt.Sprintf("    %s", srcAcct))

	return ledger
}
