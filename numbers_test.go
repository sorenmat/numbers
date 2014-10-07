package numbers

import "testing"

func TestRoundUp(t *testing.T) {
	x := Round(2.7)
	if x != 3 {
		t.Error("2.7 rounded should be 3, but was ", x)
	}

	y := Round(-2.7)
	if y != -3 {
		t.Error("-2.7 rounded should be -3, but was ", y)
	}

}

func TestRoundDown(t *testing.T) {
	x := Round(2.2)
	if x != 2 {
		t.Error("2.2 rounded should be 2")
	}

	y := Round(-2.2)
	if y != -2 {
		t.Error("-2.2 rounded should be -2")
	}

}
