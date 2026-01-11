package yingka_sdk

import (
	"encoding/binary"
	"errors"
)

// Packet represents a protocol packet
type Packet struct {
	Header uint8
	Cmd    uint8
	Data   []uint8
}

// NewPacket creates a new packet with the fixed header
func NewPacket(cmd uint8, data []uint8) *Packet {
	return &Packet{
		Header: FIXED_HEADER,
		Cmd:    cmd,
		Data:   data,
	}
}

// Serialize converts the packet to a byte slice
func (p *Packet) Serialize() []uint8 {
	res := make([]uint8, 2+len(p.Data))
	res[0] = p.Header
	res[1] = p.Cmd
	copy(res[2:], p.Data)
	return res
}

// ParsePacket parses a byte slice into a Packet
func ParsePacket(data []uint8) (*Packet, error) {
	if len(data) < 2 {
		return nil, errors.New("packet too short")
	}
	if data[0] != FIXED_HEADER {
		return nil, errors.New("invalid header")
	}
	return &Packet{
		Header: data[0],
		Cmd:    data[1],
		Data:   data[2:],
	}, nil
}

// Utility functions for big endian
func Uint32ToBytes(v uint32) []uint8 {
	b := make([]uint8, 4)
	binary.BigEndian.PutUint32(b, v)
	return b
}

func BytesToUint32(b []uint8) uint32 {
	if len(b) < 4 {
		return 0
	}
	return binary.BigEndian.Uint32(b[:4])
}

func Uint16ToBytes(v uint16) []uint8 {
	b := make([]uint8, 2)
	binary.BigEndian.PutUint16(b, v)
	return b
}

func BytesToUint16(b []uint8) uint16 {
	if len(b) < 2 {
		return 0
	}
	return binary.BigEndian.Uint16(b[:2])
}
