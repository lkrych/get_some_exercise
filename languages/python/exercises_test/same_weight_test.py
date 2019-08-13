class TestSameWeightBits(unittest.TestCase):
    def test_5(self):
        """
        Test that same_weight can handle small integers
        """
        result = same_weight(5)
        self.assertEqual(result, 6)

    def test_6(self):
        """
        Test that same_weight can handle small integers
        """
        result = same_weight(6)
        self.assertEqual(result, 5)

    def test_243(self):
        """
        Test that same_weight can handle small integers
        """
        result = same_weight(243)
        self.assertEqual(result, 245)
    def test_65536(self):
        """
        Test that same_weight can handle larger integers
        """
        result = same_weight(65536)
        self.assertEqual(result, 32768)