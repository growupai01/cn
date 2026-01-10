package warden_sdk

// DeviceBrightness 屏幕亮度功能 (0x52)

// BuildQueryBrightnessRequest 构建查询亮度请求包
func BuildQueryBrightnessRequest() []byte {
	pb := NewPacketBuilder()
	payload := []byte{ACTION_QUERY} // 0x00: 查询屏幕亮度值
	return pb.BuildSimpleRequest(CMD_BRIGHTNESS, payload)
}

// BuildSetBrightnessRequest 构建设置亮度请求包
func BuildSetBrightnessRequest(brightness byte) []byte {
	pb := NewPacketBuilder()
	payload := []byte{
		ACTION_SET, // 0x01: 设置屏幕亮度值
		brightness, // 亮度值 0-100
	}
	return pb.BuildSimpleRequest(CMD_BRIGHTNESS, payload)
}

// BrightnessResponse 亮度响应
type BrightnessResponse struct {
	ActionCmd  byte
	Brightness byte // 0-100
}

// ParseBrightnessResponse 解析亮度响应
func ParseBrightnessResponse(data []byte) (*BrightnessResponse, error) {
	packet, err := DecodePacket(data)
	if err != nil {
		return nil, err
	}

	if packet.Payload[0] == ACTION_QUERY {
		// 查询响应
		if len(packet.Payload) < 2 {
			return nil, ErrInvalidPayloadLength
		}
		return &BrightnessResponse{
			ActionCmd:  packet.Payload[0],
			Brightness: packet.Payload[1],
		}, nil
	} else {
		// 设置响应
		if len(packet.Payload) < 2 {
			return nil, ErrInvalidPayloadLength
		}
		return &BrightnessResponse{
			ActionCmd:  packet.Payload[0],
			Brightness: 0, // 设置响应不返回亮度值
		}, ValidateResponse(packet.Payload[1])
	}
}

// ParseBrightnessNotify 解析亮度主动上报
func ParseBrightnessNotify(data []byte) (byte, error) {
	packet, err := DecodePacket(data)
	if err != nil {
		return 0, err
	}

	if len(packet.Payload) < 1 {
		return 0, ErrInvalidPayloadLength
	}

	return packet.Payload[0], nil
}
