package core

import (
	"fmt"
	"math/big"
)

type Commodity string

// Amounts are full precision (rational) values of one or more commodities, e.g. ($4, 34 AAPL). Though most
// quantities in ledgers deal in a single commodity, is it simpler for any Amount to consist of multiple
// commodities. Some support functions assume a single commodity and will complain otherwise.
type Amount map[Commodity]*big.Rat

const DefaultCommodity = Commodity("$")
const StdDate = "2006/01/02"

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

func (amt Amount) Strings() []string {
	strs := []string{}

	for com, val := range amt {
		if com == "$" { // hack
			strs = append(strs, fmt.Sprintf("%s %s", com, val.FloatString(2)))
		} else {
			strs = append(strs, fmt.Sprintf("%s %s", val.FloatString(2), com))
		}
	}

	return strs
}
