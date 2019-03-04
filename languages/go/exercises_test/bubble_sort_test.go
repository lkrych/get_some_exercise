func TestBubbleSort(t *testing.T) {
	test1 := bubbleSort([]int{7, 5, 2, 4, 8, 9, 3})
	test2 := bubbleSort([]int{7, 5, 2, 4, 8, 9, 3, 1, 6, 10})
	test3 := bubbleSort([]int{7, 5, 2, 4, 8, 9, 3, 3, 3, 3})
	test4 := bubbleSort([]int{7, 5, 2, 4, 8, 9, 3, 100, 102, 6, 1, 10})
	test5 := bubbleSort([]int{7, 5, 2, 4, 8, 9, 3, 11, 10, 6, 1})
	ans1 := []int{2, 3, 4, 5, 7, 8, 9}
	ans2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ans3 := []int{2, 3, 3, 3, 3, 4, 5, 7, 8, 9}
	ans4 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 100, 102}
	ans5 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

	if !cmp.Equal(test1, ans1) {
		t.Errorf("Expected bubbleSort to sort the array %v, instead you got %v", test1, ans1)
	}
	if !cmp.Equal(test2, ans2) {
		t.Errorf("Expected bubbleSort to sort the array %v, instead you got %v", test2, ans2)
	}

	if !cmp.Equal(test3, ans3) {
		t.Errorf("Expected bubbleSort to sort the array %v, instead you got %v", test3, ans3)
	}

	if !cmp.Equal(test4, ans4) {
		t.Errorf("Expected bubbleSort to sort the array %v, instead you got %v", test4, ans4)
	}

	if !cmp.Equal(test5, ans5) {
		t.Errorf("Expected bubbleSort to sort the array %v, instead you got %v", test5, ans5)
	}
}