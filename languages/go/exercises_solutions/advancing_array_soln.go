// import "math"

// //func keep track of how far you have traveled, if you can reach the end, you made it!
// func advancingArray(arr []int) bool {
// 	farthestTraveled := arr[0]
// 	for i := 0; i <= farthestTraveled; i++ {
// 		if farthestTraveled == len(arr)-1 {
// 			return true
// 		}
// 		farthestTraveled = int(math.Max(float64(farthestTraveled), float64(arr[i]+i)))
// 	}
// 	return false
// }