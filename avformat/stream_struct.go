// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package avformat

//#cgo pkg-config: libavformat
//#include <libavformat/avformat.h>
import "C"
import (
	"unsafe"

	"github.com/xueqing/goav/avcodec"
	"github.com/xueqing/goav/avutil"
)

// CodecParameters Return codecpar
func (avs *Stream) CodecParameters() *avcodec.AvCodecParameters {
	return (*avcodec.AvCodecParameters)(unsafe.Pointer(avs.codecpar))
}

// Codec Return codec
func (avs *Stream) Codec() *CodecContext {
	return (*CodecContext)(unsafe.Pointer(avs.codec))
}

// Metadata Return metadata
func (avs *Stream) Metadata() *avutil.Dictionary {
	return (*avutil.Dictionary)(unsafe.Pointer(avs.metadata))
}

// IndexEntries Return index_entries
func (avs *Stream) IndexEntries() *AvIndexEntry {
	return (*AvIndexEntry)(unsafe.Pointer(avs.index_entries))
}

// AttachedPic Return attached_pic
func (avs *Stream) AttachedPic() avcodec.Packet {
	return *fromCPacket(&avs.attached_pic)
}

// SideData Return side_data
func (avs *Stream) SideData() *AvPacketSideData {
	return (*AvPacketSideData)(unsafe.Pointer(avs.side_data))
}

// ProbeData Return probe_data
func (avs *Stream) ProbeData() AvProbeData {
	return AvProbeData(avs.probe_data)
}

// AvgFrameRate Return avg_frame_rate
func (avs *Stream) AvgFrameRate() avcodec.Rational {
	return newRational(avs.avg_frame_rate)
}

// func (avs *Stream) DisplayAspectRatio() *Rational {
// 	return (*Rational)(unsafe.Pointer(avs.display_aspect_ratio))
// }

// RFrameRate Return r_frame_rate
func (avs *Stream) RFrameRate() avcodec.Rational {
	return newRational(avs.r_frame_rate)
}

// SampleAspectRatio Return sample_aspect_ratio
func (avs *Stream) SampleAspectRatio() avcodec.Rational {
	return newRational(avs.sample_aspect_ratio)
}

// TimeBase Return time_base
func (avs *Stream) TimeBase() avcodec.Rational {
	return newRational(avs.time_base)
}

// func (avs *Stream) RecommendedEncoderConfiguration() string {
// 	return C.GoString(avs.recommended_encoder_configuration)
// }

// Discard Return discard
func (avs *Stream) Discard() AvDiscard {
	return AvDiscard(avs.discard)
}

// NeedParsing Return need_parsing
func (avs *Stream) NeedParsing() AvStreamParseType {
	return AvStreamParseType(avs.need_parsing)
}

// CodecInfoNbFrames Return codec_info_nb_frames
func (avs *Stream) CodecInfoNbFrames() int {
	return int(avs.codec_info_nb_frames)
}

// Disposition Return disposition
func (avs *Stream) Disposition() int {
	return int(avs.disposition)
}

// EventFlags Return event_flags
func (avs *Stream) EventFlags() int {
	return int(avs.event_flags)
}

// ID Return id
func (avs *Stream) ID() int {
	return int(avs.id)
}

// Index Return index
func (avs *Stream) Index() int {
	return int(avs.index)
}

// InjectGlobalSideData Return inject_global_side_data
func (avs *Stream) InjectGlobalSideData() int {
	return int(avs.inject_global_side_data)
}

// LastIPDuration Return last_IP_duration
func (avs *Stream) LastIPDuration() int {
	return int(avs.last_IP_duration)
}

// NbDecodedFrames Return nb_decoded_frames
func (avs *Stream) NbDecodedFrames() int {
	return int(avs.nb_decoded_frames)
}

// NbIndexEntries Return nb_index_entries
func (avs *Stream) NbIndexEntries() int {
	return int(avs.nb_index_entries)
}

// NbSideData Return nb_side_data
func (avs *Stream) NbSideData() int {
	return int(avs.nb_side_data)
}

// ProbePackets Return probe_packets
func (avs *Stream) ProbePackets() int {
	return int(avs.probe_packets)
}

// PtsWrapBehavior Return pts_wrap_behavior
func (avs *Stream) PtsWrapBehavior() int {
	return int(avs.pts_wrap_behavior)
}

// RequestProbe Return request_probe
func (avs *Stream) RequestProbe() int {
	return int(avs.request_probe)
}

// SkipSamples Return skip_samples
func (avs *Stream) SkipSamples() int {
	return int(avs.skip_samples)
}

// SkipToKeyframe Return skip_to_keyframe
func (avs *Stream) SkipToKeyframe() int {
	return int(avs.skip_to_keyframe)
}

// StreamIdentifier Return stream_identifier
func (avs *Stream) StreamIdentifier() int {
	return int(avs.stream_identifier)
}

// UpdateInitialDurationsDone Return update_initial_durations_done
func (avs *Stream) UpdateInitialDurationsDone() int {
	return int(avs.update_initial_durations_done)
}

// CurDts Return cur_dts
func (avs *Stream) CurDts() int64 {
	return int64(avs.cur_dts)
}

// Duration Return duration
func (avs *Stream) Duration() int64 {
	return int64(avs.duration)
}

// func (avs *Stream) FirstDiscardSample() int64 {
// 	return int64(avs.first_discard_sample)
// }

// FirstDts Return first_dts
func (avs *Stream) FirstDts() int64 {
	return int64(avs.first_dts)
}

// InterleaverChunkDuration Return interleaver_chunk_duration
func (avs *Stream) InterleaverChunkDuration() int64 {
	return int64(avs.interleaver_chunk_duration)
}

// InterleaverChunkSize Return interleaver_chunk_size
func (avs *Stream) InterleaverChunkSize() int64 {
	return int64(avs.interleaver_chunk_size)
}

// func (avs *Stream) LastDiscardSample() int64 {
// 	return int64(avs.last_discard_sample)
// }

// LastDtsForOrderCheck Return last_dts_for_order_check
func (avs *Stream) LastDtsForOrderCheck() int64 {
	return int64(avs.last_dts_for_order_check)
}

// LastIPPts Return last_IP_pts
func (avs *Stream) LastIPPts() int64 {
	return int64(avs.last_IP_pts)
}

// MuxTsOffset Return mux_ts_offset
func (avs *Stream) MuxTsOffset() int64 {
	return int64(avs.mux_ts_offset)
}

// NbFrames Return nb_frames
func (avs *Stream) NbFrames() int64 {
	return int64(avs.nb_frames)
}

// PtsBuffer Return pts_buffer[0]
func (avs *Stream) PtsBuffer() int64 {
	return int64(avs.pts_buffer[0])
}

// PtsReorderError Return pts_reorder_error[0]
func (avs *Stream) PtsReorderError() int64 {
	return int64(avs.pts_reorder_error[0])
}

// PtsWrapReference Return pts_wrap_reference
func (avs *Stream) PtsWrapReference() int64 {
	return int64(avs.pts_wrap_reference)
}

// func (avs *Stream) StartSkipSamples() int64 {
// 	return int64(avs.start_skip_samples)
// }

// StartTime Return start_time
func (avs *Stream) StartTime() int64 {
	return int64(avs.start_time)
}

// Parser Return parser
func (avs *Stream) Parser() *CodecParserContext {
	return (*CodecParserContext)(unsafe.Pointer(avs.parser))
}

// LastInPacketBuffer Return last_in_packet_buffer
func (avs *Stream) LastInPacketBuffer() *AvPacketList {
	return (*AvPacketList)(unsafe.Pointer(avs.last_in_packet_buffer))
}

// func (avs *Stream) PrivPts() *FFFrac {
// 	return (*FFFrac)(unsafe.Pointer(avs.priv_pts))
// }

// DtsMisordered Return dts_misordered
func (avs *Stream) DtsMisordered() uint8 {
	return uint8(avs.dts_misordered)
}

// DtsOrdered Return dts_ordered
func (avs *Stream) DtsOrdered() uint8 {
	return uint8(avs.dts_ordered)
}

// PtsReorderErrorCount Return pts_reorder_error_count[0]
func (avs *Stream) PtsReorderErrorCount() uint8 {
	return uint8(avs.pts_reorder_error_count[0])
}

// IndexEntriesAllocatedSize Return index_entries_allocated_size
func (avs *Stream) IndexEntriesAllocatedSize() uint {
	return uint(avs.index_entries_allocated_size)
}

// Free ...
func (avs *Stream) Free() {
	C.av_freep(unsafe.Pointer(avs))
}
