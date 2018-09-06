package encoding

type OPCUAEncoder interface {
	Write([]byte) (int, error)
	WriteByte(byte) error
	WriteUint32(uint32) error
	WriteString(string) error
	Bytes() []byte
}
