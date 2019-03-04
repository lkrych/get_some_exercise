func TestIncrementArbitrary(t *testing.T) {
	test1 := incrementArb([]int{1, 2, 9})
	test2 := incrementArb([]int{1, 2, 9, 8, 2, 3, 4, 1, 3, 5, 0})
	test3 := incrementArb([]int{1, 2, 5, 3, 2, 4, 6, 7, 8, 2})
	test4 := incrementArb([]int{5, 3, 6, 3, 7, 9, 5, 0})
	test5 := incrementArb([]int{1, 2, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9})
	ans1 := []int{1, 3, 0}
	ans2 := []int{1, 2, 9, 8, 2, 3, 4, 1, 3, 5, 1}
	ans3 := []int{1, 2, 5, 3, 2, 4, 6, 7, 8, 3}
	ans4 := []int{5, 3, 6, 3, 7, 9, 5, 1}
	ans5 := []int{1, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if !cmp.Equal(test1, ans1) {
		t.Errorf("Expected incrementArb to increment the array %v, instead you got %v", ans1, test1)
	}
	if !cmp.Equal(test2, ans2) {
		t.Errorf("Expected incrementArb to increment the array %v, instead you got %v", ans2, test2)
	}

	if !cmp.Equal(test3, ans3) {
		t.Errorf("Expected incrementArb to increment the array %v, instead you got %v", ans3, test3)
	}

	if !cmp.Equal(test4, ans4) {
		t.Errorf("Expected incrementArb to increment the array %v, instead you got %v", ans4, test4)
	}

	if !cmp.Equal(test5, ans5) {
		t.Errorf("Expected incrementArb to increment the array %v, instead you got %v", ans5, test5)
	}
}
