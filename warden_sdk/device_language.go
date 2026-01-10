package warden_sdk

// DeviceLanguage 设备语言功能 (0x53)

// BuildQueryLanguageRequest 构建查询语言请求包
func BuildQueryLanguageRequest() []byte {
	pb := NewPacketBuilder()
	payload := []byte{ACTION_QUERY} // 0x00: 查询设备语言
	return pb.BuildSimpleRequest(CMD_LANGUAGE, payload)
}

// BuildSetLanguageRequest 构建设置语言请求包
func BuildSetLanguageRequest(langType byte) []byte {
	pb := NewPacketBuilder()
	payload := []byte{
		ACTION_SET, // 0x01: 设置设备语言
		langType,   // 语言类型
	}
	return pb.BuildSimpleRequest(CMD_LANGUAGE, payload)
}

// LanguageResponse 语言响应
type LanguageResponse struct {
	ActionCmd    byte
	LanguageType byte
}

// ParseLanguageResponse 解析语言响应
func ParseLanguageResponse(data []byte) (*LanguageResponse, error) {
	packet, err := DecodePacket(data)
	if err != nil {
		return nil, err
	}

	if packet.Payload[0] == ACTION_QUERY {
		// 查询响应
		if len(packet.Payload) < 2 {
			return nil, ErrInvalidPayloadLength
		}
		return &LanguageResponse{
			ActionCmd:    packet.Payload[0],
			LanguageType: packet.Payload[1],
		}, nil
	} else {
		// 设置响应
		if len(packet.Payload) < 2 {
			return nil, ErrInvalidPayloadLength
		}
		return &LanguageResponse{
			ActionCmd: packet.Payload[0],
		}, ValidateResponse(packet.Payload[1])
	}
}

// ParseLanguageNotify 解析语言主动上报
func ParseLanguageNotify(data []byte) (byte, error) {
	packet, err := DecodePacket(data)
	if err != nil {
		return 0, err
	}

	if len(packet.Payload) < 1 {
		return 0, ErrInvalidPayloadLength
	}

	return packet.Payload[0], nil
}
