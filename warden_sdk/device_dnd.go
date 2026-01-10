package warden_sdk

// DeviceDND 勿扰功能 (0x57)

// DNDSettings 勿扰设置
type DNDSettings struct {
	Enable      bool // 定时勿扰开关
	StartHour   byte // 开始时间：小时
	StartMinute byte // 开始时间：分钟
	EndHour     byte // 结束时间：小时
	EndMinute   byte // 结束时间：分钟
}

// BuildQueryDNDRequest 构建查询勿扰请求包
func BuildQueryDNDRequest() []byte {
	pb := NewPacketBuilder()
	payload := []byte{ACTION_QUERY} // 0x00: 查询
	return pb.BuildSimpleRequest(CMD_DND, payload)
}

// BuildSetDNDRequest 构建设置勿扰请求包
func BuildSetDNDRequest(settings *DNDSettings) []byte {
	pb := NewPacketBuilder()

	switchByte := byte(0)
	if settings.Enable {
		switchByte = 0x01
	}

	payload := []byte{
		ACTION_SET,           // 0x01: 设置
		switchByte,           // [0]: 定时勿扰开关
		settings.StartHour,   // 开始时间：小时
		settings.StartMinute, // 开始时间：分钟
		settings.EndHour,     // 结束时间：小时
		settings.EndMinute,   // 结束时间：分钟
	}

	return pb.BuildSimpleRequest(CMD_DND, payload)
}

// ParseDNDResponse 解析勿扰响应
func ParseDNDResponse(data []byte) (*DNDSettings, error) {
	packet, err := DecodePacket(data)
	if err != nil {
		return nil, err
	}

	if packet.Payload[0] == ACTION_QUERY {
		// 查询响应
		if len(packet.Payload) < 6 {
			return nil, ErrInvalidPayloadLength
		}

		return &DNDSettings{
			Enable:      (packet.Payload[1] & 0x01) != 0,
			StartHour:   packet.Payload[2],
			StartMinute: packet.Payload[3],
			EndHour:     packet.Payload[4],
			EndMinute:   packet.Payload[5],
		}, nil
	} else {
		// 设置响应
		if len(packet.Payload) < 2 {
			return nil, ErrInvalidPayloadLength
		}
		return nil, ValidateResponse(packet.Payload[1])
	}
}
