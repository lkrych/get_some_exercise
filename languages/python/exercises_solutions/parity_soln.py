# The parity of a binary value is 1 if the number of 1s in the word is odd 

# brute-force O(n)
#
# Algorithm uses XOR and AND to track the number of 1s seen so far
# the result will only flip when a 1 is found

# def parity(x):
#     result = 0
#     while x:
#         result ^= x & 1
#         x >>= 1
#     return result


# similar approach with bit-erasing O(k), where k is the number of set bits

# def parity(x):
#     result = 0
#     while x:
#         result ^= 1
#         x &= x - 1 # drops the lowest set bit of x
#     return result