package binary

import (
	"reflect"
	"testing"
)

func TestPutString(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want []byte
	}{
		{"Empty string", "", []byte{0, 0, 0, 0}},
		{"Some value", "hello", []byte{5, 0, 0, 0, 104, 101, 108, 108, 111}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := make([]byte, 4+len(tt.str))
			PutString(b, tt.str)

			if !reflect.DeepEqual(b, tt.want) {
				t.Errorf("PutString() = %v, want %v", b, tt.want)
			}
		})
	}
}

func TestPutUint32(t *testing.T) {
	tests := []struct {
		name string
		i    uint32
		want []byte
	}{
		{"Zero", 0, []byte{0, 0, 0, 0}},
		{"55", 55, []byte{55, 0, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := make([]byte, 4)
			PutUint32(b, tt.i)

			if !reflect.DeepEqual(b, tt.want) {
				t.Errorf("PutUint32() = %v, want %v", b, tt.want)
			}
		})
	}
}
