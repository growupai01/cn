package warden_sdk

// RemoteCamera 远程拍照功能 (0x91)

// 相机动作
const (
	CAMERA_ACTION_TAKE_PHOTO  byte = 0x00 // 远程拍照
	CAMERA_ACTION_EXIT_CAMERA byte = 0x01 // 退出APP相机
	CAMERA_ACTION_OPEN_CAMERA byte = 0x02 // 呼出APP相机
)

// BuildQueryCameraStatusRequest 构建查询拍照状态请求包
func BuildQueryCameraStatusRequest() []byte {
	pb := NewPacketBuilder()
	payload := []byte{ACTION_QUERY} // 0x00: 查询
	return pb.BuildSimpleRequest(CMD_REMOTE_CAMERA, payload)
}

// BuildSetCameraControlRequest 构建设置拍照开关请求包
func BuildSetCameraControlRequest(enable bool) []byte {
	pb := NewPacketBuilder()

	cameraCtrl := byte(0x00)
	if enable {
		cameraCtrl = 0x01
	}

	payload := []byte{
		ACTION_SET, // 0x01: 设置
		cameraCtrl, // 0x00: 关闭; 0x01: 打开
	}
	return pb.BuildSimpleRequest(CMD_REMOTE_CAMERA, payload)
}

// CameraStatusResponse 拍照状态响应
type CameraStatusResponse struct {
	ActionCmd byte
	Status    bool // true: 打开; false: 关闭
}

// ParseCameraResponse 解析拍照响应
func ParseCameraResponse(data []byte) (*CameraStatusResponse, error) {
	packet, err := DecodePacket(data)
	if err != nil {
		return nil, err
	}

	if packet.Payload[0] == ACTION_QUERY {
		// 查询响应
		if len(packet.Payload) < 2 {
			return nil, ErrInvalidPayloadLength
		}
		return &CameraStatusResponse{
			ActionCmd: packet.Payload[0],
			Status:    packet.Payload[1] == 0x01,
		}, nil
	} else {
		// 设置响应
		if len(packet.Payload) < 2 {
			return nil, ErrInvalidPayloadLength
		}
		return nil, ValidateResponse(packet.Payload[1])
	}
}

// CameraNotify 拍照通知
type CameraNotify struct {
	Action byte // 0x00: 远程拍照; 0x01: 退出APP相机; 0x02: 呼出APP相机
}

// ParseCameraNotify 解析拍照通知
func ParseCameraNotify(data []byte) (*CameraNotify, error) {
	packet, err := DecodePacket(data)
	if err != nil {
		return nil, err
	}

	if len(packet.Payload) < 1 {
		return nil, ErrInvalidPayloadLength
	}

	return &CameraNotify{
		Action: packet.Payload[0],
	}, nil
}
