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
	transactions := reader.Read(NewRootAccount())

	fmt.Print(balanceReport(transactions))
}

func FormatCmd(config Config) {
	r, err := os.Open(config.inputFile)
	if err != nil {
		log.Fatal(err)
	}

	reader := NewCSVTxReader(r)

	j := NewJournal()
	j.Load(reader)

	r.Close()

	os.Rename(config.inputFile, config.inputFile+".bak")
	w, err := os.Create(config.inputFile)
	defer w.Close()

	if err == nil {
		writer := NewCSVTxWriter(w)
		j.Store(writer)
	}
}
