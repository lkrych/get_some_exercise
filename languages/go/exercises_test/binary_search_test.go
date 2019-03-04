func TestBinarySearch(t *testing.T) {
	test1 := binarySearch([]int{0, 1, 2, 3, 4, 5}, 3)
	test2 := binarySearch([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 9)
	test3 := binarySearch([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 0)
	test4 := binarySearch([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 10)
	test5 := binarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 0)

	if test1 != 3 {
		t.Errorf("Expected binarySearch to find the index of 3 at 4, not %v", test1)
	}
	if test2 != 9 {
		t.Errorf("Expected binarySearch to find the index of 9 at 10, not %v", test2)
	}

	if test3 != 0 {
		t.Errorf("Expected binarySearch to find the index of 0 at 0, not %v", test3)
	}

	if test4 != -1 {
		t.Errorf("Expected binarySearch to find the index of 10 at -1, not %v", test4)
	}

	if test5 != -1 {
		t.Errorf("Expected binarySearch to find the index of 0 at -1, not %v", test5)
	}
}