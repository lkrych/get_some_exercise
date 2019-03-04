// //take sentence and reverse it, do not use library method
// func reverse(sentence string) string {
// 	r := []rune(sentence)
// 	for i, j := 0, len(sentence)-1; i < j; i, j = i+1, j-1 {
// 		r[i], r[j] = r[j], r[i]
// 	}
// 	return string(r)
// }