package ledger

import (
	"testing"

	"gopkg.in/tylerb/is.v1"
)

func TestAccount(test *testing.T) {
	is := is.New(test)

	root := NewRootAccount()

	is.Equal("", root.Name)

	acct := root.FindOrAddAccount("this:is:a:test")
	is.Equal("test", root.children["this"].children["is"].children["a"].children["test"].Name)
	is.Equal("test", acct.Name)
	is.Equal("a", acct.parent.Name)

	acct2 := root.FindOrAddAccount("this:is:a:test")
	is.Equal(acct, acct2)
}
