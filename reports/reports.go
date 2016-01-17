package main

import (
	"bytes"
	"fmt"
	"sort"
	"strings"
	"text/tabwriter"

	"github.com/kalafut/clerk/core"
	"github.com/kalafut/clerk/ledger"
)

type MultiBalance map[*ledger.Account]core.Amount

func (m MultiBalance) Add(acct *ledger.Account, amt core.Amount) {
	if m[acct] == nil {
		m[acct] = core.Amount{}
	}

	m[acct].Add(amt)
}

func (m MultiBalance) AddUp(acct *ledger.Account, amt core.Amount) {
	m.Add(acct, amt)

	if acct.Parent().Parent() != nil {
		m.AddUp(acct.Parent(), amt)
	}
}

var w = new(tabwriter.Writer)

func balanceReport(tranactions []ledger.Transaction) string {
	var b bytes.Buffer
	w.Init(&b, 0, 0, 1, ' ', 0)
	balances := MultiBalance{}

	for _, t := range tranactions {
		for _, p := range t.Postings {
			balances.AddUp(p.Account, p.Amount)
		}
	}

	traverse(ledger.RootAccount, balances)
	w.Flush()
	return b.String()
}

func traverse(acct *ledger.Account, balances MultiBalance) string {
	var result string

	//for commodity, value := range balances[acct] {
	valstrs := balances[acct].Strings()

	for _, str := range valstrs {
		fmt.Fprintf(w, "%s%s\t%s\n", strings.Repeat(" ", 2*(acct.Level()-1)), acct.Name, str)
	}
	//}

	children := make([]string, 0, len(acct.Children()))
	for child, _ := range acct.Children() {
		children = append(children, child)
	}

	sort.Strings(children)
	for _, child := range children {
		traverse(acct.Children()[child], balances)
	}

	return result
}
