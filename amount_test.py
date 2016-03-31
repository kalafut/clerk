import decimal
import pytest

from amount import Amount as A

def test_basic():
    a = A(("0.30", "$"))
    assert '$ 0.30' == str(a)

    a = A({"$": decimal.Decimal(4)})
    assert '$ 4', str(a)

def test_add():
    a = A(("2.34", "$"))
    b = A(("5.97", "$"))
    assert "$ 8.31" == str(a+b)

    c = A(("9.01", "CAD")) + A(("15.56", "$"))
    assert "CAD 9.01" == str(c.get("CAD"))
    assert "$ 15.56" == str(c.get("$"))

    d = a + c
    assert "$ 17.90" == str(d.get("$"))
    assert "CAD 9.01" == str(d.get("CAD"))
