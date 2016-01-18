package clerk

import (
	"os"
	"testing"

	"github.com/kalafut/is"
)

func TestImportExport(test *testing.T) {
	is := is.New(test)

	ledger := NewLedgerFromFile("test_data/test1.csv")
	ledger.Export(os.Stdout)

	_ = is
}
