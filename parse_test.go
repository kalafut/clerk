package main

/*

import (
	"testing"

	"gopkg.in/tylerb/is.v1"
)

func TestRegex(t *testing.T) {
	is := is.New(t)

	tests := []struct {
		match bool
		input string
		acct  string
		amt   string
	}{
		{true, "    Everyday Expenses:Medical                 $10.00", "Everyday Expenses:Medical", "10.00"},
		{true, "    Everyday Expenses:Medical Stuff           5,000", "Everyday Expenses:Medical Stuff", "5,000"},
		{true, "    Everyday Expenses:Medical Stuff           -5,000", "Everyday Expenses:Medical Stuff", "5,000"},
		{true, "    Taxes 123           600.45", "Taxes 123", "600.45"},
		{true, "    Taxes          ", "Taxes", ""},
		{true, "    Taxes 123      ", "Taxes 123", ""},
		{false, "Not Indented           600.45", "", ""},
	}

	for _, test := range tests {
		matches := acctAmtRegex.FindStringSubmatch(test.input)
		is.Equal(len(matches) > 0, test.match)
		if test.match {
			is.Equal(matches[1], test.acct)
			is.Equal(matches[2], test.amt)
		}
	}
}
*/
