from decimal import Decimal

class Amount:
    """Amounts are full precision (rational) values of one or more commodities, e.g. ($4, 34 AAPL). Though most
    quantities in ledgers deal in a single commodity, is it simpler for any Amount to consist of multiple
    commodities. Some support functions assume a single commodity and will complain otherwise."""
    def __init__(self, amt, commodity):
        self.primary = commodity
        self.qtys = { commodity: Decimal(amt) }

    def __add__(a,b):
        ka = a.qtys.keys()
        kb = b.qtys.keys()
        common = set(ka).union(set(kb))

        for k in common:
            print(k)

        return a

    def __str__(self):
        return str(self.qtys[self.primary])

a = Amount("4.54", "$")
b = Amount("2", "$")
print(a)
print(a+b)
