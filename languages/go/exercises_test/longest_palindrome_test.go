func TestLongestPalindrome(t *testing.T) {
	test1 := longestPalindrome("abcbd")
	test2 := longestPalindrome("abba")
	test3 := longestPalindrome("abcbdeffe")

	if test1 != "bcb" {
		t.Errorf("Expected longestPalindrome of abcbd to be bcb, not %v", test1)
	}

	if test2 != "abba" {
		t.Errorf("Expected longestPalindrome of abba to be abba, not %v", test2)
	}

	if test3 != "effe" {
		t.Errorf("Expected longestPalindrome of abcbdeffe to be effe, not %v", test3)
	}
}