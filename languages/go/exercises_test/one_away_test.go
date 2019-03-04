import "testing"

func TestOneAway(t *testing.T) {
	test1 := oneAway("pale", "ple")
	test2 := oneAway("pales", "pale")
	test3 := oneAway("pale", "bale")
	test4 := oneAway("pale", "bake")
	test5 := oneAway("baleen", "baleen")

	if test1 != true {
		t.Errorf("Expected oneAway of pale to ple to be true, not %v", test1)
	}

	if test2 != true {
		t.Errorf("Expected oneAway of pales to pale to be true, not %v", test2)
	}

	if test3 != true {
		t.Errorf("Expected oneAway of pale to bale to be true, not %v", test3)
	}

	if test4 != false {
		t.Errorf("Expected oneAway of pale to bake to be false, not %v", test4)
	}

	if test5 != false {
		t.Errorf("Expected oneAway of baleen to baleen to be false, not %v", test5)
	}
}