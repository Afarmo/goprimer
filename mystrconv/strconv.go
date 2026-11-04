package mystrconv

const (
	Lookup   = "0123456789abcdefghijklmnopqrstuvwxyz"
	maxInt64 = 1<<63 - 1
	minInt64 = 1 << 63

	test = -1 << 63
)

func FormatUint(u uint64, base int) (string, bool) {
	if base > 36 || base < 2 {
		return "", false
	}
	if u == 0 {
		return "0", true
	}
	var res string
	for u > 0 {
		res = string(rune(Lookup[u%uint64(base)])) + res
		u /= uint64(base)
	}

	return res, true
}

func FormatInt(i int64, base int) (string, bool) {
	if base > 36 || base < 2 {
		return "", false
	}

	if i == 0 {
		return "0", true
	}
	var newI uint64
	var neg bool
	if i < 0 {
		newI = uint64(-i)
		neg = true
	} else {
		newI = uint64(i)
	}

	var res string

	for newI > 0 {
		res = string(rune(Lookup[newI%uint64(base)])) + res
		newI /= uint64(base)
	}

	if neg {

		res = "-" + res
	}

	return res, true
}

func ParseUint(s string, base int, bitSize int) (uint64, bool) {
	if !Validator(s, base, &bitSize) {
		return 0, false
	}

	var res uint64
	for _, v := range s {
		temp := res
		res *= uint64(base)
		if res < temp {
			return 0, false
		}
		temp = res

		res += Extract(v)
		if res < temp {
			return 0, false
		}
	}

	if res>>bitSize != 0 {
		return 0, false
	}

	return res, true
}

func Validator(s string, base int, bitSize *int) bool {
	if base > 36 ||
		base < 2 ||
		s == "" ||
		*bitSize > 64 ||
		*bitSize < 0 {
		return false
	}

	if *bitSize == 0 {
		*bitSize = 64
	}

	for _, v := range s {
		if v >= 'A' && v <= 'Z' {
			v = v + 32
		}

		var valid bool
		for _, x := range Lookup[:base] {
			if x == v {
				valid = true
			}
		}

		if !valid {
			return false
		}
	}

	return true
}

func Extract(b rune) uint64 {
	switch {
	case b >= '0' && b <= '9':
		return uint64(b - '0')
	case b >= 'a' && b <= 'z':
		return uint64(b - 'a' + 10)
	case b >= 'A' && b <= 'Z':
		return uint64(b - 'A' + 10)
	default:
		return 0
	}
}

func Atoi(s string) (int, bool) {
	if s == "" {
		return 0, false
	}

	if len(s) == 1 && s[0] == '0' {
		return 0, true
	}

	var neg bool
	if s[0] == '-' || s[0] == '+' {
		if s[0] == '-' {
			neg = true
		}
		s = s[1:]
		if s == "" {
			return 0, false
		}
	}

	var res uint64
	for _, v := range s {
		if v < '0' || v > '9' {
			return 0, false
		}
		res *= 10
		res += uint64(v - '0')
	}

	if (neg && res > minInt64) || (!neg && res > maxInt64) {
		return 0, false
	}

	if neg {
		res = -res
	}

	return int(res), true
}
