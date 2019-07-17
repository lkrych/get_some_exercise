def quick_sort(arr):
  if len(arr) < 2:
    return arr

  pivot = arr[0]
  left = []
  right = []

  for el in arr[1:]:
    if el <= pivot:
      left.append(el)
    else:
      right.append(el)
  
  sorted_left = quick_sort(left)
  sorted_right = quick_sort(right)

  return sorted_left + [pivot] + sorted_right