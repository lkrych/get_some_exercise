func TestScrambleString(t *testing.T) {
	test1 := scrambleString("abcd", []int{3, 1, 2, 0})
	test2 := scrambleString("markov", []int{5, 3, 1, 4, 2, 0})

	if test1 != "dbca" {
		t.Errorf("Expected scrambleString of 'abcd' with inputs [3,1,2,0] to be 'dcba', not %v", test1)
	}

	if test2 != "vkaorm" {
		t.Errorf("Expected scrambleString of 'markov' with inputs [5,3,1,4,2,0] to be 'vkaorm', not %v", test2)
	}
}