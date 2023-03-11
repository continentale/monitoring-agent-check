package icinga

import (
	"testing"
)

func TestImplicitEndWithoutColon(t *testing.T) {
	icinga := NewIcinga("OK", "10", "20")
	icinga.Evaluate(5, "", "", "")
	if icinga.GetStatus() != 0 {
		t.Fatalf("Should be OK but is instead %d", icinga.GetStatus())
	}

	icinga = NewIcinga("OK", "30", "40")
	icinga.Evaluate(31, "", "", "")
	if icinga.GetStatus() != 1 {
		t.Fatalf("Should be Warning but is instead %d", icinga.GetStatus())
	}

	icinga = NewIcinga("OK", "50", "100")
	icinga.Evaluate(101, "", "", "")
	if icinga.GetStatus() != 2 {
		t.Fatalf("Should be Critical but is instead %d", icinga.GetStatus())
	}
}

func TestImplicitEndWithColon(t *testing.T) {
	icinga := NewIcinga("OK", "10:", "20:")
	icinga.Evaluate(21, "", "", "")
	if icinga.GetStatus() != 0 {
		t.Fatalf("Should be OK but is instead %d", icinga.GetStatus())
	}

	icinga = NewIcinga("OK", "30:", "20:")
	icinga.Evaluate(29, "", "", "")
	if icinga.GetStatus() != 1 {
		t.Fatalf("Should be Warning but is instead %d", icinga.GetStatus())
	}

	icinga = NewIcinga("OK", "50:", "25:")
	icinga.Evaluate(24, "", "", "")
	if icinga.GetStatus() != 2 {
		t.Fatalf("Should be Critical but is instead %d", icinga.GetStatus())
	}
}

func TestGivenMaximumAndMinimum(t *testing.T) {
	icinga := NewIcinga("OK", "10:25", "20:30")
	icinga.Evaluate(24, "", "", "")
	if icinga.GetStatus() != 0 {
		t.Fatalf("Should be OK but is instead %d", icinga.GetStatus())
	}

	icinga = NewIcinga("OK", "10:20", "20:30")
	icinga.Evaluate(23, "", "", "")
	if icinga.GetStatus() != 1 {
		t.Fatalf("Should be Warning but is instead %d", icinga.GetStatus())
	}

	icinga = NewIcinga("OK", "50:100", "25:50")
	icinga.Evaluate(0, "", "", "")
	if icinga.GetStatus() != 2 {
		t.Fatalf("Should be Critical but is instead %d", icinga.GetStatus())
	}
}

func TestNegativeDown(t *testing.T) {
	icinga := NewIcinga("OK", "-5:10", "5:20")
	icinga.Evaluate(6, "", "", "")
	if icinga.GetStatus() != 0 {
		t.Fatalf("Should be OK but is instead %d", icinga.GetStatus())
	}

	icinga = NewIcinga("OK", "-10:-5", "1:20")
	icinga.Evaluate(5, "", "", "")
	if icinga.GetStatus() != 1 {
		t.Fatalf("Should be Warning but is instead %d", icinga.GetStatus())
	}

	icinga = NewIcinga("OK", "-100:-50", "-50:0")
	icinga.Evaluate(1, "", "", "")
	if icinga.GetStatus() != 2 {
		t.Fatalf("Should be Critical but is instead %d", icinga.GetStatus())
	}
}

func TestNegativeUnlimited(t *testing.T) {
	icinga := NewIcinga("OK", "~:5", "~:10")
	icinga.Evaluate(2, "", "", "")
	if icinga.GetStatus() != 0 {
		t.Fatalf("Should be OK but is instead %d", icinga.GetStatus())
	}

	icinga = NewIcinga("OK", "~:5", "~:10")
	icinga.Evaluate(6, "", "", "")
	if icinga.GetStatus() != 1 {
		t.Fatalf("Should be Warning but is instead %d", icinga.GetStatus())
	}

	icinga = NewIcinga("OK", "~:5", "~:10")
	icinga.Evaluate(11, "", "", "")
	if icinga.GetStatus() != 2 {
		t.Fatalf("Should be Critical but is instead %d", icinga.GetStatus())
	}
}

func TestNegation(t *testing.T) {
	icinga := NewIcinga("OK", "@0:5", "@0:10")
	icinga.Evaluate(11, "", "", "")
	if icinga.GetStatus() != 0 {
		t.Fatalf("Should be OK but is instead %d", icinga.GetStatus())
	}

	icinga = NewIcinga("OK", "@0:5", "@5:10")
	icinga.Evaluate(2, "", "", "")
	if icinga.GetStatus() != 1 {
		t.Fatalf("Should be Warning but is instead %d", icinga.GetStatus())
	}

	icinga = NewIcinga("OK", "@0:5", "@0:10")
	icinga.Evaluate(7, "", "", "")
	if icinga.GetStatus() != 2 {
		t.Fatalf("Should be Critical but is instead %d", icinga.GetStatus())
	}
}
