package reverse

func ReverseBits(n uint32) uint32 {
	n = (n >> 16) | (n << 16)
	n = ((n & 0xFF00FF00) >> 8) | ((n & 0x00FF00FF) << 8)
	n = ((n & 0xF0F0F0F0) >> 4) | ((n & 0x0F0F0F0F) << 4)
	n = ((n & 0xCCCCCCCC) >> 2) | ((n & 0x33333333) << 2)
	n = ((n & 0xAAAAAAAA) >> 1) | ((n & 0x55555555) << 1)
	return n
}

func byteByByte(n uint32) uint32 {
	var ans uint32
	power := uint32(31)
	for n != 0 {
		ans += (n&1)<<power
		n >>= 1
		power--
	}
	return ans
}
