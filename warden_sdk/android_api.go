package warden_sdk

// Android 友好的接口包装
// gomobile 要求函数返回值最多为：一个值 + 一个 error

// WardenSDK Android SDK 主接口
type WardenSDK struct {
	builder *PacketBuilder
	merger  *PacketMerger
}

// NewWardenSDK 创建新的 SDK 实例
func NewWardenSDK() *WardenSDK {
	return &WardenSDK{
		builder: NewPacketBuilder(),
		merger:  NewPacketMerger(),
	}
}

// === 时间同步 ===

// SyncTime 同步时间
func (sdk *WardenSDK) SyncTime(timeZone byte, utc int64) []byte {
	return BuildSyncTimeRequest(timeZone, uint32(utc))
}

// === 电量管理 ===

// QueryBattery 查询电量
func (sdk *WardenSDK) QueryBattery() []byte {
	return BuildQueryBatteryRequest()
}

// GetBatteryLevel 从响应中获取电量值
func GetBatteryLevel(responseData []byte) (int, error) {
	info, err := ParseBatteryResponse(responseData)
	if err != nil {
		return 0, err
	}
	return int(info.Level), nil
}

// IsBatteryCharging 从响应中获取充电状态
func IsBatteryCharging(responseData []byte) (bool, error) {
	info, err := ParseBatteryResponse(responseData)
	if err != nil {
		return false, err
	}
	return info.IsCharging, nil
}

// === 屏幕亮度 ===

// QueryBrightness 查询屏幕亮度
func (sdk *WardenSDK) QueryBrightness() []byte {
	return BuildQueryBrightnessRequest()
}

// SetBrightness 设置屏幕亮度 (0-100)
func (sdk *WardenSDK) SetBrightness(value int) []byte {
	return BuildSetBrightnessRequest(byte(value))
}

// === 设备语言 ===

// QueryLanguage 查询设备语言
func (sdk *WardenSDK) QueryLanguage() []byte {
	return BuildQueryLanguageRequest()
}

// SetLanguage 设置设备语言
func (sdk *WardenSDK) SetLanguage(langType int) []byte {
	return BuildSetLanguageRequest(byte(langType))
}

// === 勿扰模式 ===

// QueryDND 查询勿扰设置
func (sdk *WardenSDK) QueryDND() []byte {
	return BuildQueryDNDRequest()
}

// SetDND 设置勿扰模式
func (sdk *WardenSDK) SetDND(enable bool, startHour, startMin, endHour, endMin int) []byte {
	settings := &DNDSettings{
		Enable:      enable,
		StartHour:   byte(startHour),
		StartMinute: byte(startMin),
		EndHour:     byte(endHour),
		EndMinute:   byte(endMin),
	}
	return BuildSetDNDRequest(settings)
}

// === 寻找手机 ===

// StopFindPhone 停止寻找手机
func (sdk *WardenSDK) StopFindPhone() []byte {
	return BuildStopFindPhoneRequest()
}

// === 天气单位 ===

// QueryWeatherUnit 查询天气单位
func (sdk *WardenSDK) QueryWeatherUnit() []byte {
	return BuildQueryWeatherUnitRequest()
}

// SetWeatherUnit 设置天气单位 (0:摄氏度 1:华氏度)
func (sdk *WardenSDK) SetWeatherUnit(unit int) []byte {
	return BuildSetWeatherUnitRequest(byte(unit))
}

// === 时间制 ===

// QueryTimeFormat 查询时间制
func (sdk *WardenSDK) QueryTimeFormat() []byte {
	return BuildQueryTimeFormatRequest()
}

// SetTimeFormat 设置时间制 (0:12h 1:24h)
func (sdk *WardenSDK) SetTimeFormat(format int) []byte {
	return BuildSetTimeFormatRequest(byte(format))
}

// === 设备信息 ===

// QueryDeviceInfo 查询设备信息
// dataType: 0=手表类型 1=支持语言 2=序列号 3=固件版本
func (sdk *WardenSDK) QueryDeviceInfo(dataType int) []byte {
	return BuildQueryDeviceInfoRequest(byte(dataType))
}

// === 应用信息 ===

// SetAppInfo 设置应用信息 (0:安卓 1:苹果)
func (sdk *WardenSDK) SetAppInfo(phoneType int) []byte {
	return BuildSetAppInfoRequest(byte(phoneType))
}

// === 开关设置 ===

// QuerySwitch 查询开关状态
func (sdk *WardenSDK) QuerySwitch() []byte {
	return BuildQuerySwitchRequest()
}

// SetSwitch 设置开关
func (sdk *WardenSDK) SetSwitch(flags int64) []byte {
	return BuildSetSwitchRequest(uint32(flags))
}

// === 设备绑定 ===

// StartBinding 开始绑定
func (sdk *WardenSDK) StartBinding() []byte {
	return BuildStartBindingRequest()
}

// EndBindingData 绑定数据结束
func (sdk *WardenSDK) EndBindingData() []byte {
	return BuildEndBindingDataRequest()
}

// Unbind 解除绑定
func (sdk *WardenSDK) Unbind() []byte {
	return BuildUnbindRequest()
}

// === 扩展开关 ===

// QueryExtSwitch 查询扩展开关
func (sdk *WardenSDK) QueryExtSwitch(switchType int) []byte {
	return BuildQueryExtSwitchRequest(byte(switchType))
}

// SetExtSwitch 设置扩展开关
func (sdk *WardenSDK) SetExtSwitch(switchType int, flags int64) []byte {
	return BuildSetExtSwitchRequest(byte(switchType), uint32(flags))
}

// === 远程拍照 ===

// QueryCameraStatus 查询拍照状态
func (sdk *WardenSDK) QueryCameraStatus() []byte {
	return BuildQueryCameraStatusRequest()
}

// SetCameraControl 设置拍照开关
func (sdk *WardenSDK) SetCameraControl(enable bool) []byte {
	return BuildSetCameraControlRequest(enable)
}

// === 消息推送 ===

// SetMessageControl 设置消息开关
func (sdk *WardenSDK) SetMessageControl(total, screen, vibrate bool) []byte {
	ctrl := &MessageControl{
		TotalSwitch:   total,
		ScreenSwitch:  screen,
		VibrateSwitch: vibrate,
	}
	return BuildSetMessageControlRequest(ctrl)
}

// PushMessage 推送消息（单包）
// 如果消息过长需要拆包，请使用 PushMessageWithMTU
func (sdk *WardenSDK) PushMessage(msgType int, content string) []byte {
	packets := BuildPushMessagePackets(byte(msgType), []byte(content), DEFAULT_MTU)
	if len(packets) > 0 {
		return packets[0]
	}
	return nil
}

// === 天气信息 ===

// SetSimpleWeather 设置简单天气信息
func (sdk *WardenSDK) SetSimpleWeather(dateType int, weatherType int, currentTemp int, minTemp int, maxTemp int, location string) []byte {
	weather := &WeatherData{
		DateType:    byte(dateType),
		WeatherType: int8(weatherType),
		CurrentTemp: int8(currentTemp),
		MinTemp:     int8(minTemp),
		MaxTemp:     int8(maxTemp),
		Location:    location,
	}
	return BuildSetWeatherRequest(weather)
}

// SetDetailedWeather 设置详细天气信息
func (sdk *WardenSDK) SetDetailedWeather(
	dateType int,
	weatherType int,
	currentTemp int,
	minTemp int,
	maxTemp int,
	location string,
	humidity int,
	windSpeed int,
	uvIndex int,
) []byte {
	weather := &WeatherData{
		DateType:    byte(dateType),
		WeatherType: int8(weatherType),
		CurrentTemp: int8(currentTemp),
		MinTemp:     int8(minTemp),
		MaxTemp:     int8(maxTemp),
		Location:    location,
		Humidity:    byte(humidity),
		WindSpeed:   byte(windSpeed),
		UVIndex:     byte(uvIndex),
	}
	return BuildSetWeatherRequest(weather)
}

// === 工具函数 ===

// ValidateResponseCode 验证响应码
func ValidateResponseCode(code int) error {
	return ValidateResponse(byte(code))
}

// DecodePacketData 解码数据包（通用）
func DecodePacketData(data []byte) (*Packet, error) {
	return DecodePacket(data)
}
