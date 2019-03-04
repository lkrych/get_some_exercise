// func thirdGreatest(arr []int) int {
// 	var first int
// 	var second int
// 	var third int

// 	for _, el := range arr {
// 		if el >= first {
// 			third = second
// 			second = first
// 			first = el
// 			continue
// 		}
// 		if el >= second {
// 			third = second
// 			second = el
// 			continue
// 		}
// 		if el >= third {
// 			third = el
// 		}
// 	}
// 	return third
// }