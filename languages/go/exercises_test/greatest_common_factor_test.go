func TestGreatestCommonFactor(t *testing.T) {
	test1 := greatestCommonFactor(3, 9)
	test2 := greatestCommonFactor(16, 24)
	test3 := greatestCommonFactor(7, 38)

	if test1 != 3 {
		t.Errorf("Expected greatestCommonFactor of 3 and 9 to be 3, not %v", test1)
	}

	if test2 != 8 {
		t.Errorf("Expected greatestCommonFactor of 16 and 24 to be 8, not %v", test2)
	}

	if test3 != 1 {
		t.Errorf("Expected greatestCommonFactor of 7 and 38 to be 1, not %v", test3)
	}
}