package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math/big"
	"regexp"
	"sort"
	"strings"
	"time"
)

const multiplier = 100000
const stdDate = "2006/01/02"
const defaultCommodity = "$"

var rootAccount = NewRootAccount()

type Amount int64
type Account struct {
	name     string
	children map[string]*Account
	parent   *Account
}

type Commodity string

type Posting struct {
	account   *Account
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
	amt  Amount
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
			comm = Commodity(c1)
		case c2 != "":
			comm = Commodity(c2)
		default:
			comm = defaultCommodity
		}

		r := new(big.Rat)
		r.SetString(result["amount"])
		p := Posting{
			account:   rootAccount.findOrAddAccount(result["account"]),
			amount:    r,
			commodity: comm,
		}

		postings = append(postings, p)

	}
	return postings
}

type AccountBalance struct {
	account  Account
	balances map[Commodity]*big.Rat
}

type AccountTree struct {
	node     *Account // nil for root
	children []*Account
}

var root AccountTree

func NewRootAccount() *Account {
	return &Account{children: make(map[string]*Account)}
}

func (acct Account) allSorted() []string {
	childAccts := make([]string, 0, len(acct.children))
	for a := range acct.children {
		childAccts = append(childAccts, a)
	}

	sort.Strings(childAccts)
	return childAccts
}

func (acct *Account) level() int {
	var p *Account
	level := 0

	p = acct
	for p.parent != nil {
		level++
		p = p.parent
	}

	return level
}
func (acct *Account) findOrAddAccount(acctName string) *Account {
	var child *Account
	var name string
	var ok bool

	idx := strings.Index(acctName, ":")
	if idx == -1 {
		name = acctName
	} else {
		name = acctName[0:idx]
	}

	if child, ok = acct.children[name]; !ok {
		child = &Account{
			name:     name,
			children: make(map[string]*Account),
			parent:   acct,
		}
		acct.children[name] = child
	}

	if idx != -1 {
		child = child.findOrAddAccount(acctName[idx+1:])
	}

	return child

}
