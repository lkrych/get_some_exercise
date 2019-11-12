class TestReverse(unittest.TestCase):
    def test_small_list(self):
        """
        Test that it can sort a small list of integers
        """
        data = [5, 6, 4, 3, 2, 1,]
        reverse(data) #mutates data
        self.assertEqual(data, [1,2,3,4,6,5])

    def test_bigger_list(self):
        """
        Test that it can sort a bigger list of integers
        """
        data = [5, 6, 4, 3, 7, 8, 9, 2, 1, 10]
        reverse(data) #mutates data
        self.assertEqual(data, [10, 1, 2, 9, 8, 7, 3, 4, 6, 5])