package yingka_sdk

// SyncFile1Request (0x06)
func NewSyncFile1Request(mode uint8, ssid string) []uint8 {
	nameBytes := []uint8(ssid)
	if len(nameBytes) > 255 {
		nameBytes = nameBytes[:255]
	}
	data := make([]uint8, 2+len(nameBytes))
	data[0] = mode
	data[1] = uint8(len(nameBytes))
	copy(data[2:], nameBytes)
	return NewPacket(CMD_SYNC_FILE_1, data).Serialize()
}

// SyncFile2Response (0x07)
func NewSyncFile2Response(success bool) []uint8 {
	status := STATUS_FAIL
	if success {
		status = STATUS_SUCCESS
	}
	return NewPacket(CMD_SYNC_FILE_2, []uint8{uint8(status)}).Serialize()
}

// SyncFile3Request (0x08)
func NewSyncFile3Request(success bool) []uint8 {
	status := STATUS_FAIL
	if success {
		status = STATUS_SUCCESS
	}
	return NewPacket(CMD_SYNC_FILE_3, []uint8{uint8(status)}).Serialize()
}

// NewFileNotifyRequest (0x0B) - Fetch new file count
func NewGetNewFileCountRequest() []uint8 {
	return NewPacket(CMD_NEW_FILE_NOTIFY, nil).Serialize()
}

// LiveMode1Request (0x10)
func NewLiveMode1Request(mode uint8) []uint8 {
	return NewPacket(CMD_LIVE_MODE_1, []uint8{mode}).Serialize()
}

// LiveMode2Response (0x11)
func NewLiveMode2Response(success bool) []uint8 {
	status := STATUS_FAIL
	if success {
		status = STATUS_SUCCESS
	}
	return NewPacket(CMD_LIVE_MODE_2, []uint8{uint8(status)}).Serialize()
}

// LiveMode3Request (0x12)
func NewLiveMode3Request(success bool) []uint8 {
	status := STATUS_FAIL
	if success {
		status = STATUS_SUCCESS
	}
	return NewPacket(CMD_LIVE_MODE_3, []uint8{uint8(status)}).Serialize()
}
