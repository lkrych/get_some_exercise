from typing import List
#reverse the list in place using recursion
def reverse( s: List[str]) -> None:
    length = len(s)
    def helper(left, right): 
        if left < right:
            s[left], s[right] = s[right], s[left]
            helper(left + 1, right - 1)
    helper(0, length - 1)
