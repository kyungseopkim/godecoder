package godecoder

import (
	"encoding/binary"
	"encoding/hex"
	"errors"
)

type Bitset []byte

func NewBitset() Bitset {
	return make(Bitset, 8)
}

func NewBitsetWithSize(size int32) Bitset {
	return make(Bitset, size)
}

func (bit Bitset) Set(index int32) {
	bit[index>>3] |= 1 << byte(index%8)
}

func (bit Bitset) Unset(index int32) {
	bit[index>>3] &= ^(1 << byte(index%8))
}

func (bit Bitset) Sets(indexs ...int32) {
	for _, i := range indexs {
		bit.Set(i)
	}
}

func (bit Bitset) From(data []byte) {
	var minSize int = 0
	if len(data) > len(bit) {
		minSize = len(bit)
	} else {
		minSize = len(data)
	}

	for i := 0; i < minSize; i++ {
		bit[i] = data[i]
	}
}

func (bit Bitset) String() string {
	return hex.EncodeToString(bit[:])
}

func (bit Bitset) ToUint64() uint64 {
	return binary.LittleEndian.Uint64(bit)
}

func (bit Bitset) Get(index int32) byte {
	return bit[index>>3] & (1 << byte(index%8))
}

func (bit Bitset) GetRange(start int32, len int32) Bitset {
	ret := NewBitset()
	bit.Rearrange(start, start+len, ret)
	return ret
}

func (bit Bitset) Rearrange(start int32, last int32, ret Bitset) {
	var index = int32(0)

	if last > bit.Size() {
		last = bit.Size()
	}

	for i := start; i < last; i++ {
		if bit.Get(i) != 0 {
			ret.Set(index)
		}
		index++
	}
}

func (bit Bitset) GetStringType(start int32, len int32) (error, string) {
	if len%8 != 0 {
		return errors.New("invalid alignment"), ""
	}
	ret := NewBitsetWithSize(int32(len >> 3))
	bit.Rearrange(start, start+len, ret)
	return nil, string(ret)
}

func (bit Bitset) Size() int32 {
	return int32(len(bit) * 8)
}
