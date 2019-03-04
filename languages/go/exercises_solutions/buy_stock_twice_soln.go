// import "fmt"

// //iterate forward O(n) and then backward O(n), combining the results in one array.
// func buyStockTwice(arr []int) int {
// 	maxProfit := 0
// 	maxProfits := make([]int, len(arr))
// 	low := 1000000
// 	for i := 0; i < len(arr); i++ {
// 		if arr[i] < low {
// 			low = arr[i]
// 		}
// 		if (arr[i] - low) > maxProfit {
// 			maxProfit = arr[i] - low
// 		}
// 		maxProfits[i] = maxProfit

// 	}
// 	fmt.Printf("Before backwards iteration: %v \n", maxProfits)
// 	totalMaxProfit := 0
// 	maxProfit = 0
// 	high := -1
// 	for j := len(arr) - 1; j > 0; j-- {
// 		if arr[j] > high {
// 			high = arr[j]
// 		}
// 		if (high - arr[j]) > maxProfit {
// 			maxProfit = high - arr[j]
// 		}
// 		maxProfits[j] = maxProfit + maxProfits[j-1]
// 		if maxProfits[j] > totalMaxProfit {
// 			totalMaxProfit = maxProfits[j]
// 		}
// 	}
// 	fmt.Printf("After backwards iteration: %v \n", maxProfits)
// 	return totalMaxProfit
// }