// import "strings"

// func numRepeats(word string) int {
// 	lettersMap := map[string]int{}
// 	for _, ch := range strings.Split(word, "") {
// 		if ch == " " {
// 			continue
// 		}
// 		if lettersMap[ch] >= 1 {
// 			lettersMap[ch]++
// 		} else {
// 			lettersMap[ch] = 1
// 		}
// 	}
// 	count := 0
// 	for _, v := range lettersMap {
// 		if v > 1 {
// 			count++
// 		}
// 	}
// 	return count
// }