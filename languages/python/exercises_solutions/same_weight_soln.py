# One not bad idea is to swap the LSB with the rightmost bit that differs with it
# This yields the correct answer for some results
# A variation on this technique has us flip two bits and minimize the difference between these two
# Thus we swap the two rightmost consecutive bits that differ
# def same_weight(x: int) -> int:
#     NUM_UNSIGNED_BITS = 64
#     for i in range(NUM_UNSIGNED_BITS - 1):
#         if (x >> i) & 1 != (x >> (i + 1)) & 1:
#             x ^= (1 << i) | (1 << (i + 1)) #swap bits i and i+1
#             return x
#     raise ValueError('All bits are 0 or 1')