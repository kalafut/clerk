package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"testing"
	"time"

	"gopkg.in/tylerb/is.v1"
)

func TestTransactionCreation(test *testing.T) {
	is := is.New(test)

	date := time.Now()
	summary := "Yay!"
	postings := []Posting{
		{NewRootAccount(), NewAmount("$", "4.30")},
		{NewRootAccount(), NewAmount("AAPL", "14.00")},
	}
	note := "Noted."

	t1 := NewTransaction(
		date,
		summary,
		postings,
		note,
	)

	t2 := new(Transaction).SetDate(date).SetSummary(summary).SetPostings(postings).SetNote(note)
	t3 := new(Transaction).SetDate(date).SetSummary(summary).SetPostings(postings).SetNote(note + " ")

	is.Equal(t1, t2)
	is.NotEqual(t1, t3)
}

/*
func TestEntry(test *testing.T) {
	is := is.New(test)

	e1 := NewEntry("Savings", "$50.00")
	is.Equal(e1.acct, "Savings")
	is.Equal(e1.amt, toAmount(50))

	e2 := NewEntry("Savings", "50.00")
	is.Equal(e2.amt, toAmount(50))
	is.Equal(e1, e2)

	e3 := NewEntry("Savings", "50.01")
	e4 := NewEntry("Checking", "50.00")
	is.NotEqual(e1, e3)
	is.NotEqual(e1, e4)
}
func TestTransaction(test *testing.T) {
	is := is.New(test)

	e1 := NewEntry("Savings", "$50.00")
	e2 := NewEntry("Checking", "-$50.00")
	e3 := NewEntry("Credit", "-$25.00")
	e4 := NewEntry("Cash", "0.00")
	e5 := NewEntry("Cash", "0.01")

	t := NewTransaction("2015-10-13", "Test Summary", []Entry{})
	is.Equal(t.summary, "Test Summary")
	is.True(t.balanced())

	t = NewTransaction("2015-10-13", "Test Summary", []Entry{e1, e2})
	is.True(t.balanced())

	t = NewTransaction("2015-10-13", "Test Summary", []Entry{e1, e3, e3})
	is.True(t.balanced())

	t = NewTransaction("2015-10-13", "Test Summary", []Entry{e1, e3, e3, e4})
	is.True(t.balanced())

	t = NewTransaction("2015-10-13", "Test Summary", []Entry{e1, e3, e3, e4, e5})
	is.False(t.balanced())
}
*/

func TestParse2(test *testing.T) {
	is := is.New(test)

	p := parsePostings("  Assets:Checking   AAPL 50.00 &  Credit  -34.24  ")

	is.Equal(2, len(p))

	input, _ := ioutil.ReadFile("test_data/test1.csv")
	r := bytes.NewReader(input)

	transactions := ParseTransactions(r)

	is.Equal(3, len(transactions))
	is.Equal(date("2015/12/31"), transactions[0].Date())
	is.Equal("Payee or summary", transactions[0].Summary())
	is.Equal(date("2015/12/31"), transactions[1].Date())
}

func date2(s string) time.Time {
	date, err := time.Parse(StdDate, s)
	if err != nil {
		log.Fatal(err)
	}

	return date
}
