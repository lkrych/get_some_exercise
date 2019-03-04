// func incrementArb(arr []int) []int {
// 	lastIdx := len(arr) - 1
// 	arr[lastIdx] = (arr[lastIdx] + 1) % 10
// 	if arr[lastIdx] == 0 {
// 		arr = append(incrementArb(arr[:lastIdx]), arr[lastIdx])
// 	}
// 	return arr
// }