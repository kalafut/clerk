package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

//type ynabRow struct {

//"Account","Flag","Check Number","Date","Payee","Category","Master Category","Sub Category","Memo","Outflow","Inflow","Cleared","Running Balance"

func main() {
	flag.Parse()
	filename := flag.Arg(0)
	rows := readCSV(filename)

	for _, row := range rows {
		conv := ynabRowConv(row)
		for _, r := range conv {
			fmt.Println(r)
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
