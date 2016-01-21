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
		{NewRootAccount(), NewAmount("4.30", "$")},
		{NewRootAccount(), NewAmount("14.00", "AAPL")},
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
