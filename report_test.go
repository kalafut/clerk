package main

import (
	"os"
	"strings"
	"testing"

	"gopkg.in/tylerb/is.v1"
)

func TestBalance(test *testing.T) {
	acct := NewRootAccount()
	is := is.New(test)

	f, _ := os.Open("test_data/test2.csv")
	r := NewCSVTxReader(f)
	transactions := r.Read(acct)
	//fmt.Print(balanceReport(transactions))

	rptLines := strings.Split(balanceReport(transactions), "\n")
	is.Equal(rptLines[0], "A                  $ 200.00")
	is.Equal(rptLines[9], "ETrade             34.00 AAPL")
}
