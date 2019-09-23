from typing import List

def merge_sort(arr: List[int]) -> List[int]:
    if len(arr) <= 1:
        return arr
    mid_idx = len(arr) // 2
    left_arr = merge_sort(arr[:mid_idx])
    right_arr = merge_sort(arr[mid_idx:])
    return merge(left_arr, right_arr)

    

def merge(arr1: List[int], arr2: List[int]) -> List[int]:
    merged = []
    i,j = 0,0
    while i < len(arr1) and j < len(arr2):
        if arr1[i] <= arr2[j]:
            merged.append(arr1[i])
            i += 1
        else:
            merged.append(arr2[j])
            j += 1
    merged.extend(arr1[i:])
    merged.extend(arr2[j:])
    return merged
