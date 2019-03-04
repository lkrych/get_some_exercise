// import "strings"

// func nearbyAZ(word string) bool {
// 	split := strings.Split(word, "")
// 	foundA := -2000
// 	for i, char := range split {
// 		if char == "a" {
// 			foundA = i
// 		}
// 		if char == "z" {
// 			if i-foundA <= 3 {
// 				return true
// 			}
// 		}
// 	}
// 	return false
// }