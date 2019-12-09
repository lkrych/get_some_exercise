#take a 64-bit word and return the 64-bit word consisting of the bits of the input word in reverse order

def reverse_bits(x: int) -> int:
    rev = 0
    
    while x > 0:
        rev = rev << 1
        if (x & 1 == 1):
            rev = rev ^ 1

        x = x >> 1
    
    return rev
 
