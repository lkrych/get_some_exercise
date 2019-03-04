func TestIsPrime(t *testing.T) {
	test1 := isPrime(2)
	test2 := isPrime(3)
	test3 := isPrime(4)
	test4 := isPrime(9)
	test5 := isPrime(227)

	if test1 != true {
		t.Errorf("Expected isPrime of 2 to be true, not false")
	}

	if test2 != true {
		t.Errorf("Expected isPrime of 3 to be true, not false")
	}

	if test3 != false {
		t.Errorf("Expected isPrime of 4 to be false, not true")
	}

	if test4 != false {
		t.Errorf("Expected isPrime of 9 to be false, not true")
	}

	if test5 != true {
		t.Errorf("Expected isPrime of 227 to be true, not false")
	}
}