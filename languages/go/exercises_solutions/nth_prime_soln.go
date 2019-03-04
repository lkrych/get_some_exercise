// func isPrime(n int) bool {
// 	for i := 2; i < n/2+1; i++ {
// 		if n%i == 0 {
// 			return false
// 		}
// 	}
// 	return true
// }

// func nthPrime(nth int) int {
// 	primeCount := 0
// 	n := 2
// 	for primeCount < nth {
// 		if isPrime(n) {
// 			primeCount++
// 		}
// 		n++
// 	}
// 	return n - 1
// }