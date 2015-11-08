package main

import (
	"regexp"
	"sort"
	"time"
)

// Blocks are a literal encapsulation of a ledger transaction. They
// are not called transcactions because the actual ledger file strings
// and comments are preserves. A ledger file is a sequence of blocks.
//
// Textually, a block is defined as:
//    <0+ comment lines>
//    <0 or 1 summary line: a) left justified  b) starting with a yyyy/mm/dd date>
//    <0+ acccount lines or comments: a) indented at least one space>
//
// Whitespace between blocks is ignored.

// Note: value will not have '-', intentionally
var acctAmtRegex = regexp.MustCompile(`^\s+(.*?\S)(?:\s{2,}.*?([\d,\.]+))?\s*$`)

type Block struct {
	lines []string
	date  time.Time
}

type ByDate []Block

func (a ByDate) Len() int           { return len(a) }
func (a ByDate) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByDate) Less(i, j int) bool { return a[i].date.Before(a[j].date) }

func (b *Block) Empty() bool {
	return len(b.lines) == 0
}

func (b Block) Accounts() []string {
	var ret []string
	for _, l := range b.lines {
		m := acctAmtRegex.FindStringSubmatch(l)
		if len(m) > 0 {
			ret = append(ret, m[1])
		}
	}
	sort.Strings(ret)
	return ret
}

func (b Block) Amounts() []string {
	var ret []string
	for _, l := range b.lines {
		m := acctAmtRegex.FindStringSubmatch(l)
		if len(m) > 0 {
			ret = append(ret, m[2])
		}
	}
	sort.Strings(ret)
	return ret
}

// IsDupe returns true if other is a likely duplicate based on:
//   date
//   affected accounts
//   amounts
func (b Block) IsDupe(other Block, tolerance time.Duration) bool {
	// Check time
	timeDiff := b.date.Sub(other.date)
	if timeDiff < 0 {
		timeDiff = -timeDiff
	}

	if timeDiff > tolerance {
		return false
	}

	// Check affected accounts
	accts := b.Accounts()
	acctsOther := other.Accounts()

	if len(accts) != len(acctsOther) {
		return false
	}

	for i := range accts {
		if accts[i] != acctsOther[i] {
			return false
		}
	}

	// Check affected accounts
	amts := b.Amounts()
	amtsOther := other.Amounts()

	if len(amts) != len(amtsOther) {
		return false
	}

	for i := range amts {
		if amts[i] != amtsOther[i] {
			return false
		}
	}

	return true
}
