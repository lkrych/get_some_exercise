import collections
Rectangle = collections.namedtuple('Rectangle', ('x', 'y', 'width', 'height'))
# given two rectangles, return the rectangle of their intersection
# if no intersection exists, return Rectangle(0, 0, -1, -1)
def rectangle_intersection(r1: Rectangle, r2: Rectangle) -> Rectangle:
    return Rectangle(0, 0, -1, -1)