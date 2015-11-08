package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"
	"time"
)

//type ynabRow struct {

//"Account","Flag","Check Number","Date","Payee","Category","Master Category","Sub Category","Memo","Outflow","Inflow","Cleared","Running Balance"

var blocks []Block
var blocksByDate = map[time.Time][]*Block{}

func main() {
	flag.Parse()
	filename := flag.Arg(0)

	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	blocks = parse(f)

	//sort.Sort(sort.Reverse(ByDate(blocks)))
	sort.Sort(ByDate(blocks))

	//for _, b := range blocks {
	//	for _, l := range b.lines {
	//		fmt.Println(l)
	//	}
	//	fmt.Println("")
	//}
	findDupes(blocks)
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
				blocksByDate[block.date] = append(blocksByDate[block.date], &block)
				block = Block{}
			}
		} else {
			t, err := time.Parse("2006/01/02", line[0:10])
			if err == nil {
				// Start a new block
				if !block.Empty() {
					blocks = append(blocks, block)
					blocksByDate[block.date] = append(blocksByDate[block.date], &block)
					block = Block{}
				}
				block.date = t
			}
			block.lines = append(block.lines, line)
		}
	}

	if !block.Empty() {
		blocks = append(blocks, block)
		blocksByDate[block.date] = append(blocksByDate[block.date], &block)
	}

	return blocks
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
