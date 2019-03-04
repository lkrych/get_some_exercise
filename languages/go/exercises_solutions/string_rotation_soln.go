// import "strings"

// // If there is a rotation, there is a pivot point
// //regardless of where the pivot is s2 will always be a substring of s1s1
// func stringRotation(s1, s2 string) bool {
// 	if len(s1) == len(s2) && len(s1) > 0 {
// 		s1 = s1 + s1
// 		return strings.Index(s1, s2) > 0
// 	}
// }