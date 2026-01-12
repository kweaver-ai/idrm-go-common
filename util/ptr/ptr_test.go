package ptr

import (
	"fmt"
	"testing"
)

func TestAllPtrFieldsNil(t *testing.T) {
	testCases := []struct {
		obj      interface{}
		expected bool
	}{
		{struct{}{}, true},
		{struct{ Foo int }{12345}, true},
		{&struct{ Foo int }{12345}, true},
		{struct{ Foo *int }{nil}, true},
		{&struct{ Foo *int }{nil}, true},
		{struct {
			Foo int
			Bar *int
		}{12345, nil}, true},
		{&struct {
			Foo int
			Bar *int
		}{12345, nil}, true},
		{struct {
			Foo *int
			Bar *int
		}{nil, nil}, true},
		{&struct {
			Foo *int
			Bar *int
		}{nil, nil}, true},
		{struct{ Foo *int }{new(int)}, false},
		{&struct{ Foo *int }{new(int)}, false},
		{struct {
			Foo *int
			Bar *int
		}{nil, new(int)}, false},
		{&struct {
			Foo *int
			Bar *int
		}{nil, new(int)}, false},
		{(*struct{})(nil), true},
	}
	for i, tc := range testCases {
		name := fmt.Sprintf("case[%d]", i)
		t.Run(name, func(t *testing.T) {
			if actual := AllPtrFieldsNil(tc.obj); actual != tc.expected {
				t.Errorf("%s: expected %t, got %t", name, tc.expected, actual)
			}
		})
	}
}

func TestRef(t *testing.T) {
	type T int

	val := T(0)
	pointer := To(val)
	if *pointer != val {
		t.Errorf("expected %d, got %d", val, *pointer)
	}

	val = T(1)
	pointer = To(val)
	if *pointer != val {
		t.Errorf("expected %d, got %d", val, *pointer)
	}
}

func TestDeref(t *testing.T) {
	type T int

	var val, def T = 1, 0

	out := Deref(&val, def)
	if out != val {
		t.Errorf("expected %d, got %d", val, out)
	}

	out = Deref(nil, def)
	if out != def {
		t.Errorf("expected %d, got %d", def, out)
	}
}

func TestEqual(t *testing.T) {
	type T int

	if !Equal[T](nil, nil) {
		t.Errorf("expected true (nil == nil)")
	}
	if !Equal(To(T(123)), To(T(123))) {
		t.Errorf("expected true (val == val)")
	}
	if Equal(nil, To(T(123))) {
		t.Errorf("expected false (nil != val)")
	}
	if Equal(To(T(123)), nil) {
		t.Errorf("expected false (val != nil)")
	}
	if Equal(To(T(123)), To(T(456))) {
		t.Errorf("expected false (val != val)")
	}
}
