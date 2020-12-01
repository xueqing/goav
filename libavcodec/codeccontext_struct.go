// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package libavcodec

//#cgo pkg-config: libavcodec
//#include <libavcodec/avcodec.h>
import "C"

import (
	"reflect"
	"unsafe"
)

// Release ...
func (cctx *AvCodecContext) Release() {
	C.avcodec_close((*C.struct_AVCodecContext)(unsafe.Pointer(cctx)))
	C.av_freep(unsafe.Pointer(cctx))
}

// LogLevelOffset Return log_level_offset
func (cctx *AvCodecContext) LogLevelOffset() int {
	return int(cctx.log_level_offset)
}

// CodecType Return codec_type
func (cctx *AvCodecContext) CodecType() AvMediaType {
	return (AvMediaType)(cctx.codec_type)
}

// SetCodecType Set codec_type
func (cctx *AvCodecContext) SetCodecType(ctype AvMediaType) {
	cctx.codec_type = C.enum_AVMediaType(ctype)
}

// Codec Return codec
func (cctx *AvCodecContext) Codec() *AvCodec {
	return (*AvCodec)(cctx.codec)
}

// CodecID Return codec_id
func (cctx *AvCodecContext) CodecID() AvCodecID {
	return (AvCodecID)(cctx.codec_id)
}

// SetCodecID Set codec_id
func (cctx *AvCodecContext) SetCodecID(codecID AvCodecID) {
	cctx.codec_id = C.enum_AVCodecID(codecID)
}

// ExtraData Return extradata
func (cctx *AvCodecContext) ExtraData() []byte {
	header := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(cctx.extradata)),
		Len:  int(cctx.extradata_size),
		Cap:  int(cctx.extradata_size),
	}

	return *((*[]byte)(unsafe.Pointer(&header)))
}

// SetExtraData Return data
func (cctx *AvCodecContext) SetExtraData(data []byte) {
	header := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	cctx.extradata = (*C.uint8_t)(unsafe.Pointer(header.Data))
	cctx.extradata_size = C.int(header.Len)
}

// BitRate Return bit_rate
func (cctx *AvCodecContext) BitRate() int {
	return int(cctx.bit_rate)
}

// SetBitRate Set bit_rate
func (cctx *AvCodecContext) SetBitRate(br int64) {
	cctx.bit_rate = C.int64_t(br)
}

// BitRateTolerance Return bit_rate_tolerance
func (cctx *AvCodecContext) BitRateTolerance() int {
	return int(cctx.bit_rate_tolerance)
}

// GlobalQuality Return global_quality
func (cctx *AvCodecContext) GlobalQuality() int {
	return int(cctx.global_quality)
}

// CompressionLevel Return compression_level
func (cctx *AvCodecContext) CompressionLevel() int {
	return int(cctx.compression_level)
}

// Flags Return flags
func (cctx *AvCodecContext) Flags() int {
	return int(cctx.flags)
}

// SetFlags Set flags
func (cctx *AvCodecContext) SetFlags(flags int) {
	cctx.flags = C.int(flags)
}

// Flags2 Return flags2
func (cctx *AvCodecContext) Flags2() int {
	return int(cctx.flags2)
}

// ExtradataSize Return extradata_size
func (cctx *AvCodecContext) ExtradataSize() int {
	return int(cctx.extradata_size)
}

// TimeBase Return time_base
func (cctx *AvCodecContext) TimeBase() AvRational {
	return NewAvRational(int(cctx.time_base.num), int(cctx.time_base.den))
}

// SetTimebase Set time_base
func (cctx *AvCodecContext) SetTimebase(timeBase AvRational) {
	cctx.time_base.num = C.int(timeBase.Num())
	cctx.time_base.den = C.int(timeBase.Den())
}

// TicksPerFrame Return ticks_per_frame
func (cctx *AvCodecContext) TicksPerFrame() int {
	return int(cctx.ticks_per_frame)
}

// Delay Return delay
func (cctx *AvCodecContext) Delay() int {
	return int(cctx.delay)
}

// Width Return width
func (cctx *AvCodecContext) Width() int {
	return int(cctx.width)
}

// SetWidth Set width
func (cctx *AvCodecContext) SetWidth(w int) {
	cctx.width = C.int(w)
}

// Height Return height
func (cctx *AvCodecContext) Height() int {
	return int(cctx.height)
}

// SetHeight Set height
func (cctx *AvCodecContext) SetHeight(h int) {
	cctx.height = C.int(h)
}

// CodedWidth Return coded_width
func (cctx *AvCodecContext) CodedWidth() int {
	return int(cctx.coded_width)
}

// CodedHeight Return coded_height
func (cctx *AvCodecContext) CodedHeight() int {
	return int(cctx.coded_height)
}

// GopSize Return gop_size
func (cctx *AvCodecContext) GopSize() int {
	return int(cctx.gop_size)
}

// PixFmt Return pix_fmt
func (cctx *AvCodecContext) PixFmt() AvPixelFormat {
	return (AvPixelFormat)(cctx.pix_fmt)
}

// SetPixelFormat Set pix_fmt
func (cctx *AvCodecContext) SetPixelFormat(fmt AvPixelFormat) {
	cctx.pix_fmt = C.enum_AVPixelFormat(C.int(fmt))
}

// MaxBFrames Return max_b_frames
func (cctx *AvCodecContext) MaxBFrames() int {
	return int(cctx.max_b_frames)
}

// BQuantFactor Return b_quant_factor
func (cctx *AvCodecContext) BQuantFactor() float64 {
	return float64(cctx.b_quant_factor)
}

// BFrameStrategy Return b_frame_strategy
func (cctx *AvCodecContext) BFrameStrategy() int {
	return int(cctx.b_frame_strategy)
}

// BQuantOffset Return b_quant_offset
func (cctx *AvCodecContext) BQuantOffset() float64 {
	return float64(cctx.b_quant_offset)
}

// HasBFrames Return has_b_frames
func (cctx *AvCodecContext) HasBFrames() int {
	return int(cctx.has_b_frames)
}

// MpegQuant Return mpeg_quant
func (cctx *AvCodecContext) MpegQuant() int {
	return int(cctx.mpeg_quant)
}

// IQuantFactor Return i_quant_factor
func (cctx *AvCodecContext) IQuantFactor() float64 {
	return float64(cctx.i_quant_factor)
}

// IQuantOffset Return i_quant_offset
func (cctx *AvCodecContext) IQuantOffset() float64 {
	return float64(cctx.i_quant_offset)
}

// LumiMasking Return lumi_masking
func (cctx *AvCodecContext) LumiMasking() float64 {
	return float64(cctx.lumi_masking)
}

// TemporalCplxMasking Return temporal_cplx_masking
func (cctx *AvCodecContext) TemporalCplxMasking() float64 {
	return float64(cctx.temporal_cplx_masking)
}

// SpatialCplxMasking Return spatial_cplx_masking
func (cctx *AvCodecContext) SpatialCplxMasking() float64 {
	return float64(cctx.spatial_cplx_masking)
}

// PMasking Return p_masking
func (cctx *AvCodecContext) PMasking() float64 {
	return float64(cctx.p_masking)
}

// DarkMasking Return dark_masking
func (cctx *AvCodecContext) DarkMasking() float64 {
	return float64(cctx.dark_masking)
}

// SliceCount Return slice_count
func (cctx *AvCodecContext) SliceCount() int {
	return int(cctx.slice_count)
}

// PredictionMethod Return prediction_method
func (cctx *AvCodecContext) PredictionMethod() int {
	return int(cctx.prediction_method)
}

// SampleAspectRatio Return sample_aspect_ratio
func (cctx *AvCodecContext) SampleAspectRatio() AvRational {
	return NewAvRational(int(cctx.sample_aspect_ratio.num), int(cctx.sample_aspect_ratio.den))
}

// SetSampleAspectRatio Set sample_aspect_ratio
func (cctx *AvCodecContext) SetSampleAspectRatio(sampleAspectRatio AvRational) {
	cctx.sample_aspect_ratio.num = C.int(sampleAspectRatio.Num())
	cctx.sample_aspect_ratio.den = C.int(sampleAspectRatio.Den())
}

// MeCmp Return me_cmp
func (cctx *AvCodecContext) MeCmp() int {
	return int(cctx.me_cmp)
}

// MeSubCmp Return me_sub_cmp
func (cctx *AvCodecContext) MeSubCmp() int {
	return int(cctx.me_sub_cmp)
}

// MbCmp Return mb_cmp
func (cctx *AvCodecContext) MbCmp() int {
	return int(cctx.mb_cmp)
}

// IldctCmp Return ildct_cmp
func (cctx *AvCodecContext) IldctCmp() int {
	return int(cctx.ildct_cmp)
}

// DiaSize Return dia_size
func (cctx *AvCodecContext) DiaSize() int {
	return int(cctx.dia_size)
}

// LastPredictorCount Return last_predictor_count
func (cctx *AvCodecContext) LastPredictorCount() int {
	return int(cctx.last_predictor_count)
}

// PreMe Return pre_me
func (cctx *AvCodecContext) PreMe() int {
	return int(cctx.pre_me)
}

// MePreCmp Return me_pre_cmp
func (cctx *AvCodecContext) MePreCmp() int {
	return int(cctx.me_pre_cmp)
}

// PreDiaSize Return pre_dia_size
func (cctx *AvCodecContext) PreDiaSize() int {
	return int(cctx.pre_dia_size)
}

// MeSubpelQuality Return me_subpel_quality
func (cctx *AvCodecContext) MeSubpelQuality() int {
	return int(cctx.me_subpel_quality)
}

// MeRange Return me_range
func (cctx *AvCodecContext) MeRange() int {
	return int(cctx.me_range)
}

// SetMeRange Set me_range
func (cctx *AvCodecContext) SetMeRange(r int) {
	cctx.me_range = C.int(r)
}

// SliceFlags Return slice_flags
func (cctx *AvCodecContext) SliceFlags() int {
	return int(cctx.slice_flags)
}

// MbDecision Return mb_decision
func (cctx *AvCodecContext) MbDecision() int {
	return int(cctx.mb_decision)
}

// ScenechangeThreshold Return scenechange_threshold
func (cctx *AvCodecContext) ScenechangeThreshold() int {
	return int(cctx.scenechange_threshold)
}

// NoiseReduction Return noise_reduction
func (cctx *AvCodecContext) NoiseReduction() int {
	return int(cctx.noise_reduction)
}

// IntraDcPrecision Return intra_dc_precision
func (cctx *AvCodecContext) IntraDcPrecision() int {
	return int(cctx.intra_dc_precision)
}

// SkipTop Return skip_top
func (cctx *AvCodecContext) SkipTop() int {
	return int(cctx.skip_top)
}

// SkipBottom Return skip_bottom
func (cctx *AvCodecContext) SkipBottom() int {
	return int(cctx.skip_bottom)
}

// MbLmin Return mb_lmin
func (cctx *AvCodecContext) MbLmin() int {
	return int(cctx.mb_lmin)
}

// MbLmax Return mb_lmax
func (cctx *AvCodecContext) MbLmax() int {
	return int(cctx.mb_lmax)
}

// MePenaltyCompensation Return me_penalty_compensation
func (cctx *AvCodecContext) MePenaltyCompensation() int {
	return int(cctx.me_penalty_compensation)
}

// BidirRefine Return bidir_refine
func (cctx *AvCodecContext) BidirRefine() int {
	return int(cctx.bidir_refine)
}

// BrdScale Return brd_scale
func (cctx *AvCodecContext) BrdScale() int {
	return int(cctx.brd_scale)
}

// KeyintMin Return keyint_min
func (cctx *AvCodecContext) KeyintMin() int {
	return int(cctx.keyint_min)
}

// Refs Return refs
func (cctx *AvCodecContext) Refs() int {
	return int(cctx.refs)
}

// Chromaoffset Return chromaoffset
func (cctx *AvCodecContext) Chromaoffset() int {
	return int(cctx.chromaoffset)
}

// Mv0Threshold Return mv0_threshold
func (cctx *AvCodecContext) Mv0Threshold() int {
	return int(cctx.mv0_threshold)
}

// BSensitivity Return b_sensitivity
func (cctx *AvCodecContext) BSensitivity() int {
	return int(cctx.b_sensitivity)
}

// DebugMv Return debug_mv
func (cctx *AvCodecContext) DebugMv() int {
	return int(cctx.debug_mv)
}

// ColorPrimaries Return color_primaries
func (cctx *AvCodecContext) ColorPrimaries() AvColorPrimaries {
	return (AvColorPrimaries)(cctx.color_primaries)
}

// ColorTrc Return color_trc
func (cctx *AvCodecContext) ColorTrc() AvColorTransferCharacteristic {
	return (AvColorTransferCharacteristic)(cctx.color_trc)
}

// Colorspace Return colorspace
func (cctx *AvCodecContext) Colorspace() AvColorSpace {
	return (AvColorSpace)(cctx.colorspace)
}

// ColorRange Return color_range
func (cctx *AvCodecContext) ColorRange() AvColorRange {
	return (AvColorRange)(cctx.color_range)
}

// ChromaSampleLocation Return chroma_sample_location
func (cctx *AvCodecContext) ChromaSampleLocation() AvChromaLocation {
	return (AvChromaLocation)(cctx.chroma_sample_location)
}

// Slices Return slices
func (cctx *AvCodecContext) Slices() int {
	return int(cctx.slices)
}

// FieldOrder Return field_order
func (cctx *AvCodecContext) FieldOrder() AvFieldOrder {
	return (AvFieldOrder)(cctx.field_order)
}

// SampleRate Return sample_rate
func (cctx *AvCodecContext) SampleRate() int {
	return int(cctx.sample_rate)
}

// SetSampleRate Set sample_rate
func (cctx *AvCodecContext) SetSampleRate(sampleRate int) {
	cctx.sample_rate = C.int(sampleRate)
}

// Channels Return channels
func (cctx *AvCodecContext) Channels() int {
	return int(cctx.channels)
}

// SetChannels Set channels
func (cctx *AvCodecContext) SetChannels(nc int) {
	cctx.channels = C.int(nc)
}

// SampleFmt Return sample_fmt
func (cctx *AvCodecContext) SampleFmt() AvSampleFormat {
	return (AvSampleFormat)(cctx.sample_fmt)
}

// SetSampleFmt Set sample_fmt
func (cctx *AvCodecContext) SetSampleFmt(sf AvSampleFormat) {
	cctx.sample_fmt = C.enum_AVSampleFormat(sf)
}

// FrameSize Return frame_size
func (cctx *AvCodecContext) FrameSize() int {
	return int(cctx.frame_size)
}

// FrameNumber Return frame_number
func (cctx *AvCodecContext) FrameNumber() int {
	return int(cctx.frame_number)
}

// BlockAlign Return block_align
func (cctx *AvCodecContext) BlockAlign() int {
	return int(cctx.block_align)
}

// Cutoff Return cutoff
func (cctx *AvCodecContext) Cutoff() int {
	return int(cctx.cutoff)
}

// ChannelLayout Return channel_layout
func (cctx *AvCodecContext) ChannelLayout() uint64 {
	return uint64(cctx.channel_layout)
}

// SetChannelLayout Set channel_layout
func (cctx *AvCodecContext) SetChannelLayout(cl uint64) {
	cctx.channel_layout = C.uint64_t(cl)
}

// RequestChannelLayout Return request_channel_layout
func (cctx *AvCodecContext) RequestChannelLayout() uint64 {
	return uint64(cctx.request_channel_layout)
}

// AudioServiceType Return audio_service_type
func (cctx *AvCodecContext) AudioServiceType() AvAudioServiceType {
	return (AvAudioServiceType)(cctx.audio_service_type)
}

// RequestSampleFmt Return request_sample_fmt
func (cctx *AvCodecContext) RequestSampleFmt() AvSampleFormat {
	return (AvSampleFormat)(cctx.request_sample_fmt)
}

// RefcountedFrames Return refcounted_frames
func (cctx *AvCodecContext) RefcountedFrames() int {
	return int(cctx.refcounted_frames)
}

// Qcompress Return qcompress
func (cctx *AvCodecContext) Qcompress() float64 {
	return float64(cctx.qcompress)
}

// SetQCompress Set qcompress
func (cctx *AvCodecContext) SetQCompress(v float32) {
	cctx.qcompress = C.float(v)
}

// Qblur Return qblur
func (cctx *AvCodecContext) Qblur() float64 {
	return float64(cctx.qblur)
}

// Qmin Return qmin
func (cctx *AvCodecContext) Qmin() int {
	return int(cctx.qmin)
}

// SetQMin Set qmin
func (cctx *AvCodecContext) SetQMin(v int) {
	cctx.qmin = C.int(v)
}

// Qmax Return qmax
func (cctx *AvCodecContext) Qmax() int {
	return int(cctx.qmax)
}

// SetQMax Set qmax
func (cctx *AvCodecContext) SetQMax(v int) {
	cctx.qmax = C.int(v)
}

// MaxQdiff Return max_qdiff
func (cctx *AvCodecContext) MaxQdiff() int {
	return int(cctx.max_qdiff)
}

// SetMaxQDiff Set max_qdiff
func (cctx *AvCodecContext) SetMaxQDiff(v int) {
	cctx.max_qdiff = C.int(v)
}

// RcBufferSize Return rc_buffer_size
func (cctx *AvCodecContext) RcBufferSize() int {
	return int(cctx.rc_buffer_size)
}

// RcOverrideCount Return rc_override_count
func (cctx *AvCodecContext) RcOverrideCount() int {
	return int(cctx.rc_override_count)
}

// RcMaxRate Return rc_max_rate
func (cctx *AvCodecContext) RcMaxRate() int {
	return int(cctx.rc_max_rate)
}

// RcMinRate Return rc_min_rate
func (cctx *AvCodecContext) RcMinRate() int {
	return int(cctx.rc_min_rate)
}

// RcMaxAvailableVbvUse Return rc_max_available_vbv_use
func (cctx *AvCodecContext) RcMaxAvailableVbvUse() float64 {
	return float64(cctx.rc_max_available_vbv_use)
}

// RcMinVbvOverflowUse Return rc_min_vbv_overflow_use
func (cctx *AvCodecContext) RcMinVbvOverflowUse() float64 {
	return float64(cctx.rc_min_vbv_overflow_use)
}

// RcInitialBufferOccupancy Return rc_initial_buffer_occupancy
func (cctx *AvCodecContext) RcInitialBufferOccupancy() int {
	return int(cctx.rc_initial_buffer_occupancy)
}

// CoderType Return coder_type
func (cctx *AvCodecContext) CoderType() int {
	return int(cctx.coder_type)
}

// FrameSkipThreshold Return frame_skip_threshold
func (cctx *AvCodecContext) FrameSkipThreshold() int {
	return int(cctx.frame_skip_threshold)
}

// FrameSkipFactor Return frame_skip_factor
func (cctx *AvCodecContext) FrameSkipFactor() int {
	return int(cctx.frame_skip_factor)
}

// FrameSkipExp Return frame_skip_exp
func (cctx *AvCodecContext) FrameSkipExp() int {
	return int(cctx.frame_skip_exp)
}

// FrameSkipCmp Return frame_skip_cmp
func (cctx *AvCodecContext) FrameSkipCmp() int {
	return int(cctx.frame_skip_cmp)
}

// MinPredictionOrder Return min_prediction_order
func (cctx *AvCodecContext) MinPredictionOrder() int {
	return int(cctx.min_prediction_order)
}

// MaxPredictionOrder Return max_prediction_order
func (cctx *AvCodecContext) MaxPredictionOrder() int {
	return int(cctx.max_prediction_order)
}

// Trellis Return trellis
func (cctx *AvCodecContext) Trellis() int {
	return int(cctx.trellis)
}

// RtpPayloadSize Return rtp_payload_size
func (cctx *AvCodecContext) RtpPayloadSize() int {
	return int(cctx.rtp_payload_size)
}

// MvBits Return mv_bits
func (cctx *AvCodecContext) MvBits() int {
	return int(cctx.mv_bits)
}

// HeaderBits Return header_bits
func (cctx *AvCodecContext) HeaderBits() int {
	return int(cctx.header_bits)
}

// ITexBits Return i_tex_bits
func (cctx *AvCodecContext) ITexBits() int {
	return int(cctx.i_tex_bits)
}

// PTexBits Return p_tex_bits
func (cctx *AvCodecContext) PTexBits() int {
	return int(cctx.p_tex_bits)
}

// ICount Return i_count
func (cctx *AvCodecContext) ICount() int {
	return int(cctx.i_count)
}

// PCount Return p_count
func (cctx *AvCodecContext) PCount() int {
	return int(cctx.p_count)
}

// SkipCount Return skip_count
func (cctx *AvCodecContext) SkipCount() int {
	return int(cctx.skip_count)
}

// MiscBits Return misc_bits
func (cctx *AvCodecContext) MiscBits() int {
	return int(cctx.misc_bits)
}

// FrameBits Return frame_bits
func (cctx *AvCodecContext) FrameBits() int {
	return int(cctx.frame_bits)
}

// WorkaroundBugs Return workaround_bugs
func (cctx *AvCodecContext) WorkaroundBugs() int {
	return int(cctx.workaround_bugs)
}

// StrictStdCompliance Return strict_std_compliance
func (cctx *AvCodecContext) StrictStdCompliance() int {
	return int(cctx.strict_std_compliance)
}

// ErrorConcealment Return error_concealment
func (cctx *AvCodecContext) ErrorConcealment() int {
	return int(cctx.error_concealment)
}

// Debug Return debug
func (cctx *AvCodecContext) Debug() int {
	return int(cctx.debug)
}

// ErrRecognition Return err_recognition
func (cctx *AvCodecContext) ErrRecognition() int {
	return int(cctx.err_recognition)
}

// DctAlgo Return dct_algo
func (cctx *AvCodecContext) DctAlgo() int {
	return int(cctx.dct_algo)
}

// IdctAlgo Return idct_algo
func (cctx *AvCodecContext) IdctAlgo() int {
	return int(cctx.idct_algo)
}

// BitsPerCodedSample Return bits_per_coded_sample
func (cctx *AvCodecContext) BitsPerCodedSample() int {
	return int(cctx.bits_per_coded_sample)
}

// BitsPerRawSample Return bits_per_raw_sample
func (cctx *AvCodecContext) BitsPerRawSample() int {
	return int(cctx.bits_per_raw_sample)
}

// Lowres Return lowres
func (cctx *AvCodecContext) Lowres() int {
	return int(cctx.lowres)
}

// ThreadCount Return thread_count
func (cctx *AvCodecContext) ThreadCount() int {
	return int(cctx.thread_count)
}

// ThreadType Return active_thread_type
func (cctx *AvCodecContext) ThreadType() int {
	return int(cctx.thread_type)
}

// ActiveThreadType Return active_thread_type
func (cctx *AvCodecContext) ActiveThreadType() int {
	return int(cctx.active_thread_type)
}

// ThreadSafeCallbacks Return thread_safe_callbacks
func (cctx *AvCodecContext) ThreadSafeCallbacks() int {
	return int(cctx.thread_safe_callbacks)
}

// NsseWeight Return nsse_weight
func (cctx *AvCodecContext) NsseWeight() int {
	return int(cctx.nsse_weight)
}

// Profile Return profile
func (cctx *AvCodecContext) Profile() int {
	return int(cctx.profile)
}

// Level Return level
func (cctx *AvCodecContext) Level() int {
	return int(cctx.level)
}

// SkipLoopFilter Return skip_loop_filter
func (cctx *AvCodecContext) SkipLoopFilter() AvDiscard {
	return (AvDiscard)(cctx.skip_loop_filter)
}

// SkipIdct Return skip_idct
func (cctx *AvCodecContext) SkipIdct() AvDiscard {
	return (AvDiscard)(cctx.skip_idct)
}

// SkipFrame Return skip_frame
func (cctx *AvCodecContext) SkipFrame() AvDiscard {
	return (AvDiscard)(cctx.skip_frame)
}

// SubtitleHeaderSize Return subtitle_header_size
func (cctx *AvCodecContext) SubtitleHeaderSize() int {
	return int(cctx.subtitle_header_size)
}

// SideDataOnlyPackets Return side_data_only_packets
func (cctx *AvCodecContext) SideDataOnlyPackets() int {
	return int(cctx.side_data_only_packets)
}

// Framerate Return framerate
func (cctx *AvCodecContext) Framerate() AvRational {
	return NewAvRational(int(cctx.framerate.num), int(cctx.framerate.den))
}

// SetFramerate Set framerate
func (cctx *AvCodecContext) SetFramerate(framerate AvRational) {
	cctx.framerate.num = C.int(framerate.Num())
	cctx.framerate.den = C.int(framerate.Den())
}

// CodecDescriptor Return codec_descriptor
func (cctx *AvCodecContext) CodecDescriptor() *AvCodecDescriptor {
	return (*AvCodecDescriptor)(cctx.codec_descriptor)
}

// SubCharencMode Return sub_charenc_mode
func (cctx *AvCodecContext) SubCharencMode() int {
	return int(cctx.sub_charenc_mode)
}

// SkipAlpha Return skip_alpha
func (cctx *AvCodecContext) SkipAlpha() int {
	return int(cctx.skip_alpha)
}

// SeekPreroll Return seek_preroll
func (cctx *AvCodecContext) SeekPreroll() int {
	return int(cctx.seek_preroll)
}

// SetEncodeParams2 ...
func (cctx *AvCodecContext) SetEncodeParams2(width int, height int, pxlFmt AvPixelFormat, hasBframes bool, gopSize int) {
	cctx.width = C.int(width)
	cctx.height = C.int(height)
	// cctx.bit_rate = 1000000
	cctx.gop_size = C.int(gopSize)
	// cctx.max_b_frames = 2
	if hasBframes {
		cctx.has_b_frames = 1
	} else {
		cctx.has_b_frames = 0
	}
	// cctx.extradata = nil
	// cctx.extradata_size = 0
	// cctx.channels = 0
	cctx.pix_fmt = int32(pxlFmt)
	// C.av_opt_set(cctx.priv_data, "preset", "ultrafast", 0)
}
