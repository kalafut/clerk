package main

import (
	"fmt"
	"math/big"
	"sort"
	"strings"
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

func balanceReport(tranactions []Transaction) {
	balances := MultiBalance{}

	for _, t := range tranactions {
		for _, p := range t.postings {
			balances.AddUp(p.account, p.commodity, p.amount)
		}
	}

	traverse(rootAccount, balances)
}

func traverse(acct *Account, balances MultiBalance) {
	for c, v := range balances[acct] {
		fmt.Printf("%s%v  %v  %v\n", strings.Repeat(" ", 2*(acct.level()-1)), acct.name, c, v.FloatString(2))
	}

	children := make([]string, 0, len(acct.children))
	for child, _ := range acct.children {
		children = append(children, child)
	}

	sort.Strings(children)
	for _, child := range children {
		traverse(acct.children[child], balances)
	}
}
