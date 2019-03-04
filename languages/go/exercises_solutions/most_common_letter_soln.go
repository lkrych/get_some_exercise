// import (
// 	"strconv"
// 	"strings"
// )

// func mostCommonLetter(sentence string) []string {
// 	lettersMap := map[string]int{}
// 	for _, char := range strings.Split(sentence, "") {
// 		if char == " " {
// 			continue
// 		}
// 		if lettersMap[char] >= 1 {
// 			lettersMap[char]++
// 		} else {
// 			lettersMap[char] = 1
// 		}
// 	}
// 	count := 0
// 	mostCommon := []string{}
// 	for key, val := range lettersMap {
// 		if val > count {
// 			count = val
// 			mostCommon = []string{key, strconv.Itoa(val)}
// 		}
// 	}
// 	return mostCommon
// }