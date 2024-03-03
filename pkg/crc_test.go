package pkg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestXor(t *testing.T) {
	tests := []struct {
		name   string
		a      string
		b      string
		result string
	}{
		{
			name:   "xor1",
			a:      "1010",
			b:      "1010",
			result: "000",
		}, {
			name:   "xor2",
			a:      "10101010",
			b:      "01010101",
			result: "1111111",
		}, {
			name:   "xor3",
			a:      "10101010",
			b:      "10101011",
			result: "0000001",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := xor(test.a, test.b)
			assert.Equal(t, test.result, got)
		})
	}
}

func TestCRC(t *testing.T) {
	tests := []struct {
		name   string
		data   string
		key    string
		result string
	}{
		{
			name:   "crc1",
			data:   "1001",
			key:    "1011",
			result: "1001110",
		}, {
			name:   "crc2",
			data:   "1101011011",
			key:    "10011",
			result: "11010110111110",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := CRC(test.data, test.key)
			assert.Equal(t, test.result, got)
		})
	}
}

func TestCheckSum(t *testing.T) {
	tests := []struct {
		name   string
		frame  string
		key    string
		result string
	}{
		{
			name:   "checksum1",
			frame:  "1001110",
			key:    "1011",
			result: "000",
		}, {
			name:   "checksum2",
			frame:  "11010110111110",
			key:    "10011",
			result: "0000",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := CheckSum(test.frame, test.key)
			assert.Equal(t, test.result, got)
		})
	}
}
