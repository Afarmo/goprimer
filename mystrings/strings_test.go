package mystrings

import (
	"math"
	"strings"
	"testing"
)

func TestIndex(t *testing.T) {
	tests := []struct {
		Name      string
		Value     string
		Substring string
	}{
		{"Empty String", "", "abc"},
		{"Empty Substring", "abc", ""},
		{"Empty String", "abcabc", "abc"},
		{"First Index", "abcdef", "a"},
		{"Last Index", "abcdef", "f"},
		{"Emojis", "🍎🍎🍎", "🍎"},
		{"Space", "Hex agon", " "},
		{"Leading Space", " Hexagon", " "},
		{"Trailing Space", "Hexagon ", " "},
		{"Trailing Space", "Hexagon ", " "},
		{"Long Substring", "Octagon", "Octagonn"},
		{"Partial Match", "  x  ", "   "},
		{"Digits", "1234567890", "89"},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			got := Index(tt.Value, tt.Substring)
			want := strings.Index(tt.Value, tt.Substring)
			if got != want {
				t.Errorf("Index(%s, %s) = %d; want %d", tt.Value, tt.Substring, got, want)
			}
		})
	}
}

func TestHasPrefix(t *testing.T) {
	tests := []struct {
		Name   string
		Value  string
		Prefix string
	}{
		{"Empty Value", "", "a"},
		{"Empty Prefix", "abc", ""},
		{"Empty Prefix", "abc", ""},
		{"Long Prefix", "abc", "abcd"},
		{"Exact Match", "abc", "abc"},
		{"Space", "   ", " "},
		{"Case Sensitivity", "ABC", "abc"},
		{"Leading Space ", " Go", "Go"},
		{"Emoji ", "🍎", "🍎"},
		{"Newline", "\n", "\n"},
		{"Rune Error", "\uFFFD", "\uFFFD"},
		{"Invalid Rune", "\xff", "\xff"},
		{"Highest Code Point ", "\U0010FFFF", "\U0010FFFF"},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			got := HasPrefix(tt.Value, tt.Prefix)
			want := strings.HasPrefix(tt.Value, tt.Prefix)
			if got != want {
				t.Errorf("HasPrefic(%s, %s) = %t; want %t", tt.Value, tt.Prefix, got, want)
			}
		})
	}
}

func TestCut(t *testing.T) {
	tests := []struct {
		Name  string
		Value string
		Sep   string
	}{
		{"Both Empty", "", ""},
		{"Separator At Start", "abc", "a"},
		{"Separator At End", "abc", "c"},
		{"Separator In Middle", "abc", "b"},
		{"Multiple Occurrences", "a-b-c", "-"},
		{"Overlapping Pattern", "aaaa", "aa"},
		{"Null Byte", "a\x00b", "\x00"},
		{"Whitespace Variants", "\t\r\n", "\r"},
		{"CRLF Split", "line1\r\nline2", "\r\n"},
		{"Multi-Byte Split", "🍎", "\xf0\x9f"},
		{"Combining Diacritic", "e\u0301", "e"},
		{"Precomposed vs Decomposed", "\u00e9", "e\u0301"},
		{"Zero Width Joiner", "👨‍👩‍👧‍👦", "\u200d"},
		{"BOM Prefix", "\ufeffhello", "\ufeff"},
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			before, after, found := Cut(tt.Value, tt.Sep)
			stdBefore, stdAfter, stdFound := strings.Cut(tt.Value, tt.Sep)
			if before != stdBefore || after != stdAfter || found != stdFound {
				t.Errorf("Cut(%s, %s) = %s, %s, %t; want %s, %s, %t",
					tt.Value,
					tt.Sep,
					before,
					after,
					found,
					stdBefore,
					stdAfter,
					stdFound,
				)
			}
		})
	}
}

func TestCount(t *testing.T) {
	tests := []struct {
		Name      string
		Value     string
		Substring string
	}{
		{"Both Empty", "", ""},
		{"Empty Value", "", "a"},
		{"Empty Substring", "abc", ""},
		{"Empty Substring Multi-Byte", "🍎🍌", ""},
		{"Exact Match", "abc", "abc"},
		{"Long Substring", "abc", "abcd"},
		{"No Match", "abc", "z"},
		{"Case Sensitivity", "ABC", "abc"},
		{"Single Char Multiple", "banana", "a"},
		{"Overlapping Pattern", "aaaa", "aa"},
		{"Overlapping Complex", "abababa", "aba"},
		{"Adjacent Matches", "ababab", "ab"},
		{"Leading And Trailing", "-a-b-", "-"},
		{"Whitespace", "   ", " "},
		{"Newlines", "\n\n\n", "\n"},
		{"Null Byte", "a\x00b\x00c", "\x00"},
		{"Emoji", "🍎🍎🍎", "🍎"},
		{"Partial Multi-Byte", "🍎", "\xf0\x9f"},
		{"Combining Diacritic", "e\u0301e\u0301", "e"},
		{"Precomposed vs Decomposed", "\u00e9", "e\u0301"},
		{"Zero Width Joiner", "👨‍👩‍👧‍👦", "\u200d"},
		{"Rune Error", "\uFFFD\uFFFD", "\uFFFD"},
		{"Invalid Rune", "\xff\xff", "\xff"},
		{"Highest Code Point", "\U0010FFFF\U0010FFFF", "\U0010FFFF"},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			got := Count(tt.Value, tt.Substring)
			want := strings.Count(tt.Value, tt.Substring)
			if got != want {
				t.Errorf("Count(%s, %s) = %d; want %d", tt.Value, tt.Substring, got, want)
			}

		})
	}
}

func TestIndexRune(t *testing.T) {
	tests := []struct {
		Name  string
		Value string
		Rune  rune
	}{
		{"Empty String", "", 'a'},
		{"Zero Rune", "", 0},
		{"Null Byte In String", "a\x00b", '\x00'},
		{"First Character", "abc", 'a'},
		{"Middle Character", "abc", 'b'},
		{"Last Character", "abc", 'c'},
		{"Multiple Occurrences", "banana", 'a'},
		{"No Match", "abc", 'z'},
		{"Case Sensitivity", "ABC", 'a'},
		{"Space", "   ", ' '},
		{"Newline", "a\nb", '\n'},
		{"Two Byte Rune", "café", 'é'},
		{"Three Byte Rune", "€100", '€'},
		{"Four Byte Emoji", "a🍎b", '🍎'},
		{"Multi-Byte Byte Offset", "🍎🍌", '🍌'},
		{"Combining Diacritic", "e\u0301", '\u0301'},
		{"Zero Width Joiner", "👨‍👩‍👧‍👦", '\u200d'},
		{"Rune Error Literal", "a\uFFFDb", '\uFFFD'},
		{"Invalid UTF-8 Byte", "a\xffb", '\uFFFD'},
		{"Negative Rune", "abc", -1},
		{"Surrogate Half", "abc", 0xD800},
		{"Beyond MaxRune", "abc", 0x110000},
		{"Highest Code Point", "a\U0010FFFFb", '\U0010FFFF'},
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			got := IndexRune(tt.Value, tt.Rune)
			want := strings.IndexRune(tt.Value, tt.Rune)
			if got != want {
				t.Errorf("IndexRune(%s, %v) = %d; want %d", tt.Value, tt.Rune, got, want)
			}
		})
	}
}

func TestTrimSpace(t *testing.T) {
	tests := []struct {
		Name  string
		Value string
	}{
		{"Empty String", ""},
		{"Single Space", " "},
		{"Multiple Spaces", "     "},
		{"Leading Space", "   hello"},
		{"Trailing Space", "hello   "},
		{"Both Sides", "   hello   "},
		{"Preserve Internal Space", "  hello   world  "},
		{"All ASCII Whitespace", "\t\n\v\f\r hello \r\f\v\n\t"},
		{"Only ASCII Whitespace", "\t\n\v\f\r "},
		{"No Whitespace", "hello"},
		{"Null Byte Not Trimmed", "\x00hello\x00"},
		{"Zero Width Space Not Trimmed", "\u200Bhello\u200B"},
		{"BOM Not Trimmed", "\uFEFFhello\uFEFF"},
		{"Emoji With Spaces", "  🍎 🍌  "},
		{"Rune Error", " \uFFFD "},
		{"Invalid UTF-8 Byte", " \xff "},
		{"Only Invalid UTF-8", "\xff\xfe"},
		{"Highest Code Point", " \U0010FFFF "},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			got := TrimSpace(tt.Value)
			want := strings.TrimSpace(tt.Value)
			if got != want {
				t.Errorf("TrimSpace(%s) = %s; want %s", tt.Value, got, want)
			}
		})
	}
}

func TestToLower(t *testing.T) {
	tests := []struct {
		Name  string
		Value string
		Want  string
	}{
		{"Empty String", "", ""},
		{"standard mixed ASCII", "Gopher", "gopher"},
		{"Uppercase Alphabet", "ABCDEFGHIJKLMNOPQRSTUVWXYZ", "abcdefghijklmnopqrstuvwxyz"},
		{"Lowercase", "gopher", "gopher"},
		{"ASCII boundary neighbors", "@AZ[", "@az["},
		{"Non ASCII Uppercase Preserved", "HÉLLO ÜBER Ω", "hÉllo Über Ω"},
		{"Digits, Symbols and Whitespace", "123!@# \t\n\x00", "123!@# \t\n\x00"},
		{"Invalid UTF-8 bytes Preserved", "A\xffB\xfeZ", "a\xffb\xfez"},
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			got := ToLower(tt.Value)
			if got != tt.Want {
				t.Errorf("ToLower(%s) = %q; want %q", tt.Value, got, tt.Want)
			}
		})
	}
}

func TestRepeat(t *testing.T) {
	tests := []struct {
		Name   string
		Value  string
		Count  int
		Want   string
		WantOK bool
	}{
		{"Standard Repeat", "na", 2, "nana", true},
		{"Count Zero", "na", 0, "", true},
		{"Count One", "gopher", 1, "gopher", true},
		{"Negative Count", "na", -1, "", false},
		{"Min Int Count", "na", math.MinInt, "", false},
		{"Empty String (Zero Count)", "", 0, "", true},
		{"Empty String (Positive Count)", "", 10, "", true},
		{"Empty String (Negative Count)", "", -5, "", false},
		{"Empty String (Max Int Count)", "", math.MaxInt, "", true},
		{"Overflow", "ab", math.MaxInt, "", false},
		{"Unicode and Emojis", "🚀é", 3, "🚀é🚀é🚀é", true},
		{"NUll Bytes", "\x00a", 2, "\x00a\x00a", true},
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			got, ok := Repeat(tt.Value, tt.Count)
			if got != tt.Want || ok != tt.WantOK {
				t.Errorf("Repeat(%s, %d) = %s, %t; want %s, %t", tt.Value,
					tt.Count,
					got,
					ok,
					tt.Want,
					tt.WantOK)
			}
		})
	}
}
