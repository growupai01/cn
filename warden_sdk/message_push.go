package warden_sdk

// MessagePush 消息推送功能 (0xA0)

// MessageControl 消息控制
type MessageControl struct {
	TotalSwitch   bool // 消息提醒总开关
	ScreenSwitch  bool // 消息提醒亮屏开关
	VibrateSwitch bool // 消息提醒震动开关
}

// BuildSetMessageControlRequest 构建设置消息开关请求包
func BuildSetMessageControlRequest(ctrl *MessageControl) []byte {
	pb := NewPacketBuilder()

	controlByte := byte(0)
	if ctrl.TotalSwitch {
		controlByte |= 0x01
	}
	if ctrl.ScreenSwitch {
		controlByte |= 0x02
	}
	if ctrl.VibrateSwitch {
		controlByte |= 0x04
	}

	payload := []byte{
		0x00,        // Action cmd: 0x00 设置
		controlByte, // 控制字节
	}

	return pb.BuildSimpleRequest(CMD_MESSAGE_PUSH, payload)
}

// BuildPushMessagePackets 构建推送消息数据包（支持拆包）
func BuildPushMessagePackets(msgType byte, content []byte, mtu int) [][]byte {
	pb := NewPacketBuilder()

	// 构建完整payload
	payload := make([]byte, 2+len(content))
	payload[0] = 0x01    // Action cmd: 0x01 推送内容
	payload[1] = msgType // 消息类型
	copy(payload[2:], content)

	return pb.BuildMultiFramePackets(CMD_MESSAGE_PUSH, CMD_TYPE_REQUEST, payload, mtu)
}

// ParseMessagePushResponse 解析消息推送响应
func ParseMessagePushResponse(data []byte) error {
	packet, err := DecodePacket(data)
	if err != nil {
		return err
	}

	if len(packet.Payload) < 2 {
		return ErrInvalidPayloadLength
	}

	return ValidateResponse(packet.Payload[1])
}
