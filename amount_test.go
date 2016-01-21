package main

import (
	"testing"

	"github.com/kalafut/is"
)

func TestFormatting(test *testing.T) {
	is := is.New(test)

	amt := NewAmount("4.25", "$")
	is.Equal("$ 4.25", amt.Strings()["$"])

	amt = NewAmount("4.259", "$")
	is.Equal("$ 4.26", amt.Strings()["$"])

	amt = NewAmount("4.25", "AAPL")
	is.Equal("4.25 AAPL", amt.Strings()["AAPL"])
}
