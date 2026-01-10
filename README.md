# Warden 协议 Go SDK

Warden 协议 Go SDK 是基于 Warden 蓝牙协议文档实现的完整通信库，用于 APP 端与蓝牙智能设备之间的数据交互。

## 特性

- ✅ **完整协议支持**: 实现了所有 Warden 协议指令（0x50-0xA1）
- ✅ **类型安全**: 使用 Go 强类型系统，避免协议错误
- ✅ **自动拆包/合包**: 支持 MTU 受限情况下的数据分片传输
- ✅ **易于集成**: 可编译为 Android AAR 库，直接在 Android 项目中使用
- ✅ **文档完善**: 提供详细的 API 文档和使用示例

## 功能模块

### 设备属性类 (0x50-0x5D)
- 时间同步
- 电量管理
- 屏幕亮度控制
- 语言设置
- 勿扰模式
- 寻找手机
- 天气单位设置
- 时间制切换（12h/24h）
- 设备信息查询
- 应用信息设置

### 状态控制类 (0x80-0x91)
- 开关设置（14种开关类型）
- 设备绑定流程
- 开关表扩展
- 远程拍照控制

### 消息推送类 (0xA0-0xA1)
- 消息推送（支持30+种消息类型）
- 天气信息推送

## 快速开始

### 安装

```bash
go get github.com/jiangsu/warden_sdk/warden_sdk
```

### 基本使用

```go
package main

import (
    "fmt"
    "time"
    "github.com/jiangsu/warden_sdk/warden_sdk"
)

func main() {
    // 1. 同步时间
    timeZone := warden_sdk.TIMEZONE_EAST_8
    utc := uint32(time.Now().Unix())
    timeRequest := warden_sdk.BuildSyncTimeRequest(timeZone, utc)
    // 发送 timeRequest 到蓝牙设备...
    
    // 2. 查询电量
    batteryRequest := warden_sdk.BuildQueryBatteryRequest()
    // 发送 batteryRequest 到蓝牙设备...
    
    // 3. 设置亮度
    brightnessRequest := warden_sdk.BuildSetBrightnessRequest(80) // 80%亮度
    // 发送 brightnessRequest 到蓝牙设备...
    
    // 4. 推送微信消息
    msgPackets := warden_sdk.BuildPushMessagePackets(
        warden_sdk.MSG_WECHAT,
        []byte("您有新的微信消息"),
        244, // MTU
    )
    // 逐个发送 msgPackets...
}
```

### 解析响应

```go
// 解析电量响应
batteryInfo, err := warden_sdk.ParseBatteryResponse(responseData)
if err != nil {
    fmt.Printf("解析失败: %v\n", err)
    return
}
fmt.Printf("电量: %d%%, 充电中: %v\n", 
    batteryInfo.Level, batteryInfo.IsCharging)
```

## Android 集成

### 编译为 Android AAR

```bash
# 确保已安装 gomobile
go install golang.org/x/mobile/cmd/gomobile@latest
gomobile init

# 编译 AAR
cd e:\go\jiangsu_sdk
gomobile bind -target=android -androidapi=31 -o warden_sdk.aar ./warden_sdk
```

### 在 Android 项目中使用

1. 将生成的 `warden_sdk.aar` 放入 Android 项目的 `libs` 目录

2. 在 `build.gradle` 中添加依赖：

```gradle
dependencies {
    implementation files('libs/warden_sdk.aar')
}
```

3. 在 Kotlin/Java 代码中使用：

```kotlin
import warden_sdk.*

// 同步时间
val timeRequest = Warden_sdk.buildSyncTimeRequest(
    Warden_sdk.TIMEZONE_EAST_8.toByte(),
    System.currentTimeMillis() / 1000
)

// 查询电量
val batteryRequest = Warden_sdk.buildQueryBatteryRequest()

// 发送到蓝牙设备...
```

## 项目结构

```
warden_sdk/
├── constants.go           # 协议常量定义
├── protocol.go            # 核心协议结构
├── packet_utils.go        # 数据包工具
├── errors.go             # 错误定义
├── device_time.go        # 时间同步
├── device_battery.go     # 电量管理
├── device_brightness.go  # 亮度控制
├── device_language.go    # 语言设置
├── device_dnd.go         # 勿扰模式
├── device_find_phone.go  # 寻找手机
├── device_weather_unit.go # 天气单位
├── device_time_format.go # 时间制
├── device_info.go        # 设备信息
├── app_info.go           # 应用信息
├── switch_control.go     # 开关控制
├── device_binding.go     # 设备绑定
├── switch_extend.go      # 开关扩展
├── remote_camera.go      # 远程拍照
├── message_push.go       # 消息推送
└── weather_info.go       # 天气信息
```

## API 文档

详细的 API 文档请查看 [function_list.md](function_list.md)

## 协议规范

本 SDK 基于 Warden 协议文档 v1.0.0 实现，支持：

- **蓝牙 UUID**:
  - Service: `6e400001-b5a3-f393-e0a9-e50e24dcca9d`
  - Write: `6e400002-b5a3-f393-e0a9-e50e24dcca9d`
  - Notify: `6e400003-b5a3-f393-e0a9-e50e24dcca9d`

- **数据包格式**: Header (5 bytes) + Payload (0-N bytes)
- **字节序**: 小端格式（Little Endian）
- **MTU**: 默认 244 bytes
- **拆包/合包**: 自动处理

## 常见问题

### 1. 如何处理多帧数据？

```go
// 发送端：自动拆包
packets := warden_sdk.BuildPushMessagePackets(msgType, longContent, mtu)
for _, pkt := range packets {
    // 发送每一帧
}

// 接收端：使用 PacketMerger 合并
merger := warden_sdk.NewPacketMerger()
for _, frameData := range receivedFrames {
    packet, _ := warden_sdk.DecodePacket(frameData)
    complete, merged, err := merger.AddFrame(packet)
    if complete {
        // 所有帧已接收，处理 merged 数据
    }
}
```

### 2. 如何设置多个开关？

```go
var flags uint32 = 0

// 设置防丢开关
flags = warden_sdk.SetSwitchBit(flags, warden_sdk.SWITCH_ANTI_LOST, true)

// 设置抬手亮屏
flags = warden_sdk.SetSwitchBit(flags, warden_sdk.SWITCH_RAISE_WRIST, true)

// 发送设置请求
request := warden_sdk.BuildSetSwitchRequest(flags)
```

### 3. 支持哪些消息类型？

SDK 支持 30+ 种消息类型，包括：
- 来电、未接来电、短信
- 微信、QQ、WhatsApp
- Facebook、Instagram、Twitter
- 钉钉、支付宝、抖音
- 等等...

完整列表见 `constants.go` 中的 `MSG_*` 常量。

## 开发与测试

### 运行测试

```bash
cd warden_sdk
go test -v
```

### 编译检查

```bash
go build ./warden_sdk
```

## 许可证

本项目基于 MIT 许可证开源。

## 技术支持

如有问题，请查阅：
1. [API 文档](function_list.md)
2. [Warden 协议文档](Warden协议文档.md)

---

**版本**: 1.0.0  
**更新日期**: 2026-01-10
