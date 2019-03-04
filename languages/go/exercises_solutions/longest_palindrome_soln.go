
// func longestPalindrome(sentence string) string {
// 	startIdx := 0
// 	longest := ""
// 	for startIdx < len(sentence)-1 {
// 		for endIdx := len(sentence) - 1; endIdx > startIdx; endIdx-- {
// 			subString := sentence[startIdx : endIdx+1]
// 			if subString == reverse(subString) && len(subString) > len(longest) {
// 				longest = subString
// 			}
// 		}
// 		startIdx++
// 	}
// 	return longest
// }