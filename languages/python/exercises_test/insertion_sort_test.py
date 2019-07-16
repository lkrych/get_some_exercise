class TestInsertionSort(unittest.TestCase):
    def test_small_list(self):
        """
        Test that it can sort a small list of integers
        """
        data = [5, 6, 4, 3, 2, 1,]
        result = insertion_sort(data)
        self.assertEqual(result, [1,2,3,4,5,6])

    def test_bigger_list(self):
        """
        Test that it can sort a bigger list of integers
        """
        data = [5, 6, 4, 3, 7, 8, 9, 2, 1, 10]
        result = insertion_sort(data)
        self.assertEqual(result, [1,2,3,4,5,6,7,8,9,10])