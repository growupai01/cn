package warden_sdk

// DeviceBinding 设备绑定功能 (0x81)

// 绑定控制类型
const (
	BIND_CONTROL_START    byte = 0x00 // 绑定开始
	BIND_CONTROL_DATA_END byte = 0x01 // 绑定数据结束
	BIND_CONTROL_UNBIND   byte = 0x02 // 断开绑定
	BIND_CONTROL_QR_START byte = 0x03 // 二维码绑定开始(预留)
	BIND_CONTROL_QR_END   byte = 0x04 // 二维码绑定结束(预留)
)

// BuildStartBindingRequest 构建绑定开始请求包
func BuildStartBindingRequest() []byte {
	pb := NewPacketBuilder()
	payload := []byte{
		ACTION_SET,         // 0x01: 设置
		BIND_CONTROL_START, // 0x00: 绑定开始
	}
	return pb.BuildSimpleRequest(CMD_DEVICE_BINDING, payload)
}

// BuildEndBindingDataRequest 构建绑定数据结束请求包
func BuildEndBindingDataRequest() []byte {
	pb := NewPacketBuilder()
	payload := []byte{
		ACTION_SET,            // 0x01: 设置
		BIND_CONTROL_DATA_END, // 0x01: 绑定数据结束
	}
	return pb.BuildSimpleRequest(CMD_DEVICE_BINDING, payload)
}

// BuildUnbindRequest 构建断开绑定请求包
func BuildUnbindRequest() []byte {
	pb := NewPacketBuilder()
	payload := []byte{
		ACTION_SET,          // 0x01: 设置
		BIND_CONTROL_UNBIND, // 0x02: 断开绑定
	}
	return pb.BuildSimpleRequest(CMD_DEVICE_BINDING, payload)
}

// BindingResponse 绑定响应
type BindingResponse struct {
	ActionCmd    byte
	Control      byte
	BindFirst    bool // 当Control=0x00时有效：0x00:未被绑定过 0x01:已被绑定过
	BindComplete bool // 当Control=0x01时有效：绑定是否完成
}

// ParseBindingResponse 解析绑定响应
func ParseBindingResponse(data []byte) (*BindingResponse, error) {
	packet, err := DecodePacket(data)
	if err != nil {
		return nil, err
	}

	if len(packet.Payload) < 2 {
		return nil, ErrInvalidPayloadLength
	}

	resp := &BindingResponse{
		ActionCmd: packet.Payload[0],
		Control:   packet.Payload[1],
	}

	if len(packet.Payload) >= 3 {
		switch resp.Control {
		case BIND_CONTROL_START:
			// 绑定开始响应
			resp.BindFirst = packet.Payload[2] == 0x00

		case BIND_CONTROL_DATA_END:
			// 绑定数据结束响应
			resp.BindComplete = (packet.Payload[2] & 0x01) != 0

		case BIND_CONTROL_UNBIND:
			// 断开绑定响应
			return resp, ValidateResponse(packet.Payload[2])
		}
	}

	return resp, nil
}

// BindingCompleteNotify 绑定完成通知
type BindingCompleteNotify struct {
	State byte // 0x00: 绑定完成
}

// ParseBindingCompleteNotify 解析绑定完成通知（Cmd=0x82）
func ParseBindingCompleteNotify(data []byte) (*BindingCompleteNotify, error) {
	packet, err := DecodePacket(data)
	if err != nil {
		return nil, err
	}

	if len(packet.Payload) < 1 {
		return nil, ErrInvalidPayloadLength
	}

	return &BindingCompleteNotify{
		State: packet.Payload[0],
	}, nil
}
