package protocol

import (
	"errors"

	"github.com/ByteFlinger/opcua-go/protocol/encoding"
)

// An OPCUAMessage is an OPC UA protocol message
type OPCUAMessage interface {
	marshal(encoding.OPCUAEncoder) ([]byte, error)
	// unmarshal(b []byte) error
	Type() string
}

var _ OPCUAMessage = &HelloMessage{}

// HelloMessage is an OPC UA Hello message type
type HelloMessage struct {
	ProtocolVersion   uint32
	ReceiveBufferSize uint32
	SendBufferSize    uint32
	MaxMessageSize    uint32
	MaxChunkCount     uint32
	EndpointURL       string
}

func (m *HelloMessage) Type() string {
	return "HEL"
}

func (m *HelloMessage) marshal(enc encoding.OPCUAEncoder) ([]byte, error) {

	if m.ReceiveBufferSize < 8192 {
		return []byte{}, errors.New("ReceiveBufferSize must be at least 8192 bytes")
	}

	if m.SendBufferSize < 8192 {
		return []byte{}, errors.New("ReceiveBufferSize must be at least 8192 bytes")
	}

	if len(m.EndpointURL) > 4096 {
		return []byte{}, errors.New("EndpointURL length cannot be greater than 4096 bytes")
	}

	var err error

	err = enc.WriteUint32(m.ProtocolVersion)

	if err == nil {
		err = enc.WriteUint32(m.ReceiveBufferSize)
	}

	if err == nil {
		err = enc.WriteUint32(m.SendBufferSize)
	}

	if err == nil {
		err = enc.WriteUint32(m.MaxMessageSize)
	}

	if err == nil {
		err = enc.WriteUint32(m.MaxChunkCount)
	}

	if err == nil {
		err = enc.WriteString(m.EndpointURL)
	}

	return enc.Bytes(), err
}

var _ OPCUAMessage = &AckMessage{}

// AckMessage is an OPC UA Acknowledge message type
type AckMessage struct {
	ProtocolVersion   uint32
	ReceiveBufferSize uint32
	SendBufferSize    uint32
	MaxMessageSize    uint32
	MaxChunkCount     uint32
}

func (m *AckMessage) marshal(enc encoding.OPCUAEncoder) ([]byte, error) {

	if m.ReceiveBufferSize < 8192 {
		return []byte{}, errors.New("ReceiveBufferSize must be at least 8192 bytes")
	}

	if m.SendBufferSize < 8192 {
		return []byte{}, errors.New("ReceiveBufferSize must be at least 8192 bytes")
	}

	var err error

	err = enc.WriteUint32(m.ProtocolVersion)

	if err == nil {
		err = enc.WriteUint32(m.ReceiveBufferSize)
	}

	if err == nil {
		err = enc.WriteUint32(m.SendBufferSize)
	}

	if err == nil {
		err = enc.WriteUint32(m.MaxMessageSize)
	}

	if err == nil {
		err = enc.WriteUint32(m.MaxChunkCount)
	}

	return enc.Bytes(), err
}

func (m *AckMessage) Type() string {
	return "ACK"
}

var _ OPCUAMessage = &ErrorMessage{}

// AckMessage is an OPC UA Error message type
type ErrorMessage struct {
	Error  uint32
	Reason string
}

func (m *ErrorMessage) marshal(enc encoding.OPCUAEncoder) ([]byte, error) {

	var err error

	err = enc.WriteUint32(m.Error)

	if err == nil {
		err = enc.WriteString(m.Reason)
	}

	return enc.Bytes(), err
}

func (m *ErrorMessage) Type() string {
	return "ERR"
}
