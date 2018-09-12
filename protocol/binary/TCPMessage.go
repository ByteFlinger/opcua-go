package binary

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

const headerLen = 8

type TCPMessage struct {
	Type        string
	ChunkType   byte
	PayloadSize uint32
	Payload     []byte
}

// EncodeTCP returns a byte encoded OPCUA TCP message of the given OPCUAMessage
// to be sent over the wirte
func MarshalMessage(msg OPCUAMessage) ([]byte, error) {

	var payload []byte
	var buf bytes.Buffer
	_, err := buf.Write([]byte(msg.Type()))

	if err == nil {
		err = buf.WriteByte(byte('F'))
	}

	if err == nil {
		payload, err = msg.marshal()
	}

	if err == nil {
		size := make([]byte, 4)
		binary.LittleEndian.PutUint32(size, uint32(8+len(payload)))
		_, err = buf.Write(size)
	}

	if err == nil {
		_, err = buf.Write(payload)
	}

	return buf.Bytes(), err
}

// ParseMessage parses a Message from its binary form after determining its
// type from a leading message header.
func ParseMessage(b []byte) (OPCUAMessage, error) {
	if len(b) < headerLen {
		return nil, io.ErrUnexpectedEOF
	}

	var m OPCUAMessage
	switch t := OPCUAType(b[0:3]); t {
	case OPCUAHello:
		m = new(HelloMessage)
	case OPCUAAcknowledge:
		m = new(AckMessage)
	case OPCUAReverseHello:
		return nil, fmt.Errorf("Reverse Hello not supported yet")
	case OPCUAError:
		m = new(ErrorMessage)
	default:
		return nil, fmt.Errorf("opcua: unrecognized OPCUA message type: %s", t)
	}

	if err := m.unmarshal(b[headerLen:]); err != nil {
		return nil, err
	}

	return m, nil
}
