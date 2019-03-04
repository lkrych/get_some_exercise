func TestDigitalRoot(t *testing.T) {
	test1 := digitalRoot(10)
	test2 := digitalRoot(83)
	test3 := digitalRoot(123)
	test4 := digitalRoot(789)
	test5 := digitalRoot(452982)

	if test1 != 1 {
		t.Errorf("Expected digitalRoot of 10 to be 1, not %v", test1)
	}
	if test2 != 2 {
		t.Errorf("Expected digitalRoot of 83 to be 2, not %v", test2)
	}

	if test3 != 6 {
		t.Errorf("Expected digitalRoot of 123 to be 6, not %v", test3)
	}

	if test4 != 6 {
		t.Errorf("Expected digitalRoot of 789 to be 6, not %v", test4)
	}

	if test5 != 3 {
		t.Errorf("Expected digitalRoot of 452982 to be 3, not %v", test5)
	}
}