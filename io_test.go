package main

import (
	"log"
	"os"
	"testing"

	"github.com/kalafut/is"
)

func TestImportExport(test *testing.T) {
	is := is.New(test)

	f, err := os.Open("test_data/test1.csv")
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}

	ledger := NewJournal()

	ledger.Export(os.Stdout)

	_ = is
}
