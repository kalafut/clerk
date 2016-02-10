package main

import (
	"testing"
	"time"

	"github.com/kalafut/is"
)

var (
	t1 = NewTransaction(time.Unix(200, 0), "AAA", []Posting{}, "")
	t2 = NewTransaction(time.Unix(100, 0), "ZZZ", []Posting{}, "")
	t3 = NewTransaction(time.Unix(100, 0), "YYY", []Posting{}, "")
	t4 = NewTransaction(time.Unix(300, 0), "YYY", []Posting{}, "")
)

func TestJournalAdd(test *testing.T) {
	var all []*Tx

	is := is.New(test)

	ldg := NewJournal()
	all = ldg.All()
	is.Equal(0, len(all))

	ldg.Add(t1)
	all = ldg.All()
	is.Equal(1, len(all))
	is.Equal(all[0], t1)

	ldg.Add(t2)
	all = ldg.All()
	is.Equal(2, len(all))
	is.Equal(all[0], t2)
	is.Equal(all[1], t1)

	ldg.Add(t3)
	all = ldg.All()
	is.Equal(3, len(all))
	is.Equal(all[0], t3)
	is.Equal(all[1], t2)
	is.Equal(all[2], t1)
}

func TestJournalDel(test *testing.T) {
	var all []*Tx

	is := is.New(test)

	ldg := NewJournal()

	ldg.Add(t1)
	ldg.Add(t2)
	ldg.Add(t3)
	all = ldg.All()
	is.Equal(3, len(all))

	is.True(ldg.Del(t1))
	all = ldg.All()
	is.Equal(2, len(all))
	is.Equal(all[0], t3)
	is.Equal(all[1], t2)

	is.False(ldg.Del(t1))
	all = ldg.All()
	is.Equal(2, len(all))
	is.Equal(all[1], t2)

	is.True(ldg.Del(t2))
	is.True(ldg.Del(t3))
	all = ldg.All()
	is.Equal(0, len(all))
}

func TestJournalReplace(test *testing.T) {
	var all []*Tx
	is := is.New(test)

	ldg := NewJournal()

	ldg.Add(t1)
	ldg.Add(t2)

	is.True(ldg.Replace(t2, t3))
	all = ldg.All()
	is.Equal(2, len(all))
	is.Equal(all[0], t3)
	is.Equal(all[1], t1)

	is.False(ldg.Replace(t4, t3))
	all = ldg.All()
	is.Equal(2, len(all))
	is.Equal(all[0], t3)
	is.Equal(all[1], t1)
}
