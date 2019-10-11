from typing import List
def reconstruct_queue(people: List[List[int]]) -> List[List[int]]:
    #First, sort the array desc order by height and asc order by position.
    people = sorted(people,key=lambda x:(x[0]*-1,x[1]))
    ans = []
    for p in people:
        ans.insert(p[1],p)
    return ans