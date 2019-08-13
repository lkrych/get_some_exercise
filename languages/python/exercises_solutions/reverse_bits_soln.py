# this solution is brute-force, but not bad, it is constant time
# if we need to do this operation repeatedly, a cleverer solution is to use a cache
# this cache will be created with 16-bit words, that are created by splitting up x
# into 4 parts: y3, y2, y1, y0.
# Now we just need to do four cache lookups and put them in the correct order
#
#
# here though is the brute force solution
#  def reverse_bits(x: int) -> int:
#     rev = 0
#     # traversing bits of 'n' from the right 
#     while (n > 0) : 
          
#         # bitwise left shift 'rev' by 1 
#         rev = rev << 1
          
#         # if current bit is '1' 
#         if (n & 1 == 1) : 
#             rev = rev ^ 1
          
#         # bitwise right shift 'n' by 1 
#         n = n >> 1
          
#     # required number 
#     return rev 