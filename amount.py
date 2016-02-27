from decimal import Decimal
from collections import OrderedDict

class Amount:
    """Amounts are full precision (rational) values of one or more commodities, e.g. ($4, 34 AAPL). Though most
    quantities in ledgers deal in a single commodity, is it simpler for any Amount to consist of multiple
    commodities. Some support functions assume a single commodity and will complain otherwise."""
    def __init__(self, param):
        if isinstance(param, tuple):
            amt, commodity = param
            self.qtys = OrderedDict({ commodity: Decimal(amt) })
        else:
            assert isinstance(param, dict)
            self.qtys = param

    def __add__(self, addend):
        a,b  = self.qtys, addend.qtys
        ka = set(a.keys())
        kb = set(b.keys())

        qtys = {}

        for k in ka & kb:
            qtys[k] = a[k] + b[k]

        qtys.update({k:a[k] for k in (ka - kb)})
        qtys.update({k:b[k] for k in (kb - ka)})

        return Amount(qtys)

    def __str__(self):
        item = next(iter(self.qtys.items()))
        return item[0]+ " " + str(item[1])

    def get(self, cmdty):
        return Amount((self.qtys[cmdty], cmdty))
