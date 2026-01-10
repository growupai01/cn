package warden_sdk

// DeviceTimeFormat 12h/24h时间制切换功能 (0x5B)

// BuildQueryTimeFormatRequest 构建查询时间制请求包
func BuildQueryTimeFormatRequest() []byte {
	pb := NewPacketBuilder()
	payload := []byte{ACTION_QUERY} // 0x00: 查询
	return pb.BuildSimpleRequest(CMD_TIME_FORMAT, payload)
}

// BuildSetTimeFormatRequest 构建设置时间制请求包
func BuildSetTimeFormatRequest(format byte) []byte {
	pb := NewPacketBuilder()
	payload := []byte{
		ACTION_SET, // 0x01: 设置
		format,     // 0x00: 12h; 0x01: 24h
	}
	return pb.BuildSimpleRequest(CMD_TIME_FORMAT, payload)
}

// TimeFormatResponse 时间制响应
type TimeFormatResponse struct {
	ActionCmd byte
	Format    byte // 0x00: 12h; 0x01: 24h
}

// ParseTimeFormatResponse 解析时间制响应
func ParseTimeFormatResponse(data []byte) (*TimeFormatResponse, error) {
	packet, err := DecodePacket(data)
	if err != nil {
		return nil, err
	}

	if packet.Payload[0] == ACTION_QUERY {
		// 查询响应
		if len(packet.Payload) < 2 {
			return nil, ErrInvalidPayloadLength
		}
		return &TimeFormatResponse{
			ActionCmd: packet.Payload[0],
			Format:    packet.Payload[1],
		}, nil
	} else {
		// 设置响应
		if len(packet.Payload) < 2 {
			return nil, ErrInvalidPayloadLength
		}
		return nil, ValidateResponse(packet.Payload[1])
	}
}
