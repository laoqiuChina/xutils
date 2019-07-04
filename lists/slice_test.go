package lists

import (
	"fmt"
	"github.com/gogf/gf/g/test/gtest"
	"testing"
)

func TestSliceContains(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	t.Log(Contains(slice, 10))
	t.Log(Contains(slice, 2))
}

func TestSliceContains2(t *testing.T) {
	slice := []string{"a", "b", "c", "d", "e"}
	t.Log(Contains(slice, "f"))
	t.Log(Contains(slice, "e"))
}

func TestSliceContainsAll(t *testing.T) {
	gtest.Case(t, func() {
		gtest.Assert(ContainsAll([]string{"a", "b", "c"}), false)
		gtest.Assert(ContainsAll([]string{"a", "b", "c"}, "a", "b"), true)
		gtest.Assert(ContainsAll([]string{"a", "b", "c"}, "a", "b", "c", "d"), false)
		gtest.Assert(ContainsAll([]string{"a", "b", "c"}, "d"), false)
		gtest.Assert(ContainsAll([]string{"a", "b", "c"}, "b"), true)
	})
}

func TestSliceContainsAny(t *testing.T) {
	gtest.Case(t, func() {
		gtest.Assert(ContainsAny([]string{"a", "b", "c"}), false)
		gtest.Assert(ContainsAny([]string{"a", "b", "c"}, "a", "b"), true)
		gtest.Assert(ContainsAny([]string{"a", "b", "c"}, "a", "b", "c", "d"), true)
		gtest.Assert(ContainsAny([]string{"a", "b", "c"}, "d"), false)
		gtest.Assert(ContainsAny([]string{"a", "b", "c"}, "b"), true)
	})
}

func TestSliceMap(t *testing.T) {
	s := []string{"a", "b", "c"}
	t.Log(Map(s, func(k int, v interface{}) interface{} {
		return fmt.Sprintf("%d:%s", k, v)
	}))
}

func TestSliceDelete(t *testing.T) {
	slice := []string{"a", "b", "c", "d", "e"}
	n := Delete(slice, "b").([]string)
	t.Log(n)

	slice2 := []int{1, 2, 3, 4, 5}
	n2 := Delete(slice2, 3)
	t.Log(n2)
}

func TestSliceReverse(t *testing.T) {
	a := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	t.Log(a)

	Reverse(a)
	t.Log(a)
}
