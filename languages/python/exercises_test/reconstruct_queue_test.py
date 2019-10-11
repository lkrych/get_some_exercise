class TestReconstructQueue(unittest.TestCase):
    def test_5(self):
        """
        Test that reconstruct_queue can handle small input
        """
        result = reconstruct_queue([[7,0], [4,4], [7,1], [5,0], [6,1], [5,2]])
        self.assertEqual(result, 1)

    def test_6(self):
        """
        Test that reconstruct_queue can handle larger input
        """
        result = reconstruct_queue([2, 10], [12, 0], [6, 4], [10,0], [8,1], [8,0], [5,2], [3,2], [1,2], [7,0], [4,0])
        self.assertEqual(result, 3)

    def test_243(self):
        """
        Test that reconstruct_queue can handle even larger input
        """
        result = reconstruct_queue([2, 10], [12, 0], [6, 4], [10,0], [8,1], [8,0], [5,2], [3,2], [1,2], [7,0], [4,0], [15,0], [7,6])
        self.assertEqual(result, 207)
