// import "math"

// func multiplyArb(arr1 []int, arr2 []int) []int {
// 	sign := 1
// 	//check if numbers are negative
// 	if (arr1[0] * arr2[0]) > 0 {
// 		sign = 1
// 	} else {
// 		sign = -1
// 	}
// 	//set to absolute val
// 	arr1[0] = int(math.Abs(float64(arr1[0])))
// 	arr2[0] = int(math.Abs(float64(arr2[0])))

// 	total := make([]int, len(arr1)+len(arr2))
// 	for i := len(arr1) - 1; i >= 0; i-- {
// 		for j := len(arr2) - 1; j >= 0; j-- {
// 			total[i+j+1] += arr1[i] * arr2[j]
// 			total[i+j] += total[i+j+1] / 10
// 			total[i+j+1] %= 10
// 		}
// 	}
// 	zeroIdx := 0
// 	for {
// 		if total[zeroIdx] != 0 {
// 			break
// 		}
// 		zeroIdx++
// 	}
// 	total = total[zeroIdx:]
// 	total[0] = total[0] * sign

// 	return total
// }