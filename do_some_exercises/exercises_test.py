import unittest

from exercises import *

class TestMultiply(unittest.TestCase):
    def test_5(self):
        """
        Test that multiply can handle small integers
        """
        result = multiply(5, 6)
        self.assertEqual(result, 30)

    def test_6(self):
        """
        Test that multiply can handle small integers
        """
        result = multiply(3, 8)
        self.assertEqual(result, 24)

    def test_8(self):
        """
        Test that multiply can handle small integers
        """
        result = multiply(8, 245)
        self.assertEqual(result, 1960)
    def test_65536(self):
        """
        Test that multiply can handle larger integers
        """
        result = multiply(65536, 4)
        self.assertEqual(result, 262144)

 


if __name__ == '__main__':
	unittest.main()