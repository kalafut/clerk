package clerk

import (
	"log"
	"strings"
	"testing"
	"time"

	"gopkg.in/tylerb/is.v1"
)

const test1 = `

2015/12/31, Payee or summary , A:B:C:D  $200 & Income  $-200, This is my New Years Eve??
2015/12/31, Payee or summary , Income  $-1351.32 & Assets:Bank:Chase Checking  $-1351.32, This is my New Years Eve??
2016/01/16, Stock purchase,   ETrade  -351.32 & ETrade  34 AAPL  &  ETrade  $151.33  ,
`

func TestParse(test *testing.T) {
	is := is.New(test)

	r := strings.NewReader(test1)

	reader := NewCSVTxReader(r)
	root := NewRootAccount()
	transactions := reader.Read(root)

	is.Equal(3, len(transactions))

	is.Equal(date("2015/12/31"), transactions[0].Date)
	is.Equal("Payee or summary", transactions[0].Summary)
	is.Equal("D", transactions[0].Postings[0].Account.Name)
	is.Equal(NewAmount("200", "$"), transactions[0].Postings[0].Amount)
	//is.Equal(, transactions[0].Postings[0].Account.Name)

	is.Equal(date("2016/01/16"), transactions[2].Date)
	is.Equal("Stock purchase", transactions[2].Summary)
}

func date(s string) time.Time {
	date, err := time.Parse(StdDate, s)
	if err != nil {
		log.Fatal(err)
	}

	return date
}
