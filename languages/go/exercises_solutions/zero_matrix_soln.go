//the naive way to do this would be to keep track of where the zeros are and then go back
//and zero out the respective columns and rows with either a separate matrix, or arrays
//this uses extra space, the space-efficient way to do this is to set the zeroth column and row
//to 0 when you find one, and then go back and nullify the rows and columns, this uses the matrix you
//already have O(1)
// func zeroMatrix(matrix [][]int) [][]int {
// 	rowHasZero := false
// 	colHasZero := false

// 	//check if first row has zero
// 	for j := 0; j < len(matrix[0]); j++ {
// 		if matrix[0][j] == 0 {
// 			rowHasZero = true
// 			break //you don't need to keep iterating
// 		}
// 	}

// 	//check if first row has zero
// 	for i := 0; i < len(matrix[0]); i++ {
// 		if matrix[i][0] == 0 {
// 			colHasZero = true
// 			break //you don't need to keep iterating
// 		}
// 	}

// 	//check for zeros in the rest of the matrix
// 	for i := 1; i < len(matrix); i++ {
// 		for j := 1; j < len(matrix[0]); j++ {
// 			if matrix[i][j] == 0 {
// 				matrix[i][0] = 0 //fill in zeros on the outermost edge
// 				matrix[0][j] = 0
// 			}
// 		}
// 	}

// 	for i := 1; i < len(matrix); i++ {
// 		if matrix[i][0] == 0 {
// 			nullifyRow(matrix, i)
// 		}
// 	}

// 	for j := 1; j < len(matrix[j]); j++ {
// 		if matrix[0][j] == 0 {
// 			nullifyColumn(matrix, j)
// 		}
// 	}

// 	if rowHasZero {
// 		nullifyRow(matrix, 0)
// 	}

// 	if colHasZero {
// 		nullifyColumn(matrix, 0)
// 	}
// }

// func nullifyColumn(matrix [][]int, colIdx int) {
// 	for i := 0; i < len(matrix); i++ {
// 		matrix[i][colIdx] = 0
// 	}
// }

// func nullifyRow(matrix [][]int, rowIdx int) {
// 	for j := 0; j < len(matrix[0]); j++ {
// 		matrix[rowIdx][j] = 0
// 	}
// }