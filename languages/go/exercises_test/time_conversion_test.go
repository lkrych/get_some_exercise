func TestTimeConversion(t *testing.T) {
	test1 := timeConversion(15)
	test2 := timeConversion(150)
	test3 := timeConversion(360)

	if test1 != "0:15" {
		t.Errorf("Expected TimeConversion of 15 to be 0:15, not %v", test1)
	}

	if test2 != "2:30" {
		t.Errorf("Expected TimeConversion of 150 to be 2:30, not %v", test2)
	}

	if test3 != "6:00" {
		t.Errorf("Expected TimeConversion of 360 to be 6:00, not %v", test3)
	}
}