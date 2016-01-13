package main

import (
	"bytes"
	"fmt"
	"math/big"
	"sort"
	"strings"
	"text/tabwriter"
)

type MultiBalance map[*Account]map[Commodity]*big.Rat

func (m MultiBalance) Add(acct *Account, comm Commodity, amt *big.Rat) {
	if m[acct] == nil {
		m[acct] = map[Commodity]*big.Rat{}
	}

	if m[acct][comm] == nil {
		m[acct][comm] = big.NewRat(0, 1)
	}

	bal := m[acct][comm]
	bal.Add(bal, amt)
}

func (m MultiBalance) AddUp(acct *Account, comm Commodity, amt *big.Rat) {
	m.Add(acct, comm, amt)
	//	fmt.Printf("On: %v\n", acct.name) // (parent: %v)\n", acct.name, acct.parent.name)

	// Because the single root account is not used, look two levels up
	if acct.parent.parent != nil {
		m.AddUp(acct.parent, comm, amt)
	}
}

var w = new(tabwriter.Writer)

func balanceReport(tranactions []Transaction) string {
	var b bytes.Buffer
	w.Init(&b, 0, 0, 1, ' ', 0)
	balances := MultiBalance{}

	for _, t := range tranactions {
		for _, p := range t.postings {
			balances.AddUp(p.account, p.commodity, p.amount)
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
