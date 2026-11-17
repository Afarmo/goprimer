package mystrings

import "goprimer/myunicode"

func Index(s, substr string) int {
	var res = -1
	for i := range s {
		if i+len(substr) <= len(s) && s[i:i+len(substr)] == substr {
			res = i
			break
		}
	}
	return res
}

func HasPrefix(s, prefix string) bool {
	if len(prefix) > len(s) {
		return false
	}
	return s[:len(prefix)] == prefix
}

func Cut(s, sep string) (before, after string, found bool) {
	if sep == "" {
		return "", "", true
	}

	index := Index(s, sep)
	if index == -1 {
		return s, "", false
	}

	return s[:index], s[index+len(sep):], true
}

func Count(s, substr string) int {
	if substr == "" {
		return 1 + len([]rune(s))
	}
	var res int
	for i := 0; i < len(s); i++ {
		if i+len(substr) <= len(s) && s[i:i+len(substr)] == substr {
			res++
			i += len(substr) - 1
		}
	}
	return res
}

func IndexRune(s string, r rune) int {
	for i, v := range s {
		if v == r {
			return i
		}
	}

	return -1
}

func TrimSpace(s string) string {
	for s != "" {
		if myunicode.IsSpace(rune(s[0])) {
			s = s[1:]
			continue
		}

		if myunicode.IsSpace(rune(s[len(s)-1])) {
			s = s[:len(s)-1]
			continue
		}

		break
	}

	return s
}
func ToLower(s string) string {
	var res = make([]byte, len(s))

	for i := 0; i < len(s); i++ {
		v := s[i]
		if v >= 'A' && v <= 'Z' {
			res[i] += v + 32
			continue
		}
		res[i] += v
	}

	return string(res)
}

func Repeat(s string, count int) (string, bool) {
	if count == 0 {
		return "", true
	}

	if count < 0 || len(s) != (len(s)*count)/count {
		return "", false
	}

	if s == "" {
		return "", true
	}

	var res = make([]byte, len(s)*count)
	bytePosition := copy(res, s)
	for bytePosition < len(res) {
		copy(res[bytePosition:], res[:bytePosition])
		bytePosition *= 2
	}

	return string(res), true
}
