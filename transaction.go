package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math/big"
	"regexp"
	"strings"
	"time"
)

const multiplier = 100000
const stdDate = "2006/01/02"
const defaultCommodity = "$"

type Amount int64
type Account string
type Commodity string

type Posting struct {
	account   Account
	amount    *big.Rat
	commodity Commodity
}

type Transaction struct {
	date     time.Time
	summary  string
	postings []Posting
	note     string
}

type Entry struct {
	acct Account
	//amt  *big.Rat
	amt Amount
}

func NewEntry(acct Account, amt string) Entry {
	e := Entry{
		acct: acct,
		amt:  parseAmount(amt),
	}
	return e
}

func parseAmount(amt string) Amount {
	var val float64
	//r := new(big.Rat)
	cleaned := strings.Replace(amt, "$", "", -1)

	_, err := fmt.Sscanf(cleaned, "%f", &val)
	if err != nil {
		log.Fatal("error scanning value:", err)
	}

	return toAmount(val)
}

func toAmount(v float64) Amount {
	return Amount(int64(v * float64(multiplier)))
}

type _Transaction struct {
	date    time.Time
	summary string
	entries []Entry
}

func NewTransaction(date, summary string, entries []Entry) _Transaction {
	parsedDate, err := time.Parse("2006-01-02", date)

	if err != nil {
		log.Fatal("error parsing date:", err)
	}

	t := _Transaction{
		date:    parsedDate,
		summary: summary,
		entries: entries,
	}

	return t
}

func (t _Transaction) balanced() bool {
	var total Amount
	for _, e := range t.entries {
		total += e.amt
	}

	return total == 0
}

// Equal tests whether two transactions are equal according to the given
// level of strictness:
//
//
func Equal(a, b _Transaction, strictness int) bool {
	/*
		balancesA := map[Account]*big.Rat{}
		balancesB := map[Account]*big.Rat{}

		for _,t:=range a.entries {
		}
	*/
	return true
}

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

		/*
			for k, v := range config.Columns {
				t.set(k, record[v])
			}

			out = append(out, t)
		*/
	}

	return trans
}

var rePosting2 = regexp.MustCompile(`^(?P<account>.*?)\s{2,}(?P<comm1>[^-.0-9]*?)\s?(?P<amount>-?[.0-9]+)\s?(?P<comm2>[^-.0-9]*)$`)

func parsePostings(p string) []Posting {
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

		var comm string
		if result["comm1"] != "" && result["comm2"] != "" {
			log.Fatalf("Multiple commmdities in posting: %s", posting)
		} else if result["comm1"] != "" {
			comm = result["comm1"]
		} else if result["comm2"] != "" {
			comm = result["comm2"]
		} else {
			comm = defaultCommodity
		}

		r := new(big.Rat)
		r.SetString(result["amount"])
		p := Posting{
			account:   Account(result["account"]),
			amount:    r,
			commodity: Commodity(comm),
		}

		postings = append(postings, p)

	}
	return postings
}

func balance(tranactions []Transaction) {
	balances := map[Account]map[Commodity]*big.Rat{}

	for _, t := range tranactions {
		for _, p := range t.postings {
			if balances[p.account] == nil {
				balances[p.account] = map[Commodity]*big.Rat{}
			}

			if balances[p.account][p.commodity] == nil {
				balances[p.account][p.commodity] = big.NewRat(0, 1)
			}

			bal := balances[p.account][p.commodity]
			bal.Add(bal, p.amount)
		}
	}

	for a := range balances {
		for c, v := range balances[a] {
			fmt.Printf("%v  %v  %v\n", a, c, v.FloatString(2))
		}
	}
}
