package warden_sdk

// AppInfo 应用端信息功能 (0x5D)

// BuildSetAppInfoRequest 构建设置应用端信息请求包
func BuildSetAppInfoRequest(phoneType byte) []byte {
	pb := NewPacketBuilder()
	payload := []byte{
		ACTION_SET, // 0x01: 设置
		0x00,       // Data type: 手机类型
		phoneType,  // 0x00: 安卓; 0x01: 苹果
	}
	return pb.BuildSimpleRequest(CMD_APP_INFO, payload)
}

// ParseSetAppInfoResponse 解析设置应用信息响应
func ParseSetAppInfoResponse(data []byte) error {
	packet, err := DecodePacket(data)
	if err != nil {
		return err
	}

	if len(packet.Payload) < 2 {
		return ErrInvalidPayloadLength
	}

	return ValidateResponse(packet.Payload[1])
}
