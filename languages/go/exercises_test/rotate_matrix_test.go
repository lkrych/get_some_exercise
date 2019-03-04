import "testing"

func TestRotateMatrix(t *testing.T) {
	test1 := rotateMatrix([][]int{{1, 2, 3}, {1, 2, 3}, {1, 2, 3}})
	test2 := rotateMatrix([][]int{{1, 2, 3, 4}, {1, 2, 3, 4}, {1, 2, 3, 4}, {1, 2, 3, 4}})
	test3 := rotateMatrix([][]int{{1, 1}, {2, 2}})

	ans1 := [][]int{{1, 1, 1}, {2, 2, 2}, {3, 3, 3}}
	ans2 := [][]int{{1, 1, 1, 1}, {2, 2, 2, 2}, {3, 3, 3, 3}, {4, 4, 4, 4}}
	ans3 := [][]int{{2, 1}, {2, 1}}
	if !cmp.Equal(test1, ans1) {
		t.Errorf("Expected rotateMatrix to rotate the matrix %v into %v", test1, ans1)
	}

	if !cmp.Equal(test2, ans2) {
		t.Errorf("Expected rotateMatrix to rotate the matrix %v into %v", test1, ans1)
	}

	if !cmp.Equal(test3, ans3) {
		t.Errorf("Expected rotateMatrix to rotate the matrix %v into %v", test1, ans1)
	}

}