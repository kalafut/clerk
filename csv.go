// Common csv importer support
package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/naoina/toml"
)

type columnConfig struct {
	date        string
	posted      string
	code        string
	description string
	amount      string
	cost        string
	total       string
	note        string
}

type csvConfig struct {
	multiaccount bool
	invertAmount bool
	columns      columnConfig
}

//"Account","Flag","Check Number","Date","Payee","Category","Master Category","Sub Category","Memo","Outflow","Inflow","Cleared","Running Balance"
//Type,Trans Date,Post Date,Description,Amount

func CSVImport(csvfile string, config AccountConfig) []transaction {
	var out = []transaction{}

	f, err := os.Open(csvfile)
	defer f.Close()

	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(f)

	if config.SkipFirstLine {
		r.Read()
	}

	for {
		t := transaction{}
		record, err := r.Read()
		//fmt.Printf("%v\n", record)
		if err == io.EOF {
			break
		}

		for k, v := range config.Columns {
			t.set(k, record[v])
		}

		out = append(out, t)
	}

	return out
}

var configToml = `
[accounts."Chase Checking"]
    skipFirstLine = true
	invertAmount = true

	[accounts."Chase Checking".columns]
	amount = 4
	date = 1
	description = 3
`

type ImportConfig struct {
	Accounts map[string]AccountConfig
}

type AccountConfig struct {
	TargetAccount string
	SkipFirstLine bool
	InvertAmount  bool
	Columns       map[string]int
}

/*
type transaction struct {
	date        string
	posted      string
	code        string
	description string
	amount      string
	cost        string
	total       string
	note        string
}
*/
func test() {
	var config ImportConfig
	if err := toml.Unmarshal([]byte(configToml), &config); err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", config)
}
