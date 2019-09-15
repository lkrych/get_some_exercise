# The parity of a binary value is 1 if the number of 1s in the word is odd 
def parity(x:int) -> int:
    result = 0
    while x > 0:
        result ^= x & 1
        x = x >> 1
    return result
 
