func TestMultiplyArbitrary(t *testing.T) {
	test1 := multiplyArb([]int{1, 9, 3, 7, 0, 7, 7, 2, 1}, []int{-7, 6, 1, 8, 3, 8, 2, 5, 7, 2, 8, 7})
	test2 := multiplyArb([]int{1, 2, 9, 8, 2, 3, 4, 1, 3, 5, 0}, []int{4, 5})
	test3 := multiplyArb([]int{1, 2, 5, 3, 2, 4, 6, 7, 8, 2}, []int{-4, 0, 0, 2, 3})
	test4 := multiplyArb([]int{5, 3, 6, 3, 7, 9, 5, 0}, []int{3, 5, 2})
	test5 := multiplyArb([]int{1, 2}, []int{5})
	test6 := multiplyArb([]int{1, 2, 3}, []int{9, 8, 7})
	ans1 := []int{-1, 4, 7, 5, 7, 3, 9, 5, 2, 5, 8, 9, 6, 7, 6, 4, 1, 2, 9, 2, 7}
	ans2 := []int{5, 8, 4, 2, 0, 5, 3, 6, 0, 7, 5, 0}
	ans3 := []int{-5, 0, 1, 5, 8, 6, 9, 5, 9, 5, 5, 9, 8, 6}
	ans4 := []int{1, 8, 8, 8, 0, 5, 5, 8, 4, 0, 0}
	ans5 := []int{6, 0}
	ans6 := []int{1, 2, 1, 4, 0, 1}

	if !cmp.Equal(test1, ans1) {
		t.Errorf("Expected multiplyArb to multiply the array %v, instead you got %v", ans1, test1)
	}
	if !cmp.Equal(test2, ans2) {
		t.Errorf("Expected multiplyArb to multiply the array %v, instead you got %v", ans2, test2)
	}

	if !cmp.Equal(test3, ans3) {
		t.Errorf("Expected multiplyArb to multiply the array %v, instead you got %v", ans3, test3)
	}

	if !cmp.Equal(test4, ans4) {
		t.Errorf("Expected multiplyArb to multiply the array %v, instead you got %v", ans4, test4)
	}

	if !cmp.Equal(test5, ans5) {
		t.Errorf("Expected multiplyArb to multiply the array %v, instead you got %v", ans5, test5)
	}

	if !cmp.Equal(test6, ans6) {
		t.Errorf("Expected multiplyArb to multiply the array %v, instead you got %v", ans6, test6)
	}
}
