package clerk

/*

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
	type ss map[string]string
	is := is.New(t)

	tests := []struct {
		line  string
		class int
		data  ss
	}{
		{"", clsBlank, nil},
		{";Comment", clsComment, nil},
		{"2015/03/25 Simple summary", clsSummary, ss{"date": "2015/03/25", "cleared": "", "code": ""}},
		{"2015/03/25Simple summary", clsInvalid, nil},
		{"2015/03/25 ! Simple summary", clsSummary, ss{"date": "2015/03/25", "cleared": "!", "code": ""}},
		{"2015/03/25 * Simple summary", clsSummary, ss{"date": "2015/03/25", "cleared": "*", "code": ""}},
		{"2015/03/25 * (523) Simple summary", clsSummary, ss{"date": "2015/03/25", "cleared": "*", "code": "523"}},
		{"2015/03/25   !   (1a2)  Simple summary", clsSummary, ss{"date": "2015/03/25", "cleared": "!", "code": "1a2"}},
		{" a posting", clsPosting, nil},
		{" #a comment", clsTxnComment, nil},
	}

	for _, t := range tests {
		//println(t.line)
		cls, data := classifyLine(t.line)
		is.Equal(cls, t.class)
		is.Equal(len(data), len(t.data))
		for k, _ := range data {
			is.Equal(data[k], t.data[k])
		}
	}
}

func TestNewBlock(t *testing.T) {
	is := is.New(t)

	acctCfg := AccountConfig{
		TargetAccount: "Assets:Bank:Chase Checking",
	}
	transaction := transaction{
		date:        "2006/12/01",
		description: "Test description",
		amount:      "$65.00",
	}

	block := NewBlock(transaction, acctCfg)

	is.Equal("2006/12/01   Test description", block.lines[0])
	is.Equal("    __Uncategorized__          $65.00", block.lines[1])
	is.Equal("    Assets:Bank:Chase Checking", block.lines[2])
}

func cloneBlock(orig Block) Block {
	b := Block{
		date:  orig.date,
		lines: make([]string, len(orig.lines)),
	}
	copy(b.lines, orig.lines)

	return b
}
*/
