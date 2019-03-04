// import (
// 	"strconv"
// 	"strings"
// )

// func dasherizeNumber(n int) string {
// 	dasherized := []string{}
// 	toString := strconv.Itoa(n)
// 	for i, el := range strings.Split(toString, "") {

// 		if int, _ := strconv.Atoi(el); int%2 != 0 {
// 			if i == 0 {
// 				dasherized = append(dasherized, el, "-")
// 			} else if i == len(toString)-1 {
// 				dasherized = append(dasherized, "-", el)
// 			} else {
// 				dasherized = append(dasherized, "-", el, "-")
// 			}
// 		} else {
// 			dasherized = append(dasherized, el)
// 		}
// 	}
// 	joined := strings.Join(dasherized, "")
// 	return strings.Replace(joined, "--", "-", -1)
// }