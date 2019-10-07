def multiply_strings(num1: str, num2: str) -> str:
    products = []
    bigger = '0'
    smaller = '0'
    if len(num1) > len(num2):
        bigger = num1
        smaller = num2
    else:
        smaller = num1
        bigger = num2
    for idx1, i in enumerate(bigger[::-1]):
        for idx2, j in enumerate(smaller[::-1]):
            x = i
            y = j
            zeroes = ('0' * idx1) + ('0' * idx2) 
            products.append(multiply(x, y, zeroes))
    ans = 0
    for el in products:
        ans += int(el)
    return ans

#helper function for multiplying two strings
def multiply(x: str, y: str, z: str = '') -> str:
    return str(int(x) * int(y)) + z