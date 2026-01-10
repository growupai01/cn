# Warden SDK

Warden 协议 Go SDK - 用于蓝牙智能设备通信

[![Build Android AAR](https://github.com/你的用户名/warden_sdk/actions/workflows/build-android.yml/badge.svg)](https://github.com/你的用户名/warden_sdk/actions/workflows/build-android.yml)

## 快速开始

### 下载编译好的 SDK

从 [GitHub Actions](https://github.com/你的用户名/warden_sdk/actions) 下载最新编译的：
- **Android AAR**: `warden-sdk-aar`
- **iOS Framework**: `warden-sdk-ios-framework`

### 使用 Go 包

```go
import "github.com/jiangsu/warden_sdk/warden_sdk"

sdk := warden_sdk.NewWardenSDK()
timeData := sdk.SyncTime(warden_sdk.TIMEZONE_EAST_8, time.Now().Unix())
```

## 功能特性

- ✅ 完整的 Warden 协议实现（19个指令模块）
- ✅ 85+ 个函数接口
- ✅ 自动拆包/合包支持
- ✅ Android 和 iOS 跨平台
- ✅ GitHub Actions 自动编译

## 文档

- [README.md](README.md) - 完整使用说明
- [function_list.md](function_list.md) - API 函数列表
- [GITHUB_ACTIONS_GUIDE.md](GITHUB_ACTIONS_GUIDE.md) - 自动编译指南

## 许可证

MIT License
