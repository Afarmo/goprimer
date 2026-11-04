package mybits

func OnesCount32(x uint32) int {
	var count uint32
	for x > 0 {
		count += x & 1
		x >>= 1
	}
	return int(count)
}

func Len32(x uint32) int {
	var count int

	for count < 32 {
		if x < 1<<count {
			return count
		}
		count++
	}
	return count
}

func RotateLeft32(x uint32, k int) uint32 {
	k %= 32
	if k < 0 {
		k = -k
		return x>>k | x<<(32-k)
	}

	return x<<k | x>>(32-k)
}

func Reverse32(x uint32) uint32 {
	var res uint32

	for i := range 32 {
		res = res | (x&1)<<(31-i)
		x >>= 1
	}

	return res
}

func ReverseBytes32(x uint32) (res uint32) {
	for i := 1; i <= 4; i++ {
		res <<= 8
		res = (x & 0b11111111) | res
		x >>= 8
	}

	return res
}

func ReverseBytes32v2(x uint32) (res uint32) {
	res = (x & 0x000000FF) << 24
	res = ((x>>8)&0x000000FF)<<16 | res
	res = ((x>>16)&0x000000FF)<<8 | res
	res = ((x >> 24) & 0x000000FF) | res

	return res
}

func ReverseBytes32v3(x uint32) (res uint32) {
	res = (x & 0x000000FF) << 24
	res = (x&0x0000FF00)<<8 | res
	res = (x&0x00FF0000)>>8 | res
	res = (x&0xFF000000)>>24 | res

	// res += (x & 0x000000FF) << 24
	// res += (x & 0x0000FF00) << 8
	// res += (x & 0x00FF0000) >> 8
	// res += (x & 0xFF000000) >> 24

	return res
}
