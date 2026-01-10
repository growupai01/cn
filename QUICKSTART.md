# Warden SDK - GitHub 自动编译快速入门

## ⚡ 5 分钟快速开始

### 步骤 1: 上传代码到 GitHub

```bash
cd e:\go\jiangsu_sdk

# 初始化 Git 仓库
git init

# 添加所有文件
git add .

# 提交
git commit -m "Warden SDK v1.0.0"

# 添加远程仓库（替换为你的仓库地址）
git remote add origin https://github.com/你的用户名/warden_sdk.git

# 推送到 GitHub
git push -u origin main
```

### 步骤 2: 查看自动编译

1. 打开 https://github.com/你的用户名/warden_sdk
2. 点击顶部的 **Actions** 标签
3. 会看到 `Build Android AAR` 正在运行 ⚡
4. 等待 5-10 分钟完成 ✅

### 步骤 3: 下载编译结果

1. 点击成功的运行（绿色勾号）
2. 页面底部 **Artifacts** 区域
3. 点击 `warden-sdk-aar` 下载
4. 解压得到 `warden_sdk.aar`

### 步骤 4: 在 Android 项目中使用

```gradle
// build.gradle
dependencies {
    implementation files('libs/warden_sdk.aar')
}
```

```kotlin
// Kotlin 代码
import warden_sdk.*

val sdk = Warden_sdk.NewWardenSDK()
val timeData = sdk.SyncTime(
    Warden_sdk.TIMEZONE_EAST_8.toByte(),
    System.currentTimeMillis() / 1000
)
```

## 🎉 完成！

你的 SDK 已经成功编译并可以使用了！

---

## 💡 手动触发编译

如果你想手动触发编译（不推送代码）：

1. 进入 Actions 页面
2. 左侧选择 `Build Android AAR`
3. 右侧点击 `Run workflow`
4. 点击绿色的 `Run workflow` 按钮

---

## 📚 更多信息

- [完整 GitHub Actions 指南](GITHUB_ACTIONS_GUIDE.md)
- [API 文档](function_list.md)
- [使用说明](README.md)
