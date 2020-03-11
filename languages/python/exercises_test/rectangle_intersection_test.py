import collections
Rectangle = collections.namedtuple('Rectangle', ('x', 'y', 'width', 'height'))

class TestRectangeIntersection(unittest.TestCase):
    def test_1(self):
        """
        Test that it can find an intersection for the same rectangle 
        """
        r1 = Rectangle(1, 2, 2, 2)
        r2 = Rectangle(1, 2, 2, 2)
        result = rectangle_intersection(r1, r2)
        self.assertEqual(result.x, 1)
        self.assertEqual(result.y, 2)
        self.assertEqual(result.width, 2)
        self.assertEqual(result.height, 2)

    def test_2(self):
        """
        Test that it can find an intersection
        """
        r1 = Rectangle(1, 2, 2, 2)
        r2 = Rectangle(3, 1, 2, 2)
        result = rectangle_intersection(r1, r2)
        self.assertEqual(result.x, 3)
        self.assertEqual(result.y, 2)
        self.assertEqual(result.width, 0)
        self.assertEqual(result.height, 1)

    def test_3(self):
        """
        TTest that it can find not find an intersection
        """
        r1 = Rectangle(1, 2, 2, 2)
        r2 = Rectangle(4, 2, 2, 2)
        result = rectangle_intersection(r1, r2)
        self.assertEqual(result.x, 0)
        self.assertEqual(result.y, 0)
        self.assertEqual(result.width, -1)
        self.assertEqual(result.height, -1)