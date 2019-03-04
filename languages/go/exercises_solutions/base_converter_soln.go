// import (
// 	"fmt"
// 	"strconv"
// 	"strings"
// )

// func baseConverter(n int, base int) string {
// 	base16 := map[int]string{
// 		10: "a",
// 		11: "b",
// 		12: "c",
// 		13: "d",
// 		14: "e",
// 		15: "f"}
// 	answer := []string{}
// 	for n > 0 {
// 		answer = append(answer, strconv.Itoa(n%base))
// 		n = n / base
// 	}
// 	if base == 16 {
// 		for i, el := range answer {
// 			intEl, err := strconv.Atoi(el)
// 			if err != nil {
// 				fmt.Println("There was an error!")
// 			}
// 			if hex, exists := base16[intEl]; exists {
// 				answer[i] = hex
// 			}
// 		}
// 	}
// 	return reverse(strings.Join(answer, ""))
// }