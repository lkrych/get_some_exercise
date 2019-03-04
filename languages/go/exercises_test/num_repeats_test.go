func TestNumRepeats(t *testing.T) {
	test1 := numRepeats("abdbc")
	test2 := numRepeats("aaa")
	test3 := numRepeats("abab")
	test4 := numRepeats("cadac")
	test5 := numRepeats("abcde")

	if test1 != 1 {
		t.Errorf("Expected numRepeats of abdbc to be 1, not %v", test1)
	}

	if test2 != 1 {
		t.Errorf("Expected numRepeats of aaa to be 1, not %v", test2)
	}

	if test3 != 2 {
		t.Errorf("Expected numRepeats of abab to be 2, not %v", test3)
	}

	if test4 != 2 {
		t.Errorf("Expected numRepeats of cadac to be 2, not %v", test4)
	}

	if test5 != 0 {
		t.Errorf("Expected numRepeats of abcde to be 0, not %v", test5)
	}
}