// func mergeSort(arr []int) []int {
// 	//make copy of array so you don't mutate input
// 	tmp := make([]int, len(arr))
// 	copy(tmp, arr)
// 	if len(tmp) <= 1 {
// 		return tmp
// 	}
// 	mid := len(tmp) / 2
// 	left := mergeSort(tmp[0:mid])
// 	right := mergeSort(tmp[mid:])

// 	return merge(left, right)
// }

// func merge(left []int, right []int) []int {
// 	sorted := []int{}
// 	for len(left) > 0 && len(right) > 0 {
// 		if left[0] < right[0] {
// 			sorted = append(sorted, left[0])
// 			left = left[1:]
// 		} else {
// 			sorted = append(sorted, right[0])
// 			right = right[1:]
// 		}
// 	}
// 	sorted = append(sorted, left...)
// 	sorted = append(sorted, right...)
// 	return sorted
// }