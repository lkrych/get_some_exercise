#given two positive integers, compute their quotient using only the addition, subtraction and shifting operators
def divide(x: int, y: int) -> int:
    quotient = 1
    while x > y:
        x -= y
        quotient += 1
    return quotient
 
