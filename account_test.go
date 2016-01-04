package main

import (
	"testing"

	"gopkg.in/tylerb/is.v1"
)

func TestAccount(test *testing.T) {
	is := is.New(test)

	root := NewRootAccount()

	is.Equal("", root.name)

	acct := root.findOrAddAccount("this:is:a:test")
	is.Equal("test", root.children["this"].children["is"].children["a"].children["test"].name)
	is.Equal("test", acct.name)
	is.Equal("a", acct.parent.name)

	acct2 := root.findOrAddAccount("this:is:a:test")
	is.Equal(acct, acct2)
}
