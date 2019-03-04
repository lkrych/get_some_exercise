func TestDasherizeNumber(t *testing.T) {
	test1 := dasherizeNumber(203)
	test2 := dasherizeNumber(303)
	test3 := dasherizeNumber(333)
	test4 := dasherizeNumber(3223)

	if test1 != "20-3" {
		t.Errorf("Expected dasherizeNumber of 203  to be 20-3, not %v", test1)
	}

	if test2 != "3-0-3" {
		t.Errorf("Expected dasherizeNumber of 303  to be 3-0-3, not %v", test2)
	}

	if test3 != "3-3-3" {
		t.Errorf("Expected dasherizeNumber of 333 to be 3-3-3, not %v", test3)
	}

	if test4 != "3-22-3" {
		t.Errorf("Expected dasherizeNumber of 3223  to be 3-22-3, not %v", test4)
	}
}