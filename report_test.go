package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/kalafut/is"
)

func TestBalance(test *testing.T) {
	is := is.New(test)

	input, _ := ioutil.ReadFile("test_data/test2.csv")
	r := bytes.NewReader(input)
	transactions := ParseTransactions(r)
	fmt.Print(balanceReport(transactions))

	//fmt.Println(transactions)
	_ = "breakpoint"
	rpt_lines := strings.Split(balanceReport(transactions), "\n")
	is.Equal(rpt_lines[0], "A                  $ 200.00")
	is.Equal(rpt_lines[9], "ETrade             34.00 AAPL")
}
