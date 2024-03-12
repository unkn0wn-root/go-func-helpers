package nullchecker

import (
	"testing"
)

func TestIsNull(t *testing.T) {
	// Test cases
	var tests = []struct {
		name   string
		value  interface{}
		result bool
	}{
		{"EmptySlice", []int{}, true},
		{"EmptyString", "", true},
		{"EmptyMap", map[string]int{}, true},
		{"NilInterface", interface{}(nil), true},
		{"NonNilInterface", interface{}(42), false},
		{"NilChannel", chan int(nil), true},
		{"NilPointer", (*int)(nil), true},
		{"EmptyArray", [0]int{}, true},
		{"NonNilInteger", 10, false},
		{"NonNilBoolean", true, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsNull(tt.value)
			if got != tt.result {
				t.Errorf("IsNull(%v) = %v; want %v", tt.value, got, tt.result)
			}
		})
	}
}
