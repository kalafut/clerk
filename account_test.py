from unittest import TestCase, main
import account

class AccountTestCase(TestCase):
    def test_basic(self):
        a = account.find_or_create("A")
