# Warden SDK Android 编译脚本 (Windows PowerShell)
# 用途：将 Go SDK 编译为 Android AAR 格式

Write-Host "==========================================" -ForegroundColor Cyan
Write-Host "Warden SDK Android 编译脚本" -ForegroundColor Cyan
Write-Host "==========================================" -ForegroundColor Cyan
Write-Host ""

# 检查 Go 环境
try {
    $goVersion = go version
    Write-Host "✓ Go 版本: $goVersion" -ForegroundColor Green
    Write-Host ""
} catch {
    Write-Host "错误: 未找到 Go 环境，请先安装 Go 1.16+" -ForegroundColor Red
    exit 1
}

# 检查 gomobile
try {
    gomobile version 2>$null
} catch {
    Write-Host "未找到 gomobile，正在安装..." -ForegroundColor Yellow
    go install golang.org/x/mobile/cmd/gomobile@latest
    Write-Host "✓ gomobile 安装完成" -ForegroundColor Green
    Write-Host ""
    
    Write-Host "正在初始化 gomobile..." -ForegroundColor Yellow
    gomobile init
    Write-Host "✓ gomobile 初始化完成" -ForegroundColor Green
    Write-Host ""
}

# 设置输出文件名
$OUTPUT_AAR = "warden_sdk.aar"
$PACKAGE_PATH = "./warden_sdk"

Write-Host "开始编译 Android AAR..." -ForegroundColor Yellow
Write-Host "  目标平台: Android"
Write-Host "  API 等级: 31 (Android 12+)"
Write-Host "  输出文件: $OUTPUT_AAR"
Write-Host ""

# 编译 AAR
try {
    gomobile bind -target=android -androidapi=31 -o $OUTPUT_AAR $PACKAGE_PATH
    
    # 检查编译结果
    if (Test-Path $OUTPUT_AAR) {
        $fileInfo = Get-Item $OUTPUT_AAR
        $fileSize = "{0:N2} MB" -f ($fileInfo.Length / 1MB)
        
        Write-Host ""
        Write-Host "==========================================" -ForegroundColor Green
        Write-Host "✓ 编译成功!" -ForegroundColor Green
        Write-Host "==========================================" -ForegroundColor Green
        Write-Host "文件: $OUTPUT_AAR"
        Write-Host "大小: $fileSize"
        Write-Host ""
        Write-Host "使用方法:" -ForegroundColor Cyan
        Write-Host "  1. 将 $OUTPUT_AAR 复制到 Android 项目的 libs 目录"
        Write-Host "  2. 在 build.gradle 中添加依赖:"
        Write-Host "     implementation files('libs/$OUTPUT_AAR')"
        Write-Host "  3. 在 Kotlin/Java 代码中导入:"
        Write-Host "     import warden_sdk.*"
        Write-Host ""
    } else {
        throw "AAR 文件未生成"
    }
} catch {
    Write-Host ""
    Write-Host "==========================================" -ForegroundColor Red
    Write-Host "✗ 编译失败" -ForegroundColor Red
    Write-Host "==========================================" -ForegroundColor Red
    Write-Host "错误信息: $_" -ForegroundColor Red
    exit 1
}
