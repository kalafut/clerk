package main

import (
	"bytes"
	"encoding/csv"
	"io"
	"log"
	"math/big"
	"regexp"
	"strings"
	"time"
)

const multiplier = 100000
const stdDate = "2006/01/02"

type Posting struct {
	account *Account
	amount  Amount
}

type Transaction struct {
	date     time.Time
	summary  string
	postings []Posting
	note     string
}

func (t Transaction) toCSV() string {
	var buf bytes.Buffer
	var postings bytes.Buffer

	for _, p := range t.postings {
		postings.WriteString(p.account.name)
		postings.WriteString("  &  ")
	}

	w := csv.NewWriter(&buf)
	record := []string{
		t.date.Format(stdDate),
		t.summary,
		postings.String(),
		t.note,
	}

	w.Write(record)
	w.Flush()

	return buf.String()
}

// Equal tests whether two transactions are equal according to the given
// level of strictness:
//
//
//func Equal(a, b _Transaction, strictness int) bool {
//	/*
//		balancesA := map[Account]*big.Rat{}
//		balancesB := map[Account]*big.Rat{}
//
//		for _,t:=range a.entries {
//		}
//	*/
//	return true
//}

func ParseTransactions(in io.Reader) []Transaction {
	trans := []Transaction{}
	r := csv.NewReader(in)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		date, err := time.Parse(stdDate, record[0])
		if err != nil {
			log.Fatal(err)
		}

		t := Transaction{
			date:     date,
			summary:  strings.TrimSpace(record[1]),
			postings: parsePostings(record[2]),
		}
		trans = append(trans, t)
	}

	return trans
}

var rePosting2 = regexp.MustCompile(`^(?P<account>.*?)\s{2,}(?P<comm1>[^-.0-9]*?)\s?(?P<amount>-?[.0-9]+)\s?(?P<comm2>[^-.0-9]*)$`)

func parsePostings(p string) []Posting {
	var comm Commodity
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
			comm = Commodity{abbr: c1} // TODO: use a commodity pool instead, else "$ 1" is different than "1 $"
		case c2 != "":
			comm = Commodity{abbr: c2, postfix: true}
		default:
			comm = defaultCommodity
		}

		r := new(big.Rat)
		r.SetString(result["amount"])
		p := Posting{
			account: rootAccount.findOrAddAccount(result["account"]),
			amount:  NewAmount(result["amount"], comm),
		}

		postings = append(postings, p)
	}
	checkBalance(postings)
	return postings
}

func checkBalance(postings []Posting) bool {
	sum := Amount{}

	for _, p := range postings {
		sum.Add(p.amount)
	}

	if !sum.Zero() {
		// do something
	}
	return true
}
