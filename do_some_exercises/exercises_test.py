import unittest

from exercises import *

class TestParity(unittest.TestCase):
    def test_5(self):
        """
        Test that parity can handle small integers
        """
        result = parity(5)
        self.assertEqual(result, 0)

    def test_6(self):
        """
        Test that parity can handle small integers
        """
        result = parity(6)
        self.assertEqual(result, 0)

    def test_8(self):
        """
        Test that parity can handle small integers
        """
        result = parity(8)
        self.assertEqual(result, 1)
    def test_65536(self):
        """
        Test that parity can handle larger integers
        """
        result = parity(65536)
        self.assertEqual(result, 1)
 


if __name__ == '__main__':
	unittest.main()