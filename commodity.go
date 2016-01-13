package main

var defaultCommodity = Commodity{abbr: "$"}

type Commodity struct {
	abbr    string
	postfix bool
}

func (c Commodity) String() string {
	return c.abbr
}
