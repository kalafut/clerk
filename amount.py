from decimal import Decimal
from collections import OrderedDict

class Amount:
    """Amounts are full precision (rational) values of one or more commodities, e.g. ($4, 34 AAPL). Though most
    quantities in ledgers deal in a single commodity, is it simpler for any Amount to consist of multiple
    commodities. Some support functions assume a single commodity and will complain otherwise."""
    def __init__(self, amt, commodity):
        self.qtys = OrderedDict({ commodity: Decimal(amt) })

    def __add__(self, b):
        ka = set(self.qtys.keys())
        kb = set(b.qtys.keys())
        common =  set(ka).union(set(kb))

        for k in common:
            print(k)

        return self

    def __str__(self):
        item = next(iter(self.qtys.items()))
        return item[0]+ " " + str(item[1])
