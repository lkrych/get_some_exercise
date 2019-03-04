// import "strings"

// func caesarCipher(shift int, sentence string) string {
// 	shifted := []string{}
// 	for _, ascii := range []byte(sentence) {
// 		if ascii == 32 {
// 			shifted = append(shifted, " ")
// 		} else {
// 			shifted = append(shifted, shiftChar(shift, ascii))
// 		}
// 	}
// 	return strings.Join(shifted, "")
// }

// func shiftChar(shift int, ascii byte) string {
// 	newChar := int(ascii) + (shift % 26)
// 	if newChar > 122 {
// 		newChar = 96 + (newChar % 122)
// 	}
// 	return string(newChar)
// }