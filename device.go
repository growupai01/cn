package yingka_sdk

// GetKeyRequest (0x01)
func NewGetKeyRequest() []uint8 {
	return NewPacket(CMD_GET_KEY, nil).Serialize()
}

// GetBTStatusRequest (0x09)
func NewGetBTStatusRequest() []uint8 {
	return NewPacket(CMD_GET_BT_STATUS, nil).Serialize()
}

// GetConnectedNameRequest (0x0A)
func NewGetConnectedNameRequest() []uint8 {
	return NewPacket(CMD_GET_CONNECTED_NAME, nil).Serialize()
}

// SyncTimeRequest (0x0F)
func NewSyncTimeRequest(timestamp uint32) []uint8 {
	return NewPacket(CMD_SYNC_TIME, Uint32ToBytes(timestamp)).Serialize()
}

// VoiceRecogSwitchRequest (0x14) - Get/Set
func NewVoiceRecogSwitchRequest(isSet bool, on bool) []uint8 {
	data := make([]uint8, 2)
	if isSet {
		data[0] = 0x02
		if on {
			data[1] = 0x01
		} else {
			data[1] = 0x00
		}
	} else {
		data[0] = 0x01
	}
	return NewPacket(CMD_VOICE_RECOG_SWITCH, data).Serialize()
}

// GetBatteryRequest (0x15)
func NewGetBatteryRequest() []uint8 {
	return NewPacket(CMD_SYNC_BATTERY, nil).Serialize()
}

// VideoDurationRequest (0x16) - Get/Set
func NewVideoDurationRequest(isSet bool, minutes uint8) []uint8 {
	data := make([]uint8, 2)
	if isSet {
		data[0] = 0x02
		data[1] = minutes
	} else {
		data[0] = 0x01
	}
	return NewPacket(CMD_VIDEO_DURATION, data).Serialize()
}

// GetDevStatusRequest (0x17)
func NewGetDevStatusRequest() []uint8 {
	return NewPacket(CMD_GET_DEV_STATUS, nil).Serialize()
}

// ControlMediaRequest (0x18)
func NewControlMediaRequest(ctrlType uint8) []uint8 {
	return NewPacket(CMD_CONTROL_MEDIA, []uint8{ctrlType}).Serialize()
}

// SyncLanguageRequest (0x1A)
func NewSyncLanguageRequest(lang uint8) []uint8 {
	return NewPacket(CMD_SYNC_LANGUAGE, []uint8{lang}).Serialize()
}

// GetDevInfoRequest (0x1B)
func NewGetDevInfoRequest() []uint8 {
	return NewPacket(CMD_GET_DEV_INFO, nil).Serialize()
}
