package godecoder

import (
	"strings"

	"github.com/kyungseopkim/goarxml"
)

type BitDecoder struct {
	Data   []byte
	Signal goarxml.Signal
}

func NewBitDecoder(payload []byte, signal goarxml.Signal) *BitDecoder {
	return &BitDecoder{Data: payload, Signal: signal}
}

func reverseByteArray(data []byte) []byte {
	if len(data) == 0 {
		return data
	}
	return append(reverseByteArray(data[1:]), data[0])
}

func (decoder BitDecoder) GetString() string {
	if strings.Compare(decoder.Signal.DataType, "string") == 0 {
		startByte, startBit, last := decoder.normalize()
		var unit = decoder.Data[startByte:last]
		data := NewBitsetWithSize(last - startByte)
		data.From(unit)
		if err, value := data.GetStringType(startBit, decoder.Signal.Length); err == nil {
			return value
		}
	}
	return ""
}

func (decoder BitDecoder) GetValue() float64 {
	if strings.Compare(decoder.Signal.DataType, "number") == 0 {
		startByte, startBit, last := decoder.normalize()
		var unit = decoder.Data[startByte:last]
		if decoder.Signal.Endian == goarxml.BIG_ENDIAN {
			unit = reverseByteArray(unit)
			startBit = 64 - startBit - decoder.Signal.Length
		}
		data := NewBitset()
		data.From(unit)
		value := data.GetRange(startBit, decoder.Signal.Length)

		return (float64(value.ToUint64()) * decoder.Signal.Slope) + decoder.Signal.Intercept
	}
	return 0.0
}

func (decoder BitDecoder) normalize() (int32, int32, int32) {
	startByte := decoder.Signal.StartBit >> 3
	startBit := decoder.Signal.StartBit - (startByte * 8)
	lenthByte := decoder.Signal.Length >> 3
	if (decoder.Signal.Length % 8) > 0 {
		lenthByte += 1
	}

	last := startByte + lenthByte
	if int(last) > len(decoder.Data) {
		last = int32(len(decoder.Data))
	}
	return startByte, startBit, last
}
