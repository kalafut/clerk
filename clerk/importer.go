package clerk

/*
import (
	"fmt"

	"github.com/naoina/toml"
)

const importAcct = "__Uncategorized__"

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

func (t *transaction) set(field, val string) {
	switch field {
	case "date":
		t.date = val
	case "description":
		t.description = val
	case "amount":
		t.amount = val
	}
}

type Importer interface {
	Import() (transaction, error)
}

func importAll(in Importer) []Block {
	var blocks []Block

	for t, err := in.Import(); err == nil; {
		_ = t
		lines := []string{}

		b := Block{lines: lines}
		blocks = append(blocks, b)
	}

	return blocks
}

func Import(ledgerFile, importFile string) {
	ledger := NewLedgerFromFile(ledgerFile)

	var config ImportConfig
	if err := toml.Unmarshal([]byte(configToml), &config); err != nil {
		panic(err)
	}

	csv := CSVImport(importFile, config.Accounts["Chase Checking"])
	fmt.Println(csv)

	_ = ledger
	test()
}
*/
