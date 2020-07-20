package propertybasedtestexample

import (
	"testing"

	"pgregory.net/rapid"
)

func TestAdd(t *testing.T) {
	t.Run("commutativity", rapid.MakeCheck(func(t *rapid.T) {
		x := rapid.Int().Draw(t, "x").(int)
		y := rapid.Int().Draw(t, "y").(int)
		result1 := add(x, y)
		result2 := add(y, x)
		if result1 != result2 {
			t.Fatalf("add must be commutative, x=%d, y=%d, result1=%d, result2=%d", x, y, result1, result2)
		}
	}))
	t.Run("associativity", rapid.MakeCheck(func(t *rapid.T) {
		x := rapid.Int().Draw(t, "x").(int)
		result1 := add(add(x, 1), 1)
		result2 := add(x, 2)
		if result1 != result2 {
			t.Fatalf("add 1 twice must be equal to add 2, x=%d, result1=%d, result2=%d", x, result1, result2)
		}
	}))
	t.Run("identity", rapid.MakeCheck(func(t *rapid.T) {
		x := rapid.Int().Draw(t, "x").(int)
		result1 := add(x, 0)
		result2 := x
		if result1 != result2 {
			t.Fatalf("add 0 must be equal to original, x=%d, result1=%d, result2=%d", x, result1, result2)
		}
	}))
}
