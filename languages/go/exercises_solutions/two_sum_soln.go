// func twoSum(arr []int, find int) []int {
// 	idxs := []int{}
// 	for i, val1 := range arr {
// 		for j, val2 := range arr[i+1:] {
// 			if val1+val2 == find {
// 				//bc you are iterating with modified slice
// 				//you need to take into account correct j index
// 				idxs = append(idxs, i, j+i+1)
// 			}
// 		}
// 	}
// 	return idxs
// }

// func twoSumOptimized(arr []int, find int) []int {
// 	sums := []int{}
// 	idxs := map[int]int{}
// 	//fill map
// 	for i, val := range arr {
// 		idxs[val] = i
// 	}
// 	//search for correct factors to negate the key
// 	for key, val := range idxs {
// 		search := find - key
// 		if searchIdx, exists := idxs[search]; exists {
// 			sums = append(sums, val, searchIdx)
// 			return sums
// 		}
// 	}
// 	return sums
// }
