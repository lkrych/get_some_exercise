func TestPartitionEvenOdd(t *testing.T) {
	test1 := partitionEvenOdd([]int{7, 5, 2, 4, 8, 9, 3})
	test2 := partitionEvenOdd([]int{7, 5, 2, 4, 8, 9, 3, 1, 6, 10})
	test3 := partitionEvenOdd([]int{7, 5, 2, 4, 8, 9, 3, 3, 3, 3})
	test4 := partitionEvenOdd([]int{7, 5, 2, 4, 8, 9, 3, 100, 102, 6, 1, 10})
	test5 := partitionEvenOdd([]int{7, 5, 2, 4, 8, 9, 3, 11, 10, 6, 1})
	ans1 := []int{2, 4, 8, 3, 5, 7, 9}
	ans2 := []int{2, 4, 6, 8, 10, 1, 3, 5, 7, 9}
	ans3 := []int{2, 4, 8, 3, 3, 3, 3, 5, 7, 9}
	ans4 := []int{2, 4, 6, 8, 10, 100, 102, 1, 3, 5, 7, 9}
	ans5 := []int{2, 4, 6, 8, 10, 1, 3, 5, 7, 9, 11}

	if !cmp.Equal(test1, ans1) {
		t.Errorf("Expected partitionEvenOdd to separate the array %v, instead you got %v", ans1, test1)
	}
	if !cmp.Equal(test2, ans2) {
		t.Errorf("Expected partitionEvenOdd to separate the array %v, instead you got %v", ans2, test2)
	}

	if !cmp.Equal(test3, ans3) {
		t.Errorf("Expected partitionEvenOdd to separate the array %v, instead you got %v", ans3, test3)
	}

	if !cmp.Equal(test4, ans4) {
		t.Errorf("Expected partitionEvenOdd to separate the array %v, instead you got %v", ans4, test4)
	}

	if !cmp.Equal(test5, ans5) {
		t.Errorf("Expected partitionEvenOdd to separate the array %v, instead you got %v", ans5, test5)
	}
}
