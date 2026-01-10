package warden_sdk

// WeatherInfo 天气信息功能 (0xA1)

// 天气日期类型
const (
	WEATHER_DATE_TODAY     byte = 0x00 // 当天天气
	WEATHER_DATE_TOMORROW  byte = 0x01 // 明天天气
	WEATHER_DATE_AFTER_TOM byte = 0x02 // 后天天气
)

// WeatherData 天气数据
type WeatherData struct {
	DateType      byte   // 天气日期类型
	WeatherType   int8   // 天气类型
	CurrentTemp   int8   // 当前温度（℃）
	MinTemp       int8   // 最低温度（℃）
	MaxTemp       int8   // 最高温度（℃）
	Location      string // 地点位置（UTF-8）
	Humidity      byte   // 湿度（%）
	WindSpeed     byte   // 风速（km/h）
	UVIndex       byte   // 紫外线指数
	Pressure      uint32 // 气压（pa）
	SunriseHour   byte   // 日出时间：小时
	SunriseMinute byte   // 日出时间：分钟
	SunsetHour    byte   // 日落时间：小时
	SunsetMinute  byte   // 日落时间：分钟
	RainProb      byte   // 降雨概率（%）
	Rainfall      uint16 // 降雨量（0.1mm）
	Visibility    uint32 // 能见度（m）
	AirQuality    uint16 // 空气质量指数
}

// BuildWeatherTLVPayload 构建天气TLV数据
func BuildWeatherTLVPayload(data *WeatherData) []byte {
	var tlvList []*TLV

	// 添加天气类型
	if data.WeatherType != 0 {
		tlvList = append(tlvList, &TLV{
			Type:   WEATHER_CMD_TYPE,
			Length: 1,
			Value:  []byte{byte(data.WeatherType)},
		})
	}

	// 添加当前温度
	if data.CurrentTemp != 0 {
		tlvList = append(tlvList, &TLV{
			Type:   WEATHER_CMD_CURRENT_TEMP,
			Length: 1,
			Value:  []byte{EncodeInt8(data.CurrentTemp)},
		})
	}

	// 添加最低温度
	if data.MinTemp != 0 {
		tlvList = append(tlvList, &TLV{
			Type:   WEATHER_CMD_MIN_TEMP,
			Length: 1,
			Value:  []byte{EncodeInt8(data.MinTemp)},
		})
	}

	// 添加最高温度
	if data.MaxTemp != 0 {
		tlvList = append(tlvList, &TLV{
			Type:   WEATHER_CMD_MAX_TEMP,
			Length: 1,
			Value:  []byte{EncodeInt8(data.MaxTemp)},
		})
	}

	// 添加地点位置（UTF-8编码）
	if data.Location != "" {
		locationBytes := []byte(data.Location)
		tlvList = append(tlvList, &TLV{
			Type:   WEATHER_CMD_LOCATION,
			Length: byte(len(locationBytes)),
			Value:  locationBytes,
		})
	}

	// 添加湿度
	if data.Humidity != 0 {
		tlvList = append(tlvList, &TLV{
			Type:   WEATHER_CMD_HUMIDITY,
			Length: 1,
			Value:  []byte{data.Humidity},
		})
	}

	// 添加风速
	if data.WindSpeed != 0 {
		tlvList = append(tlvList, &TLV{
			Type:   WEATHER_CMD_WIND_SPEED,
			Length: 1,
			Value:  []byte{data.WindSpeed},
		})
	}

	// 添加紫外线指数
	if data.UVIndex != 0 {
		tlvList = append(tlvList, &TLV{
			Type:   WEATHER_CMD_UV_INDEX,
			Length: 1,
			Value:  []byte{data.UVIndex},
		})
	}

	// 添加气压
	if data.Pressure != 0 {
		tlvList = append(tlvList, &TLV{
			Type:   WEATHER_CMD_PRESSURE,
			Length: 4,
			Value:  EncodeUint32LE(data.Pressure),
		})
	}

	// 添加日出时间
	if data.SunriseHour != 0 || data.SunriseMinute != 0 {
		tlvList = append(tlvList, &TLV{
			Type:   WEATHER_CMD_SUNRISE,
			Length: 2,
			Value:  []byte{data.SunriseHour, data.SunriseMinute},
		})
	}

	// 添加日落时间
	if data.SunsetHour != 0 || data.SunsetMinute != 0 {
		tlvList = append(tlvList, &TLV{
			Type:   WEATHER_CMD_SUNSET,
			Length: 2,
			Value:  []byte{data.SunsetHour, data.SunsetMinute},
		})
	}

	// 添加降雨概率
	if data.RainProb != 0 {
		tlvList = append(tlvList, &TLV{
			Type:   WEATHER_CMD_RAIN_PROB,
			Length: 1,
			Value:  []byte{data.RainProb},
		})
	}

	// 添加降雨量
	if data.Rainfall != 0 {
		tlvList = append(tlvList, &TLV{
			Type:   WEATHER_CMD_RAINFALL,
			Length: 2,
			Value:  EncodeUint16LE(data.Rainfall),
		})
	}

	// 添加能见度
	if data.Visibility != 0 {
		tlvList = append(tlvList, &TLV{
			Type:   WEATHER_CMD_VISIBILITY,
			Length: 4,
			Value:  EncodeUint32LE(data.Visibility),
		})
	}

	// 添加空气质量指数
	if data.AirQuality != 0 {
		tlvList = append(tlvList, &TLV{
			Type:   WEATHER_CMD_AIR_QUALITY,
			Length: 2,
			Value:  EncodeUint16LE(data.AirQuality),
		})
	}

	return EncodeTLVList(tlvList)
}

// BuildSetWeatherRequest 构建设置天气请求包
func BuildSetWeatherRequest(data *WeatherData) []byte {
	pb := NewPacketBuilder()

	// 构建TLV数据
	tlvPayload := BuildWeatherTLVPayload(data)

	// 组装完整payload
	payload := make([]byte, 2+len(tlvPayload))
	payload[0] = ACTION_SET    // 0x01: 设置
	payload[1] = data.DateType // 日期类型
	copy(payload[2:], tlvPayload)

	return pb.BuildSimpleRequest(CMD_WEATHER_INFO, payload)
}

// ParseSetWeatherResponse 解析设置天气响应
func ParseSetWeatherResponse(data []byte) error {
	packet, err := DecodePacket(data)
	if err != nil {
		return err
	}

	if len(packet.Payload) < 2 {
		return ErrInvalidPayloadLength
	}

	return ValidateResponse(packet.Payload[1])
}
