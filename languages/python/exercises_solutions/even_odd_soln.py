from typing import List
#order the arr so that even integers come first
def even_odd(arr: List[int]) -> List[int]:
    next_even, next_odd = 0, len(arr) -1
    while next_even < next_odd:
        if arr[next_even] % 2 == 0:
            next_even += 1
        else:
            arr[next_even], arr[next_odd] = arr[next_odd], arr[next_even]
            next_odd -= 1
