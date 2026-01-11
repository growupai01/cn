package yingka_sdk

// SDK is a wrapper for Android/iOS access
type SDK struct {
}

func NewSDK() *SDK {
	return &SDK{}
}

// Command Creation Methods (Commonly used by APP)

func (s *SDK) CreateGetKeyRequest() []uint8 {
	return NewGetKeyRequest()
}

func (s *SDK) CreateAIWakeupResponse(status int32) []uint8 {
	return NewAIWakeupResponse(uint8(status))
}

func (s *SDK) CreateAIEndCommand(status int32) []uint8 {
	return NewAIEndCommand(uint8(status))
}

func (s *SDK) CreateSimultaneousTranslateRequest(on bool) []uint8 {
	return NewSimultaneousTranslateRequest(on)
}

func (s *SDK) CreateMeetingAssistantRequest(on bool) []uint8 {
	return NewMeetingAssistantRequest(on)
}

func (s *SDK) CreateAIPhotoRecognizeRequest(on bool) []uint8 {
	return NewAIPhotoRecognizeRequest(on)
}

func (s *SDK) CreateSyncFile1Request(mode int32, ssid string) []uint8 {
	return NewSyncFile1Request(uint8(mode), ssid)
}

func (s *SDK) CreateSyncFile2Response(success bool) []uint8 {
	return NewSyncFile2Response(success)
}

func (s *SDK) CreateSyncFile3Request(success bool) []uint8 {
	return NewSyncFile3Request(success)
}

func (s *SDK) CreateSyncTimeRequest(timestamp int64) []uint8 {
	return NewSyncTimeRequest(uint32(timestamp))
}

func (s *SDK) CreateGetBatteryRequest() []uint8 {
	return NewGetBatteryRequest()
}

func (s *SDK) CreateControlMediaRequest(ctrlType int32) []uint8 {
	return NewControlMediaRequest(uint8(ctrlType))
}

func (s *SDK) CreateGetBTStatusRequest() []uint8 {
	return NewGetBTStatusRequest()
}

func (s *SDK) CreateGetConnectedNameRequest() []uint8 {
	return NewGetConnectedNameRequest()
}

func (s *SDK) CreateVoiceRecogSwitchRequest(isSet bool, on bool) []uint8 {
	return NewVoiceRecogSwitchRequest(isSet, on)
}

func (s *SDK) CreateVideoDurationRequest(isSet bool, minutes int32) []uint8 {
	return NewVideoDurationRequest(isSet, uint8(minutes))
}

func (s *SDK) CreateGetDevStatusRequest() []uint8 {
	return NewGetDevStatusRequest()
}

func (s *SDK) CreateDevTriggerPhotoResponse(status int32) []uint8 {
	return NewDevTriggerPhotoResponse(uint8(status))
}

func (s *SDK) CreateGetNewFileCountRequest() []uint8 {
	return NewGetNewFileCountRequest()
}

func (s *SDK) CreateLiveMode1Request(mode int32) []uint8 {
	return NewLiveMode1Request(uint8(mode))
}

func (s *SDK) CreateLiveMode2Response(success bool) []uint8 {
	return NewLiveMode2Response(success)
}

func (s *SDK) CreateLiveMode3Request(success bool) []uint8 {
	return NewLiveMode3Request(success)
}

func (s *SDK) CreateSyncLanguageRequest(lang int32) []uint8 {
	return NewSyncLanguageRequest(uint8(lang))
}

func (s *SDK) CreateGetDevInfoRequest() []uint8 {
	return NewGetDevInfoRequest()
}

// Parser
type ParsedPacket struct {
	Cmd  uint32
	Data []uint8
}

func (s *SDK) Parse(data []uint8) *ParsedPacket {
	p, err := ParsePacket(data)
	if err != nil {
		return nil
	}
	return &ParsedPacket{
		Cmd:  uint32(p.Cmd),
		Data: p.Data,
	}
}
