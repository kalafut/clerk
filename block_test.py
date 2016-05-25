from datetime import date
import block

def test_parse():
    sample = """
2006/10/15 McDonald's
    Expenses:Dining   $5.36
    Assets:Checking


2007/01/01 * Burger King
    Expenses:Dining   $5.36
    Assets:Checking
    """

    blocks = block.parse(sample.split("\n"))

    assert len(blocks) == 2
    assert blocks[0].date == date(2006, 10, 15)
    assert blocks[1].date == date(2007, 1, 1)
    assert blocks[0].summary == "McDonald's"
    assert blocks[1].summary == "Burger King"
    assert blocks[0].cleared is None
    assert blocks[1].cleared == "*"
    assert blocks[0].postings[0].account == "Expenses:Dining"
    assert blocks[0].postings[0].amount == "$5.36"
    assert blocks[0].postings[1].account == "Assets:Checking"
    assert blocks[0].postings[1].amount is None
