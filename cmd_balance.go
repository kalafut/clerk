package main

import (
	"fmt"
	"log"
	"os"
)

func BalanceCmd(config Config) {
	r, err := os.Open(config.inputFile)
	if err != nil {
		log.Fatal(err)
	}

	reader := NewCSVTxReader(r)
	transactions := reader.Read(RootAccount) // get rid of RootAccount

	fmt.Print(balanceReport(transactions))
}
