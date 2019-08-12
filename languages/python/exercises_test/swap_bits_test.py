class TestSwapBits(unittest.TestCase):
    def test_5(self):
        """
        Test that swap_bits can handle small integers
        """
        result = swap_bits(5, 5, 12)
        self.assertEqual(result, 4097)

    def test_6(self):
        """
        Test that swap_bits can handle small integers
        """
        result = swap_bits(6, 6, 1)
        self.assertEqual(result, 68)

    def test_8(self):
        """
        Test that swap_bits can handle small integers
        """
        result = swap_bits(8, 3 , 8)
        self.assertEqual(result, 256)
    def test_65536(self):
        """
        Test that swap_bits can handle larger integers
        """
        result = swap_bits(65536, 16, 17)
        self.assertEqual(result, 131702)