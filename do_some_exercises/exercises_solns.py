from typing import List
import bisect
import itertools
import random
# 1. create an array of cumulative sums of the probabilities
# 2. generate a random number and use binary search to return where it would be placed in array
# 3. use index to profit

def nonuniform_random_number(values: List[int], probabilities: List[float]) -> int:
    cumulative_sums = [0.0] + list(itertools.accumulate(probabilities))
    idx = bisect.bisect(cumulative_sums, random.random()) - 1
    return values[idx]
 
