// func partitionEvenOdd(arr []int) []int {
// 	partition := 0
// 	for i := 0; i < len(arr); i++ {
// 		if arr[i]%2 == 0 {
// 			arr[i], arr[partition] = arr[partition], arr[i]
// 			partition++
// 		}
// 	}
// 	front := quickSort(arr[:partition])
// 	back := quickSort(arr[partition:])
// 	return append(front, back...)
// }

// func quickSort(arr []int) []int {
// 	if len(arr) <= 1 {
// 		return arr
// 	}

// 	pivot := arr[0]
// 	left := []int{}
// 	right := []int{}

// 	for _, val := range arr[1:] {
// 		if val >= pivot {
// 			right = append(right, val)
// 		} else {
// 			left = append(left, val)
// 		}
// 	}
// 	answer := []int{}
// 	sortedLeft := quickSort(left)
// 	sortedRight := quickSort(right)
// 	answer = append(answer, sortedLeft...)
// 	answer = append(answer, pivot)
// 	answer = append(answer, sortedRight...)
// 	return answer
// }