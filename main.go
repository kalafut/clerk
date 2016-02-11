package main

import (
	"fmt"
	"os"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app = kingpin.New("clerk", "clerk Helper")
	//ledgerFile = app.Flag("filename", "clerk filename").Short('f').Default("master.dat").String()
	//importFile = app.Flag("csv", "CSV filename").String()
	//inplace    = app.Flag("inplace", "Edit file in place").Short('i').Bool()
	//outfile    = app.Flag("outfile", "Output file").Short('o').String()
	//sortCmd    = app.Command("sort", "Sort the ledger by date.")
	//dedupeCmd  = app.Command("dedupe", "Deduplicate the ledger.")
	//importCmd  = app.Command("import", "Import from external sources.")
	balanceCmd = app.Command("balance", "Show journal balance.")
	formatCmd  = app.Command("format", "Format the journal")

	devconfig = Config{inputFile: "test_data/clerk.dat"}
)

func fatalf(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "ERROR: "+format+"\n", args...)
	os.Exit(1)
}

func main() {
	//var f *os.File
	//var output *bufio.Writer
	//var tempBuffer bytes.Buffer

	cmd := kingpin.MustParse(app.Parse(os.Args[1:]))
	//f, err := os.Open(*ledgerFile)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//ledger := NewLedgerReader(f)
	//f.Close()

	//if *inplace {
	//	output = bufio.NewWriter(&tempBuffer)
	//} else if *outfile != "" {
	//	f, err = os.Create(*outfile)
	//	defer f.Close()
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	output = bufio.NewWriter(f)
	//} else {
	//	output = bufio.NewWriter(os.Stdout)
	//}

	//defer output.Flush()

	switch cmd {
	case balanceCmd.FullCommand():
		BalanceCmd(devconfig)
	case formatCmd.FullCommand():
		FormatCmd(devconfig)
	}

	//if *inplace {
	//	f, err = os.Create(*ledgerFile)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	output.Flush()
	//	f.Write(tempBuffer.Bytes())
	//	f.Close()
	//}
}
