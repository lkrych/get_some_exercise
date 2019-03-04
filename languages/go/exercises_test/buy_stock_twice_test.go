
i
func TestBuyStockTwice(t *testing.T) {
	test1 := []int{12, 11, 13, 9, 12, 8, 14, 13, 15}
	test2 := []int{3, 2, 0, 0, 2, 0, 1, 5}
	test3 := []int{5, 3, 1, 5, 2, 0, 1}
	test4 := []int{5, 3, 2, 4, 1}
	test5 := []int{3, 3, 1, 0, 2, 0, 1, 3, 1, 0, 6}
	test6 := []int{7, 6, 5, 4, 3, 2}
	ans1 := buyStockTwice(test1)
	ans2 := buyStockTwice(test2)
	ans3 := buyStockTwice(test3)
	ans4 := buyStockTwice(test4)
	ans5 := buyStockTwice(test5)
	ans6 := buyStockTwice(test6)
	if ans1 != 10 {
		t.Errorf("Expected array %v to be 10, not %v", test1, ans1)
	}
	if ans2 != 7 {
		t.Errorf("Expected array %v to be 7, not %v", test2, ans2)
	}

	if ans3 != 5 {
		t.Errorf("Expected array %v to be 5, not %v", test3, ans3)
	}

	if ans4 != 2 {
		t.Errorf("Expected array %v to be 2, not %v", test4, ans4)
	}

	if ans5 != 9 {
		t.Errorf("Expected array %v to be 9, not %v", test5, ans5)
	}

	if ans6 != 0 {
		t.Errorf("Expected array %v to be 0, not %v", test6, ans6)
	}
}