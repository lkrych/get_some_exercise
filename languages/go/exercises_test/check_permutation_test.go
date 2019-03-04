import "testing"

func TestCheckPermutation(t *testing.T) {
	test1 := checkPermutation("blah", "halb")
	test2 := checkPermutation("baleen", "neelab")
	test3 := checkPermutation("cactus", "flower")
	test4 := checkPermutation("ballsy", "heptagon")
	test5 := checkPermutation("winky face", "kin wafeyc")

	if test1 != true {
		t.Errorf("Expected checkPermutation to detect that blah and halb are permutations of eachother")
	}

	if test2 != true {
		t.Errorf("Expected checkPermutation to detect that baleen and neelab are permutations of eachother")
	}

	if test3 != false {
		t.Errorf("Expected checkPermutation to detect that cactus and flower are not permutations of eachother")
	}

	if test4 != false {
		t.Errorf("Expected checkPermutation to detect that ballsy and heptagon are not permutations of eachother")
	}

	if test5 != true {
		t.Errorf("Expected checkPermutation to detect that winky face and kin wafeyc are permutations of eachother")
	}
}