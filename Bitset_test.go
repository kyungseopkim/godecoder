package godecoder

import (
	"fmt"
	"testing"
	//"encoding/ascii85"
)

func TestBitset_GetStringType(t *testing.T) {
	raw := []byte("Hello, playground")
	//dst := make([]byte, 25, 25)
	//ascii85.Encode(dst, raw)

	numberOfbytes := int32(len(raw))
	fmt.Printf("len(raw) = %d\n", numberOfbytes)
	bitset := NewBitsetWithSize(numberOfbytes)
	bitset.From(raw)

	if err, value := bitset.GetStringType(0, numberOfbytes*8); err == nil {
		fmt.Println(value)
	}
}
