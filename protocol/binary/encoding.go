package binary

import (
	"encoding/binary"
)

func PutUint32(b []byte, i uint32) {
	binary.LittleEndian.PutUint32(b, i)
}

func PutString(b []byte, str string) {
	binary.LittleEndian.PutUint32(b, uint32(len(str)))

	for i := 0; i < len(str); i++ {
		b[4+i] = str[i]
	}
}

func Uint32(b []byte) uint32 {
	return binary.LittleEndian.Uint32(b)
}
