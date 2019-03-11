package goExercises

import "strings"

//assume you have a method isSubstring which checks if one word is asubstring of another
//given two strings s1 and s2, write code to check if s2 is a rotation of s1 using
//only 1 call to isSubstring
//ex: waterbottle is a rotation of erbottlewat
func stringRotation(s1, s2 string) bool {
	if len(s1) == len(s2) && len(s1) > 0 {
		s1 = s1 + s1
		return strings.Index(s1, s2) > 0
	}
	return false
}

//write an algorithm such that if an element in an mxn matrix is 0
//its entire row and column are set to zero
func zeroMatrix(matrix [][]int) [][]int {
	firstRowZero := false
	firstColumnZero := false

	for j := 0; j < len(matrix[0]); j++ {
		if matrix[0][j] == 0 {
			firstColumnZero = true
			break // you can break because the rest of the column will have to be zero
		}
	}

	for i := 0; i < len(matrix); i++ {
		if matrix[i][0] == 0 {
			firstRowZero = true
			break
		}
	}

	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0 //fill in the zeros on the outermost edge
				matrix[0][j] = 0
			}
		}
	}

	//now go back through and nullify all the rows and columns

	for j := 1; j < len(matrix[0]); j++ {
		if matrix[0][j] == 0 {
			nullifyColumn(matrix, j)
		}
	}

	for i := 0; i < len(matrix); i++ {
		if matrix[i][0] == 0 {
			nullifyRow(matrix, i)
		}
	}

	if firstColumnZero {
		nullifyColumn(matrix, 0)
	}

	if firstRowZero {
		nullifyRow(matrix, 0)
	}

	return matrix

}

func nullifyColumn(matrix [][]int, colIdx int) {
	for i := 1; i < len(matrix); i++ {
		matrix[i][colIdx] = 0
	}
}

func nullifyRow(matrix [][]int, rowIdx int) {
	for j := 1; j < len(matrix[0]); j++ {
		matrix[rowIdx][j] = 0
	}
}
