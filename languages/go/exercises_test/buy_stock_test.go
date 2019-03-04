
func TestBuyStock(t *testing.T) {
	test1 := []int{3, 3, 1, 0, 2, 0, 1}
	test2 := []int{3, 2, 0, 0, 2, 0, 1, 5}
	test3 := []int{5, 3, 1, 5, 2, 0, 1}
	test4 := []int{5, 3, 2, 4, 1}
	test5 := []int{3, 3, 1, 0, 2, 0, 1, 3, 1, 0, 6}
	test6 := []int{7, 6, 5, 4, 3, 2}
	if buyStock(test1) != 2 {
		t.Errorf("Expected array %v to be 2", test1)
	}
	if buyStock(test2) != 5 {
		t.Errorf("Expected array %v to be 5", test2)
	}

	if buyStock(test3) != 4 {
		t.Errorf("Expected array %v to be 4", test3)
	}

	if buyStock(test4) != 2 {
		t.Errorf("Expected array %v to be 2", test4)
	}

	if buyStock(test5) != 6 {
		t.Errorf("Expected array %v to be 6", test5)
	}

	if buyStock(test6) != 0 {
		t.Errorf("Expected array %v to be 0", test6)
	}
}