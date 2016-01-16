package ledger

import (
	"bytes"
	"io/ioutil"
	"log"
	"testing"
	"time"

	"gopkg.in/tylerb/is.v1"
)

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

func TestParse(test *testing.T) {
	is := is.New(test)

	p := parsePostings("  Assets:Checking   AAPL 50.00 &  Credit  -34.24  ")

	is.Equal(2, len(p))

	input, _ := ioutil.ReadFile("test_data/test1.csv")
	r := bytes.NewReader(input)

	transactions := ParseTransactions(r)

	is.Equal(3, len(transactions))
	is.Equal(date("2015/12/31"), transactions[0].Date)
	is.Equal("Payee or summary", transactions[0].Summary)
	is.Equal(date("2015/12/31"), transactions[1].Date)
}

func date(s string) time.Time {
	date, err := time.Parse(stdDate, s)
	if err != nil {
		log.Fatal(err)
	}

	return date
}
