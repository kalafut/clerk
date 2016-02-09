package main

import (
	"bytes"
	"fmt"
	"log"
	"math/big"
	"reflect"
)

const DefaultCommodity = "$"
const StdDate = "2006/01/02"

// Amounts are full precision (rational) values of one or more commodities, e.g. ($4, 34 AAPL). Though most
// quantities in ledgers deal in a single commodity, is it simpler for any Amount to consist of multiple
// commodities. Some support functions assume a single commodity and will complain otherwise.
type Amount map[string]*big.Rat

func NewAmount(qty interface{}, cmdty string) Amount {
	r := new(big.Rat)

	switch qty.(type) {
	case int:
		r.SetInt64(int64(qty.(int)))
	case float64:
		r.SetFloat64(qty.(float64))
	case string:
		if _, ok := r.SetString(qty.(string)); !ok {
			log.Fatalf("Invalid Amount(qty: %s  cmdty: %s)", qty, cmdty)
		}
	default:
		log.Fatalf("Invalid NewAmount type: %v", reflect.TypeOf(qty))
	}

	return Amount{cmdty: r}
}

func (amt Amount) Add(incr Amount) Amount {
	for cmdty, val := range incr {
		curval, ok := amt[cmdty]
		if !ok {
			curval = big.NewRat(0, 1)
		}
		curval.Add(curval, val)
		amt[cmdty] = curval
	}

	return amt
}

func (amt Amount) Zero() bool {
	for _, val := range amt {
		if val.Num().Int64() != 0 {
			return false
		}
	}

	return true
}

func (amt Amount) String() string {
	var b bytes.Buffer

	s := amt.Strings()
	for _, v := range s {
		fmt.Fprint(&b, v)
	}

	return fmt.Sprintf("(%s)", b.String())
}

func (amt Amount) Strings() map[string]string {
	strs := map[string]string{}

	for com, val := range amt {
		if com == "$" { // hack
			strs[com] = fmt.Sprintf("%s %s", com, val.FloatString(2))
		} else {
			strs[com] = fmt.Sprintf("%s %s", val.FloatString(2), com)
		}
	}

	return strs
}

// Equal returns true if the two Amounts are equal. This tests all values
// of all commodities. If lenient is true, amounts that have different
// commodities but all of those have a value of 0 are still considered
// equal (e.g. ($1, 0 AAPL) == ($1) only in lenient mode)
func (a1 Amount) Equal(a2 Amount, lenient bool) bool {
	checked := make(map[string]struct{})

	equal := func(a1, a2 Amount) bool {
		for cmdty, val := range a1 {
			if _, alreadyChecked := checked[cmdty]; !alreadyChecked {
				if v2, ok := a2[cmdty]; ok {
					if val.Cmp(v2) != 0 {
						return false
					}
				} else {
					if !lenient || val.Num().Int64() != 0 {
						return false
					}
				}
				checked[cmdty] = struct{}{}
			}
		}

		return true
	}

	return equal(a1, a2) && equal(a2, a1)
}
