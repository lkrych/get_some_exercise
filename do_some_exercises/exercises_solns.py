# #bubble sort
# def bubble_sort(arr):
#   sorted = False
#   while not sorted: #start bubbling
#     sorted = True #sort until there are no more swaps
#     for i, el in enumerate(arr):
#       #avoid bounds error with this check
#       if i < len(arr) - 1:
#         if el > arr[i+1]:
#           swap(arr, i, i+1)
#           sorted = False
#   return arr

# #helper function for bubble sort
# def swap(arr, i, j):
#   temp = arr[i]
#   arr[i] = arr[j]
#   arr[j] = temp
#   return arr
 
