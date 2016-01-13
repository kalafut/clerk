package main

import (
	"math/big"
	"sort"
	"strings"
)

var rootAccount = NewRootAccount()

type Account struct {
	name     string
	children map[string]*Account
	parent   *Account
}

type AccountBalance struct {
	account  Account
	balances map[Commodity]*big.Rat
}

func NewRootAccount() *Account {
	return &Account{children: make(map[string]*Account)}
}

func (acct Account) allSorted() []string {
	childAccts := make([]string, 0, len(acct.children))
	for a := range acct.children {
		childAccts = append(childAccts, a)
	}

	sort.Strings(childAccts)
	return childAccts
}

func (acct *Account) level() int {
	var p *Account
	level := 0

	p = acct
	for p.parent != nil {
		level++
		p = p.parent
	}

	return level
}
func (acct *Account) findOrAddAccount(acctName string) *Account {
	var child *Account
	var name string
	var ok bool

	idx := strings.Index(acctName, ":")
	if idx == -1 {
		name = acctName
	} else {
		name = acctName[0:idx]
	}

	if child, ok = acct.children[name]; !ok {
		child = &Account{
			name:     name,
			children: make(map[string]*Account),
			parent:   acct,
		}
		acct.children[name] = child
	}

	if idx != -1 {
		child = child.findOrAddAccount(acctName[idx+1:])
	}

	return child

}
