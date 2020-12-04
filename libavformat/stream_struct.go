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

// Free ...
func (st *AvStream) Free() {
	C.av_freep(unsafe.Pointer(st))
}

// Index Return index
func (st *AvStream) Index() int {
	return int(st.index)
}

// ID Return id
func (st *AvStream) ID() int {
	return int(st.id)
}

// Codec Return codec
func (st *AvStream) Codec() *AvCodecContext {
	return (*AvCodecContext)(unsafe.Pointer(st.codec))
}

// PrivData Return priv_data
func (st *AvStream) PrivData() unsafe.Pointer {
	return st.priv_data
}

// SetPrivData Set priv_data
func (st *AvStream) SetPrivData(pData unsafe.Pointer) {
	st.priv_data = pData
}

// TimeBase Return time_base
func (st *AvStream) TimeBase() libavcodec.AvRational {
	return newAvRational(st.time_base)
}

// SetTimeBase Set time_base
func (st *AvStream) SetTimeBase(timeBase libavcodec.AvRational) {
	st.time_base.num = C.int(timeBase.Num())
	st.time_base.den = C.int(timeBase.Den())
}

// StartTime Return start_time
func (st *AvStream) StartTime() int64 {
	return int64(st.start_time)
}

// SetStartTime Set start_time
func (st *AvStream) SetStartTime(starttime int64) {
	st.start_time = C.int64_t(starttime)
}

// Duration Return duration
func (st *AvStream) Duration() int64 {
	return int64(st.duration)
}

// SetDuration Set duration
func (st *AvStream) SetDuration(dur int64) {
	st.duration = C.int64_t(dur)
}

// NbFrames Return nb_frames
func (st *AvStream) NbFrames() int64 {
	return int64(st.nb_frames)
}

// SetNbFrames Set nb_frames
func (st *AvStream) SetNbFrames(nbframes int64) {
	st.nb_frames = C.int64_t(nbframes)
}

// Disposition Return disposition
func (st *AvStream) Disposition() int {
	return int(st.disposition)
}

// SetDisposition Set disposition
func (st *AvStream) SetDisposition(disposition int) {
	st.disposition = C.int(disposition)
}

// Discard Return discard
func (st *AvStream) Discard() AvDiscard {
	return AvDiscard(st.discard)
}

// SetDiscard Set discard
func (st *AvStream) SetDiscard(discard AvDiscard) {
	st.discard = C.enum_AVDiscard(discard)
}

// SampleAspectRatio Return sample_aspect_ratio
func (st *AvStream) SampleAspectRatio() libavcodec.AvRational {
	return newAvRational(st.sample_aspect_ratio)
}

// Metadata Return metadata
func (st *AvStream) Metadata() *libavutil.AvDictionary {
	return (*libavutil.AvDictionary)(unsafe.Pointer(st.metadata))
}

// AvgFrameRate Return avg_frame_rate
func (st *AvStream) AvgFrameRate() libavcodec.AvRational {
	return newAvRational(st.avg_frame_rate)
}

// AttachedPic Return attached_pic
func (st *AvStream) AttachedPic() libavcodec.AvPacket {
	return *fromCPacket(&st.attached_pic)
}

// SideData Return side_data
func (st *AvStream) SideData() *AvPacketSideData {
	return (*AvPacketSideData)(unsafe.Pointer(st.side_data))
}

// NbSideData Return nb_side_data
func (st *AvStream) NbSideData() int {
	return int(st.nb_side_data)
}

// EventFlags Return event_flags
func (st *AvStream) EventFlags() int {
	return int(st.event_flags)
}

// RFrameRate Return r_frame_rate
func (st *AvStream) RFrameRate() libavcodec.AvRational {
	return newAvRational(st.r_frame_rate)
}

// SetRFrameRate Set r_frame_rate
func (st *AvStream) SetRFrameRate(frameRate libavcodec.AvRational) {
	st.r_frame_rate.num = C.int(frameRate.Num())
	st.r_frame_rate.den = C.int(frameRate.Den())
}

// CodecParameters Return codecpar
func (st *AvStream) CodecParameters() *libavcodec.AvCodecParameters {
	return (*libavcodec.AvCodecParameters)(unsafe.Pointer(st.codecpar))
}

// SetCodecParameters Set codecpar
func (st *AvStream) SetCodecParameters(pCodecPara *libavcodec.AvCodecParameters) {
	st.codecpar = (*C.struct_AVCodecParameters)(unsafe.Pointer(pCodecPara))
}

// FirstDts Return first_dts
func (st *AvStream) FirstDts() int64 {
	return int64(st.first_dts)
}

// CurDts Return cur_dts
func (st *AvStream) CurDts() int64 {
	return int64(st.cur_dts)
}

// LastIPPts Return last_IP_pts
func (st *AvStream) LastIPPts() int64 {
	return int64(st.last_IP_pts)
}

// LastIPDuration Return last_IP_duration
func (st *AvStream) LastIPDuration() int {
	return int(st.last_IP_duration)
}

// ProbePackets Return probe_packets
func (st *AvStream) ProbePackets() int {
	return int(st.probe_packets)
}

// CodecInfoNbFrames Return codec_info_nb_frames
func (st *AvStream) CodecInfoNbFrames() int {
	return int(st.codec_info_nb_frames)
}

// NeedParsing Return need_parsing
func (st *AvStream) NeedParsing() AvStreamParseType {
	return AvStreamParseType(st.need_parsing)
}

// Parser Return parser
func (st *AvStream) Parser() *AvCodecParserContext {
	return (*AvCodecParserContext)(unsafe.Pointer(st.parser))
}

// LastInPacketBuffer Return last_in_packet_buffer
func (st *AvStream) LastInPacketBuffer() *AvPacketList {
	return (*AvPacketList)(unsafe.Pointer(st.last_in_packet_buffer))
}

// ProbeData Return probe_data
func (st *AvStream) ProbeData() AvProbeData {
	return AvProbeData(st.probe_data)
}

// PtsBuffer Return pts_buffer[0]
func (st *AvStream) PtsBuffer() int64 {
	return int64(st.pts_buffer[0])
}

// IndexEntries Return index_entries
func (st *AvStream) IndexEntries() *AvIndexEntry {
	return (*AvIndexEntry)(unsafe.Pointer(st.index_entries))
}

// NbIndexEntries Return nb_index_entries
func (st *AvStream) NbIndexEntries() int {
	return int(st.nb_index_entries)
}

// IndexEntriesAllocatedSize Return index_entries_allocated_size
func (st *AvStream) IndexEntriesAllocatedSize() uint {
	return uint(st.index_entries_allocated_size)
}

// StreamIdentifier Return stream_identifier
func (st *AvStream) StreamIdentifier() int {
	return int(st.stream_identifier)
}

// InterleaverChunkSize Return interleaver_chunk_size
func (st *AvStream) InterleaverChunkSize() int64 {
	return int64(st.interleaver_chunk_size)
}

// InterleaverChunkDuration Return interleaver_chunk_duration
func (st *AvStream) InterleaverChunkDuration() int64 {
	return int64(st.interleaver_chunk_duration)
}

// RequestProbe Return request_probe
func (st *AvStream) RequestProbe() int {
	return int(st.request_probe)
}

// SkipToKeyframe Return skip_to_keyframe
func (st *AvStream) SkipToKeyframe() int {
	return int(st.skip_to_keyframe)
}

// SkipSamples Return skip_samples
func (st *AvStream) SkipSamples() int {
	return int(st.skip_samples)
}

// StartSkipSamples Return start_skip_samples
func (st *AvStream) StartSkipSamples() int64 {
	return int64(st.start_skip_samples)
}

// FirstDiscardSample Return first_discard_sample
func (st *AvStream) FirstDiscardSample() int64 {
	return int64(st.first_discard_sample)
}

// LastDiscardSample Return last_discard_sample
func (st *AvStream) LastDiscardSample() int64 {
	return int64(st.last_discard_sample)
}

// NbDecodedFrames Return nb_decoded_frames
func (st *AvStream) NbDecodedFrames() int {
	return int(st.nb_decoded_frames)
}

// MuxTsOffset Return mux_ts_offset
func (st *AvStream) MuxTsOffset() int64 {
	return int64(st.mux_ts_offset)
}

// PtsWrapReference Return pts_wrap_reference
func (st *AvStream) PtsWrapReference() int64 {
	return int64(st.pts_wrap_reference)
}

// PtsWrapBehavior Return pts_wrap_behavior
func (st *AvStream) PtsWrapBehavior() int {
	return int(st.pts_wrap_behavior)
}

// UpdateInitialDurationsDone Return update_initial_durations_done
func (st *AvStream) UpdateInitialDurationsDone() int {
	return int(st.update_initial_durations_done)
}

// PtsReorderError Return pts_reorder_error[0]
func (st *AvStream) PtsReorderError() int64 {
	return int64(st.pts_reorder_error[0])
}

// PtsReorderErrorCount Return pts_reorder_error_count[0]
func (st *AvStream) PtsReorderErrorCount() uint8 {
	return uint8(st.pts_reorder_error_count[0])
}

// LastDtsForOrderCheck Return last_dts_for_order_check
func (st *AvStream) LastDtsForOrderCheck() int64 {
	return int64(st.last_dts_for_order_check)
}

// DtsOrdered Return dts_ordered
func (st *AvStream) DtsOrdered() uint8 {
	return uint8(st.dts_ordered)
}

// DtsMisordered Return dts_misordered
func (st *AvStream) DtsMisordered() uint8 {
	return uint8(st.dts_misordered)
}

// InjectGlobalSideData Return inject_global_side_data
func (st *AvStream) InjectGlobalSideData() int {
	return int(st.inject_global_side_data)
}

// DisplayAspectRatio Return display_aspect_ratio
func (st *AvStream) DisplayAspectRatio() libavcodec.AvRational {
	return newAvRational(st.display_aspect_ratio)
}
