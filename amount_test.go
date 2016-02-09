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

	amt = NewAmount(4.25, "AAPL")
	is.Equal("4.25 AAPL", amt.Strings()["AAPL"])
}

func TestEqual(test *testing.T) {
	is := is.New(test)

	amt := NewAmount(4.25, "$")
	is.True(amt.Equal(NewAmount("4.25", "$"), false))
	is.False(amt.Equal(NewAmount("4.251", "$"), false))
	is.False(amt.Equal(NewAmount("4.25", "X"), false))

	amt2 := NewAmount(4.25, "$").Add(NewAmount(1, "X"))
	is.False(amt.Equal(amt2, false))

	amt3 := NewAmount(4.25, "$").Add(NewAmount(0, "X"))
	is.False(amt.Equal(amt3, false))
	is.True(amt.Equal(amt3, true))

}
