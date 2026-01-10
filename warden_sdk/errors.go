package warden_sdk

import "errors"

// 通用错误定义
var (
	ErrInvalidPayloadLength = errors.New("invalid payload length")
	ErrInvalidHeader        = errors.New("invalid header")
	ErrInvalidTLV           = errors.New("invalid TLV data")
	ErrPacketTooShort       = errors.New("packet too short")
	ErrInvalidCmd           = errors.New("invalid command")
	ErrInvalidCmdType       = errors.New("invalid command type")
)
