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

	if acct.parent.parent != nil {
		m.AddUp(acct.parent, amt)
	}
}

var w = new(tabwriter.Writer)

func balanceReport(tranactions []Transaction) string {
	var b bytes.Buffer
	w.Init(&b, 0, 0, 1, ' ', 0)
	balances := MultiBalance{}

	for _, t := range tranactions {
		for _, p := range t.postings {
			balances.AddUp(p.account, p.amount)
		}
	}

	traverse(rootAccount, balances)
	w.Flush()
	return b.String()
}

func traverse(acct *Account, balances MultiBalance) string {
	var result string

	for commodity, value := range balances[acct] {
		var valstr string
		if commodity.postfix {
			valstr = value.FloatString(2) + " " + commodity.abbr
		} else {
			valstr = commodity.abbr + " " + value.FloatString(2)
		}
		fmt.Fprintf(w, "%s%v\t%v\n", strings.Repeat(" ", 2*(acct.level()-1)), acct.name, valstr)
	}

	children := make([]string, 0, len(acct.children))
	for child, _ := range acct.children {
		children = append(children, child)
	}

	sort.Strings(children)
	for _, child := range children {
		traverse(acct.children[child], balances)
	}

	return result
}
