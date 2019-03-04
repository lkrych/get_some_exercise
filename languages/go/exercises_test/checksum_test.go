func TestCheckSum(t *testing.T) {
	test1 := checkSum([][]int{{1, 2, 3}, {1, 2, 3}, {1, 2, 3}})
	test2 := checkSum([][]int{{1, 2, 3, 4}, {1, 2, 3}, {1, 2, 3}})
	test3 := checkSum([][]int{{1, 2, 3, 10}, {1, 2, 3, 12}, {1, 2, 3, 15}})
	test4 := checkSum([][]int{{3, 22, 124}, {3, 70}, {1, 2, 5}})
	test5 := checkSum([][]int{{1, 2, 10}, {1, 11, 3}, {500, 2, 3}})
	test6 := checkSum(checkSumAdvent())

	if test1 != 6 {
		t.Errorf("Expected answer to be 6, not %v", test1)
	}
	if test2 != 7 {
		t.Errorf("Expected answer to be 7, not %v", test2)
	}

	if test3 != 34 {
		t.Errorf("Expected answer to be 34, not %v", test3)
	}

	if test4 != 192 {
		t.Errorf("Expected answer to be 192, not %v", test4)
	}

	if test5 != 517 {
		t.Errorf("Expected answer to be 517, not %v", test5)
	}

	if test6 != 53978 {
		t.Errorf("Expected answer to be 53978, not %v", test6)
	}
}