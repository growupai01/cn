package warden_sdk

// DeviceWeatherUnit 天气单位设置功能 (0x5A)

// BuildQueryWeatherUnitRequest 构建查询天气单位请求包
func BuildQueryWeatherUnitRequest() []byte {
	pb := NewPacketBuilder()
	payload := []byte{ACTION_QUERY} // 0x00: 查询
	return pb.BuildSimpleRequest(CMD_WEATHER_UNIT, payload)
}

// BuildSetWeatherUnitRequest 构建设置天气单位请求包
func BuildSetWeatherUnitRequest(unit byte) []byte {
	pb := NewPacketBuilder()
	payload := []byte{
		ACTION_SET, // 0x01: 设置
		unit,       // 0x00: 摄氏度; 0x01: 华氏度
	}
	return pb.BuildSimpleRequest(CMD_WEATHER_UNIT, payload)
}

// WeatherUnitResponse 天气单位响应
type WeatherUnitResponse struct {
	ActionCmd byte
	Unit      byte // 0x00: 摄氏度; 0x01: 华氏度
}

// ParseWeatherUnitResponse 解析天气单位响应
func ParseWeatherUnitResponse(data []byte) (*WeatherUnitResponse, error) {
	packet, err := DecodePacket(data)
	if err != nil {
		return nil, err
	}

	if packet.Payload[0] == ACTION_QUERY {
		// 查询响应
		if len(packet.Payload) < 2 {
			return nil, ErrInvalidPayloadLength
		}
		return &WeatherUnitResponse{
			ActionCmd: packet.Payload[0],
			Unit:      packet.Payload[1],
		}, nil
	} else {
		// 设置响应
		if len(packet.Payload) < 2 {
			return nil, ErrInvalidPayloadLength
		}
		return nil, ValidateResponse(packet.Payload[1])
	}
}
