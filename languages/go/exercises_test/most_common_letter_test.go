func TestMostCommonLetter(t *testing.T) {
	test1 := mostCommonLetter("abca")
	ans1 := []string{
		"a", "2",
	}
	test2 := mostCommonLetter("abbab")
	ans2 := []string{
		"b", "3",
	}
	test3 := mostCommonLetter("guernsey was here")
	ans3 := []string{
		"e", "4",
	}

	if !cmp.Equal(test1, ans1) {
		t.Errorf("Expected mostCommonLetter of abca to be a with a count of 2, not %v", test1)
	}

	if !cmp.Equal(test2, ans2) {
		t.Errorf("Expected mostCommonLetter of abbab to be b with a count of 3, not %v", test2)
	}

	if !cmp.Equal(test3, ans3) {
		t.Errorf("Expected mostCommonLetter of guernsey was here to be e with a count of 4, not %v", test3)
	}
}