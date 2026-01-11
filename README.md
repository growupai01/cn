# Yingka iOS SDK

This project contains the Go source code to generate the Yingka SDK for iOS as an XCFramework using `gomobile`.

## How to Build

### Using GitHub Actions (Recommended)

Since `gomobile bind -target=ios` requires macOS and Xcode, the easiest way to build is using the included GitHub Actions workflow.

1.  Push this entire `yingka_ios_sdk` folder to a GitHub repository.
2.  Go to the "Actions" tab in your repository.
3.  You should see the "Build iOS SDK" workflow running.
4.  Once completed, download the `YingkaSDK.xcframework` artifact.

### Local Build (macOS only)

If you have a Mac with Xcode installed:

1.  Install Go 1.21+.
2.  Install gomobile:
    ```bash
    go install golang.org/x/mobile/cmd/gomobile@latest
    gomobile init
    ```
3.  Run the build command:
    ```bash
    gomobile bind -target=ios -o YingkaSDK.xcframework .
    ```

## File Structure

- `sdk.go`: Main SDK entry point (adapted from Android SDK).
- `constants.go`: Command constants.
- `protocol.go`: Packet definition and serialization.
- `device.go`, `ai.go`, `sync.go`: Features implementations.
- `.github/workflows/ios.yml`: CI configuration.
