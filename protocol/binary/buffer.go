package binary

import (
	"bytes"
)

// A Buffer is a variable-sized buffer of bytes with Read and Write methods.
// The zero value for Buffer is an empty buffer ready to use.
// It makes use of a bytes.Buffer under the hood so all Write/Read functions have
// the same limitations and constraints as byte.Buffer
type Buffer struct {
	buf bytes.Buffer
}

func (e *Buffer) WriteUint32(n uint32) error {
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

func (e *Buffer) WriteString(s string) error {
	err := e.WriteUint32(uint32(len(s)))

	if err == nil {
		_, err = e.buf.Write([]byte(s))
	}

	return err
}

// WriteByte appends the byte c to the buffer, growing the buffer as needed.
// The returned error is always nil, but is included to match bufio.Writer's
// WriteByte. If the buffer becomes too large, WriteByte will panic with
// ErrTooLarge.
func (e *Buffer) WriteByte(b byte) error {
	return e.buf.WriteByte(b)
}

// Write appends the contents of p to the buffer, growing the buffer as
// needed. The return value n is the length of p; err is always nil. If the
// buffer becomes too large, Write will panic with ErrTooLarge.
func (e *Buffer) Write(bb []byte) (int, error) {
	return e.buf.Write(bb)
}

// Bytes returns a byte array representation of the current state
// of the encoder. See documentation on bytes.Buffer.Bytes() for more details
func (e *Buffer) Bytes() []byte {
	return e.buf.Bytes()
}
