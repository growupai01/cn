package yingka_sdk

// AIWakeupRequest (0x02) - APP Response to Dev or Stop
func NewAIWakeupResponse(status uint8) []uint8 {
	return NewPacket(CMD_AI_WAKEUP, []uint8{status}).Serialize()
}

// AIEndCommand (0x02 - Status 0x05/0x06)
func NewAIEndCommand(status uint8) []uint8 {
	return NewPacket(CMD_AI_WAKEUP, []uint8{status}).Serialize()
}

// SimultaneousTranslateRequest (0x03)
func NewSimultaneousTranslateRequest(on bool) []uint8 {
	status := STATUS_STOP
	if on {
		status = STATUS_START
	}
	return NewPacket(CMD_SIMULTANEOUS_TRANSLATE, []uint8{uint8(status)}).Serialize()
}

// MeetingAssistantRequest (0x04)
func NewMeetingAssistantRequest(on bool) []uint8 {
	status := STATUS_STOP
	if on {
		status = STATUS_START
	}
	return NewPacket(CMD_MEETING_ASSISTANT, []uint8{uint8(status)}).Serialize()
}

// AIPhotoRecognizeRequest (0x05)
func NewAIPhotoRecognizeRequest(on bool) []uint8 {
	status := STATUS_STOP
	if on {
		status = STATUS_START
	}
	return NewPacket(CMD_AI_PHOTO_RECOGNIZE, []uint8{uint8(status)}).Serialize()
}

// DevTriggerPhotoResponse (0x13 - Status 0x01/0x02)
func NewDevTriggerPhotoResponse(status uint8) []uint8 {
	return NewPacket(CMD_DEV_TRIGGER_PHOTO, []uint8{status}).Serialize()
}
