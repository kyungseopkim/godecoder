package godecoder

import (
	"fmt"
	"testing"

	"github.com/kyungseopkim/goarxml"
)

var data []byte

func init() {
	data = []byte{82, 187, 1, 0, 0, 0, 4, 194, 1, 0, 0, 0}
}

func TestBitDecoder_GetValue(t *testing.T) {
	sig := goarxml.Signal{
		"IiBESP2_pRunout_InfoT_Chassis",
		1,
		59,
		8,
		1,
		0,
		250,
		0,
		"bar",
		false,
		"number",
	}
	decorder := NewBitDecoder(data, sig)
	fmt.Println(decorder.GetValue())

}

// func TestBitset_GetRange(t *testing.T) {
//     t.Error("hello")

// }
