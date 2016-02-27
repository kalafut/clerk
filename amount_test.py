import decimal

from unittest import TestCase, main
from amount import Amount as A

class AmountTestCase(TestCase):
    def test_basic(self):
        a = A(("0.30", "$"))
        self.assertEqual('$ 0.30', str(a))

        a = A({"$": decimal.Decimal(4)})
        self.assertEqual('$ 4', str(a))

    def test_add(self):
        a = A(("2.34", "$"))
        b = A(("5.97", "$"))
        self.assertEqual("$ 8.31", str(a+b))

        c = A(("9.01", "CAD")) + A(("15.56", "$"))
        self.assertEqual("CAD 9.01", str(c.get("CAD")))
        self.assertEqual("$ 15.56", str(c.get("$")))

        d = a + c
        self.assertEqual("$ 17.90", str(d.get("$")))
        self.assertEqual("CAD 9.01", str(d.get("CAD")))

if __name__ == '__main__':
    main()
