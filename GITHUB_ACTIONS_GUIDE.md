# GitHub Actions 自动编译指南

本项目已配置 GitHub Actions 自动化编译流程，可自动构建 Android 和 iOS SDK。

## 📁 工作流文件

- `.github/workflows/build-android.yml` - Android AAR 自动编译
- `.github/workflows/build-ios.yml` - iOS Framework 编译（手动触发）

## 🚀 使用方法

### 方法 1：推送代码自动触发（Android）

将代码推送到 GitHub 仓库的 `main` 或 `master` 分支时，会自动触发 Android 编译：

```bash
git init
git add .
git commit -m "Initial commit"
git remote add origin https://github.com/你的用户名/warden_sdk.git
git push -u origin main
```

编译完成后，在 GitHub Actions 页面的 Artifacts 中下载 `warden-sdk-aar` 文件。

### 方法 2：手动触发编译

1. 访问你的 GitHub 仓库页面
2. 点击 **Actions** 标签
3. 选择左侧的工作流：
   - `Build Android AAR` - 编译 Android SDK
   - `Build iOS Framework` - 编译 iOS SDK
4. 点击右侧的 **Run workflow** 按钮
5. 选择分支，点击 **Run workflow**

### 方法 3：通过 Pull Request 触发

创建 Pull Request 时也会自动触发 Android 编译，用于验证代码变更。

## 📦 下载编译结果

### 从 GitHub Actions 下载

1. 进入 **Actions** 页面
2. 点击最近的成功运行（绿色勾号）
3. 在页面底部的 **Artifacts** 部分找到：
   - `warden-sdk-aar` - Android AAR 文件
   - `warden-sdk-ios-framework` - iOS Framework

### 通过 GitHub CLI 下载

```bash
# 安装 GitHub CLI
# https://cli.github.com/

# 下载最新的 Android AAR
gh run list --workflow=build-android.yml --limit 1
gh run download <run-id> -n warden-sdk-aar
```

## 🔧 配置说明

### Android 编译配置

- **Go 版本**: 1.24
- **Java 版本**: OpenJDK 17
- **Android API**: 31 (Android 12+)
- **NDK 版本**: 28.2.13676358
- **构建工具**: gomobile

### iOS 编译配置

- **Go 版本**: 1.24
- **最低 iOS 版本**: 15.0
- **运行环境**: macOS (GitHub 提供)
- **构建工具**: gomobile

### 触发条件

**Android 编译触发条件**：
- Push 到 `main` 或 `master` 分支
- Pull Request 到 `main` 或 `master` 分支
- 手动触发（workflow_dispatch）

**iOS 编译触发条件**：
- 仅手动触发（因为需要 macOS runner）

## 📊 编译流程

### Android 编译流程

1. 检出代码
2. 安装 Go 1.24
3. 安装 Java 17
4. 配置 Android SDK
5. 安装 Android NDK 28.2
6. 安装 gomobile 和 gobind
7. 初始化 gomobile
8. 执行编译生成 AAR
9. 上传 AAR 到 Artifacts

**预计编译时间**: 5-10 分钟

### iOS 编译流程

1. 检出代码
2. 安装 Go 1.24（macOS）
3. 安装 gomobile 和 gobind
4. 初始化 gomobile
5. 执行编译生成 Framework
6. 上传 Framework 到 Artifacts

**预计编译时间**: 3-5 分钟

## 🛠️ 本地测试工作流

可以使用 [act](https://github.com/nektos/act) 在本地测试 GitHub Actions：

```bash
# 安装 act
# Windows: choco install act-cli
# macOS: brew install act

# 测试 Android 编译工作流
act -W .github/workflows/build-android.yml

# 测试 iOS 编译工作流（需要 macOS）
act -W .github/workflows/build-ios.yml
```

## ⚠️ 注意事项

1. **首次使用**需要在 GitHub 仓库中启用 Actions
2. **私有仓库**可能受 Actions 使用分钟数限制
3. **iOS 编译**需要 macOS runner（GitHub 免费提供）
4. 编译产物保留 **30 天**，请及时下载
5. 确保仓库中有正确的 `go.mod` 文件

## 🔍 故障排查

### 编译失败

查看 Actions 日志：
1. 进入 Actions 页面
2. 点击失败的运行
3. 查看具体步骤的错误信息

### 常见问题

**问题 1**: `gomobile init` 失败
- 检查 NDK 版本是否正确安装
- 查看环境变量配置

**问题 2**: `gomobile bind` 失败
- 检查 Go 代码是否有编译错误
- 确认 gomobile 兼容性

**问题 3**: Artifacts 未生成
- 检查编译步骤是否成功
- 确认文件路径正确

## 📝 自定义配置

### 修改 NDK 版本

编辑 `.github/workflows/build-android.yml`:

```yaml
- name: Install Android NDK
  run: |
    sdkmanager --install "ndk;27.2.12479018"  # 修改版本号
    echo "ANDROID_NDK_HOME=$ANDROID_SDK_ROOT/ndk/27.2.12479018" >> $GITHUB_ENV
```

### 修改 Android API 级别

```yaml
- name: Build Android AAR
  run: |
    gomobile bind \
      -target=android \
      -androidapi=33 \  # 修改 API 级别
      -o warden_sdk.aar \
      ./warden_sdk
```

### 添加发布步骤

可以在工作流末尾添加自动发布到 GitHub Releases：

```yaml
- name: Create Release
  uses: softprops/action-gh-release@v1
  if: startsWith(github.ref, 'refs/tags/')
  with:
    files: warden_sdk.aar
```

## 🎯 最佳实践

1. **标签发布**: 使用 Git 标签触发正式版本编译
2. **语义版本**: 采用语义化版本号（v1.0.0, v1.1.0）
3. **变更日志**: 在 Release 中添加变更说明
4. **测试验证**: 在合并前通过 PR 验证编译

---

## 总结

GitHub Actions 提供了稳定可靠的自动化编译环境，解决了本地 gomobile 工具链的问题。通过 Actions 编译的 SDK 可以直接下载使用，无需配置复杂的本地环境。
