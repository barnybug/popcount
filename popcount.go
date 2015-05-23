package popcount

func Slow32(i uint32) (n uint8) {
	for ; i > 0; i &= i - 1 {
		n += 1
	}
	return
}

func Slow64(i uint64) (n uint8) {
	for ; i > 0; i &= i - 1 {
		n += 1
	}
	return
}

func Hamming32(i uint32) uint8 {
	i -= ((i >> 1) & 0x55555555)
	i = (i & 0x33333333) + ((i >> 2) & 0x33333333)
	i = ((i + (i >> 4)) & 0x0F0F0F0F)
	i = (i * 0x01010101) >> 24
	return uint8(i)
}

func Hamming64(i uint64) uint8 {
	i = i - ((i >> 1) & 0x5555555555555555)
	i = (i & 0x3333333333333333) + ((i >> 2) & 0x3333333333333333)
	i = (i + (i >> 4)) & 0x0F0F0F0F0F0F0F0F
	i = (i * 0x0101010101010101) >> 56
	return uint8(i)
}

// defined in popcount_amd64.s
func Fast32(i uint32) uint8
func Fast64(i uint64) uint8
func PopCnt32(i uint32) uint8

var btable [1 << 8]byte

func init() {
	for i := range btable {
		var n byte
		for x := i; x != 0; x >>= 1 {
			if x&1 != 0 {
				n++
			}
		}
		btable[i] = n
	}
}

func ByteTable32(i uint32) uint8 {
	return uint8(btable[byte(i>>(0*8))] +
		btable[byte(i>>(1*8))] +
		btable[byte(i>>(2*8))] +
		btable[byte(i>>(3*8))])
}

func ByteTable64(i uint64) uint8 {
	return uint8(btable[byte(i>>(0*8))] +
		btable[byte(i>>(1*8))] +
		btable[byte(i>>(2*8))] +
		btable[byte(i>>(3*8))] +
		btable[byte(i>>(4*8))] +
		btable[byte(i>>(5*8))] +
		btable[byte(i>>(6*8))] +
		btable[byte(i>>(7*8))])
}
