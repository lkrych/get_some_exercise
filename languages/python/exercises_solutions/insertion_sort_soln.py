# #insertion sort
# def insertion_sort(arr):
#   insertion_idx = 0
#   while insertion_idx < len(arr):
#     insertion_el = arr[insertion_idx]
#     for i, el in enumerate(arr[:insertion_idx]): #loop through every el within the range of insertion_idx
#       if i + 1 >= insertion_idx: #compare el to the elements within the range of insertion_idx
#         if arr[insertion_idx] < el:
#           #remove insertion el and place it where it needs to go
#           del arr[insertion_idx]
#           arr.insert(i, insertion_el)
#           #you don't need to check anymore elements
#           break
#     insertion_idx += 1
#   return arr
          