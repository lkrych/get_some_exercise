// func buyStock(arr []int) int {
// 	low := 1000000
// 	diff := 0
// 	for i := 0; i < len(arr); i++ {
// 		if arr[i] < low {
// 			low = arr[i]
// 		}
// 		if (arr[i] - low) > diff {
// 			diff = arr[i] - low
// 		}
// 	}
// 	return diff
// }