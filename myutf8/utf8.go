package myutf8

const (
	SurrogateMin = 0xD800
	SurrogateMax = 0xDFFF
	MaxRune      = 0x0010FFFF
	rune1Max     = 0x7F
	rune2Max     = 0x7FF
	rune3Max     = 0xFFFF
	runeError    = 0xFFFD
	mask         = 0b00111111
)

func ValidRune(r rune) bool {
	return r >= 0 && r < SurrogateMin || r > SurrogateMax && r <= MaxRune
}

func RuneLen(r rune) int {
	if !ValidRune(r) {
		return -1
	}

	switch {
	case r <= rune1Max:
		return 1
	case r <= rune2Max:
		return 2
	case r <= rune3Max:
		return 3
	default:
		return 4
	}
}

func EncodeRune(p []byte, r rune) (int, bool) {
	byteLen := RuneLen(r)
	if byteLen == -1 {
		r = runeError
		byteLen = 3
	}

	if len(p) < byteLen {
		return 0, false
	}

	switch byteLen {
	case 1:
		p[0] = byte(r)
		return 1, true
	case 2:
		p[0] = byte(0b11000000 | r>>6)
		p[1] = byte(0b10000000 | r&mask)
		return 2, true
	case 3:
		p[0] = byte(0b11100000 | r>>12)
		p[1] = byte(0b10000000 | r>>6&mask)
		p[2] = byte(0b10000000 | r&mask)
		return 3, true
	case 4:
		p[0] = byte(0b11110000 | r>>18)
		p[1] = byte(0b10000000 | r>>12&mask)
		p[2] = byte(0b10000000 | r>>6&mask)
		p[3] = byte(0b10000000 | r&mask)
		return 4, true
	default:
		return 0, false

	}
}

func AppendRune(p []byte, r rune) []byte {
	var buf [4]byte
	bytesWritten, _ := EncodeRune(buf[:], r)
	return append(p, buf[:bytesWritten]...)
}
