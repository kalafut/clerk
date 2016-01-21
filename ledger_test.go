package main

import (
	"testing"

	"github.com/kalafut/is"
)

func TestLedgerAdd(test *testing.T) {
	var all []*Transaction

	is := is.New(test)

	ldg := NewLedger()
	all = ldg.All()
	is.Equal(0, len(all))

	t1 := new(Transaction)
	ldg.Add(t1)
	all = ldg.All()
	is.Equal(1, len(all))
	is.Equal(all[0], t1)

	t2 := new(Transaction)
	ldg.Add(t2)
	all = ldg.All()
	is.Equal(2, len(all))
	is.Equal(all[1], t2)
}

func TestLedgerDel(test *testing.T) {
	var all []*Transaction

	is := is.New(test)

	ldg := NewLedger()
	t1 := new(Transaction)
	t2 := new(Transaction)

	ldg.Add(t1)
	ldg.Add(t2)
	all = ldg.All()
	is.Equal(2, len(all))

	r := ldg.Del(t1)
	all = ldg.All()
	is.True(r)
	is.Equal(1, len(all))
	is.Equal(all[0], t2)

	r = ldg.Del(t1)
	all = ldg.All()
	is.False(r)
	is.Equal(1, len(all))
	is.Equal(all[0], t2)

	r = ldg.Del(t2)
	all = ldg.All()
	is.True(r)
	is.Equal(0, len(all))
}

func TestLedgerReplace(test *testing.T) {
	var all []*Transaction

	is := is.New(test)

	ldg := NewLedger()
	t1 := new(Transaction)
	t2 := new(Transaction)
	t3 := new(Transaction)
	t4 := new(Transaction)

	ldg.Add(t1)
	ldg.Add(t2)

	r := ldg.Replace(t2, t3)
	all = ldg.All()
	is.True(r)
	is.Equal(2, len(all))
	is.Equal(all[1], t3)

	r = ldg.Replace(t4, t3)
	is.False(r)
	all = ldg.All()
	is.Equal(2, len(all))
	is.Equal(all[1], t3)
}
