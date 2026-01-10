package warden_sdk

// MergeResult 合并结果（gomobile 友好）
type MergeResult struct {
	Complete bool
	Data     []byte
	Error    error
}

// AddFrameData 添加帧数据并返回合并结果（gomobile 友好版本）
func (pm *PacketMerger) AddFrameData(frameData []byte) (*MergeResult, error) {
	packet, err := DecodePacket(frameData)
	if err != nil {
		return &MergeResult{Error: err}, err
	}

	complete, merged, err := pm.addFrame(packet)
	return &MergeResult{
		Complete: complete,
		Data:     merged,
		Error:    err,
	}, err
}
