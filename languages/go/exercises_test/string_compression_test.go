import "testing"

func TestStringCompression(t *testing.T) {
	test1 := stringCompression("abcd")
	test2 := stringCompression("aaabbbbccccccdddd")
	test3 := stringCompression("aabbbbcccccddd")

	if test1 != "abcd" {
		t.Errorf("Expected stringCompression of 'abcd' to be 'abcd', not %v", test1)
	}

	if test2 != "a3b4c6d4" {
		t.Errorf("Expected stringCompression of 'aaabbbbccccccdddd' to be 'a3b4c6d4', not %v", test2)
	}

	if test3 != "a2b4c5d3" {
		t.Errorf("Expected stringCompression of 'aabbbbcccccddd' to be 'a2b4c5d3', not %v", test3)
	}
}