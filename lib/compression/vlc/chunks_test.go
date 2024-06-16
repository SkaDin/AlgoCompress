package vlc

import (
	"reflect"
	"testing"
)

func TestBinaryChunks_Join(t *testing.T) {
	tests := []struct {
		name string
		bcs  BinaryChunks
		want string
	}{
		{
			name: "base test",
			bcs:  BinaryChunks{"10100000", "10000000"},
			want: "1010000010000000",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.bcs.Join(); got != tt.want {
				t.Errorf("Join() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewBinChunks(t *testing.T) {

	tests := []struct {
		name string
		data []byte
		want BinaryChunks
	}{
		{
			name: "base test",
			data: []byte{20, 30, 60, 18},
			want: BinaryChunks{"00010100", "00011110", "00111100", "00010010"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBinChunks(tt.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBinChunks() = %v, want %v", got, tt.want)
			}
		})
	}
}
