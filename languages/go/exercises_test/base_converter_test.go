func TestBaseConverter(t *testing.T) {
	test1 := baseConverter(5, 2)
	test2 := baseConverter(1239449, 16)
	test3 := baseConverter(112, 2)
	test4 := baseConverter(112, 16)

	if test1 != "101" {
		t.Errorf("Expected baseConverter of 5 to base 2 to equal 101, not %v", test1)
	}
	if test2 != "12e999" {
		t.Errorf("Expected baseConverter of 1239449 to base 16 to be 12e999, not %v", test2)
	}

	if test3 != "1110000" {
		t.Errorf("Expected baseConverter of 112 to base 2 to be 1110000, not %v", test3)
	}

	if test4 != "70" {
		t.Errorf("Expected baseConverter of 112 to base 16 to be 70, not %v", test4)
	}
}