package warden_sdk

// SwitchControl 开关设置功能 (0x80)

// BuildQuerySwitchRequest 构建查询开关请求包
func BuildQuerySwitchRequest() []byte {
	pb := NewPacketBuilder()
	payload := []byte{ACTION_QUERY} // 0x00: 查询开关
	return pb.BuildSimpleRequest(CMD_SWITCH_CONTROL, payload)
}

// BuildSetSwitchRequest 构建设置开关请求包
func BuildSetSwitchRequest(switchFlags uint32) []byte {
	pb := NewPacketBuilder()

	payload := make([]byte, 5)
	payload[0] = ACTION_SET // 0x01: 设置开关

	// 开关设置（小端）
	switchBytes := EncodeUint32LE(switchFlags)
	copy(payload[1:5], switchBytes)

	return pb.BuildSimpleRequest(CMD_SWITCH_CONTROL, payload)
}

// SwitchResponse 开关响应
type SwitchResponse struct {
	ActionCmd   byte
	SwitchFlags uint32 // 开关位图（见表7-4-2）
}

// ParseSwitchResponse 解析开关响应
func ParseSwitchResponse(data []byte) (*SwitchResponse, error) {
	packet, err := DecodePacket(data)
	if err != nil {
		return nil, err
	}

	if packet.Payload[0] == ACTION_QUERY {
		// 查询响应
		if len(packet.Payload) < 5 {
			return nil, ErrInvalidPayloadLength
		}
		return &SwitchResponse{
			ActionCmd:   packet.Payload[0],
			SwitchFlags: DecodeUint32LE(packet.Payload[1:5]),
		}, nil
	} else {
		// 设置响应
		if len(packet.Payload) < 2 {
			return nil, ErrInvalidPayloadLength
		}
		return nil, ValidateResponse(packet.Payload[1])
	}
}

// Helper functions for switch operations

// SetSwitchBit 设置开关位
func SetSwitchBit(flags uint32, switchType byte, enable bool) uint32 {
	if enable {
		return flags | (1 << switchType)
	}
	return flags &^ (1 << switchType)
}

// GetSwitchBit 获取开关位状态
func GetSwitchBit(flags uint32, switchType byte) bool {
	return (flags & (1 << switchType)) != 0
}
