// import "strconv"

// func timeConversion(timeInMinutes int) string {
// 	minutes := timeInMinutes % 60
// 	hours := timeInMinutes / 60
// 	var time string
// 	if minutes < 10 {
// 		time = strconv.Itoa(hours) + ":0" + strconv.Itoa(minutes)
// 	} else {
// 		time = strconv.Itoa(hours) + ":" + strconv.Itoa(minutes)
// 	}
// 	return time
// }