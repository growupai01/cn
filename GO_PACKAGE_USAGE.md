# 方式 2: 直接作为 Go 包使用

Warden SDK 可以直接在任何 Go 项目中作为包使用，无需编译为 Android/iOS SDK。

## 📦 安装

### 方式 A: 本地引用（当前项目）

```go
import "github.com/jiangsu/warden_sdk/warden_sdk"
```

在 `go.mod` 中添加：
```
replace github.com/jiangsu/warden_sdk => ./
```

### 方式 B: 从 GitHub 安装（公开后）

```bash
go get github.com/你的用户名/warden_sdk/warden_sdk
```

## 🚀 基本使用

### 1. 创建 SDK 实例

```go
package main

import (
    "fmt"
    "time"
    "github.com/jiangsu/warden_sdk/warden_sdk"
)

func main() {
    // 创建 SDK 实例
    sdk := warden_sdk.NewWardenSDK()
    
    // 使用 SDK 功能...
}
```

### 2. 同步时间

```go
// 构建时间同步请求
timeData := sdk.SyncTime(
    warden_sdk.TIMEZONE_EAST_8,  // 东八区
    time.Now().Unix(),             // UTC 时间戳
)

// 发送 timeData 到蓝牙设备
// bluetoothDevice.Write(timeData)
```

### 3. 查询电量

```go
// 构建查询请求
requestData := sdk.QueryBattery()

// 发送到蓝牙设备
// bluetoothDevice.Write(requestData)

// 接收响应后解析
// responseData := bluetoothDevice.Read()
// level, err := warden_sdk.GetBatteryLevel(responseData)
// if err == nil {
//     fmt.Printf("电量: %d%%\n", level)
// }
```

### 4. 设置亮度

```go
// 设置屏幕亮度为 80%
brightnessData := sdk.SetBrightness(80)

// 发送到蓝牙设备
// bluetoothDevice.Write(brightnessData)
```

### 5. 推送消息

```go
// 推送微信消息
msgData := sdk.PushMessage(
    int(warden_sdk.MSG_WECHAT),
    "您有新的微信消息",
)

// 发送到蓝牙设备
// bluetoothDevice.Write(msgData)
```

### 6. 设置天气

```go
// 设置今天的天气
weatherData := sdk.SetSimpleWeather(
    int(warden_sdk.WEATHER_DATE_TODAY),
    int(warden_sdk.WEATHER_SUNSHINE),  // 晴天
    25,  // 当前温度
    18,  // 最低温度
    28,  // 最高温度
    "北京",
)

// 发送到蓝牙设备
// bluetoothDevice.Write(weatherData)
```

## 🔧 完整示例

查看 `example/demo_usage.go` 获取完整的使用示例。

运行演示：
```bash
cd e:\go\jiangsu_sdk
go run example/demo_usage.go
```

## 📚 可用的常量

SDK 提供了丰富的常量定义：

### 语言常量
```go
warden_sdk.LANG_ENGLISH
warden_sdk.LANG_CHINESE_SIMPLIFIED
warden_sdk.LANG_JAPANESE
// ... 共 33 种语言
```

### 消息类型常量
```go
warden_sdk.MSG_WECHAT
warden_sdk.MSG_QQ
warden_sdk.MSG_WHATSAPP
warden_sdk.MSG_FACEBOOK
// ... 共 30+ 种消息类型
```

### 天气类型常量
```go
warden_sdk.WEATHER_SUNSHINE   // 晴天
warden_sdk.WEATHER_RAIN       // 雨天
warden_sdk.WEATHER_SNOW       // 雪天
// ... 共 8 种天气类型
```

### 其他常量
```go
warden_sdk.TIMEZONE_EAST_8    // 东八区
warden_sdk.PHONE_TYPE_ANDROID // Android 手机
warden_sdk.TEMP_UNIT_CELSIUS  // 摄氏度
```

## 🎯 实际集成示例

### 与蓝牙库集成

```go
package main

import (
    "github.com/jiangsu/warden_sdk/warden_sdk"
    "your-bluetooth-library" // 你的蓝牙库
)

type WardenDevice struct {
    sdk *warden_sdk.WardenSDK
    ble *bluetooth.Device
}

func NewWardenDevice(bleDevice *bluetooth.Device) *WardenDevice {
    return &WardenDevice{
        sdk: warden_sdk.NewWardenSDK(),
        ble: bleDevice,
    }
}

// 同步时间
func (w *WardenDevice) SyncTime() error {
    data := w.sdk.SyncTime(
        warden_sdk.TIMEZONE_EAST_8,
        time.Now().Unix(),
    )
    return w.ble.Write(data)
}

// 查询电量
func (w *WardenDevice) GetBattery() (int, error) {
    requestData := w.sdk.QueryBattery()
    if err := w.ble.Write(requestData); err != nil {
        return 0, err
    }
    
    responseData, err := w.ble.Read()
    if err != nil {
        return 0, err
    }
    
    return warden_sdk.GetBatteryLevel(responseData)
}

// 推送消息
func (w *WardenDevice) PushNotification(msgType int, content string) error {
    data := w.sdk.PushMessage(msgType, content)
    return w.ble.Write(data)
}
```

## 📖 API 文档

完整的 API 文档请查看：
- [function_list.md](../function_list.md) - 所有函数详细说明
- [README.md](../README.md) - 项目概览

## 💡 优势

使用 Go 包的优势：

1. **类型安全** - Go 强类型系统
2. **无需编译** - 直接引用即可使用
3. **调试方便** - 可以直接查看源码
4. **性能优秀** - 原生 Go 代码
5. **跨平台** - 可用于服务器端、CLI 工具等

## 🎓 适用场景

- ✅ 服务器端协议处理
- ✅ 蓝牙通信库
- ✅ 测试和调试工具
- ✅ 命令行工具
- ✅ 协议转换服务

---

**更多示例和详细文档请查看项目目录下的其他文档！**
