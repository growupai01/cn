# Warden 协议 SDK 函数功能列表

本文档列出了 Warden SDK 中所有可用的函数，按协议指令分类。

## 目录

- [设备属性类 (0x50-0x5D)](#设备属性类)
- [状态控制类 (0x80-0x91)](#状态控制类)
- [消息推送类 (0xA0-0xA1)](#消息推送类)
- [工具函数](#工具函数)

---

## 设备属性类

### 0x50: 同步时间 (device_time.go)

| 函数名 | 参数 | 返回值 | 说明 |
|--------|------|--------|------|
| `BuildSyncTimeRequest` | `timeZone byte, utc uint32` | `[]byte` | 构建同步时间请求包 |
| `ParseSyncTimeResponse` | `data []byte` | `*SyncTimeResponse, error` | 解析同步时间响应 |

**时区取值**: 0x00~0x0B(西十二区~西一区), 0x0C(零时区), 0x0D~0x18(东一区~东十二区)

---

### 0x51: 电量获取 (device_battery.go)

| 函数名 | 参数 | 返回值 | 说明 |
|--------|------|--------|------|
| `BuildQueryBatteryRequest` | - | `[]byte` | 构建查询电量请求包 |
| `ParseBatteryResponse` | `data []byte` | `*BatteryInfo, error` | 解析电量查询响应 |
| `ParseBatteryNotify` | `data []byte` | `*BatteryInfo, error` | 解析电量主动上报 |

**BatteryInfo 结构**:
- `Level byte`: 电量值 0-100
- `IsCharging bool`: 是否充电中

---

### 0x52: 屏幕亮度 (device_brightness.go)

| 函数名 | 参数 | 返回值 | 说明 |
|--------|------|--------|------|
| `BuildQueryBrightnessRequest` | - | `[]byte` | 构建查询亮度请求包 |
| `BuildSetBrightnessRequest` | `brightness byte` | `[]byte` | 构建设置亮度请求包 |
| `ParseBrightnessResponse` | `data []byte` | `*BrightnessResponse, error` | 解析亮度响应 |
| `ParseBrightnessNotify` | `data []byte` | `byte, error` | 解析亮度主动上报 |

**亮度值范围**: 0-100

---

### 0x53: 设备语言 (device_language.go)

| 函数名 | 参数 | 返回值 | 说明 |
|--------|------|--------|------|
| `BuildQueryLanguageRequest` | - | `[]byte` | 构建查询语言请求包 |
| `BuildSetLanguageRequest` | `langType byte` | `[]byte` | 构建设置语言请求包 |
| `ParseLanguageResponse` | `data []byte` | `*LanguageResponse, error` | 解析语言响应 |
| `ParseLanguageNotify` | `data []byte` | `byte, error` | 解析语言主动上报 |

**支持的语言**: 见 `constants.go` 中的 `LANG_*` 常量（33种语言）

---

### 0x57: 勿扰功能 (device_dnd.go)

| 函数名 | 参数 | 返回值 | 说明 |
|--------|------|--------|------|
| `BuildQueryDNDRequest` | - | `[]byte` | 构建查询勿扰请求包 |
| `BuildSetDNDRequest` | `settings *DNDSettings` | `[]byte` | 构建设置勿扰请求包 |
| `ParseDNDResponse` | `data []byte` | `*DNDSettings, error` | 解析勿扰响应 |

**DNDSettings 结构**:
- `Enable bool`: 定时勿扰开关
- `StartHour, StartMinute byte`: 开始时间
- `EndHour, EndMinute byte`: 结束时间

---

### 0x59: 寻找手机 (device_find_phone.go)

| 函数名 | 参数 | 返回值 | 说明 |
|--------|------|--------|------|
| `ParseFindPhoneNotify` | `data []byte` | `*FindPhoneNotify, error` | 解析寻找手机通知（设备->手机） |
| `BuildStopFindPhoneRequest` | - | `[]byte` | 构建停止寻找手机请求包 |
| `ParseStopFindPhoneResponse` | `data []byte` | `error` | 解析停止寻找响应 |

---

### 0x5A: 天气单位设置 (device_weather_unit.go)

| 函数名 | 参数 | 返回值 | 说明 |
|--------|------|--------|------|
| `BuildQueryWeatherUnitRequest` | - | `[]byte` | 构建查询天气单位请求包 |
| `BuildSetWeatherUnitRequest` | `unit byte` | `[]byte` | 构建设置天气单位请求包 |
| `ParseWeatherUnitResponse` | `data []byte` | `*WeatherUnitResponse, error` | 解析天气单位响应 |

**单位**: 0x00=摄氏度, 0x01=华氏度

---

### 0x5B: 12h/24h 时间制切换 (device_time_format.go)

| 函数名 | 参数 | 返回值 | 说明 |
|--------|------|--------|------|
| `BuildQueryTimeFormatRequest` | - | `[]byte` | 构建查询时间制请求包 |
| `BuildSetTimeFormatRequest` | `format byte` | `[]byte` | 构建设置时间制请求包 |
| `ParseTimeFormatResponse` | `data []byte` | `*TimeFormatResponse, error` | 解析时间制响应 |

**format**: 0x00=12小时制, 0x01=24小时制

---

### 0x5C: 设备端信息 (device_info.go)

| 函数名 | 参数 | 返回值 | 说明 |
|--------|------|--------|------|
| `BuildQueryDeviceInfoRequest` | `dataType byte` | `[]byte` | 构建查询设备信息请求包 |
| `ParseDeviceInfoResponse` | `data []byte` | `*DeviceInfoData, error` | 解析设备信息响应 |
| `ParseDeviceCapabilitiesNotify` | `data []byte` | `*DeviceCapabilities, error` | 解析设备能力上报 |

**dataType**:
- 0x00: 手表类型
- 0x01: 支持的语言
- 0x02: 序列号
- 0x03: 固件版本

---

### 0x5D: 应用端信息 (app_info.go)

| 函数名 | 参数 | 返回值 | 说明 |
|--------|------|--------|------|
| `BuildSetAppInfoRequest` | `phoneType byte` | `[]byte` | 构建设置应用端信息请求包 |
| `ParseSetAppInfoResponse` | `data []byte` | `error` | 解析设置应用信息响应 |

**phoneType**: 0x00=安卓, 0x01=苹果

---

## 状态控制类

### 0x80: 开关设置 (switch_control.go)

| 函数名 | 参数 | 返回值 | 说明 |
|--------|------|--------|------|
| `BuildQuerySwitchRequest` | - | `[]byte` | 构建查询开关请求包 |
| `BuildSetSwitchRequest` | `switchFlags uint32` | `[]byte` | 构建设置开关请求包 |
| `ParseSwitchResponse` | `data []byte` | `*SwitchResponse, error` | 解析开关响应 |
| `SetSwitchBit` | `flags uint32, switchType byte, enable bool` | `uint32` | 设置开关位（辅助函数） |
| `GetSwitchBit` | `flags uint32, switchType byte` | `bool` | 获取开关位状态（辅助函数） |

**支持的开关**: 见 `constants.go` 中的 `SWITCH_*` 常量

---

### 0x81: 设备绑定 (device_binding.go)

| 函数名 | 参数 | 返回值 | 说明 |
|--------|------|--------|------|
| `BuildStartBindingRequest` | - | `[]byte` | 构建绑定开始请求包 |
| `BuildEndBindingDataRequest` | - | `[]byte` | 构建绑定数据结束请求包 |
| `BuildUnbindRequest` | - | `[]byte` | 构建断开绑定请求包 |
| `ParseBindingResponse` | `data []byte` | `*BindingResponse, error` | 解析绑定响应 |
| `ParseBindingCompleteNotify` | `data []byte` | `*BindingCompleteNotify, error` | 解析绑定完成通知 |

**绑定流程**: 开始绑定 -> 同步数据 -> 绑定数据结束 -> (设备初始化) -> 绑定完成通知

---

### 0x86: 开关表扩展 (switch_extend.go)

| 函数名 | 参数 | 返回值 | 说明 |
|--------|------|--------|------|
| `BuildQueryExtSwitchRequest` | `switchType byte` | `[]byte` | 构建查询扩展开关请求包 |
| `BuildSetExtSwitchRequest` | `switchType byte, flags uint32` | `[]byte` | 构建设置扩展开关请求包 |
| `ParseExtSwitchResponse` | `data []byte` | `*ExtSwitchResponse, error` | 解析扩展开关响应 |

**switchType**: 0x00=社交开关类型（可扩展）

---

### 0x91: 远程拍照 (remote_camera.go)

| 函数名 | 参数 | 返回值 | 说明 |
|--------|------|--------|------|
| `BuildQueryCameraStatusRequest` | - | `[]byte` | 构建查询拍照状态请求包 |
| `BuildSetCameraControlRequest` | `enable bool` | `[]byte` | 构建设置拍照开关请求包 |
| `ParseCameraResponse` | `data []byte` | `*CameraStatusResponse, error` | 解析拍照响应 |
| `ParseCameraNotify` | `data []byte` | `*CameraNotify, error` | 解析拍照通知 |

**拍照动作**:
- 0x00: 远程拍照
- 0x01: 退出APP相机
- 0x02: 呼出APP相机

---

## 消息推送类

### 0xA0: 消息推送 (message_push.go)

| 函数名 | 参数 | 返回值 | 说明 |
|--------|------|--------|------|
| `BuildSetMessageControlRequest` | `ctrl *MessageControl` | `[]byte` | 构建设置消息开关请求包 |
| `BuildPushMessagePackets` | `msgType byte, content []byte, mtu int` | `[][]byte` | 构建推送消息数据包（支持拆包） |
| `ParseMessagePushResponse` | `data []byte` | `error` | 解析消息推送响应 |

**MessageControl 结构**:
- `TotalSwitch bool`: 消息提醒总开关
- `ScreenSwitch bool`: 消息提醒亮屏开关
- `VibrateSwitch bool`: 消息提醒震动开关

**支持的消息类型**: 见 `constants.go` 中的 `MSG_*` 常量（30+种）

---

### 0xA1: 天气信息 (weather_info.go)

| 函数名 | 参数 | 返回值 | 说明 |
|--------|------|--------|------|
| `BuildWeatherTLVPayload` | `data *WeatherData` | `[]byte` | 构建天气TLV数据 |
| `BuildSetWeatherRequest` | `data *WeatherData` | `[]byte` | 构建设置天气请求包 |
| `ParseSetWeatherResponse` | `data []byte` | `error` | 解析设置天气响应 |

**WeatherData 支持的字段**:
- 基础: 天气类型、当前/最低/最高温度、地点
- 扩展: 湿度、风速、紫外线、气压、日出日落、降雨、能见度、空气质量

**日期类型**:
- 0x00: 当天天气
- 0x01: 明天天气
- 0x02: 后天天气

---

## 工具函数

### 协议核心 (protocol.go)

| 函数/类型 | 说明 |
|----------|------|
| `Header` | 协议头结构体 |
| `Packet` | 数据包结构体 |
| `TLV` | TLV数据结构体 |
| `NewHeader()` | 创建新协议头 |
| `DecodeHeader()` | 解码协议头 |
| `DecodePacket()` | 解码数据包 |
| `EncodeTLV()` | 编码TLV数据 |
| `DecodeTLV()` | 解码TLV数据 |
| `EncodeTLVList()` | 编码TLV列表 |
| `DecodeTLVList()` | 解码TLV列表 |
| `EncodeUint16LE()` | 编码16位整数（小端） |
| `DecodeUint16LE()` | 解码16位整数（小端） |
| `EncodeUint32LE()` | 编码32位整数（小端） |
| `DecodeUint32LE()` | 解码32位整数（小端） |
| `EncodeInt8()` | 编码有符号8位整数 |
| `DecodeInt8()` | 解码有符号8位整数 |

### 数据包工具 (packet_utils.go)

| 函数/类型 | 说明 |
|----------|------|
| `PacketBuilder` | 数据包构建器 |
| `NewPacketBuilder()` | 创建新的构建器 |
| `BuildSimpleRequest()` | 构建简单请求包 |
| `BuildSimpleResponse()` | 构建简单响应包 |
| `BuildSimpleNotify()` | 构建简单通知包 |
| `BuildMultiFramePackets()` | 构建多帧数据包（拆包） |
| `PacketMerger` | 数据包合并器 |
| `NewPacketMerger()` | 创建新的合并器 |
| `AddFrame()` | 添加一帧数据 |
| `ValidateResponse()` | 验证响应结果 |

### 常量定义 (constants.go)

包含所有协议相关的常量定义：
- 命令常量 (CMD_*)
- 命令类型常量 (CMD_TYPE_*)
- 返回值常量 (CMD_RESULT_*)
- 语言常量 (LANG_*)
- 消息类型常量 (MSG_*)
- 开关类型常量 (SWITCH_*)
- 天气类型常量 (WEATHER_*)
- BLE UUID 常量

---

## 使用示例

### 示例 1: 同步时间

```go
import "github.com/jiangsu/warden_sdk/warden_sdk"

// 构建同步时间请求
timeZone := warden_sdk.TIMEZONE_EAST_8  // 东八区
utc := uint32(time.Now().Unix())
requestData := warden_sdk.BuildSyncTimeRequest(timeZone, utc)

// 发送 requestData 到蓝牙设备...

// 解析响应
resp, err := warden_sdk.ParseSyncTimeResponse(responseData)
if err != nil {
    // 处理错误
}
```

### 示例 2: 查询电量

```go
// 构建查询请求
requestData := warden_sdk.BuildQueryBatteryRequest()

// 发送到蓝牙设备...

// 解析响应
batteryInfo, err := warden_sdk.ParseBatteryResponse(responseData)
if err != nil {
    // 处理错误
}
fmt.Printf("电量: %d%%, 充电中: %v\n", batteryInfo.Level, batteryInfo.IsCharging)
```

### 示例 3: 推送消息

```go
// 构建消息推送（支持拆包）
msgType := warden_sdk.MSG_WECHAT
content := []byte("您有新的微信消息")
mtu := 244  // 蓝牙MTU值

packets := warden_sdk.BuildPushMessagePackets(msgType, content, mtu)

// 逐个发送数据包
for _, pkt := range packets {
    // 发送 pkt 到蓝牙设备...
}
```

### 示例 4: 设置天气

```go
// 构建天气数据
weather := &warden_sdk.WeatherData{
    DateType:    warden_sdk.WEATHER_DATE_TODAY,
    WeatherType: int8(warden_sdk.WEATHER_SUNSHINE),
    CurrentTemp: 25,
    MinTemp:     18,
    MaxTemp:     28,
    Location:    "北京",
}

requestData := warden_sdk.BuildSetWeatherRequest(weather)

// 发送到蓝牙设备...
```

---

## 注意事项

1. **字节序**: 协议使用小端格式传输多字节字段
2. **拆包**: 当数据超过 MTU 时，使用 `BuildMultiFramePackets` 自动拆包
3. **合包**: 接收端使用 `PacketMerger` 合并多帧数据
4. **错误处理**: 所有 Parse 函数都返回 error，请务必检查
5. **序列号**: `PacketBuilder` 自动管理序列号递增

---

## 版本信息

- **SDK 版本**: 1.0.0
- **协议版本**: 1.0.0
- **Go 版本要求**: 1.16+
