#given two positive integers, compute their quotient using only the addition, subtraction and shifting operators
def divide(x: int, y: int) -> int:
    q = 0
    while x >= y:
        x -= y
        q += 1
    return q
 
