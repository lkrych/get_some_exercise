# write a program that multiples two nonnegative integers, the only operators you are allowed
# to use are assignment, bitwise operators, equality checks and Boolean combinations
def multiply(x: int, y: int) -> int:
    def add(a: int, b: int) -> int:
        running_sum, carryin, k, temp_a, temp_b = 0, 0, 1, a, b
        while temp_a or temp_b:
            ak, bk = a & k, b & k #check what the 1s place digit are of running_sum and y
            carryout = (ak & bk) | (ak & carryin) | (bk & carryin) #calculate what needs to be carried
            running_sum |= ak ^ bk ^ carryin
            carryin, k, temp_a, temp_b = (carryout << 1, k << 1, temp_a >> 1, temp_b >> 1)
        return running_sum | carryin

    running_sum = 0
    while x:
        if x & 1:
            running_sum = add(running_sum, y)
        x, y = x >> 1, y << 1
    return running_sum
 
