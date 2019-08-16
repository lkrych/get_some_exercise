# a brute force approach would iteratively subtract y from x until what remains is less than y
# the number of the subtractions is the quotient
# this is really inefficient if say y equals 1, it will take 2 ^ x iterations

# a better approach is to subtract more in each iteration
# we can compute the largest k such that 2^k*y is less than or equal to x, subtract from x
# we can then add 2^k to the quotient

def divide(x: int, y: int) -> int:
    result, power = 0, 32
    y_power = y << power
    while x >= y:
        while y_power > x:
            y_power >>= 1
            power -=1
        
        result += 1 << power
        x -= y_power
    return result


