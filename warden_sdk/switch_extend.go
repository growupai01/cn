package warden_sdk

// SwitchExtend 开关表扩展功能 (0x86)

// 扩展开关类型
const (
	EXT_SWITCH_TYPE_SOCIAL byte = 0x00 // 社交开关类型
)

// BuildQueryExtSwitchRequest 构建查询扩展开关请求包
func BuildQueryExtSwitchRequest(switchType byte) []byte {
	pb := NewPacketBuilder()
	payload := []byte{
		ACTION_QUERY, // 0x00: 查询开关
		switchType,   // 开关类型
	}
	return pb.BuildSimpleRequest(CMD_SWITCH_EXTEND, payload)
}

// BuildSetExtSwitchRequest 构建设置扩展开关请求包
func BuildSetExtSwitchRequest(switchType byte, flags uint32) []byte {
	pb := NewPacketBuilder()

	payload := make([]byte, 6)
	payload[0] = ACTION_SET // 0x01: 设置开关
	payload[1] = switchType // 开关类型

	// 开关设置（小端）
	flagsBytes := EncodeUint32LE(flags)
	copy(payload[2:6], flagsBytes)

	return pb.BuildSimpleRequest(CMD_SWITCH_EXTEND, payload)
}

// ExtSwitchResponse 扩展开关响应
type ExtSwitchResponse struct {
	ActionCmd  byte
	SwitchType byte
	Flags      uint32
}

// ParseExtSwitchResponse 解析扩展开关响应
func ParseExtSwitchResponse(data []byte) (*ExtSwitchResponse, error) {
	packet, err := DecodePacket(data)
	if err != nil {
		return nil, err
	}

	if packet.Payload[0] == ACTION_QUERY {
		// 查询响应
		if len(packet.Payload) < 6 {
			return nil, ErrInvalidPayloadLength
		}
		return &ExtSwitchResponse{
			ActionCmd:  packet.Payload[0],
			SwitchType: packet.Payload[1],
			Flags:      DecodeUint32LE(packet.Payload[2:6]),
		}, nil
	} else {
		// 设置响应
		if len(packet.Payload) < 2 {
			return nil, ErrInvalidPayloadLength
		}
		return nil, ValidateResponse(packet.Payload[1])
	}
}
