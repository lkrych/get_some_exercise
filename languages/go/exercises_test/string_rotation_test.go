import "testing"

func TestReverse(t *testing.T) {
	test1 := stringRotation("candelabra", "elabracand")
	test2 := stringRotation("reversethissentence", "entencereversethiss")
	test3 := stringRotation("blah", "cactus")

	if test1 != true {
		t.Errorf("Expected stringRotation to be able to detect a rotation in candelabra and elabracand")
	}

	if test2 != true {
		t.Errorf("Expected stringRotation to be able to detect a rotation in reversethissentence and entencereversethiss")
	}

	if test3 != false {
		t.Errorf("Expected stringRotation to be able to detect there isn't a rotation in blah and cactus")
	}

}