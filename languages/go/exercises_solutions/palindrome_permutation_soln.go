//there are two ways to do this:
//one contruct a hash and keep a count of all the characters you see
//if there are two odd counts, this is not a possible palindrome
//big O = O(n)

//the other solution might be a little more space-efficient, you use a bit
//vector to model the alphabet, you toggle the index of the char when you see it
//you should only have one 1
// func palindromePerm(perm string) bool {
// 	charMap := map[char]int{}
// 	for i := 0; i < len(perm); i++ {
// 		char := perm[i]
// 		if count, ok := charMap[char]; ok {
// 			charMap[char]++
// 		} else {
// 			charMap[char] = 1
// 		}
// 	}
// 	oddCount := 0
// 	for key, value := range charMap {
// 		if value%2 != 0 {
// 			oddCount++
// 		}
// 	}
// 	return oddCount <= 1
// }

// func palindromePermBit(perm string) bool {
// 	bitVector := createBitVector(len(perm))
// 	alpha := "abcdefghijklmnopqrstuvwxyz"
// 	for _, char := strings.Split(perm, "") {
// 		if char == " " { //skip spaces
// 			continue
// 		}
// 		idx := strings.Index(alpha, char)
// 		bitVector[idx] = bitVector[idx] == 1 ? 0 : 1
// 	}
// 	oddCount := 0
// 	for _, el := range bitVector {
// 		if el == 1 {
// 			oddCount+++
// 		}
// 	}
// 	return oddCount <= 1

// }

// func createBitVector(len int) []int {
// 	bitVector := make([]int, len)
// 	for i := 0; i < len; i++ {
// 		bitVector[i] = 0
// 	}
// }