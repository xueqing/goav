// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package libavformat

//#cgo pkg-config: libavformat
//#include <libavformat/avformat.h>
import "C"
import (
	"unsafe"

	"github.com/xueqing/goav/libavcodec"
	"github.com/xueqing/goav/libavutil"
)

// CodecParameters Return codecpar
func (avs *AvStream) CodecParameters() *libavcodec.AvCodecParameters {
	return (*libavcodec.AvCodecParameters)(unsafe.Pointer(avs.codecpar))
}

// SetCodecParameters Set codecpar
func (avs *AvStream) SetCodecParameters(pCodecPara *libavcodec.AvCodecParameters) {
	avs.codecpar = (*C.struct_AVCodecParameters)(unsafe.Pointer(pCodecPara))
}

// Codec Return codec
func (avs *AvStream) Codec() *AvCodecContext {
	return (*AvCodecContext)(unsafe.Pointer(avs.codec))
}

// Metadata Return metadata
func (avs *AvStream) Metadata() *libavutil.AvDictionary {
	return (*libavutil.AvDictionary)(unsafe.Pointer(avs.metadata))
}

// IndexEntries Return index_entries
func (avs *AvStream) IndexEntries() *AvIndexEntry {
	return (*AvIndexEntry)(unsafe.Pointer(avs.index_entries))
}

// AttachedPic Return attached_pic
func (avs *AvStream) AttachedPic() libavcodec.AvPacket {
	return *fromCPacket(&avs.attached_pic)
}

// SideData Return side_data
func (avs *AvStream) SideData() *AvPacketSideData {
	return (*AvPacketSideData)(unsafe.Pointer(avs.side_data))
}

// ProbeData Return probe_data
func (avs *AvStream) ProbeData() AvProbeData {
	return AvProbeData(avs.probe_data)
}

// AvgFrameRate Return avg_frame_rate
func (avs *AvStream) AvgFrameRate() libavcodec.AvRational {
	return newRational(avs.avg_frame_rate)
}

// func (avs *AvStream) DisplayAspectRatio() *AvRational {
// 	return (*AvRational)(unsafe.Pointer(avs.display_aspect_ratio))
// }

// RFrameRate Return r_frame_rate
func (avs *AvStream) RFrameRate() libavcodec.AvRational {
	return newRational(avs.r_frame_rate)
}

// SetRFrameRate Set r_frame_rate
func (avs *AvStream) SetRFrameRate(frameRate libavcodec.AvRational) {
	avs.r_frame_rate.num = C.int(frameRate.Num())
	avs.r_frame_rate.den = C.int(frameRate.Den())
}

// SampleAspectRatio Return sample_aspect_ratio
func (avs *AvStream) SampleAspectRatio() libavcodec.AvRational {
	return newRational(avs.sample_aspect_ratio)
}

// TimeBase Return time_base
func (avs *AvStream) TimeBase() libavcodec.AvRational {
	return newRational(avs.time_base)
}

// SetTimeBase Set time_base
func (avs *AvStream) SetTimeBase(timeBase libavcodec.AvRational) {
	avs.time_base.num = C.int(timeBase.Num())
	avs.time_base.den = C.int(timeBase.Den())
}

// PrivData Return priv_data
func (avs *AvStream) PrivData() unsafe.Pointer {
	return avs.priv_data
}

// SetPrivData Set priv_data
func (avs *AvStream) SetPrivData(pData unsafe.Pointer) {
	avs.priv_data = pData
}

// func (avs *AvStream) RecommendedEncoderConfiguration() string {
// 	return C.GoString(avs.recommended_encoder_configuration)
// }

// Discard Return discard
func (avs *AvStream) Discard() AvDiscard {
	return AvDiscard(avs.discard)
}

// SetDiscard Set discard
func (avs *AvStream) SetDiscard(discard AvDiscard) {
	avs.discard = C.enum_AVDiscard(discard)
}

// NeedParsing Return need_parsing
func (avs *AvStream) NeedParsing() AvStreamParseType {
	return AvStreamParseType(avs.need_parsing)
}

// CodecInfoNbFrames Return codec_info_nb_frames
func (avs *AvStream) CodecInfoNbFrames() int {
	return int(avs.codec_info_nb_frames)
}

// Disposition Return disposition
func (avs *AvStream) Disposition() int {
	return int(avs.disposition)
}

// SetDisposition Set disposition
func (avs *AvStream) SetDisposition(disposition int) {
	avs.disposition = C.int(disposition)
}

// EventFlags Return event_flags
func (avs *AvStream) EventFlags() int {
	return int(avs.event_flags)
}

// ID Return id
func (avs *AvStream) ID() int {
	return int(avs.id)
}

// Index Return index
func (avs *AvStream) Index() int {
	return int(avs.index)
}

// InjectGlobalSideData Return inject_global_side_data
func (avs *AvStream) InjectGlobalSideData() int {
	return int(avs.inject_global_side_data)
}

// LastIPDuration Return last_IP_duration
func (avs *AvStream) LastIPDuration() int {
	return int(avs.last_IP_duration)
}

// NbDecodedFrames Return nb_decoded_frames
func (avs *AvStream) NbDecodedFrames() int {
	return int(avs.nb_decoded_frames)
}

// NbIndexEntries Return nb_index_entries
func (avs *AvStream) NbIndexEntries() int {
	return int(avs.nb_index_entries)
}

// NbSideData Return nb_side_data
func (avs *AvStream) NbSideData() int {
	return int(avs.nb_side_data)
}

// ProbePackets Return probe_packets
func (avs *AvStream) ProbePackets() int {
	return int(avs.probe_packets)
}

// PtsWrapBehavior Return pts_wrap_behavior
func (avs *AvStream) PtsWrapBehavior() int {
	return int(avs.pts_wrap_behavior)
}

// RequestProbe Return request_probe
func (avs *AvStream) RequestProbe() int {
	return int(avs.request_probe)
}

// SkipSamples Return skip_samples
func (avs *AvStream) SkipSamples() int {
	return int(avs.skip_samples)
}

// SkipToKeyframe Return skip_to_keyframe
func (avs *AvStream) SkipToKeyframe() int {
	return int(avs.skip_to_keyframe)
}

// StreamIdentifier Return stream_identifier
func (avs *AvStream) StreamIdentifier() int {
	return int(avs.stream_identifier)
}

// UpdateInitialDurationsDone Return update_initial_durations_done
func (avs *AvStream) UpdateInitialDurationsDone() int {
	return int(avs.update_initial_durations_done)
}

// CurDts Return cur_dts
func (avs *AvStream) CurDts() int64 {
	return int64(avs.cur_dts)
}

// Duration Return duration
func (avs *AvStream) Duration() int64 {
	return int64(avs.duration)
}

// SetDuration Set duration
func (avs *AvStream) SetDuration(dur int64) {
	avs.duration = C.int64_t(dur)
}

// func (avs *AvStream) FirstDiscardSample() int64 {
// 	return int64(avs.first_discard_sample)
// }

// FirstDts Return first_dts
func (avs *AvStream) FirstDts() int64 {
	return int64(avs.first_dts)
}

// InterleaverChunkDuration Return interleaver_chunk_duration
func (avs *AvStream) InterleaverChunkDuration() int64 {
	return int64(avs.interleaver_chunk_duration)
}

// InterleaverChunkSize Return interleaver_chunk_size
func (avs *AvStream) InterleaverChunkSize() int64 {
	return int64(avs.interleaver_chunk_size)
}

// func (avs *AvStream) LastDiscardSample() int64 {
// 	return int64(avs.last_discard_sample)
// }

// LastDtsForOrderCheck Return last_dts_for_order_check
func (avs *AvStream) LastDtsForOrderCheck() int64 {
	return int64(avs.last_dts_for_order_check)
}

// LastIPPts Return last_IP_pts
func (avs *AvStream) LastIPPts() int64 {
	return int64(avs.last_IP_pts)
}

// MuxTsOffset Return mux_ts_offset
func (avs *AvStream) MuxTsOffset() int64 {
	return int64(avs.mux_ts_offset)
}

// NbFrames Return nb_frames
func (avs *AvStream) NbFrames() int64 {
	return int64(avs.nb_frames)
}

// SetNbFrames Set nb_frames
func (avs *AvStream) SetNbFrames(nbframes int64) {
	avs.nb_frames = C.int64_t(nbframes)
}

// PtsBuffer Return pts_buffer[0]
func (avs *AvStream) PtsBuffer() int64 {
	return int64(avs.pts_buffer[0])
}

// PtsReorderError Return pts_reorder_error[0]
func (avs *AvStream) PtsReorderError() int64 {
	return int64(avs.pts_reorder_error[0])
}

// PtsWrapReference Return pts_wrap_reference
func (avs *AvStream) PtsWrapReference() int64 {
	return int64(avs.pts_wrap_reference)
}

// func (avs *AvStream) StartSkipSamples() int64 {
// 	return int64(avs.start_skip_samples)
// }

// StartTime Return start_time
func (avs *AvStream) StartTime() int64 {
	return int64(avs.start_time)
}

// SetStartTime Set start_time
func (avs *AvStream) SetStartTime(starttime int64) {
	avs.start_time = C.int64_t(starttime)
}

// Parser Return parser
func (avs *AvStream) Parser() *AvCodecParserContext {
	return (*AvCodecParserContext)(unsafe.Pointer(avs.parser))
}

// LastInPacketBuffer Return last_in_packet_buffer
func (avs *AvStream) LastInPacketBuffer() *AvPacketList {
	return (*AvPacketList)(unsafe.Pointer(avs.last_in_packet_buffer))
}

// func (avs *AvStream) PrivPts() *FFFrac {
// 	return (*FFFrac)(unsafe.Pointer(avs.priv_pts))
// }

// DtsMisordered Return dts_misordered
func (avs *AvStream) DtsMisordered() uint8 {
	return uint8(avs.dts_misordered)
}

// DtsOrdered Return dts_ordered
func (avs *AvStream) DtsOrdered() uint8 {
	return uint8(avs.dts_ordered)
}

// PtsReorderErrorCount Return pts_reorder_error_count[0]
func (avs *AvStream) PtsReorderErrorCount() uint8 {
	return uint8(avs.pts_reorder_error_count[0])
}

// IndexEntriesAllocatedSize Return index_entries_allocated_size
func (avs *AvStream) IndexEntriesAllocatedSize() uint {
	return uint(avs.index_entries_allocated_size)
}

// Free ...
func (avs *AvStream) Free() {
	C.av_freep(unsafe.Pointer(avs))
}
