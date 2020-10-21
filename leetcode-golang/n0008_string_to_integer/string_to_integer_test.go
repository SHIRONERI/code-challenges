package n0008_string_to_integer

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestMyAtoi(t *testing.T) {
	a := assert.New(t)
	a.Equal(int(math.Pow(2, 31)-1), MaxInt32)
	a.Equal(int(-math.Pow(2, 31)), MinInt32)

	digit, ok := isNumber('0')
	a.True(ok)
	a.Equal(0, digit)
	digit, ok = isNumber('3')
	a.True(ok)
	a.Equal(3, digit)
	digit, ok = isNumber('9')
	a.True(ok)
	a.Equal(9, digit)
	_, ok = isNumber(' ')
	a.False(ok)
	_, ok = isNumber('w')
	a.False(ok)

	a.Equal(0, myAtoi(""))
	a.Equal(0, myAtoi("  "))
	a.Equal(0, myAtoi("  +"))
	a.Equal(0, myAtoi("  +-"))
	a.Equal(1, myAtoi("1"))
	a.Equal(4193, myAtoi("4193 with words"))
	a.Equal(0, myAtoi("words and 987"))
	a.Equal(-42, myAtoi("  -42"))
	a.Equal(42, myAtoi("  +42"))
	a.Equal(0, myAtoi("  -+42"))
	a.Equal(0, myAtoi("  -+42   "))
	a.Equal(-2147483648, myAtoi("-91283472332"))
}
