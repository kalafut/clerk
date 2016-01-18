package clerk

import (
	"sort"
	"strings"
)

var RootAccount = NewRootAccount()

type Account struct {
	Name     string
	children map[string]*Account
	parent   *Account
}

func NewRootAccount() *Account {
	return &Account{children: make(map[string]*Account)}
}

func (a Account) Parent() *Account {
	return a.parent
}

func (a Account) Children() map[string]*Account {
	return a.children
}

func (acct Account) allSorted() []string {
	childAccts := make([]string, 0, len(acct.children))
	for a := range acct.children {
		childAccts = append(childAccts, a)
	}

	sort.Strings(childAccts)
	return childAccts
}

func (acct *Account) Level() int {
	var p *Account
	level := 0

	p = acct
	for p.parent != nil {
		level++
		p = p.parent
	}

	return level
}
func (acct *Account) FindOrAddAccount(acctName string) *Account {
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
			Name:     name,
			children: make(map[string]*Account),
			parent:   acct,
		}
		acct.children[name] = child
	}

	if idx != -1 {
		child = child.FindOrAddAccount(acctName[idx+1:])
	}

	return child

}
