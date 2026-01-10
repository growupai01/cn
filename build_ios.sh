#!/bin/bash

# Warden SDK iOS 编译脚本
# 用途：将 Go SDK 编译为 iOS Framework 格式
# 系统要求：macOS + Xcode

set -e  # 遇到错误立即退出

echo "=========================================="
echo "Warden SDK iOS 编译脚本"
echo "=========================================="
echo ""

# 检查是否在 macOS 上运行
if [[ "$OSTYPE" != "darwin"* ]]; then
    echo "错误: 此脚本只能在 macOS 系统上运行"
    echo "当前系统: $OSTYPE"
    exit 1
fi

# 检查 Go 环境
if ! command -v go &> /dev/null; then
    echo "错误: 未找到 Go 环境，请先安装 Go 1.16+"
    exit 1
fi

echo "✓ Go 版本: $(go version)"
echo ""

# 检查 Xcode
if ! command -v xcodebuild &> /dev/null; then
    echo "错误: 未找到 Xcode，请先安装 Xcode"
    exit 1
fi

echo "✓ Xcode 版本: $(xcodebuild -version | head -n 1)"
echo ""

# 检查 gomobile
if ! command -v gomobile &> /dev/null; then
    echo "未找到 gomobile，正在安装..."
    go install golang.org/x/mobile/cmd/gomobile@latest
    echo "✓ gomobile 安装完成"
    echo ""
    
    echo "正在初始化 gomobile..."
    gomobile init
    echo "✓ gomobile 初始化完成"
    echo ""
fi

# 设置输出文件名
OUTPUT_FRAMEWORK="Warden_sdk.framework"
OUTPUT_XCFRAMEWORK="Warden_sdk.xcframework"
PACKAGE_PATH="./warden_sdk"

echo "开始编译 iOS Framework..."
echo "  目标平台: iOS"
echo "  最低版本: iOS 15.0"
echo "  输出文件: $OUTPUT_XCFRAMEWORK"
echo ""

# 清理旧文件
if [ -d "$OUTPUT_FRAMEWORK" ]; then
    rm -rf "$OUTPUT_FRAMEWORK"
fi
if [ -d "$OUTPUT_XCFRAMEWORK" ]; then
    rm -rf "$OUTPUT_XCFRAMEWORK"
fi

# 编译 iOS Framework
# -target=ios 表示编译为 iOS 通用 Framework（支持模拟器和真机）
gomobile bind \
    -target=ios \
    -iosversion=15.0 \
    -o "$OUTPUT_FRAMEWORK" \
    "$PACKAGE_PATH"

# 检查编译结果
if [ -d "$OUTPUT_FRAMEWORK" ]; then
    FRAMEWORK_SIZE=$(du -sh "$OUTPUT_FRAMEWORK" | awk '{print $1}')
    echo ""
    echo "=========================================="
    echo "✓ 编译成功!"
    echo "=========================================="
    echo "Framework: $OUTPUT_FRAMEWORK"
    echo "大小: $FRAMEWORK_SIZE"
    echo ""
    
    # 可选：创建 XCFramework（同时支持模拟器和真机）
    echo "正在创建 XCFramework（可选）..."
    
    # 分别编译模拟器和真机版本
    echo "  编译模拟器版本..."
    gomobile bind \
        -target=ios/arm64,iossimulator/amd64 \
        -iosversion=15.0 \
        -o "$OUTPUT_XCFRAMEWORK" \
        "$PACKAGE_PATH" 2>/dev/null || {
        echo "  (XCFramework 创建跳过，使用标准 Framework 即可)"
    }
    
    echo ""
    echo "使用方法:"
    echo "  1. 将 $OUTPUT_FRAMEWORK 拖入 Xcode 项目"
    echo "  2. 在 Build Phases -> Link Binary With Libraries 中添加"
    echo "  3. 在 Swift 代码中导入:"
    echo "     import Warden_sdk"
    echo ""
    echo "Swift 使用示例:"
    echo "  let timeRequest = Warden_sdkBuildSyncTimeRequest("
    echo "      Warden_sdkTIMEZONE_EAST_8,"
    echo "      UInt32(Date().timeIntervalSince1970)"
    echo "  )"
    echo ""
else
    echo ""
    echo "=========================================="
    echo "✗ 编译失败"
    echo "=========================================="
    exit 1
fi
