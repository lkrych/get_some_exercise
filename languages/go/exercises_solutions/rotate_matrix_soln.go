// import "log"

// func rotateMatrix(matrix [][]int) [][]int {
// 	if matrix.length == 0 || matrix.length != matrix[0].length {
// 		log.Fatal("this matrix is not square")
// 	}
// 	n = len(matrix)
// 	for layer := 0; layer < n/2; layer++ {
// 		first := layer
// 		last = n - 1 - layer
// 		for i = first; i < last; i++ {
// 			offset := i - first
// 			top := matrix[first][i] //save top
// 			//rotate left -> top
// 			matrix[first][i] = matrix[last-offest][first]
// 			//bottom -> left
// 			matrix[last-offset][first] = matrix[last][last-offset]
// 			//right -> bottom
// 			matrix[last][last-offset] = matrix[i][last]
// 			//top -> right
// 			matrix[i][last] = top
// 		}
// 	}
// 	return matrix
// }