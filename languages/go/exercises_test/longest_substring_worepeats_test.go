import "testing"

func TestLongestSubString(t *testing.T) {
	test1 := longestSubString("abcabcbb")
	test2 := longestSubString("bbbbb")
	test3 := longestSubString("pwwkew")

	if test1 != 3 {
		t.Errorf("Expected longestSubString of abcabcbb to be 3, not %v", test1)
	}

	if test2 != 1 {
		t.Errorf("Expected longestSubString of abba to be 1, not %v", test2)
	}

	if test3 != 3 {
		t.Errorf("Expected longestSubString of pwwkew to be 3, not %v", test3)
	}
}