// func binarySearch(arr []int, toFind int) int {
// 	if len(arr) < 1 {
// 		return -1
// 	}
// 	midIdx := len(arr) / 2
// 	if toFind == arr[midIdx] {
// 		return midIdx
// 	} else if toFind > arr[midIdx] {
// 		search := binarySearch(arr[midIdx+1:], toFind)
// 		if search != -1 {
// 			return search + midIdx + 1
// 		}
// 		return search
// 	} else {
// 		return binarySearch(arr[0:midIdx], toFind)
// 	}

// }