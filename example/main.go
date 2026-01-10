package main

import (
	"fmt"
	"time"

	"github.com/jiangsu/warden_sdk/warden_sdk"
)

// 演示：如何使用 Warden SDK

func main() {
	fmt.Println("=== Warden SDK 使用示例 ===\n")

	// 1. 同步时间
	fmt.Println("1. 同步时间")
	syncTime()
	fmt.Println()

	// 2. 查询电量
	fmt.Println("2. 查询电量")
	queryBattery()
	fmt.Println()

	// 3. 设置亮度
	fmt.Println("3. 设置屏幕亮度")
	setBrightness()
	fmt.Println()

	// 4. 设置语言
	fmt.Println("4. 设置设备语言")
	setLanguage()
	fmt.Println()

	// 5. 设置勿扰模式
	fmt.Println("5. 设置勿扰模式")
	setDND()
	fmt.Println()

	// 6. 推送消息
	fmt.Println("6. 推送微信消息")
	pushMessage()
	fmt.Println()

	// 7. 设置天气
	fmt.Println("7. 设置天气信息")
	setWeather()
	fmt.Println()

	// 8. 设备绑定流程
	fmt.Println("8. 设备绑定流程")
	deviceBinding()
	fmt.Println()
}

// syncTime 演示时间同步
func syncTime() {
	timeZone := warden_sdk.TIMEZONE_EAST_8 // 东八区（北京时间）
	utc := uint32(time.Now().Unix())

	requestData := warden_sdk.BuildSyncTimeRequest(timeZone, utc)
	fmt.Printf("  构建同步时间请求: %d 字节\n", len(requestData))
	fmt.Printf("  时区: 东八区, UTC: %d\n", utc)

	// 这里应该发送 requestData 到蓝牙设备
	// bluetoothDevice.Write(requestData)

	// 模拟接收响应
	// responseData := bluetoothDevice.Read()
	// resp, err := warden_sdk.ParseSyncTimeResponse(responseData)
	// if err != nil {
	//     fmt.Printf("  错误: %v\n", err)
	// } else {
	//     fmt.Println("  同步成功")
	// }
}

// queryBattery 演示查询电量
func queryBattery() {
	requestData := warden_sdk.BuildQueryBatteryRequest()
	fmt.Printf("  构建查询电量请求: %d 字节\n", len(requestData))

	// 模拟响应数据（这里演示解析）
	// 假设设备返回：电量80%，正在充电
	mockResponse := []byte{0x00, 0x51, 0x02, 0x00, 0x02, 0x00, 0xD0} // 0xD0 = 0x80 | 0x50 (充电中 + 80%)

	batteryInfo, err := warden_sdk.ParseBatteryResponse(mockResponse)
	if err != nil {
		fmt.Printf("  错误: %v\n", err)
	} else {
		fmt.Printf("  电量: %d%%, 充电中: %v\n", batteryInfo.Level, batteryInfo.IsCharging)
	}
}

// setBrightness 演示设置亮度
func setBrightness() {
	brightness := byte(80) // 设置为 80%
	requestData := warden_sdk.BuildSetBrightnessRequest(brightness)
	fmt.Printf("  构建设置亮度请求: %d 字节\n", len(requestData))
	fmt.Printf("  目标亮度: %d%%\n", brightness)
}

// setLanguage 演示设置语言
func setLanguage() {
	lang := warden_sdk.LANG_CHINESE_SIMPLIFIED // 简体中文
	requestData := warden_sdk.BuildSetLanguageRequest(lang)
	fmt.Printf("  构建设置语言请求: %d 字节\n", len(requestData))
	fmt.Println("  设置语言: 简体中文")
}

// setDND 演示设置勿扰模式
func setDND() {
	settings := &warden_sdk.DNDSettings{
		Enable:      true,
		StartHour:   22, // 晚上10点
		StartMinute: 0,
		EndHour:     7, // 早上7点
		EndMinute:   0,
	}

	requestData := warden_sdk.BuildSetDNDRequest(settings)
	fmt.Printf("  构建设置勿扰请求: %d 字节\n", len(requestData))
	fmt.Printf("  勿扰时间: %02d:%02d - %02d:%02d\n",
		settings.StartHour, settings.StartMinute,
		settings.EndHour, settings.EndMinute)
}

// pushMessage 演示推送消息
func pushMessage() {
	msgType := warden_sdk.MSG_WECHAT
	content := []byte("您有新的微信消息：张三给您发送了一条消息")
	mtu := 244 // 蓝牙MTU

	packets := warden_sdk.BuildPushMessagePackets(msgType, content, mtu)
	fmt.Printf("  消息类型: 微信\n")
	fmt.Printf("  消息内容长度: %d 字节\n", len(content))
	fmt.Printf("  拆分为 %d 个数据包\n", len(packets))

	for i, pkt := range packets {
		fmt.Printf("    包 %d: %d 字节\n", i+1, len(pkt))
	}
}

// setWeather 演示设置天气
func setWeather() {
	weather := &warden_sdk.WeatherData{
		DateType:    warden_sdk.WEATHER_DATE_TODAY,
		WeatherType: int8(warden_sdk.WEATHER_SUNSHINE), // 晴天
		CurrentTemp: 25,                                // 当前温度 25℃
		MinTemp:     18,                                // 最低温度 18℃
		MaxTemp:     28,                                // 最高温度 28℃
		Location:    "北京",
		Humidity:    60, // 湿度 60%
		WindSpeed:   15, // 风速 15 km/h
		UVIndex:     7,  // 紫外线指数 7
	}

	requestData := warden_sdk.BuildSetWeatherRequest(weather)
	fmt.Printf("  构建设置天气请求: %d 字节\n", len(requestData))
	fmt.Printf("  地点: %s\n", weather.Location)
	fmt.Printf("  天气: 晴天\n")
	fmt.Printf("  温度: %d℃ (最低: %d℃, 最高: %d℃)\n",
		weather.CurrentTemp, weather.MinTemp, weather.MaxTemp)
	fmt.Printf("  湿度: %d%%, 风速: %d km/h, 紫外线: %d\n",
		weather.Humidity, weather.WindSpeed, weather.UVIndex)
}

// deviceBinding 演示设备绑定流程
func deviceBinding() {
	fmt.Println("  步骤 1: 发送绑定开始")
	startRequest := warden_sdk.BuildStartBindingRequest()
	fmt.Printf("    请求大小: %d 字节\n", len(startRequest))

	fmt.Println("  步骤 2: 同步数据到设备")
	fmt.Println("    - 同步时间")
	fmt.Println("    - 设置应用信息")
	fmt.Println("    - 设置开关状态")

	fmt.Println("  步骤 3: 发送绑定数据结束")
	endRequest := warden_sdk.BuildEndBindingDataRequest()
	fmt.Printf("    请求大小: %d 字节\n", len(endRequest))

	fmt.Println("  步骤 4: 等待设备初始化完成...")
	fmt.Println("  步骤 5: 接收设备绑定完成通知")
	fmt.Println("  ✓ 绑定流程完成")
}

// 演示数据包合并（多帧数据）
func demonstratePacketMerging() {
	fmt.Println("\n=== 演示：数据包合并 ===\n")

	merger := warden_sdk.NewPacketMerger()

	// 模拟接收3帧数据
	fmt.Println("接收多帧消息数据...")

	// 这里应该是从蓝牙设备接收的数据
	// for each received frame:
	//   packet, _ := warden_sdk.DecodePacket(frameData)
	//   complete, merged, err := merger.AddFrame(packet)
	//   if complete {
	//       // 所有帧已接收，处理合并后的数据
	//       fmt.Printf("接收完整消息: %s\n", string(merged))
	//   }

	fmt.Println("(实际使用时需要接收蓝牙数据)")
}

// 演示开关操作
func demonstrateSwitchOperations() {
	fmt.Println("\n=== 演示：开关操作 ===\n")

	var flags uint32 = 0

	// 设置多个开关
	flags = warden_sdk.SetSwitchBit(flags, warden_sdk.SWITCH_ANTI_LOST, true)
	flags = warden_sdk.SetSwitchBit(flags, warden_sdk.SWITCH_RAISE_WRIST, true)
	flags = warden_sdk.SetSwitchBit(flags, warden_sdk.SWITCH_MSG_TOTAL, true)

	fmt.Printf("开关标志: 0x%08X\n", flags)
	fmt.Printf("  防丢开关: %v\n", warden_sdk.GetSwitchBit(flags, warden_sdk.SWITCH_ANTI_LOST))
	fmt.Printf("  抬手亮屏: %v\n", warden_sdk.GetSwitchBit(flags, warden_sdk.SWITCH_RAISE_WRIST))
	fmt.Printf("  消息提醒: %v\n", warden_sdk.GetSwitchBit(flags, warden_sdk.SWITCH_MSG_TOTAL))

	requestData := warden_sdk.BuildSetSwitchRequest(flags)
	fmt.Printf("\n构建设置开关请求: %d 字节\n", len(requestData))
}
