package mybits

import (
	"math/bits"
	"testing"
)

func TestOnesCount32(t *testing.T) {

	tests := []struct {
		Name  string
		Value uint32
	}{
		{"Max Int", 1<<32 - 1},
		{"Zero", 0},
		{"One", 1},
		{"Two bits", 3},
		{"Three bits", 7},
		{"Highest Bit", 1 << 31},
		{"Apart", 0x80808080},
		{"Alternate", 0xAAAAAAAA},
		{"High Nibble", 0xF0F0F0F0},
		{"Low Nibble", 0x0F0F0F0F},
		{"Doubles", 0x33333333},
		{"Triples", 0x77777777},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			got := OnesCount32(tt.Value)
			want := bits.OnesCount32(tt.Value)

			if got != want {
				t.Errorf("OnesCount32(%d) = %d; want %d", tt.Value, got, want)
			}
		})
	}
}

func TestLen32(t *testing.T) {

	tests := []struct {
		Name  string
		Value uint32
	}{
		{"Zero", 0},
		{"One", 1},
		{"Two", 2},
		{"Three", 4},
		{"Max Int", 1<<32 - 1},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			got := Len32(tt.Value)
			want := bits.Len32(tt.Value)

			if got != want {
				t.Errorf("Len32(%d) = %d; want %d", tt.Value, got, want)
			}
		})
	}
}

func TestRotateLeft32(t *testing.T) {
	tests := []struct {
		Name        string
		Value       uint32
		SecondValue int
	}{
		{"Zero", 0, 0},
		{"One", 1, 1},
		{"Two", 1, 2},
		{"Highest To Lowest", 1 << 31, 1},
		{"Lowest To Highest", 1, 31},
		{"Highest To Lowest (Negative)", 1 << 31, -31},
		{"Lowest To Highest (Negative)", 1, -1},
		{"Low Nibbles To High Nibbles", 0x0F0F0F0F, 4},
		{"High Nibbles To Low Nibbles", 0xF0F0F0F0, 4},
		{"Low Nibbles To High Nibbles (Negative)", 0x0F0F0F0F, -4},
		{"High Nibbles To Low Nibbles (Negative)", 0xF0F0F0F0, -4},
		{"Alternate", 0x55555555, 1},
		{"Doubles", 0x33333333, 2},
		{"Max Int Rotations", 1<<32 - 1, 1<<63 - 1},
		{"Max Int Rotations (Negative)", 1<<32 - 1, -1 << 63},
		{"Full Rotation", 0x12345678, 32},
		{"Full Rotation (Negative)", 0x12345678, -32},
		{"Full Rotation Plus One", 1, 33},
		{"Full Rotation Minus One (Negative)", 2, -33},
		{"Double Full Rotation", 0x87654321, 64},
		{"Double Full Rotation (Negative)", 0x87654321, -64},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			defer func() {
				if err := recover(); err != nil {
					t.Errorf("RotateLeft32() panicked and recovered: %v", err)
				}
			}()
			got := RotateLeft32(tt.Value, tt.SecondValue)
			want := bits.RotateLeft32(tt.Value, tt.SecondValue)

			if got != want {
				t.Errorf("RotateLeft32(%d, %d) = %d; want %d", tt.Value, tt.SecondValue, got, want)
			}

		})
	}
}

func BenchmarkRotateLeft32(b *testing.B) {
	tests := []struct {
		Name        string
		Value       uint32
		SecondValue int
	}{
		{"Zero", 0, 0},
		{"One", 1, 1},
		{"Two", 1, 2},
		{"Highest To Lowest", 1 << 31, 1},
		{"Lowest To Highest", 1, 31},
		{"Highest To Lowest (Negative)", 1 << 31, -31},
		{"Lowest To Highest (Negative)", 1, -1},
		{"Low Nibbles To High Nibbles", 0x0F0F0F0F, 4},
		{"High Nibbles To Low Nibbles", 0xF0F0F0F0, 4},
		{"Low Nibbles To High Nibbles (Negative)", 0x0F0F0F0F, -4},
		{"High Nibbles To Low Nibbles (Negative)", 0xF0F0F0F0, -4},
		{"Alternate", 0x55555555, 1},
		{"Doubles", 0x33333333, 2},
		{"Max Int Rotations", 1<<32 - 1, 1<<63 - 1},
		{"Max Int Rotations (Negative)", 1<<32 - 1, -1 << 63},
		{"Full Rotation", 0x12345678, 32},
		{"Full Rotation (Negative)", 0x12345678, -32},
		{"Full Rotation Plus One", 1, 33},
		{"Full Rotation Minus One (Negative)", 2, -33},
		{"Double Full Rotation", 0x87654321, 64},
		{"Double Full Rotation (Negative)", 0x87654321, -64},
	}

	for _, tt := range tests {
		b.Run(tt.Name, func(b *testing.B) {
			for b.Loop() {
				RotateLeft32(tt.Value, tt.SecondValue)
			}
		})
	}
}

func TestReverse32(t *testing.T) {
	tests := []struct {
		Name  string
		Value uint32
	}{
		{"Zero", 0},
		{"One", 1},
		{"Two", 2},
		{"One", 1},
		{"Max Int", 1<<32 - 1},
		{"Alternate", 0x55555555},
		{"High Nibble", 0xF0F0F0F0},
		{"Low Nibble", 0x0F0F0F0F},
		{"Doubles", 0x33333333},
		{"Doubles v2", 0xCCCCCCCC},
		{"Triples", 0x77777777},
		{"Apart", 0x80808080},
		{"Alternate", 0xAAAAAAAA},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			got := Reverse32(tt.Value)
			want := bits.Reverse32(tt.Value)
			if got != want {
				t.Errorf("Reverse32(%d) = %d; want %d", tt.Value, got, want)
			}
		})
	}
}

func TestReverseBytes32(t *testing.T) {
	tests := []struct {
		Name  string
		Value uint32
	}{
		{"Zero", 0},
		{"One", 1},
		{"Two", 2},
		{"High Nibble", 0xF0F0F0F0},
		{"Low Nibble", 0x0F0F0F0F},
		{"Max Int", 0xFFFFFFFF},
		{"Two Bytes Apart", 0xFF0000FF},
		{"Alternate", 0x55555555},
		{"Doubles", 0x33333333},
		{"Triples", 0x77777777},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			got := ReverseBytes32(tt.Value)
			want := bits.ReverseBytes32(tt.Value)
			if got != want {
				t.Errorf("ReverseBytes32(%d) = %d; want %d", tt.Value, got, want)
			}
		})
	}
}
