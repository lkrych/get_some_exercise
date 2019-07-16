# #merge sort
# def merge_sort(arr):
#   #base case, the arr is sorted if it is of length 1 or 0
#   if len(arr) < 2: 
#     return arr
  
#   #divide the array in two
#   mid_idx = len(arr) // 2
#   left_arr = arr[: mid_idx]
#   right_arr = arr[mid_idx :]

#   #recursively sort the two arrays
#   sorted_left_arr = merge_sort(left_arr)
#   sorted_right_arr = merge_sort(right_arr)

#   #use helper function to merge the two sorted arrays
#   return merge(sorted_left_arr, sorted_right_arr)

# #helper function for merge sort
# def merge(arr1, arr2):
#   merged = []
#   i = 0 #init indices
#   j = 0
#   while len(merged) < len(arr1) + len(arr2):
#     #if either array has been completely checked, add the other array to merged
#     if i >= len(arr1):
#       merged.extend(arr2[j:])
#       return merged
#     elif j >= len(arr2):
#       merged.extend(arr1[i:])
#       return merged

#     #if neither array is empty, check which element is smaller and add to merged
#     if arr1[i] > arr2[j]:
#       merged.append(arr2[j])
#       j+=1
#     else:
#       merged.append(arr1[i])
#       i+=1

#   return merged