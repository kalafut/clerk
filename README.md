# Clerk
Personal finance, portable and terse.

# Amount inference
All transactions must balance. If 'Food' increases by $5, there must be $5 worth of decreases elsewhere.

### Commodity handling in Ledger

Ledger will examine the first use of any commodity to determine how that commodity should be printed on reports. It pays attention to whether the name of commodity was separated from the amount, whether it came before or after, the precision used in specifying the amount, whether thousand marks were used, etc. This is done so that printing the commodity looks the same as the way you use it.
