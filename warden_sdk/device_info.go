package warden_sdk

// DeviceInfo 设备端信息功能 (0x5C)

// 设备信息数据类型
const (
	DEVICE_INFO_TYPE_WATCH_TYPE    byte = 0x00 // 手表类型
	DEVICE_INFO_TYPE_SUPPORT_LANG  byte = 0x01 // 支持的设备语言
	DEVICE_INFO_TYPE_SERIAL_NUMBER byte = 0x02 // 序列号
	DEVICE_INFO_TYPE_FIRMWARE      byte = 0x03 // 固件版本信息
)

// DeviceInfoData 设备信息数据
type DeviceInfoData struct {
	DataType        byte
	WatchType       byte   // 手表类型（当DataType=0x00）
	SupportLanguage uint32 // 支持的语言位图（当DataType=0x01）
	SerialNumber    []byte // 序列号（当DataType=0x02）
	FirmwareMajor   byte   // 固件主版本号（当DataType=0x03）
	FirmwareMinor   byte   // 固件从版本号（当DataType=0x03）
}

// DeviceCapabilities 设备功能信息（Notify上报）
type DeviceCapabilities struct {
	WatchType       byte   // 手表类型
	SupportLanguage uint32 // 支持的语言
	FunctionControl uint32 // 功能控制标志
	HealthControl   uint32 // 健康功能标志
	MsgControl      uint32 // 消息控制标志
	SwitchControl   uint32 // 开关控制标志
	SportControl    uint32 // 运动控制标志
}

// BuildQueryDeviceInfoRequest 构建查询设备信息请求包
func BuildQueryDeviceInfoRequest(dataType byte) []byte {
	pb := NewPacketBuilder()
	payload := []byte{
		ACTION_QUERY, // 0x00: 查询
		dataType,     // 数据类型
	}
	return pb.BuildSimpleRequest(CMD_DEVICE_INFO, payload)
}

// ParseDeviceInfoResponse 解析设备信息响应
func ParseDeviceInfoResponse(data []byte) (*DeviceInfoData, error) {
	packet, err := DecodePacket(data)
	if err != nil {
		return nil, err
	}

	if len(packet.Payload) < 2 {
		return nil, ErrInvalidPayloadLength
	}

	info := &DeviceInfoData{
		DataType: packet.Payload[1],
	}

	switch info.DataType {
	case DEVICE_INFO_TYPE_WATCH_TYPE:
		if len(packet.Payload) < 3 {
			return nil, ErrInvalidPayloadLength
		}
		info.WatchType = packet.Payload[2] & 0x0F

	case DEVICE_INFO_TYPE_SUPPORT_LANG:
		if len(packet.Payload) < 6 {
			return nil, ErrInvalidPayloadLength
		}
		info.SupportLanguage = DecodeUint32LE(packet.Payload[2:6])

	case DEVICE_INFO_TYPE_SERIAL_NUMBER:
		if len(packet.Payload) < 3 {
			return nil, ErrInvalidPayloadLength
		}
		info.SerialNumber = packet.Payload[2:]

	case DEVICE_INFO_TYPE_FIRMWARE:
		if len(packet.Payload) < 4 {
			return nil, ErrInvalidPayloadLength
		}
		info.FirmwareMajor = packet.Payload[2]
		info.FirmwareMinor = packet.Payload[3]
	}

	return info, nil
}

// ParseDeviceCapabilitiesNotify 解析设备能力上报
func ParseDeviceCapabilitiesNotify(data []byte) (*DeviceCapabilities, error) {
	packet, err := DecodePacket(data)
	if err != nil {
		return nil, err
	}

	if len(packet.Payload) < 30 {
		return nil, ErrInvalidPayloadLength
	}

	return &DeviceCapabilities{
		WatchType:       packet.Payload[0] & 0x0F,
		SupportLanguage: DecodeUint32LE(packet.Payload[1:5]),
		FunctionControl: DecodeUint32LE(packet.Payload[5:9]),
		HealthControl:   DecodeUint32LE(packet.Payload[9:13]),
		MsgControl:      DecodeUint32LE(packet.Payload[13:17]),
		SwitchControl:   DecodeUint32LE(packet.Payload[17:21]),
		SportControl:    DecodeUint32LE(packet.Payload[21:25]),
	}, nil
}
