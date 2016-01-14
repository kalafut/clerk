package main

import "math/big"

var defaultCommodity = Commodity{abbr: "$"}

type Commodity struct {
	abbr    string
	postfix bool
}

func (c Commodity) String() string {
	return c.abbr
}

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
