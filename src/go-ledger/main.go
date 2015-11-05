package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"time"
)

//type ynabRow struct {

//"Account","Flag","Check Number","Date","Payee","Category","Master Category","Sub Category","Memo","Outflow","Inflow","Cleared","Running Balance"

type Block struct {
	lines []string
	date  time.Time
}

var blocks []Block

type ByDate []Block

func (a ByDate) Len() int           { return len(a) }
func (a ByDate) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByDate) Less(i, j int) bool { return a[i].date.Before(a[j].date) }

func main() {
	flag.Parse()
	filename := flag.Arg(0)

	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)

	block := Block{}
	for scanner.Scan() {
		line := scanner.Text()

		if len(strings.TrimSpace(line)) == 0 {
			if len(block.lines) > 0 {
				blocks = append(blocks, block)
				block = Block{}
			}
		} else {
			t, err := time.Parse("2006/01/02", line[0:10])
			if err == nil {
				block.date = t
			}
			block.lines = append(block.lines, line)
		}
	}

	if len(block.lines) > 0 {
		blocks = append(blocks, block)
	}

	sort.Sort(sort.Reverse(ByDate(blocks)))

	for _, b := range blocks {
		for _, l := range b.lines {
			fmt.Println(l)
		}
		fmt.Println("")
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
