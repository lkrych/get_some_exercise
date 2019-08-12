# def swap_bits(x: int, i: int, j: int) -> int:
#     #Extract the ith and jth bits and see if they differ
#     if (x >> i) & 1 != (x >> j) & 1:
#         # swap by flipping their values
#         # select bits to flip with bit_mask
#         bit_mask = (1 << i) | (1 << j) #bitwise or of the two masks
#         x ^= bit_mask
#     return x