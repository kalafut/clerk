package core

import "math/big"

var DefaultCommodity = Commodity{Abbr: "$"}

type Commodity struct {
	Abbr    string
	Postfix bool
}

func (c Commodity) String() string {
	return c.Abbr
}

// Amounts are full precision (rational) values of one or more commodities, e.g. ($4, 34 AAPL). Though most
// quantities in ledgers deal in a single commodity, is it simpler for any Amount to consist of multiple
// commodities. Some support functions assume a single commodity and will complain otherwise.
type Amount map[Commodity]*big.Rat

func NewAmount(qty string, cmdty Commodity) Amount {
	r := new(big.Rat)
	r.SetString(qty)

	amt := Amount{}
	amt[cmdty] = r

	return amt
}

func (amt Amount) Add(incr Amount) {
	for cmdty, val := range incr {
		curval, ok := amt[cmdty]
		if !ok {
			curval = big.NewRat(0, 1)
		}
		curval.Add(curval, val)
		amt[cmdty] = curval
	}
}

func (amt Amount) Zero() bool {
	for _, val := range amt {
		if val.Num().Int64() != 0 {
			return false
		}
	}

	return true
}
