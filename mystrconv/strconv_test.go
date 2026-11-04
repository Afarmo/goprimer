package mystrconv

import (
	"strconv"
	"testing"
)

func TestFormatUint(t *testing.T) {
	tests := []struct {
		Name  string
		Value uint64
		Base  int
	}{
		{"Zero", 0, 10},
		{"One", 1, 10},
		{"Two", 2, 10},
		{"Max Int", 1<<64 - 1, 10},
		{"Max Int (Base 2)", 1<<64 - 1, 2},
		{"Max Int (Base 36)", 1<<64 - 1, 36},
		{"Max Int (Base 36)", 1<<64 - 1, 36},
		{"Over", 1000, 37},
		{"Over", 1000, 1<<63 - 1},
		{"Under", 1000, 0},
		{"Under", 1000, -1 << 63},
		{"Binary", 1234567890, 2},
		{"Octal", 1234567890, 8},
		{"Hex", 1234567890, 16},
		{"Base 36", 1234567890, 36},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			defer func() {
				if err := recover(); err != nil {
				}
			}()

			got, ok := FormatUint(tt.Value, tt.Base)
			want := strconv.FormatUint(tt.Value, tt.Base)

			stdok := true
			if want == "" {
				stdok = false
			}

			if got != want {
				t.Errorf("FormatUint(%d, %d) = %s, %t; want %s, %t", tt.Value, tt.Base, got, ok, want, stdok)
			}

		})
	}
}

func TestFormatInt(t *testing.T) {
	tests := []struct {
		Name  string
		Value int64
		Base  int
	}{
		{"Zero", 0, 10},
		{"One", 1, 10},
		{"Two", 2, 10},
		{"Max Int (Base 2)", 1<<63 - 1, 2},
		{"Max Int (Base 8)", 1<<63 - 1, 8},
		{"Max Int (Base 10)", 1<<63 - 1, 10},
		{"Max Int (Base 36)", 1<<63 - 1, 36},
		{"Min Int (Base 2)", -1 << 63, 2},
		{"Min Int (Base 8)", -1 << 63, 8},
		{"Min Int (Base 10)", -1 << 63, 10},
		{"Min Int (Base 36)", -1 << 63, 36},
		{"Over", 1000, 37},
		{"Over", 1000, 1<<63 - 1},
		{"Under", 1000, 0},
		{"Under", 1000, -1 << 63},
		{"Negative (Base 2)", -10, 2},
		{"Negative (Base 10)", -10, 10},
		{"Negative (Base 8)", -10, 8},
		{"Negative (Base 36)", -10, 36},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			defer func() {
				if err := recover(); err != nil {
				}
			}()

			got, ok := FormatInt(tt.Value, tt.Base)
			want := strconv.FormatInt(tt.Value, tt.Base)
			stdok := true
			if want == "" {
				stdok = false
			}

			if got != want {
				t.Errorf("FormatInt(%d, %d) = %s, %t; want %s, %t", tt.Value, tt.Base, got, ok, want, stdok)
			}

		})
	}
}

func TestParseUint(t *testing.T) {
	tests := []struct {
		Name    string
		Value   string
		Base    int
		BitSize int
		Want    uint64
		WantOK  bool
	}{
		{"Zero", "0", 10, 64, 0, true},
		{"One", "1", 10, 64, 1, true},
		{"Max Uint", "18446744073709551615", 10, 64, ^uint64(0), true},
		{"Max Uint (Base 2)", "18446744073709551615", 2, 64, 0, false},
		{"Max Uint (Base 10)", "18446744073709551615", 10, 64, ^uint64(0), true},
		{"Max Uint (Base 36)", "18446744073709551615", 36, 64, 0, false},
		{"Over", "1000", 37, 64, 0, false},
		{"Under", "1000", 0, 64, 0, false},
		{"Bit Size Zero", "18446744073709551615", 10, 0, ^uint64(0), true},
		{"Large Bit Size", "1000", 10, 65, 0, false},
		{"Negative Bit Size", "1000", 10, -1, 0, false},
		{"Binary", "1000100010001000", 2, 64, 34952, true},
		{"Octal", "7777777777777777", 8, 64, 281474976710655, true},
		{"Hex", "FFFFFFFFFFFFFFFF", 16, 64, ^uint64(0), true},
		{"Base 36", "ZZZZZZZZZZZZZZZZ", 36, 64, 0, false},
		{"Negative Sign", "-1000", 10, 64, 0, false},
		{"Positive Sign", "+1000", 10, 64, 0, false},
		{"Special Characters (Symbols)", "!&%€", 10, 64, 0, false},
		{"Special Characters (Brackets)", "(){}[]", 10, 64, 0, false},
		{"Uint16 Overflow", "65536", 10, 16, 0, false},
		{"Uint8 Overflow", "256", 10, 8, 0, false},
		{"Uint32 Overflow", "4294967296", 10, 32, 0, false},
		{"Uint64 Overflow", "18446744073709551616", 10, 64, 0, false},
		{"Max Uint8", "255", 10, 8, 255, true},
		{"Leading Zeros", "0001000", 10, 64, 1000, true},
		{"Leading Space", " 1000", 10, 64, 0, false},
		{"Trailing Space", "1000 ", 10, 64, 0, false},
		{"Empty", "", 10, 64, 0, false},
		{"Underscore", "1_000", 10, 64, 0, false},
		{"Invalid Digit For Base", "9", 8, 64, 0, false},
		{"Hex Prefix", "0x1f", 16, 64, 0, false},
		{"Binary Prefix", "0b1100", 2, 64, 0, false},
		{"Lowercase Hex", "ff", 16, 64, 255, true},
		{"Minimum Bit Size", "1", 10, 1, 1, true},
		{"Minimum Bit Size Overflow", "2", 10, 1, 0, false},
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			got, ok := ParseUint(tt.Value, tt.Base, tt.BitSize)
			if got != tt.Want {
				t.Errorf("ParseUint(%s, %d, %d) = %d, %t; want %d, %t", tt.Value, tt.Base, tt.BitSize, got, ok, tt.Want, tt.WantOK)
			}

		})
	}
}

func TestAtoi(t *testing.T) {
	tests := []struct {
		Name  string
		Value string
	}{
		{"Zero", "0"},
		{"One", "1"},
		{"Max Uint", "9223372036854775807"},
		{"Positive Sign Only", "+"},
		{"Negative Sign Only", "-"},
		{"Positive Sign", "+1000"},
		{"Double Negative Sign", "--1000"},
		{"Double Positive Sign", "++1000"},
		{"Negative Sign", "-1000"},
		{"Positive Sign", "+1000"},
		{"Double Negative Sign", "--1000"},
		{"Double Positive Sign", "++1000"},
		{"Special Characters (Symbols)", "!&%€"},
		{"Special Characters (Brackets)", "(){}[]"},
		{"Leading Zeros", "0001000"},
		{"Leading Space", " 1000"},
		{"Trailing Space", "1000 "},
		{"Empty", ""},
		{"Underscore", "1_000"},
		{"Decimal Point", "1.000"},
		{"Hex Prefix", "0x1f"},
		{"Binary Prefix", "0b1100"},
		{"Lowercase Letters", "abc"},
		{"Uppercase Letters", "ABC"},
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			got, ok := Atoi(tt.Value)
			want, err := strconv.Atoi(tt.Value)
			stdok := true
			if err != nil {
				stdok = false
			}

			if got != want || ok != stdok {
				t.Errorf("Atoi(%s) = %d, %t; want %d, %t", tt.Value, got, ok, want, stdok)
			}
		})
	}
}
