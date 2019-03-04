// import "errors"

// func fibonacci(n int) ([]int, error) {
// 	if n < 1 {
// 		return nil, errors.New("n must be greater than 0")
// 	}
// 	fibArr := []int{}
// 	if n < 2 {
// 		fibArr = append(fibArr, n)
// 		return fibArr, nil
// 	}
// 	fibArr = append(fibArr, (fibArr[n-1] + fibArr[n-2]))
// 	return fibArr, nil
// }