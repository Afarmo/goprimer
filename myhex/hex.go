package myhex

const Lookup = "0123456789abcdef"

func Encode(src []byte) []byte {
	if len(src) == 0 {
		return []byte{}
	}

	// var res []byte
	// for _, v := range src {
	// 	highNibble := v >> 4
	// 	lowNibble := v & 0b00001111
	// 	res = append(res, Lookup[highNibble], Lookup[lowNibble])
	// }

	var res = make([]byte, len(src)*2)

	for i, v := range src {
		highNibble := v >> 4
		lowNibble := v & 0b00001111
		res[i*2] = Lookup[highNibble]
		res[i*2+1] = Lookup[lowNibble]
	}

	return res
}

func Extract(b byte) byte {
	switch {
	case b >= '0' && b <= '9':
		return b - '0'
	case b >= 'a' && b <= 'f':
		return b - 'a' + 10
	case b >= 'A' && b <= 'F':
		return b - 'A' + 10
	}

	return 16
}

func Decode(src []byte) ([]byte, bool) {
	if len(src)%2 != 0 {
		return nil, false
	}

	var res = make([]byte, len(src)/2)

	for i := 0; i < len(src); i += 2 {
		high := Extract(src[i])
		low := Extract(src[i+1])
		if high == 16 || low == 16 {
			return nil, false
		}

		res[i/2] = high<<4 + low
	}

	return res, true
}
