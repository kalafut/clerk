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

func (m MultiBalance) String() string {
	var b bytes.Buffer

	for k, v := range m {
		fmt.Fprintf(&b, "%s  %s\n", k.Name, v)
	}
	return b.String()
}

func balanceReport(transactions []*Tx) string {
	var w = new(tabwriter.Writer)
	var b bytes.Buffer
	var root *Account

	w.Init(&b, 0, 0, 1, ' ', 0)
	balances := MultiBalance{}

	for _, t := range transactions {
		for _, p := range t.Postings() {
			balances.AddUp(p.Acct, p.Amt)
			if root == nil {
				root = p.Acct.Root()
			}
		}
	}

	traverse(root, balances, w)
	w.Flush()
	return b.String()
}

func traverse(acct *Account, balances MultiBalance, w *tabwriter.Writer) string {
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
		traverse(acct.Children()[child], balances, w)
	}

	return result
}
