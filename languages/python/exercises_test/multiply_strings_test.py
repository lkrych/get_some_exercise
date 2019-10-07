class TestMultiplyStrings(unittest.TestCase):
    def test_5(self):
        """
        Test that multiply_strings can handle small integers
        """
        result = multiply_strings('5', '6')
        self.assertEqual(result, '30')

    def test_6(self):
        """
        Test that multiply_strings can handle small integers
        """
        result = multiply_strings('3', '8')
        self.assertEqual(result, '24')

    def test_8(self):
        """
        Test that multiply_strings can handle small integers
        """
        result = multiply_strings('8', '245')
        self.assertEqual(result, '1960')
    def test_65536(self):
        """
        Test that multiply_strings can handle larger integers
        """
        result = multiply_strings('65536', '4')
        self.assertEqual(result, '262144')