package warden_sdk

// DeviceBattery 电量功能 (0x51)

// BatteryInfo 电量信息
type BatteryInfo struct {
	Level      byte // 电量值 0-100
	IsCharging bool // 是否充电中
}

// BuildQueryBatteryRequest 构建查询电量请求包
func BuildQueryBatteryRequest() []byte {
	pb := NewPacketBuilder()
	payload := []byte{ACTION_QUERY} // 0x00: 查询电量值
	return pb.BuildSimpleRequest(CMD_BATTERY, payload)
}

// ParseBatteryResponse 解析电量查询响应
func ParseBatteryResponse(data []byte) (*BatteryInfo, error) {
	packet, err := DecodePacket(data)
	if err != nil {
		return nil, err
	}

	if len(packet.Payload) < 2 {
		return nil, ErrInvalidPayloadLength
	}

	responseValue := packet.Payload[1]

	return &BatteryInfo{
		Level:      responseValue & 0x7F,        // bit6~bit0: 电量值
		IsCharging: (responseValue & 0x80) != 0, // bit7: 充电状态
	}, nil
}

// ParseBatteryNotify 解析电量主动上报
func ParseBatteryNotify(data []byte) (*BatteryInfo, error) {
	packet, err := DecodePacket(data)
	if err != nil {
		return nil, err
	}

	if len(packet.Payload) < 1 {
		return nil, ErrInvalidPayloadLength
	}

	responseValue := packet.Payload[0]

	return &BatteryInfo{
		Level:      responseValue & 0x7F,
		IsCharging: (responseValue & 0x80) != 0,
	}, nil
}
