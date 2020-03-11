import collections
Rectangle = collections.namedtuple('Rectangle', 'x', 'y', 'width', 'height')
# There are a lot of possiblities for intersections, so the best solution
# to this problem is to focus on conditions where the rectangles DO NOT intersect. 
def rectangle_intersection(r1: Rectangle, r2: Rectangle) -> Rectangle:
    def is_intersect(r1, r2):
        return (r1.x <= r2.x + r2.width and r1.x + r1.width >= r2.x and
                r1.y <= r2.y + r2.height and r1.y + r1.height >= r2.y )
    
    if not is_intersect(r1, r2):
        return Rectangle(0, 0, -1, -1)
    return Rectangle(
        max(r1.x, r2.x),
        max(r1.y, r2.y),
        min(r1.x + r1.width, r2.x + r2.width) - max(r1.x, r2.x),
        min(r1.y + r1.height, r2.y + r2.height) - max(r1.y, r2.y),
    )
    