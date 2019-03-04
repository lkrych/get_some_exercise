// func removeDups(arr []int) []int {
// 	writeIdx := 1
// 	for i := 1; i < len(arr); i++ {
// 		if arr[writeIdx-1] != arr[i] {
// 			arr[writeIdx] = arr[i]
// 			writeIdx++
// 		}
// 	}
// 	return arr[:writeIdx]
// }