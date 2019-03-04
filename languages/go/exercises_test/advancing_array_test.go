func TestAdvancingArray(t *testing.T) {
	test1 := []int{3, 3, 1, 0, 2, 0, 1}
	test2 := []int{3, 2, 0, 0, 2, 0, 1}
	test3 := []int{5, 3, 1, 0, 2, 0, 1}
	test4 := []int{0, 3, 1, 0, 2, 0, 1}
	test5 := []int{3, 3, 1, 0, 2, 0, 1, 3, 1, 0, 6}
	if advancingArray(test1) != true {
		t.Errorf("Expected array %v to be true", test1)
	}
	if advancingArray(test2) != false {
		t.Errorf("Expected array %v to be false", test2)
	}

	if advancingArray(test3) != true {
		t.Errorf("Expected array %v to be true", test3)
	}

	if advancingArray(test4) != false {
		t.Errorf("Expected array %v to be false", test4)
	}

	if advancingArray(test5) != true {
		t.Errorf("Expected array %v to be true", test5)
	}
}