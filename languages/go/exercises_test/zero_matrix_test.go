func TestRotateMatrix(t *testing.T) {
	test1 := zeroMatrix([][]int{{1, 2, 3}, {1, 0, 3}, {1, 2, 3}})
	test2 := zeroMatrix([][]int{{1, 2, 3, 4}, {1, 0, 3, 4}, {1, 2, 3, 4}, {1, 2, 3, 4}})
	test3 := zeroMatrix([][]{1,1},{2,2})

	ans1 := [][]int{{1, 0, 3}, {0, 0, 0}, {1, 0, 3}}
	ans2 := []][]int{{1, 0, 3, 4}, {0, 0, 0, 0}, {1, 0, 3, 4}, {1, 0, 3, 4}}
	ans3 := [][]int{{2, 1}, {2, 1}}
	if !cmp.Equal(test1, ans1) {
		t.Errorf("Expected zeroMatrix to zero the matrix %v into %v", test1, ans1)
	}

	if !cmp.Equal(test2, ans2) {
		t.Errorf("Expected zeroMatrix to zero the matrix %v into %v", test1, ans1)
	}

	if !cmp.Equal(test3, ans3) {
		t.Errorf("Expected zeroMatrix to zero the matrix %v into %v", test1, ans1)
	}

}