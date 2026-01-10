package warden_sdk

import (
	"encoding/binary"
	"fmt"
)

// Header 协议头结构
type Header struct {
	SeqNum      byte // [0,3]: 序列号; [7]: 加密标志
	Cmd         byte // 命令字节
	CmdType     byte // 命令类型
	FrameSeq    byte // [0,3]: 帧序号; [4,7]: 总帧数
	FrameLength byte // 帧长度
}

// Packet 数据包结构
type Packet struct {
	Header  Header
	Payload []byte
}

// TLV TLV格式数据结构
type TLV struct {
	Type   byte
	Length byte
	Value  []byte
}

// NewHeader 创建新的协议头
func NewHeader(seqNum, cmd, cmdType, frameSeq, frameLength byte) *Header {
	return &Header{
		SeqNum:      seqNum & 0x0F, // 只保留低4位作为序列号
		Cmd:         cmd,
		CmdType:     cmdType,
		FrameSeq:    frameSeq,
		FrameLength: frameLength,
	}
}

// SetEncrypted 设置加密标志
func (h *Header) SetEncrypted(encrypted bool) {
	if encrypted {
		h.SeqNum |= 0x80
	} else {
		h.SeqNum &= 0x7F
	}
}

// IsEncrypted 检查是否加密
func (h *Header) IsEncrypted() bool {
	return (h.SeqNum & 0x80) != 0
}

// GetSeqNum 获取序列号
func (h *Header) GetSeqNum() byte {
	return h.SeqNum & 0x0F
}

// GetFrameSeq 获取帧序号
func (h *Header) GetFrameSeq() byte {
	return h.FrameSeq & 0x0F
}

// GetTotalFrames 获取总帧数
func (h *Header) GetTotalFrames() byte {
	return ((h.FrameSeq >> 4) & 0x0F) + 1
}

// SetFrameInfo 设置帧信息
func (h *Header) SetFrameInfo(frameSeq, totalFrames byte) {
	if totalFrames > 0 {
		totalFrames-- // 实际值减1存储
	}
	h.FrameSeq = (frameSeq & 0x0F) | ((totalFrames & 0x0F) << 4)
}

// Encode 编码协议头为字节数组
func (h *Header) Encode() []byte {
	return []byte{
		h.SeqNum,
		h.Cmd,
		h.CmdType,
		h.FrameSeq,
		h.FrameLength,
	}
}

// DecodeHeader 从字节数组解码协议头
func DecodeHeader(data []byte) (*Header, error) {
	if len(data) < HEADER_SIZE {
		return nil, fmt.Errorf("data too short for header: need %d bytes, got %d", HEADER_SIZE, len(data))
	}

	return &Header{
		SeqNum:      data[0],
		Cmd:         data[1],
		CmdType:     data[2],
		FrameSeq:    data[3],
		FrameLength: data[4],
	}, nil
}

// NewPacket 创建新的数据包
func NewPacket(header *Header, payload []byte) *Packet {
	return &Packet{
		Header:  *header,
		Payload: payload,
	}
}

// Encode 编码整个数据包
func (p *Packet) Encode() []byte {
	header := p.Header.Encode()
	result := make([]byte, len(header)+len(p.Payload))
	copy(result, header)
	copy(result[len(header):], p.Payload)
	return result
}

// DecodePacket 解码数据包
func DecodePacket(data []byte) (*Packet, error) {
	header, err := DecodeHeader(data)
	if err != nil {
		return nil, err
	}

	if len(data) < HEADER_SIZE+int(header.FrameLength) {
		return nil, fmt.Errorf("data too short for payload: need %d bytes, got %d",
			HEADER_SIZE+int(header.FrameLength), len(data))
	}

	payload := data[HEADER_SIZE : HEADER_SIZE+int(header.FrameLength)]
	return NewPacket(header, payload), nil
}

// EncodeTLV 编码TLV数据
func EncodeTLV(tlvType byte, value []byte) []byte {
	length := byte(len(value))
	result := make([]byte, 2+len(value))
	result[0] = tlvType
	result[1] = length
	copy(result[2:], value)
	return result
}

// decodeTLV 解码单个TLV (内部函数，不导出到 gomobile)
func decodeTLV(data []byte) (*TLV, int, error) {
	if len(data) < 2 {
		return nil, 0, fmt.Errorf("data too short for TLV header")
	}

	tlvType := data[0]
	length := data[1]

	if len(data) < 2+int(length) {
		return nil, 0, fmt.Errorf("data too short for TLV value: need %d bytes, got %d",
			2+int(length), len(data))
	}

	value := make([]byte, length)
	copy(value, data[2:2+int(length)])

	return &TLV{
		Type:   tlvType,
		Length: length,
		Value:  value,
	}, 2 + int(length), nil
}

// DecodeTLVList 解码TLV列表
func DecodeTLVList(data []byte) ([]*TLV, error) {
	var tlvList []*TLV
	offset := 0

	for offset < len(data) {
		tlv, consumed, err := decodeTLV(data[offset:])
		if err != nil {
			return nil, fmt.Errorf("failed to decode TLV at offset %d: %v", offset, err)
		}
		tlvList = append(tlvList, tlv)
		offset += consumed
	}

	return tlvList, nil
}

// EncodeTLVList 编码TLV列表
func EncodeTLVList(tlvList []*TLV) []byte {
	var result []byte
	for _, tlv := range tlvList {
		result = append(result, EncodeTLV(tlv.Type, tlv.Value)...)
	}
	return result
}

// EncodeUint16LE 编码16位无符号整数（小端）
func EncodeUint16LE(value uint16) []byte {
	buf := make([]byte, 2)
	binary.LittleEndian.PutUint16(buf, value)
	return buf
}

// DecodeUint16LE 解码16位无符号整数（小端）
func DecodeUint16LE(data []byte) uint16 {
	return binary.LittleEndian.Uint16(data)
}

// EncodeUint32LE 编码32位无符号整数（小端）
func EncodeUint32LE(value uint32) []byte {
	buf := make([]byte, 4)
	binary.LittleEndian.PutUint32(buf, value)
	return buf
}

// DecodeUint32LE 解码32位无符号整数（小端）
func DecodeUint32LE(data []byte) uint32 {
	return binary.LittleEndian.Uint32(data)
}

// EncodeInt8 编码有符号8位整数（温度等）
func EncodeInt8(value int8) byte {
	return byte(value)
}

// DecodeInt8 解码有符号8位整数
func DecodeInt8(data byte) int8 {
	return int8(data)
}
