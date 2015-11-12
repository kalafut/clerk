package main

import (
	"testing"
	"time"

	"gopkg.in/tylerb/is.v1"
)

const day = 24 * time.Hour

func TestIsDupe(t *testing.T) {
	is := is.New(t)
	now := time.Now()

	b_base := Block{
		date: now,
		lines: []string{
			"2015/10/04 Doctor Roberts",
			"    Everyday Expenses:Medical        $10.00",
			"    Assets:Bank:Chase Checking",
		},
	}

	// Test date
	b_test := cloneBlock(b_base)
	is.True(b_base.IsDupe(b_test, 0*day))

	b_test.date = now.Add(1 * day)
	is.True(b_base.IsDupe(b_test, 1*day))

	b_test.date = now.Add(2 * day)
	is.False(b_base.IsDupe(b_test, 1*day))
	is.True(b_base.IsDupe(b_test, 2*day))

	b_test.date = now.Add(-1 * day)
	is.True(b_base.IsDupe(b_test, 1*day))

	b_test.date = now.Add(-2 * day)
	is.False(b_base.IsDupe(b_test, 1*day))
	is.True(b_base.IsDupe(b_test, 2*day))

	// Test account change differences
	b_test = cloneBlock(b_base)
	is.True(b_base.IsDupe(b_test, 0*day))

	b_test.lines = append(b_test.lines, "    Assets:Bank:Chase Checking")
	is.False(b_base.IsDupe(b_test, 0*day))

	b_test = cloneBlock(b_base)
	b_test.lines[2] = "    Assets:Bank:Chase Checking"
	is.True(b_base.IsDupe(b_test, 0*day))
	b_test.lines[2] = "       Assets:Bank:Chase Checking   "
	is.True(b_base.IsDupe(b_test, 0*day))

	b_test.lines[2] = "    Assets:Bank:Zhase Checking"
	is.False(b_base.IsDupe(b_test, 0*day))

	// Test amount change differences
	b_test = cloneBlock(b_base)
	is.True(b_base.IsDupe(b_test, 0*day))

	b_test.lines[1] = "    Everyday Expenses:Medical        $-10.00"
	is.True(b_base.IsDupe(b_test, 0*day))

	b_test.lines[1] = "    Everyday Expenses:Medical        $10.01"
	is.False(b_base.IsDupe(b_test, 0*day))

	b_test.lines[2] = "    Assets:Bank:Chase Checking       $10.00"
	is.False(b_base.IsDupe(b_test, 0*day))

	b_test = cloneBlock(b_base)
	b_test.lines[1], b_test.lines[2] = b_test.lines[2], b_test.lines[1]
	is.True(b_base.IsDupe(b_test, 0*day))
}

func TestClassifyLine(t *testing.T) {
	is := is.New(t)

	tests := []struct {
		line  string
		class int
	}{
		{"", clsBlank},
		{";Comment", clsComment},
		{"2015/03/25 Simple summary", clsSummary},
		{" a posting", clsPosting},
		{" #a comment", clsTxnComment},
	}

	for _, t := range tests {
		//println(t.line)
		cls := classifyLine(t.line)
		is.Equal(cls, t.class)
	}
}

func cloneBlock(orig Block) Block {
	b := Block{
		date:  orig.date,
		lines: make([]string, len(orig.lines)),
	}
	copy(b.lines, orig.lines)

	return b
}
