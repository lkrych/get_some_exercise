class TestReverseBits(unittest.TestCase):
    def test_5(self):
        """
        Test that reverse_bits can handle small integers
        """
        result = reverse_bits(5)
        self.assertEqual(result, 5)

    def test_6(self):
        """
        Test that reverse_bits can handle small integers
        """
        result = reverse_bits(6)
        self.assertEqual(result, 3)

    def test_243(self):
        """
        Test that reverse_bits can handle small integers
        """
        result = reverse_bits(243)
        self.assertEqual(result, 207)
    def test_65536(self):
        """
        Test that reverse_bits can handle larger integers
        """
        result = reverse_bits(65536)
        self.assertEqual(result, 1)
