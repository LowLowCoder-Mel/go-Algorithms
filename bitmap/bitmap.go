package bitmap

// 一个byte是4位 uint是32位的
// uint -> 8 * 4 = 32 (8个byte)
type BitMap []byte

func New(length uint) BitMap {
	return make([]byte, length/8+1)
}

func (b BitMap) Set(value uint) {
	byteIndex := value / 8
	if byteIndex >= uint(len(b)) {
		return
	}

	bitIndex := value % 8
	BitMap(b)[byteIndex] |= 1 << bitIndex
}

func (b BitMap) Get(value uint) bool {
	byteIndex := value / 8
	if byteIndex >= uint(len(b)) {
		return false
	}

	bitIndex := value % 8
	return BitMap(b)[byteIndex]&(1<<bitIndex) != 0
}

