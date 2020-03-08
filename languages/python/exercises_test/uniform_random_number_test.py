import random
random.seed(30)

class TestUniformRandom(unittest.TestCase):
    def test_1(self):
        """
        Test that uniform_random_number can handle small integers
        """
        result = uniform_random_number(1, 10)
        self.assertEqual(result, 10)

    def test_2(self):
        """
        Test that uniform_random_number can handle small integers
        """
        result = uniform_random_number(8, 44)
        self.assertEqual(result, 33)

    def test_3(self):
        """
        Test that uniform_random_number can handle small integers
        """
        result = uniform_random_number(36, 37)
        self.assertEqual(result, 36)
