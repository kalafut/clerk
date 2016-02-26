from unittest import TestCase, main
from amount import Amount as A

class AmountTestCase(TestCase):
    def test_basic(self):
        a = A("0.30", "$")
        self.assertEqual('$ 0.30', str(a))

if __name__ == '__main__':
    main()
