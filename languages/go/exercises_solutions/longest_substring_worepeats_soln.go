// //Given a string, find the length of the longest substring without repeating characters.
// func lengthOfLongestSubstring(s string) int {
// 	bs := []byte(s)
// 	var maxLen, start int
// 	idxMap := map[byte]int{}
// 	for i := 0; i < len(bs); i++ {
// 			if _, ok := idxMap[bs[i]]; ok && start <= idxMap[bs[i]] {
// 					start = idxMap[bs[i]] + 1
// 			} else {
// 					if maxLen < i - start + 1 {
// 							maxLen = i - start + 1
// 					}
// 			}
// 			idxMap[bs[i]] = i
// 	}
// 	return maxLen
// }