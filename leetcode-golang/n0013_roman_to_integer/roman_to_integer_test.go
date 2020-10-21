// go test -v .
package n0013_roman_to_integer

import "testing"

func TestRomanToInt(t *testing.T) {
	if ans := romanToInt("III"); ans != 3 {
		t.Error("Failure: ", ans)
	}

	if ans := romanToInt("IV"); ans != 4 {
		t.Error("Failure: ", ans)
	}

	if ans := romanToInt("IX"); ans != 9 {
		t.Error("Failure: ", ans)
	}

	if ans := romanToInt("LVIII"); ans != 58 {
		t.Error("Failure: ", ans)
	}

	if ans := romanToInt("MCMXCIV"); ans != 1994 {
		t.Error("Failure: ", ans)
	}
}
