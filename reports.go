package main

import (
	"bytes"
	"fmt"
	"sort"
	"strings"
	"text/tabwriter"
)

type MultiBalance map[*Account]Amount

func (m MultiBalance) Add(acct *Account, amt Amount) {
	if m[acct] == nil {
		m[acct] = Amount{}
	}

	m[acct].Add(amt)
}

func (m MultiBalance) AddUp(acct *Account, amt Amount) {
	m.Add(acct, amt)

	if acct.Parent().Parent() != nil {
		m.AddUp(acct.Parent(), amt)
	}
}

var w = new(tabwriter.Writer)

func balanceReport(transactions []*Tx) string {
	var b bytes.Buffer
	w.Init(&b, 0, 0, 1, ' ', 0)
	balances := MultiBalance{}

	for _, t := range transactions {
		for _, p := range t.Postings() {
			balances.AddUp(p.Acct, p.Amt)
		}
	}

	traverse(RootAccount, balances)
	w.Flush()
	return b.String()
}

func traverse(acct *Account, balances MultiBalance) string {
	var result string

	valstrs := balances[acct].Strings()

	commodities := []string{}
	for c := range valstrs {
		commodities = append(commodities, c)
	}
	sort.Strings(commodities)

	for _, c := range commodities {
		fmt.Fprintf(w, "%s%s\t%s\n", strings.Repeat(" ", 2*(acct.Level()-1)), acct.Name, valstrs[c])
	}

	children := make([]string, 0, len(acct.Children()))
	for child := range acct.Children() {
		children = append(children, child)
	}

	sort.Strings(children)
	for _, child := range children {
		traverse(acct.Children()[child], balances)
	}

	return result
}
