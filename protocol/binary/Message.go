package binary

import (
	"encoding/binary"
	"errors"
	"io"
)

type OPCUAType string

// OPC UA Message types
// Described in OPC Unified Architecture 1.04, Part 6, 7.1.2.2
const (
	OPCUAHello        OPCUAType = "HEL"
	OPCUAAcknowledge  OPCUAType = "ACK"
	OPCUAError        OPCUAType = "ERR"
	OPCUAReverseHello OPCUAType = "RHE"
)

// An OPCUAMessage is an OPC UA protocol message
type OPCUAMessage interface {
	Type() OPCUAType
	marshal() ([]byte, error)
	unmarshal(b []byte) error
}

// HelloMessage is an OPC UA Hello message type
type HelloMessage struct {
	ProtocolVersion   uint32
	ReceiveBufferSize uint32
	SendBufferSize    uint32
	MaxMessageSize    uint32
	MaxChunkCount     uint32
	EndpointURL       string
}

func (m *HelloMessage) Type() OPCUAType {
	return OPCUAHello
}

func (m *HelloMessage) marshal() ([]byte, error) {

	if m.ReceiveBufferSize < 8192 {
		return []byte{}, errors.New("ReceiveBufferSize must be at least 8192 bytes")
	}

	if m.SendBufferSize < 8192 {
		return []byte{}, errors.New("ReceiveBufferSize must be at least 8192 bytes")
	}

	if len(m.EndpointURL) > 4096 {
		return []byte{}, errors.New("EndpointURL length cannot be greater than 4096 bytes")
	}

	buf := make([]byte, 24+len(m.EndpointURL))

	PutUint32(buf[0:4], m.ProtocolVersion)
	PutUint32(buf[4:8], m.ReceiveBufferSize)
	PutUint32(buf[8:12], m.SendBufferSize)
	PutUint32(buf[12:16], m.MaxMessageSize)
	PutUint32(buf[16:20], m.MaxChunkCount)
	PutString(buf[20:], m.EndpointURL)

	return buf, nil
}

func (m *HelloMessage) unmarshal(b []byte) error {
	if len(b) < 24 {
		return io.ErrUnexpectedEOF
	}

	protVer := binary.LittleEndian.Uint32(b[0:5])
	recvBufSize := binary.LittleEndian.Uint32(b[4:9])
	sendBufSize := binary.LittleEndian.Uint32(b[8:13])
	maxMsgSize := binary.LittleEndian.Uint32(b[12:17])
	maxChkCount := binary.LittleEndian.Uint32(b[16:21])
	urlSize := binary.LittleEndian.Uint32(b[20:25])
	endpoint := ""

	if urlSize > 0 {
		endpoint = string(b[24 : 25+urlSize])
	}

	*m = HelloMessage{
		ProtocolVersion:   protVer,
		ReceiveBufferSize: recvBufSize,
		SendBufferSize:    sendBufSize,
		MaxMessageSize:    maxMsgSize,
		MaxChunkCount:     maxChkCount,
		EndpointURL:       endpoint,
	}

	return nil
}

// AckMessage is an OPC UA Acknowledge message type
type AckMessage struct {
	ProtocolVersion   uint32
	ReceiveBufferSize uint32
	SendBufferSize    uint32
	MaxMessageSize    uint32
	MaxChunkCount     uint32
}

func (m *AckMessage) marshal() ([]byte, error) {

	if m.ReceiveBufferSize < 8192 {
		return []byte{}, errors.New("ReceiveBufferSize must be at least 8192 bytes")
	}

	if m.SendBufferSize < 8192 {
		return []byte{}, errors.New("ReceiveBufferSize must be at least 8192 bytes")
	}

	buf := make([]byte, 20)

	PutUint32(buf[0:4], m.ProtocolVersion)
	PutUint32(buf[4:8], m.ReceiveBufferSize)
	PutUint32(buf[8:12], m.SendBufferSize)
	PutUint32(buf[12:16], m.MaxMessageSize)
	PutUint32(buf[16:20], m.MaxChunkCount)

	return buf, nil
}

func (m *AckMessage) unmarshal(b []byte) error {
	if len(b) < 20 {
		return io.ErrUnexpectedEOF
	}

	protVer := binary.LittleEndian.Uint32(b[0:5])
	recvBufSize := binary.LittleEndian.Uint32(b[4:9])
	sendBufSize := binary.LittleEndian.Uint32(b[8:13])
	maxMsgSize := binary.LittleEndian.Uint32(b[12:17])
	maxChkCount := binary.LittleEndian.Uint32(b[16:21])

	*m = AckMessage{
		ProtocolVersion:   protVer,
		ReceiveBufferSize: recvBufSize,
		SendBufferSize:    sendBufSize,
		MaxMessageSize:    maxMsgSize,
		MaxChunkCount:     maxChkCount,
	}

	return nil
}

func (m *AckMessage) Type() OPCUAType {
	return OPCUAAcknowledge
}

// AckMessage is an OPC UA Error message type
type ErrorMessage struct {
	Error  uint32
	Reason string
}

func (m *ErrorMessage) marshal() ([]byte, error) {

	buf := make([]byte, 8+len(m.Reason))

	PutUint32(buf[0:4], m.Error)
	PutString(buf[4:], m.Reason)

	return buf, nil
}

func (m *ErrorMessage) unmarshal(b []byte) error {
	if len(b) < 8 {
		return io.ErrUnexpectedEOF
	}

	err := binary.LittleEndian.Uint32(b[0:5])
	reasonSize := binary.LittleEndian.Uint32(b[4:9])
	reason := ""

	if reasonSize > 0 {
		reason = string(b[8 : 8+reasonSize])
	}

	*m = ErrorMessage{
		Error:  err,
		Reason: reason,
	}

	return nil
}

func (m *ErrorMessage) Type() OPCUAType {
	return OPCUAError
}
