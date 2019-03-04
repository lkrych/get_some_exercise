func TestCaesarCipher(t *testing.T) {
	test1 := caesarCipher(3, "abc")
	test2 := caesarCipher(3, "abc xyz")
	test3 := caesarCipher(29, "abc xyz")

	if test1 != "def" {
		t.Errorf("Expected caesarCipher with a shift of 3 and input of abc to be def, not %v", test1)
	}

	if test2 != "def abc" {
		t.Errorf("Expected caesarCipher with a shift of 3 and input of abc xyz to be def abc, not %v", test2)
	}

	if test3 != "def abc" {
		t.Errorf("Expected caesarCipher with a shift of 29 and input of abc xyz to be def abc, not %v", test3)
	}
}