import random
# solution: flip a coin each time around the loop
# shift the value of the coin onto the result
# Once you've shifted more then number of outcome times (inside loop)
# check to see if your shifted number is less than or equal to number of outcomes
# lastly add the lower_bound. This ensures you add up to the upper bound because your range
# is number_of_outcomes
def uniform_random_number(lower: int, upper: int) -> int:
    number_of_outcomes = upper - lower + 1
    while True:
        result, i = 0, 0
        while (1 << i) < number_of_outcomes:
            result = (result << 1) | random.randint(0,1)
            i += 1
        if result < number_of_outcomes:
            break
    return result + lower