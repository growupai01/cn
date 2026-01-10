package warden_sdk

import (
	"fmt"
)

// PacketBuilder 数据包构建器
type PacketBuilder struct {
	seqNum byte
}

// NewPacketBuilder 创建新的数据包构建器
func NewPacketBuilder() *PacketBuilder {
	return &PacketBuilder{
		seqNum: 0,
	}
}

// BuildSimpleRequest 构建简单请求包（单帧，不加密）
func (pb *PacketBuilder) BuildSimpleRequest(cmd byte, payload []byte) []byte {
	header := NewHeader(pb.nextSeqNum(), cmd, CMD_TYPE_REQUEST, 0, byte(len(payload)))
	packet := NewPacket(header, payload)
	return packet.Encode()
}

// BuildSimpleResponse 构建简单响应包
func (pb *PacketBuilder) BuildSimpleResponse(cmd byte, payload []byte) []byte {
	header := NewHeader(pb.nextSeqNum(), cmd, CMD_TYPE_RESPONSE, 0, byte(len(payload)))
	packet := NewPacket(header, payload)
	return packet.Encode()
}

// BuildSimpleNotify 构建简单通知包
func (pb *PacketBuilder) BuildSimpleNotify(cmd byte, payload []byte) []byte {
	header := NewHeader(pb.nextSeqNum(), cmd, CMD_TYPE_NOTIFY, 0, byte(len(payload)))
	packet := NewPacket(header, payload)
	return packet.Encode()
}

// BuildMultiFramePackets 构建多帧数据包（用于拆包）
func (pb *PacketBuilder) BuildMultiFramePackets(cmd, cmdType byte, payload []byte, mtu int) [][]byte {
	maxPayloadPerFrame := mtu - HEADER_SIZE
	totalLength := len(payload)

	if totalLength <= maxPayloadPerFrame {
		// 不需要拆包
		header := NewHeader(pb.nextSeqNum(), cmd, cmdType, 0, byte(totalLength))
		packet := NewPacket(header, payload)
		return [][]byte{packet.Encode()}
	}

	// 需要拆包
	totalFrames := (totalLength + maxPayloadPerFrame - 1) / maxPayloadPerFrame
	if totalFrames > MAX_TOTAL_FRAMES {
		totalFrames = MAX_TOTAL_FRAMES
	}

	packets := make([][]byte, 0, totalFrames)
	seqNum := pb.nextSeqNum()

	for frameSeq := 0; frameSeq < totalFrames; frameSeq++ {
		start := frameSeq * maxPayloadPerFrame
		end := start + maxPayloadPerFrame
		if end > totalLength {
			end = totalLength
		}

		framePayload := payload[start:end]
		header := NewHeader(seqNum, cmd, cmdType, 0, byte(len(framePayload)))
		header.SetFrameInfo(byte(frameSeq), byte(totalFrames))

		packet := NewPacket(header, framePayload)
		packets = append(packets, packet.Encode())
	}

	return packets
}

// nextSeqNum 获取下一个序列号
func (pb *PacketBuilder) nextSeqNum() byte {
	current := pb.seqNum
	pb.seqNum = (pb.seqNum + 1) & 0x0F // 循环递增 0-15
	return current
}

// ResetSeqNum 重置序列号
func (pb *PacketBuilder) ResetSeqNum() {
	pb.seqNum = 0
}

// PacketMerger 数据包合并器（用于处理拆包数据）
type PacketMerger struct {
	frames map[byte]map[byte][]byte // [cmd][frameSeq] -> payload
}

// NewPacketMerger 创建新的数据包合并器
func NewPacketMerger() *PacketMerger {
	return &PacketMerger{
		frames: make(map[byte]map[byte][]byte),
	}
}

// addFrame 添加一帧数据 (内部方法，不导出到 gomobile)
func (pm *PacketMerger) addFrame(packet *Packet) (complete bool, merged []byte, err error) {
	cmd := packet.Header.Cmd
	frameSeq := packet.Header.GetFrameSeq()
	totalFrames := packet.Header.GetTotalFrames()

	// 单帧数据，直接返回
	if totalFrames == 1 {
		return true, packet.Payload, nil
	}

	// 初始化该命令的帧映射
	if pm.frames[cmd] == nil {
		pm.frames[cmd] = make(map[byte][]byte)
	}

	// 存储该帧
	pm.frames[cmd][frameSeq] = packet.Payload

	// 检查是否收集齐所有帧
	if len(pm.frames[cmd]) == int(totalFrames) {
		// 合并所有帧
		merged = pm.mergeFrames(cmd, totalFrames)
		delete(pm.frames, cmd) // 清理
		return true, merged, nil
	}

	return false, nil, nil
}

// mergeFrames 合并帧数据
func (pm *PacketMerger) mergeFrames(cmd byte, totalFrames byte) []byte {
	var result []byte
	for i := byte(0); i < totalFrames; i++ {
		if frame, ok := pm.frames[cmd][i]; ok {
			result = append(result, frame...)
		}
	}
	return result
}

// ClearCmd 清理指定命令的缓存帧
func (pm *PacketMerger) ClearCmd(cmd byte) {
	delete(pm.frames, cmd)
}

// ValidateResponse 验证响应结果
func ValidateResponse(resultCode byte) error {
	switch resultCode {
	case CMD_RESULT_OK:
		return nil
	case CMD_RESULT_FAIL:
		return fmt.Errorf("command failed")
	case CMD_RESULT_LEN_INVALID:
		return fmt.Errorf("invalid packet length")
	case CMD_RESULT_CMD_INVALID:
		return fmt.Errorf("invalid command type")
	case CMD_RESULT_IDX_INVALID:
		return fmt.Errorf("invalid index")
	case CMD_RESULT_NOT_SUPPORT:
		return fmt.Errorf("command not supported")
	case CMD_RESULT_SWI_ERROR:
		return fmt.Errorf("switch operation error")
	case CMD_RESULT_CHECK_ERROR:
		return fmt.Errorf("data check error")
	case CMD_RESULT_PACKET_LOSS:
		return fmt.Errorf("packet loss detected")
	default:
		return fmt.Errorf("unknown error code: 0x%02X", resultCode)
	}
}
