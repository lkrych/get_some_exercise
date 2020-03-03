class TestIntPal(unittest.TestCase):
    def test_5(self):
        """
        Test that int_palindrome can handle small integers
        """
        result = int_palindrome(5)
        self.assertEqual(result, True)

    def test_66(self):
        """
        Test that int_palindrome can handle small integers
        """
        result = int_palindrome(66)
        self.assertEqual(result, True)

    def test_84(self):
        """
        Test that int_palindrome can handle small integers
        """
        result = int_palindrome(84)
        self.assertEqual(result, False)
    def test_65536(self):
        """
        Test that int_palindrome can handle larger integers
        """
        result = int_palindrome(65536)
        self.assertEqual(result, False)
    def test_65556(self):
        """
        Test that int_palindrome can handle larger integers
        """
        result = int_palindrome(65556)
        self.assertEqual(result, True)

    def test_10000000000001(self):
        """
        Test that int_palindrome can handle larger integers
        """
        result = int_palindrome(10000000000001)
        self.assertEqual(result, True)