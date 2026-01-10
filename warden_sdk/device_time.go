package warden_sdk

// DeviceTime 时间同步功能 (0x50)

// SyncTimeRequest 同步时间请求
type SyncTimeRequest struct {
	TimeZone byte   // 时区
	UTC      uint32 // UTC时间戳
}

// BuildSyncTimeRequest 构建同步时间请求包
func BuildSyncTimeRequest(timeZone byte, utc uint32) []byte {
	pb := NewPacketBuilder()

	payload := make([]byte, 6)
	payload[0] = ACTION_SET // Action: 设置时间
	payload[1] = timeZone   // 时区

	// UTC时间戳（小端）
	utcBytes := EncodeUint32LE(utc)
	copy(payload[2:6], utcBytes)

	return pb.BuildSimpleRequest(CMD_SYNC_TIME, payload)
}

// SyncTimeResponse 同步时间响应
type SyncTimeResponse struct {
	ActionCmd byte
	Result    byte
}

// ParseSyncTimeResponse 解析同步时间响应
func ParseSyncTimeResponse(data []byte) (*SyncTimeResponse, error) {
	packet, err := DecodePacket(data)
	if err != nil {
		return nil, err
	}

	if len(packet.Payload) < 2 {
		return nil, ErrInvalidPayloadLength
	}

	resp := &SyncTimeResponse{
		ActionCmd: packet.Payload[0],
		Result:    packet.Payload[1],
	}

	return resp, ValidateResponse(resp.Result)
}
