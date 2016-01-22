package main

import (
	"encoding/csv"
	"io"
	"log"
	"math/big"
	"regexp"
	"strings"
	"time"
)

type CSVTxReader struct {
	reader *csv.Reader
}

func NewCSVTxReader(r io.Reader) CSVTxReader {
	return CSVTxReader{
		reader: csv.NewReader(r),
	}
}

func (r CSVTxReader) Read(root *Account) []*Tx {
	trans := []*Tx{}

	for {
		record, err := r.reader.Read()
		if err == io.EOF {
			break
		}

		date, err := time.Parse(StdDate, record[0])
		if err != nil {
			log.Fatal(err)
		}

		t := NewTransaction(
			date,
			strings.TrimSpace(record[1]),
			parsePostings2(root, record[2]),
			"",
		)
		trans = append(trans, t)
	}

	return trans
}

var rePosting2 = regexp.MustCompile(`^(?P<account>.*?)\s{2,}(?P<comm1>[^-.0-9]*?)\s?(?P<amount>-?[.0-9]+)\s?(?P<comm2>[^-.0-9]*)$`)

func parsePostings2(root *Account, p string) []Posting {
	var comm string
	postings := []Posting{}

	for _, posting := range strings.Split(p, "&") {
		posting = strings.TrimSpace(posting)
		match := rePosting2.FindStringSubmatch(posting)

		if len(match) == 0 {
			log.Fatalf("Invalid posting: %s", posting)
		}

		result := make(map[string]string)
		for i, name := range rePosting2.SubexpNames() {
			if i != 0 {
				result[name] = match[i]
			}
		}

		c1, c2 := result["comm1"], result["comm2"]

		switch {
		case c1 != "" && c2 != "":
			log.Fatalf("Multiple commmodities in posting: %s", posting)
		case c1 != "":
			comm = c1 // TODO: use a commodity pool instead, else "$ 1" is different than "1 $"
		case c2 != "":
			comm = c2
		default:
			comm = DefaultCommodity
		}

		r := new(big.Rat)
		r.SetString(result["amount"])
		p := Posting{
			Acct: root.FindOrAddAccount(result["account"]),
			Amt:  NewAmount(result["amount"], comm),
		}

		postings = append(postings, p)
	}
	//checkBalance(postings)
	return postings
}

/*

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/naoina/toml"
)

type columnConfig struct {
	date        string
	posted      string
	code        string
	description string
	amount      string
	cost        string
	total       string
	note        string
}

type csvConfig struct {
	multiaccount bool
	invertAmount bool
	columns      columnConfig
}

//"Account","Flag","Check Number","Date","Payee","Category","Master Category","Sub Category","Memo","Outflow","Inflow","Cleared","Running Balance"
//Type,Trans Date,Post Date,Description,Amount

func CSVImport(csvfile string, config AccountConfig) []transaction {
	var out = []transaction{}

	f, err := os.Open(csvfile)
	defer f.Close()

	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(f)

	if config.SkipFirstLine {
		r.Read()
	}

	for {
		t := transaction{}
		record, err := r.Read()
		//fmt.Printf("%v\n", record)
		if err == io.EOF {
			break
		}

		for k, v := range config.Columns {
			t.set(k, record[v])
		}

		out = append(out, t)
	}

	return out
}

var configToml = `
[accounts."Chase Checking"]
    skipFirstLine = true
	invertAmount = true

	[accounts."Chase Checking".columns]
	amount = 4
	date = 1
	description = 3
`

type ImportConfig struct {
	Accounts map[string]AccountConfig
}

type AccountConfig struct {
	TargetAccount string
	SkipFirstLine bool
	InvertAmount  bool
	Columns       map[string]int
}

//type transaction struct {
//	date        string
//	posted      string
//	code        string
//	description string
//	amount      string
//	cost        string
//	total       string
//	note        string
//}

func test() {
	var config ImportConfig
	if err := toml.Unmarshal([]byte(configToml), &config); err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", config)
}
*/
