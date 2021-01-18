package tools

import "testing"

func TestInt16ToFloat64(t *testing.T) {
	type m map[uint16]uint16
	type te struct {
		value float64
		reg   m
	}
	testPattern := []te{
		{value: 11, reg: m{1: 11, 2: 65535, 3: 65535, 4: 65535, 5: 65535, 6: 65535}},
		{value: -1, reg: m{1: 65535, 2: 0, 3: 65535, 4: 65535, 5: 65535, 6: 65535}},
		{value: -32768, reg: m{1: 32768, 2: 0, 3: 65535, 4: 65535, 5: 65535, 6: 65535}},
		{value: 32767, reg: m{1: 32767, 2: 0, 3: 65535, 4: 65535, 5: 65535, 6: 65535}},
		{value: 0, reg: m{1: 0, 2: 0, 3: 65535, 4: 65535, 5: 65535, 6: 65535}},
	}

	for _, v := range testPattern {
		x := Int16ToFloat64(v.reg)
		if v.value != x {
			t.Errorf("is: %v should be: %v", x, v.value)
		}
	}
}

func TestInt32ToFloat64(t *testing.T) {
	type m map[uint16]uint16
	type te struct {
		value float64
		reg   m
	}
	testPattern := []te{
		{value: 11, reg: m{1: 0, 2: 11, 3: 65535, 4: 65535, 5: 65535, 6: 65535}},
	}

	for _, v := range testPattern {
		x := Int32ToFloat64(v.reg)
		if v.value != x {
			t.Errorf("is: %v should be: %v", x, v.value)
		}
	}
}
