package godecoder

import "github.com/kyungseopkim/goarxml"

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

func (decoder BitDecoder) GetValue() float32 {
    idx := decoder.Signal.StartBit / 64
    startByte := idx * 8
    sliceStart := startByte * 8
    var startBit  = decoder.Signal.StartBit - sliceStart
    last := startByte+8
    if int(last) > len(decoder.Data) {
        last = int32(len(decoder.Data))
    }
    var unit = decoder.Data[startByte : last]
    if decoder.Signal.Endian == goarxml.BIG_ENDIAN {
        unit = reverseByteArray(unit)
        startBit = 64 - startBit - decoder.Signal.Length
    }
    data := NewBitset()
    data.From(unit)
    value := data.GetRange(startBit, decoder.Signal.Length)

    return (float32(value.ToUint64()) * decoder.Signal.Slope) + decoder.Signal.Intercept
}

