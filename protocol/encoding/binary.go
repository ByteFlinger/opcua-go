package encoding

import (
	"bytes"
)

// BinaryEncoder implements an OPCUAEncoder{} for the OPCUA Binary protocol
// It makes use of a bytes.Buffer under the hood so all Write/Read functions have
// the same limitations and constraints as byte.Buffer
type BinaryEncoder struct {
	buf bytes.Buffer
}

var _ OPCUAEncoder = &BinaryEncoder{}

func (e *BinaryEncoder) WriteUint32(n uint32) error {
	err := e.buf.WriteByte(byte(n))
	if err == nil {
		err = e.buf.WriteByte(byte(n >> 8))
	}

	if err == nil {
		err = e.buf.WriteByte(byte(n >> 16))
	}

	if err == nil {
		err = e.buf.WriteByte(byte(n >> 24))
	}

	return err
}

func (e *BinaryEncoder) WriteString(s string) error {
	err := e.WriteUint32(uint32(len(s)))

	if err == nil {
		_, err = e.buf.Write([]byte(s))
	}

	return err
}

func (e *BinaryEncoder) WriteByte(b byte) error {
	return e.buf.WriteByte(b)
}

func (e *BinaryEncoder) Write(bb []byte) (int, error) {
	return e.buf.Write(bb)
}

// Bytes returns a byte array representation of the current state
// of the encoder. See documentation on bytes.Buffer.Bytes() for more details
func (e *BinaryEncoder) Bytes() []byte {
	return e.buf.Bytes()
}
