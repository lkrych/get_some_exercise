func TestThirdGreatest(t *testing.T) {
	test1 := thirdGreatest([]int{5, 3, 7})
	test2 := thirdGreatest([]int{5, 3, 4, 7})
	test3 := thirdGreatest([]int{2, 3, 4, 7})

	if test1 != 3 {
		t.Errorf("Expected thirdGreatest of [5, 3, 7] to be 3, not %v", test1)
	}

	if test2 != 4 {
		t.Errorf("Expected thirdGreatest of [5, 3, 4, 7] to be 4, not %v", test2)
	}

	if test3 != 3 {
		t.Errorf("Expected thirdGreatest of [2, 3, 4, 7] to be 3, not %v", test3)
	}
}