package clerk

import (
	"testing"

	"gopkg.in/tylerb/is.v1"
)

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
