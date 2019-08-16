# the trick to this problem is to think about the gradeschool multiplication algorithm
# this algorithm uses left shifts for numbers beyond the least significant digit

def multiply(x: int, y: int) -> int:
    #helper function
    def add(a: int, b: int) -> int:
        running_sum, carryin, k, temp_a, temp_b = 0, 0, 1, a, b
        while temp_a or temp_b:
            ak, bk = a & k, b & k
            carryout = (ak & bk) | (ak & carryin) | (bk & carryin)
            running_sum |= ak ^ bk ^ carryin
            carryin, k, temp_a, temp_b = (carryout << 1, k << 1, temp_a >> 1, temp_b >> 1)
        return running_sum | carryin
    
    running_sum = 0
    while x: #examine each bit of x
        if x & 1:
            running_sum = add(running_sum, y)
        x, y = x >> 1, y << 1
    return running_sum