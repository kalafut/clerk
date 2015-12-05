package main

import (
	"fmt"
	"log"
	"strings"
	"time"
)

const multiplier = 100000

type Amount int64
type Account string

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
