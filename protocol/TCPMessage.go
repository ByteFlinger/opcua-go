package protocol

import (
	"encoding/binary"
	"fmt"

	"github.com/ByteFlinger/opcua-go/protocol/encoding"
)

type TCPMessage struct {
	Type        string
	ChunkType   byte
	PayloadSize uint32
	Payload     []byte
}

// EncodeTCP returns a byte encoded OPCUA TCP message of the given OPCUAMessage
// to be sent over the wirte
func EncodeTCP(msg OPCUAMessage) ([]byte, error) {

	var payload []byte
	enc := &encoding.BinaryEncoder{}
	_, err := enc.Write([]byte(msg.Type()))

	if err == nil {
		err = enc.WriteByte(byte('F'))
	}

	if err == nil {
		payload, err = msg.marshal(&encoding.BinaryEncoder{})
	}

	if err == nil {
		err = enc.WriteUint32(uint32(8 + len(payload)))
	}

	if err == nil {
		_, err = enc.Write(payload)
	}

	return enc.Bytes(), err
}

// DecodeTCP decodes a byte encoded OPCUA TCP message into a TCPMessage{}
func DecodeTCP(bytes []byte) (*TCPMessage, error) {

	if len(bytes) < 8 {
		return nil, fmt.Errorf("Invalid TCP message of size %d. Must be at least of length 8", len(bytes))
	}

	msgType := string(bytes[:3])
	size := binary.LittleEndian.Uint32(bytes[4:8]) - 8

	fSize := int(size + 8)

	if len(bytes) < fSize {
		return nil, fmt.Errorf("Invalid TCP message. Payload size smaller than expected %d", size)
	}
	payload := bytes[8:fSize]

	msg := &TCPMessage{
		Type:        msgType,
		ChunkType:   'F',
		PayloadSize: size,
		Payload:     payload,
	}

	return msg, nil
}
