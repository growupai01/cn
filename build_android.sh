#!/bin/bash

# Warden SDK Android 编译脚本
# 用途：将 Go SDK 编译为 Android AAR 格式

set -e  # 遇到错误立即退出

echo "=========================================="
echo "Warden SDK Android 编译脚本"
echo "=========================================="
echo ""

# 检查 Go 环境
if ! command -v go &> /dev/null; then
    echo "错误: 未找到 Go 环境，请先安装 Go 1.16+"
    exit 1
fi

echo "✓ Go 版本: $(go version)"
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
OUTPUT_AAR="warden_sdk.aar"
PACKAGE_PATH="./warden_sdk"

echo "开始编译 Android AAR..."
echo "  目标平台: Android"
echo "  API 等级: 31 (Android 12+)"
echo "  输出文件: $OUTPUT_AAR"
echo ""

# 编译 AAR
gomobile bind \
    -target=android \
    -androidapi=31 \
    -o "$OUTPUT_AAR" \
    "$PACKAGE_PATH"

# 检查编译结果
if [ -f "$OUTPUT_AAR" ]; then
    FILE_SIZE=$(ls -lh "$OUTPUT_AAR" | awk '{print $5}')
    echo ""
    echo "=========================================="
    echo "✓ 编译成功!"
    echo "=========================================="
    echo "文件: $OUTPUT_AAR"
    echo "大小: $FILE_SIZE"
    echo ""
    echo "使用方法:"
    echo "  1. 将 $OUTPUT_AAR 复制到 Android 项目的 libs 目录"
    echo "  2. 在 build.gradle 中添加依赖:"
    echo "     implementation files('libs/$OUTPUT_AAR')"
    echo "  3. 在 Kotlin/Java 代码中导入:"
    echo "     import warden_sdk.*"
    echo ""
else
    echo ""
    echo "=========================================="
    echo "✗ 编译失败"
    echo "=========================================="
    exit 1
fi
