package main

import (
	"fmt"
	"time"
)

// Transaction store all data for a transaction. It should be treated as immutable. Use Set*()
// function to receive a new, modified Transaction.
type Tx struct {
	date     time.Time
	summary  string
	postings []Posting
	note     string
}

// TxReader is the common interface for creating Tx objects from a persistent store.
type TxReader interface {
	Read(root *Account) []*Tx
}

type TxWriter interface {
	Write([]*Tx)
}

type Posting struct {
	Acct *Account
	Amt  Amount
}

func (p Posting) String() string {
	return fmt.Sprintf("%s %v", p.Acct.Name, p.Amt)
}

func NewTransaction(date time.Time, summary string, postings []Posting, note string) *Tx {
	// First check whether postings are balanced. This will never be false if multiple
	// commodities are involved since there are implicit conversions.
	sum := Amount{}

	for _, p := range postings {
		sum.Add(p.Amt)
	}

	if len(sum) <= 1 && !sum.Zero() {
		fatalf("Unbalanced postings: %v", postings)
	}

	t := Tx{
		date:    date,
		summary: summary,
		note:    note,
	}
	t.postings = make([]Posting, len(postings))
	copy(t.postings, postings)

	return &t
}

func (t Tx) Date() time.Time {
	return t.date
}
func (t Tx) Summary() string {
	return t.summary
}
func (t Tx) Postings() []Posting {
	return t.postings
}
func (t Tx) Note() string {
	return t.note
}

func (t Tx) SetDate(date time.Time) *Tx {
	return NewTransaction(date, t.summary, t.postings, t.note)
}
func (t Tx) SetSummary(summary string) *Tx {
	return NewTransaction(t.date, summary, t.postings, t.note)
}
func (t Tx) SetPostings(postings []Posting) *Tx {
	return NewTransaction(t.date, t.summary, postings, t.note)
}
func (t Tx) SetNote(note string) *Tx {
	return NewTransaction(t.date, t.summary, t.postings, note)
}

// Equal tests whether two transactions are equal according to the given
// level of strictness:
//
//
//func Equal(a, b _Transaction, strictness int) bool {
//	/*
//		balancesA := map[Account]*big.Rat{}
//		balancesB := map[Account]*big.Rat{}
//
//		for _,t:=range a.entries {
//		}
//	*/
//	return true
//}

//func ParseTransactions(in io.Reader) []*Tx {
//	trans := []*Tx{}
//	r := csv.NewReader(in)
//
//	for {
//		record, err := r.Read()
//		if err == io.EOF {
//			break
//		}
//
//		date, err := time.Parse(StdDate, record[0])
//		if err != nil {
//			log.Fatal(err)
//		}
//
//		t := NewTransaction(
//			date,
//			strings.TrimSpace(record[1]),
//			parsePostings(record[2]),
//			"",
//		)
//		trans = append(trans, t)
//	}
//
//	return trans
//}

/*
func parsePostings(p string) []Posting {
	var comm string
	postings := []Posting{}

	for _, posting := range strings.Split(p, "&") {
		posting = strings.TrimSpace(posting)
		match := rePosting2.FindStringSubmatch(posting)

		if len(match) == 0 {
			log.Fatalf("Invalid posting: %s", posting)
		}

		result := make(map[string]string)
		for i, name := range rePosting2.SubexpNames() {
			if i != 0 {
				result[name] = match[i]
			}
		}

		c1, c2 := result["comm1"], result["comm2"]

		switch {
		case c1 != "" && c2 != "":
			log.Fatalf("Multiple commmodities in posting: %s", posting)
		case c1 != "":
			comm = c1 // TODO: use a commodity pool instead, else "$ 1" is different than "1 $"
		case c2 != "":
			comm = c2
		default:
			comm = DefaultCommodity
		}

		r := new(big.Rat)
		r.SetString(result["amount"])
		p := Posting{
			Acct: RootAccount.FindOrAddAccount(result["account"]),
			Amt:  NewAmount(result["amount"], comm),
		}

		postings = append(postings, p)
	}
	checkBalance(postings)
	return postings
}
*/

/*
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
*/
