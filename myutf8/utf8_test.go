package myutf8

import (
	"fmt"
	"slices"
	"testing"
	"unicode"
	"unicode/utf8"
)

func TestValidRune(t *testing.T) {
	for r := range unicode.MaxRune {
		got := ValidRune(r)
		want := utf8.ValidRune(r)
		if got != want {
			t.Errorf("ValidRune(%v) = %v; want %v", r, got, want)
		}
	}
}

func TestEncodeRune(t *testing.T) {
	runes := []rune{-1, 0x0, rune1Max, rune2Max, rune3Max, utf8.MaxRune,
		utf8.MaxRune + 1, SurrogateMin, SurrogateMax,
		'🎶', '☽', '℃', '℉', '🍎', '🍊', '🎱', '🧩', '‼', '☆',
		'A', 'a', '0'}
	for _, r := range runes {
		t.Run(fmt.Sprintf("%U", r), func(t *testing.T) {
			var wantBuf [4]byte

			want := utf8.EncodeRune(wantBuf[:], r)
			for size := 0; size <= 4; size++ {
				var gotBuf [4]byte

				got, ok := EncodeRune(gotBuf[:size], r)

				if size < want {
					if got != 0 || ok {
						t.Errorf("EncodeRune (%U, size=%d) = %d, %t; want 0, false", r, size, got, ok)
					}
					continue
				}

				if got != want || !ok {
					t.Errorf("EncodeRune (%U, size=%d) = %d, %t; want %d, true", r, size, got, ok, want)

				}

				if !slices.Equal(gotBuf[:want], wantBuf[:want]) {
					t.Errorf("EncodeRune (%U) = %v; want %v", r, gotBuf, wantBuf)
				}
			}
		})
	}
}

func TestAppendRune(t *testing.T) {
	runes := []rune{-1, 0x0, rune1Max, rune2Max, rune3Max, utf8.MaxRune,
		utf8.MaxRune + 1, SurrogateMin, SurrogateMax,
		'🎶', '☽', '℃', '℉', '🍎', '🍎', '🎱', '🧩', '‼', '☆',
		'A', 'a', '0'}

	for _, r := range runes {
		t.Run(fmt.Sprintf("%U", r), func(t *testing.T) {
			var gotBuf [4]byte
			var wantBuf [4]byte

			got := AppendRune(gotBuf[:0], r)
			want := utf8.AppendRune(wantBuf[:0], r)
			if !slices.Equal(got, want) {
				t.Errorf("AppendRune(%U) = %v; want %v", r, got, want)
			}
		})
	}
	for _, r := range runes {
		t.Run(fmt.Sprintf("%U", r), func(t *testing.T) {

			got := AppendRune([]byte("🍎🎱🧩"), r)
			want := utf8.AppendRune([]byte("🍎🎱🧩"), r)
			if !slices.Equal(got, want) {
				t.Errorf("AppendRune(%U) = %v; want %v", r, got, want)
			}
		})
	}
}
