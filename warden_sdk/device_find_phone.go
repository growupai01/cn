package warden_sdk

// DeviceFindPhone 寻找手机功能 (0x59)

// FindPhoneAction 寻找手机动作
const (
	FIND_PHONE_START byte = 0x00 // 寻找手机
	FIND_PHONE_STOP  byte = 0x01 // 停止寻找
)

// FindPhoneNotify 寻找手机通知
type FindPhoneNotify struct {
	Action byte // 0x00: 寻找手机; 0x01: 停止寻找
}

// ParseFindPhoneNotify 解析寻找手机通知（设备 -> 手机）
func ParseFindPhoneNotify(data []byte) (*FindPhoneNotify, error) {
	packet, err := DecodePacket(data)
	if err != nil {
		return nil, err
	}

	if len(packet.Payload) < 1 {
		return nil, ErrInvalidPayloadLength
	}

	return &FindPhoneNotify{
		Action: packet.Payload[0],
	}, nil
}

// BuildStopFindPhoneRequest 构建停止寻找手机请求包（手机 -> 设备）
func BuildStopFindPhoneRequest() []byte {
	pb := NewPacketBuilder()
	payload := []byte{FIND_PHONE_STOP} // 0x01: 停止寻找
	return pb.BuildSimpleRequest(CMD_FIND_PHONE, payload)
}

// ParseStopFindPhoneResponse 解析停止寻找响应
func ParseStopFindPhoneResponse(data []byte) error {
	packet, err := DecodePacket(data)
	if err != nil {
		return err
	}

	if len(packet.Payload) < 2 {
		return ErrInvalidPayloadLength
	}

	return ValidateResponse(packet.Payload[1])
}
