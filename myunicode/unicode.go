package myunicode

func IsDigit(r rune) bool {
	if ok := InASCIIRange(r, false); !ok {
		return false
	}
	return r >= '0' && r <= '9'
}

func IsLetter(r rune) bool {
	if ok := InASCIIRange(r, false); !ok {
		return false
	}
	return r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z'
}

func IsSpace(r rune) bool {
	if ok := InASCIIRange(r, false); !ok {
		return false
	}
	return r == ' '
}

func IsUpper(r rune) bool {
	if ok := InASCIIRange(r, false); !ok {
		return false
	}
	return r >= 'A' && r <= 'Z'
}

func IsLower(r rune) bool {
	if ok := InASCIIRange(r, false); !ok {
		return false
	}
	return r >= 'a' && r <= 'z'
}

func IsPunct(r rune) bool {
	if ok := InASCIIRange(r, true); !ok {
		return false
	}

	if r >= 'a' && r <= 'z' ||
		r >= 'A' && r <= 'Z' ||
		r == ' ' ||
		r >= '0' && r <= '9' {
		return false
	}

	return true
}

func IsControl(r rune) bool {
	return r >= 0 && r <= 31
}

func ToUpper(r rune) rune {
	if r >= 'a' && r <= 'z' {
		r -= 'a' - 'A'
	}
	return r
}

func ToLower(r rune) rune {
	if r >= 'A' && r <= 'Z' {
		r += 'a' - 'A'
	}
	return r
}

func InASCIIRange(r rune, printable bool) bool {
	if printable {
		return r >= 32 && r <= 127
	}
	return r >= 0 && r <= 127
}
