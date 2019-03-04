func TestNearbyAZ(t *testing.T) {
	test1 := nearbyAZ("baz")
	test2 := nearbyAZ("abz")
	test3 := nearbyAZ("abcz")
	test4 := nearbyAZ("a")
	test5 := nearbyAZ("z")
	test6 := nearbyAZ("za")

	if test1 != true {
		t.Errorf("Expected nearbyAZ of baz to be true, not %v", test1)
	}

	if test2 != true {
		t.Errorf("Expected nearbyAZ of abz to be true, not %v", test2)
	}

	if test3 != true {
		t.Errorf("Expected nearbyAZ of abcz to be true, not %v", test3)
	}

	if test4 != false {
		t.Errorf("Expected nearbyAZ of a to be false, not %v", test4)
	}

	if test5 != false {
		t.Errorf("Expected nearbyAZ of z to be false, not %v", test5)
	}

	if test6 != false {
		t.Errorf("Expected nearbyAZ of za to be false, not %v", test6)
	}
}