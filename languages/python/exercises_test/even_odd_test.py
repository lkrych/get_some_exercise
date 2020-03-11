class TestEvenOdd(unittest.TestCase):
    def test_small_list(self):
        """
        Test that it can sort a small list of integers
        """
        data = [5, 6, 4, 3, 2, 1,]
        result = even_odd(data)
        self.assertEqual(result, [2, 6, 4, 3, 1, 5])

    def test_bigger_list(self):
        """
        Test that it can sort a bigger list of integers
        """
        data = [5, 6, 4, 3, 7, 8, 9, 2, 1, 10]
        result = even_odd(data)
        self.assertEqual(result, [10, 6, 4, 2, 8, 9, 7, 1, 3, 5])