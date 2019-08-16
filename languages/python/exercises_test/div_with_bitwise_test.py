class TestDivide(unittest.TestCase):
    def test_5(self):
        """
        Test that divide can handle small integers
        """
        result = divide(10, 2)
        self.assertEqual(result, 5)

    def test_6(self):
        """
        Test that divide can handle small integers
        """
        result = divide(15, 3)
        self.assertEqual(result, 5)

    def test_8(self):
        """
        Test that divide can handle small integers
        """
        result = divide(360, 8)
        self.assertEqual(result, 45)
    def test_65536(self):
        """
        Test that divide can handle larger integers
        """
        result = divide(65536, 4)
        self.assertEqual(result, 16384)