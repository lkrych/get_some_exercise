// func greatestCommonFactor(n1 int, n2 int) int {
// 	var smaller int
// 	var larger int
// 	if n1 > n2 {
// 		smaller = n2
// 		larger = n1
// 	} else {
// 		smaller = n1
// 		larger = n2
// 	}

// 	for i := 1; i <= smaller; i++ {
// 		if smaller%i == 0 {
// 			factor := smaller / i
// 			if larger%factor == 0 {
// 				return factor
// 			}
// 		}
// 	}
// 	//if no factors are found, the GCF is 1
// 	return 1
// }