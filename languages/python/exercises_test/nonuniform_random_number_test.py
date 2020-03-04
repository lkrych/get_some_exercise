import random
random.seed(30)

class TestNonUniformRandom(unittest.TestCase):
    def test_1(self):
        """
        Test that nonuniform_random_number can handle small integers
        """
        result = nonuniform_random_number([3, 5, 7, 11], [9/18,6/18,2/18,1/18])
        self.assertEqual(result, 5)

    def test_2(self):
        """
        Test that nonuniform_random_number can handle small integers
        """
        result = nonuniform_random_number([3, 5, 7, 11, 13], [9/18,6/18,1/18,1/18, 1/18])
        self.assertEqual(result, 3)

    def test_3(self):
        """
        Test that nonuniform_random_number can handle small integers
        """
        result = nonuniform_random_number([3, 5], [9/18,9/18])
        self.assertEqual(result, 3)
