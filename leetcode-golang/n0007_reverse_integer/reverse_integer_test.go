// go test -v .
package n0007_reverse_integer

import "testing"

func TestReverse(t *testing.T) {
	if ans := reverse(12345); ans != 54321 {
		t.Error("Failure: ", ans)
	}

	if ans := reverse(123450); ans != 54321 {
		t.Error("Failure: ", ans)
	}

	if ans := reverse(1234500); ans != 54321 {
		t.Error("Failure: ", ans)
	}

	if ans := reverse(12345001); ans != 10054321 {
		t.Error("Failure: ", ans)
	}

	if ans := reverse(-12345); ans != -54321 {
		t.Error("Failure: ", ans)
	}

	if ans := reverse(-123450); ans != -54321 {
		t.Error("Failure: ", ans)
	}

	if ans := reverse(-1234500); ans != -54321 {
		t.Error("Failure: ", ans)
	}

	if ans := reverse(-12345001); ans != -10054321 {
		t.Error("Failure: ", ans)
	}

	if ans := reverse(2147483647); ans != 0 {
		t.Error("Failure: ", ans)
	}

	if ans := reverse(-2147483647); ans != 0 {
		t.Error("Failure: ", ans)
	}

	if ans := reverse(-1563847412); ans != 0 {
		t.Error("Failure: ", ans)
	}
}
