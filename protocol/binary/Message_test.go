package binary

import (
	"reflect"
	"testing"
)

func TestHelloMessage_marshal(t *testing.T) {
	tests := []struct {
		name    string
		msg     HelloMessage
		want    []byte
		wantErr bool
	}{
		{"Default", HelloMessage{}, []byte{}, true},
		{"Simple Hello", HelloMessage{
			ProtocolVersion:   0,
			SendBufferSize:    8192,
			ReceiveBufferSize: 8192,
		}, []byte{0, 0, 0, 0, 0, 32, 0, 0, 0, 32, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, false},
		{"Full Hello", HelloMessage{
			ProtocolVersion:   0,
			SendBufferSize:    8192,
			ReceiveBufferSize: 8192,
			MaxMessageSize:    50,
			MaxChunkCount:     3,
			EndpointURL:       "opc.tcp://localhost",
		}, []byte{0, 0, 0, 0, 0, 32, 0, 0, 0, 32, 0, 0, 50, 0, 0, 0, 3, 0, 0, 0, 19, 0, 0, 0, 111, 112, 99, 46, 116, 99, 112, 58, 47, 47, 108, 111, 99, 97, 108, 104, 111, 115, 116}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.msg.marshal()
			if (err != nil) != tt.wantErr {
				t.Errorf("HelloMessage.marshal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HelloMessage.marshal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAckMessage_marshal(t *testing.T) {
	tests := []struct {
		name    string
		msg     AckMessage
		want    []byte
		wantErr bool
	}{
		{"Default", AckMessage{}, []byte{}, true},
		{"Simple Ack", AckMessage{
			ProtocolVersion:   0,
			SendBufferSize:    8192,
			ReceiveBufferSize: 8192,
		}, []byte{0, 0, 0, 0, 0, 32, 0, 0, 0, 32, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, false},
		{"Full Ack", AckMessage{
			ProtocolVersion:   0,
			SendBufferSize:    8192,
			ReceiveBufferSize: 8192,
			MaxMessageSize:    50,
			MaxChunkCount:     3,
		}, []byte{0, 0, 0, 0, 0, 32, 0, 0, 0, 32, 0, 0, 50, 0, 0, 0, 3, 0, 0, 0}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.msg.marshal()
			if (err != nil) != tt.wantErr {
				t.Errorf("HelloMessage.marshal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HelloMessage.marshal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestErrorMessage_marshal(t *testing.T) {
	tests := []struct {
		name    string
		msg     ErrorMessage
		want    []byte
		wantErr bool
	}{
		{"Default", ErrorMessage{}, []byte{0, 0, 0, 0, 0, 0, 0, 0}, false},
		{"Simple Error", ErrorMessage{
			Error: 50,
		}, []byte{50, 0, 0, 0, 0, 0, 0, 0}, false},
		{"Full Error", ErrorMessage{
			Error:  99,
			Reason: "Some Reasons",
		}, []byte{99, 0, 0, 0, 12, 0, 0, 0, 83, 111, 109, 101, 32, 82, 101, 97, 115, 111, 110, 115}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.msg.marshal()
			if (err != nil) != tt.wantErr {
				t.Errorf("HelloMessage.marshal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HelloMessage.marshal() = %v, want %v", got, tt.want)
			}
		})
	}
}
