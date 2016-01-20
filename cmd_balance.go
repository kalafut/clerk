package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

func BalanceCmd(config Config) {
	input, _ := ioutil.ReadFile(config.inputFile)
	r := bytes.NewReader(input)
	transactions := ParseTransactions(r)
	fmt.Print(balanceReport(transactions))
}
