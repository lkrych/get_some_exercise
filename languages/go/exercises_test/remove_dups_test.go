func TestRemoveDups(t *testing.T) {
	test1 := removeDups([]int{1, 1, 1, 2, 3, 4, 4, 4, 5})
	test2 := removeDups([]int{1, 1, 1, 2, 33, 102, 102})
	test3 := removeDups([]int{1, 1, 1, 2, 3, 3, 4, 5, 6, 7, 8, 8, 9, 9})
	test4 := removeDups([]int{1, 1, 1, 2, 3, 3, 3, 3, 3})
	test5 := removeDups([]int{1, 1, 1, 2, 3, 3, 4, 5, 5})
	ans1 := []int{1, 2, 3, 4, 5}
	ans2 := []int{1, 2, 33, 102}
	ans3 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	ans4 := []int{1, 2, 3}
	ans5 := []int{1, 2, 3, 4, 5}
	if !cmp.Equal(test1, ans1) {
		t.Errorf("Expected array %v to be %v", test1, ans1)
	}
	if !cmp.Equal(test2, ans2) {
		t.Errorf("Expected array %v to be %v", test2, ans2)
	}

	if !cmp.Equal(test3, ans3) {
		t.Errorf("Expected array %v to be %v", test4, ans3)
	}

	if !cmp.Equal(test4, ans4) {
		t.Errorf("Expected array %v to be %v", test5, ans4)
	}

	if !cmp.Equal(test5, ans5) {
		t.Errorf("Expected array %v to be %v", test5, ans5)
	}
}