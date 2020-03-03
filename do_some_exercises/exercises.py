import math
def int_palindrome(x: int) -> bool:
    if x < 0:
        return False

    while x > 0:
        number_of_digits = int(math.log10(x)) + 1
        mask = 10 ** (number_of_digits - 1)
        #compare least significant digit and most significant digit
        lsd = x % 10
        msd = x // mask
        if lsd != msd:
            return False
        
        #shave off last digit and first digit
        x = x % mask
        x = x // 10
        
    return True
