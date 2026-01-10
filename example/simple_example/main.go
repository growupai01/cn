package main

import (
	"fmt"
	"time"

	"github.com/jiangsu/warden_sdk/warden_sdk"
)

func main() {
	fmt.Println("=== Warden SDK Go 包使用演示 ===\n")

	// 创建 SDK 实例
	sdk := warden_sdk.NewWardenSDK()
	fmt.Println("✓ SDK 实例已创建\n")

	// 1. 同步时间
	fmt.Println("1️⃣ 同步时间")
	timeData := sdk.SyncTime(
		warden_sdk.TIMEZONE_EAST_8,
		time.Now().Unix(),
	)
	fmt.Printf("   请求数据: %d 字节\n", len(timeData))
	fmt.Printf("   时区: 东八区（北京时间）\n\n")

	// 2. 查询电量
	fmt.Println("2️⃣ 查询电量")
	batteryRequest := sdk.QueryBattery()
	fmt.Printf("   请求数据: %d 字节\n\n", len(batteryRequest))

	// 3. 设置亮度
	fmt.Println("3️⃣ 设置屏幕亮度")
	brightnessData := sdk.SetBrightness(80)
	fmt.Printf("   目标亮度: 80%%\n")
	fmt.Printf("   请求数据: %d 字节\n\n", len(brightnessData))

	// 4. 推送微信消息
	fmt.Println("4️⃣ 推送微信消息")
	msgData := sdk.PushMessage(
		int(warden_sdk.MSG_WECHAT),
		"您有新的微信消息",
	)
	fmt.Printf("   消息类型: 微信\n")
	fmt.Printf("   消息内容: 您有新的微信消息\n")
	fmt.Printf("   请求数据: %d 字节\n\n", len(msgData))

	// 5. 设置天气
	fmt.Println("5️⃣ 设置天气信息")
	weatherData := sdk.SetSimpleWeather(
		int(warden_sdk.WEATHER_DATE_TODAY),
		int(warden_sdk.WEATHER_SUNSHINE),
		25, // 当前温度
		18, // 最低温度
		28, // 最高温度
		"北京",
	)
	fmt.Printf("   地点: 北京\n")
	fmt.Printf("   天气: 晴天\n")
	fmt.Printf("   温度: 25℃ (18℃ ~ 28℃)\n")
	fmt.Printf("   请求数据: %d 字节\n\n", len(weatherData))

	// 6. 设置勿扰模式
	fmt.Println("6️⃣ 设置勿扰模式")
	dndData := sdk.SetDND(true, 22, 0, 7, 0)
	fmt.Printf("   勿扰时间: 22:00 - 07:00\n")
	fmt.Printf("   请求数据: %d 字节\n\n", len(dndData))

	// 7. 设备绑定流程
	fmt.Println("7️⃣ 设备绑定流程")
	bindStart := sdk.StartBinding()
	fmt.Printf("   ① 发送绑定开始: %d 字节\n", len(bindStart))

	appInfo := sdk.SetAppInfo(int(warden_sdk.PHONE_TYPE_ANDROID))
	fmt.Printf("   ② 设置应用信息: %d 字节\n", len(appInfo))

	bindEnd := sdk.EndBindingData()
	fmt.Printf("   ③ 绑定数据结束: %d 字节\n", len(bindEnd))
	fmt.Println("   ✓ 绑定流程完成\n")

	// 使用常量
	fmt.Println("8️⃣ SDK 提供的常量")
	fmt.Printf("   支持的语言数: 33种\n")
	fmt.Printf("   支持的消息类型: 30+种\n")
	fmt.Printf("   支持的天气类型: 8种\n")
	fmt.Printf("   支持的开关类型: 14种\n\n")

	fmt.Println("=== 演示完成 ===")
	fmt.Println("\n💡 在实际使用中，将这些数据包发送到蓝牙设备即可！")
}
