package myunicode

func IsDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

func IsLetter(r rune) bool {
	return r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z'
}

func IsSpace(r rune) bool {
	return r == '\t' ||
		r == '\n' ||
		r == '\v' ||
		r == '\f' ||
		r == '\r' ||
		r == ' '
}

func IsUpper(r rune) bool {
	return r >= 'A' && r <= 'Z'
}

func IsLower(r rune) bool {
	return r >= 'a' && r <= 'z'
}

func IsControl(r rune) bool {
	return r >= 0x0 && r <= 0x1F || r == 0x7F
}

func ToUpper(r rune) rune {
	if r >= 'a' && r <= 'z' {
		r -= 32
	}
	return r
}

func ToLower(r rune) rune {
	if r >= 'A' && r <= 'Z' {
		r += 32
	}
	return r
}

func IsPunct(r rune) bool {
	return r >= 32 &&
		r <= 127 &&
		!IsDigit(r) &&
		!IsLetter(r) &&
		!IsSpace(r) &&
		!IsControl(r)
}
