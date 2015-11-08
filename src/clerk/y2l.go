package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

//type ynabRow struct {

//"Account","Flag","Check Number","Date","Payee","Category","Master Category","Sub Category","Memo","Outflow","Inflow","Cleared","Running Balance"

var acct = flag.String("a", "", "filter account")

type Config struct {
	Ynab struct {
		AccountMappings map[string]string
	}
	UnknownAccount string
}

func _main() {
	flag.Parse()
	filename := flag.Arg(0)
	data, err := ioutil.ReadFile("y2l.json")

	if err != nil {
		log.Fatal(err)
	}

	cfg := parseConfig(data)

	rows := readCSV(filename)

	for _, row := range rows {
		if *acct == "" || row[0] == *acct {
			//println("Here")
			//println(row[0])
			conv := ynabRowConv(row, cfg)
			fmt.Println(conv)
		}
	}
}

func parseConfig(data []byte) Config {
	c := Config{}
	dec := json.NewDecoder(bytes.NewReader(data))
	err := dec.Decode(&c)
	if err != nil {
		log.Fatal(err)
	}

	return c
}

func readCSV(filename string) [][]string {
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

func ynabRowConv(row []string, cfg Config) string {
	var ledger string
	var amt string

	srcAcct := convertAccount(row[0], cfg)
	destAcct := convertAccount(row[5], cfg)
	payee := row[4]
	outflow := row[9]
	inflow := row[10]

	// Handle transfers
	if strings.HasPrefix(payee, "Transfer : ") {
		destAcct = convertAccount(payee[11:], cfg)
		payee = ""
	}

	if destAcct == "" {
		destAcct = cfg.UnknownAccount
	}

	if outflow != "$0.00" && inflow != "$0.00" {
		log.Fatal("Both inflow and outflow?!?")
	}

	if inflow != "$0.00" {
		amt = "-" + inflow
	} else {
		amt = outflow
	}

	date := row[3]
	dateConverted := date[6:10] + "/" + date[0:2] + "/" + date[3:5]

	ledger = fmt.Sprintf("%s  %s\n", dateConverted, payee)
	ledger += fmt.Sprintf("    %-30s    %s\n", destAcct, amt)
	ledger += fmt.Sprintf("    %s\n", srcAcct)

	return ledger
}

func convertAccount(acct string, cfg Config) string {
	if cnvt, ok := cfg.Ynab.AccountMappings[acct]; ok {
		return cnvt
	}
	return acct
}
