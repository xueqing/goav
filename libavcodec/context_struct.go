// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package libavcodec

import "C"

// ActiveThreadType Return active_thread_type
func (ctxt *AvCodecContext) ActiveThreadType() int {
	return int(ctxt.active_thread_type)
}

// BFrameStrategy Return b_frame_strategy
func (ctxt *AvCodecContext) BFrameStrategy() int {
	return int(ctxt.b_frame_strategy)
}

// BQuantFactor Return b_quant_factor
func (ctxt *AvCodecContext) BQuantFactor() float64 {
	return float64(ctxt.b_quant_factor)
}

// BQuantOffset Return b_quant_offset
func (ctxt *AvCodecContext) BQuantOffset() float64 {
	return float64(ctxt.b_quant_offset)
}

// BSensitivity Return b_sensitivity
func (ctxt *AvCodecContext) BSensitivity() int {
	return int(ctxt.b_sensitivity)
}

// BidirRefine Return bidir_refine
func (ctxt *AvCodecContext) BidirRefine() int {
	return int(ctxt.bidir_refine)
}

// BitRate Return bit_rate
func (ctxt *AvCodecContext) BitRate() int {
	return int(ctxt.bit_rate)
}

// BitRateTolerance Return bit_rate_tolerance
func (ctxt *AvCodecContext) BitRateTolerance() int {
	return int(ctxt.bit_rate_tolerance)
}

// BitsPerCodedSample Return bits_per_coded_sample
func (ctxt *AvCodecContext) BitsPerCodedSample() int {
	return int(ctxt.bits_per_coded_sample)
}

// BitsPerRawSample Return bits_per_raw_sample
func (ctxt *AvCodecContext) BitsPerRawSample() int {
	return int(ctxt.bits_per_raw_sample)
}

// BlockAlign Return block_align
func (ctxt *AvCodecContext) BlockAlign() int {
	return int(ctxt.block_align)
}

// BrdScale Return brd_scale
func (ctxt *AvCodecContext) BrdScale() int {
	return int(ctxt.brd_scale)
}

// Channels Return channels
func (ctxt *AvCodecContext) Channels() int {
	return int(ctxt.channels)
}

// Chromaoffset Return chromaoffset
func (ctxt *AvCodecContext) Chromaoffset() int {
	return int(ctxt.chromaoffset)
}

// CodedHeight Return coded_height
func (ctxt *AvCodecContext) CodedHeight() int {
	return int(ctxt.coded_height)
}

// CodedWidth Return coded_width
func (ctxt *AvCodecContext) CodedWidth() int {
	return int(ctxt.coded_width)
}

// CoderType Return coder_type
func (ctxt *AvCodecContext) CoderType() int {
	return int(ctxt.coder_type)
}

// CompressionLevel Return compression_level
func (ctxt *AvCodecContext) CompressionLevel() int {
	return int(ctxt.compression_level)
}

// ContextModel Return context_model
func (ctxt *AvCodecContext) ContextModel() int {
	return int(ctxt.context_model)
}

// Cutoff Return cutoff
func (ctxt *AvCodecContext) Cutoff() int {
	return int(ctxt.cutoff)
}

// DarkMasking Return dark_masking
func (ctxt *AvCodecContext) DarkMasking() float64 {
	return float64(ctxt.dark_masking)
}

// DctAlgo Return dct_algo
func (ctxt *AvCodecContext) DctAlgo() int {
	return int(ctxt.dct_algo)
}

// Debug Return debug
func (ctxt *AvCodecContext) Debug() int {
	return int(ctxt.debug)
}

// DebugMv Return debug_mv
func (ctxt *AvCodecContext) DebugMv() int {
	return int(ctxt.debug_mv)
}

// Delay Return delay
func (ctxt *AvCodecContext) Delay() int {
	return int(ctxt.delay)
}

// DiaSize Return dia_size
func (ctxt *AvCodecContext) DiaSize() int {
	return int(ctxt.dia_size)
}

// ErrRecognition Return err_recognition
func (ctxt *AvCodecContext) ErrRecognition() int {
	return int(ctxt.err_recognition)
}

// ErrorConcealment Return error_concealment
func (ctxt *AvCodecContext) ErrorConcealment() int {
	return int(ctxt.error_concealment)
}

// ExtradataSize Return extradata_size
func (ctxt *AvCodecContext) ExtradataSize() int {
	return int(ctxt.extradata_size)
}

// Flags Return flags
func (ctxt *AvCodecContext) Flags() int {
	return int(ctxt.flags)
}

// Flags2 Return flags2
func (ctxt *AvCodecContext) Flags2() int {
	return int(ctxt.flags2)
}

// FrameBits Return frame_bits
func (ctxt *AvCodecContext) FrameBits() int {
	return int(ctxt.frame_bits)
}

// FrameNumber Return frame_number
func (ctxt *AvCodecContext) FrameNumber() int {
	return int(ctxt.frame_number)
}

// FrameSize Return frame_size
func (ctxt *AvCodecContext) FrameSize() int {
	return int(ctxt.frame_size)
}

// FrameSkipCmp Return frame_skip_cmp
func (ctxt *AvCodecContext) FrameSkipCmp() int {
	return int(ctxt.frame_skip_cmp)
}

// FrameSkipExp Return frame_skip_exp
func (ctxt *AvCodecContext) FrameSkipExp() int {
	return int(ctxt.frame_skip_exp)
}

// FrameSkipFactor Return frame_skip_factor
func (ctxt *AvCodecContext) FrameSkipFactor() int {
	return int(ctxt.frame_skip_factor)
}

// FrameSkipThreshold Return frame_skip_threshold
func (ctxt *AvCodecContext) FrameSkipThreshold() int {
	return int(ctxt.frame_skip_threshold)
}

// GlobalQuality Return global_quality
func (ctxt *AvCodecContext) GlobalQuality() int {
	return int(ctxt.global_quality)
}

// GopSize Return gop_size
func (ctxt *AvCodecContext) GopSize() int {
	return int(ctxt.gop_size)
}

// HasBFrames Return has_b_frames
func (ctxt *AvCodecContext) HasBFrames() int {
	return int(ctxt.has_b_frames)
}

// HeaderBits Return header_bits
func (ctxt *AvCodecContext) HeaderBits() int {
	return int(ctxt.header_bits)
}

// Height Return height
func (ctxt *AvCodecContext) Height() int {
	return int(ctxt.height)
}

// SetHeight Set height
func (ctxt *AvCodecContext) SetHeight(h int) {
	ctxt.height = C.int(h)
}

// ICount Return i_count
func (ctxt *AvCodecContext) ICount() int {
	return int(ctxt.i_count)
}

// IQuantFactor Return i_quant_factor
func (ctxt *AvCodecContext) IQuantFactor() float64 {
	return float64(ctxt.i_quant_factor)
}

// IQuantOffset Return i_quant_offset
func (ctxt *AvCodecContext) IQuantOffset() float64 {
	return float64(ctxt.i_quant_offset)
}

// ITexBits Return i_tex_bits
func (ctxt *AvCodecContext) ITexBits() int {
	return int(ctxt.i_tex_bits)
}

// IdctAlgo Return idct_algo
func (ctxt *AvCodecContext) IdctAlgo() int {
	return int(ctxt.idct_algo)
}

// IldctCmp Return ildct_cmp
func (ctxt *AvCodecContext) IldctCmp() int {
	return int(ctxt.ildct_cmp)
}

// IntraDcPrecision Return intra_dc_precision
func (ctxt *AvCodecContext) IntraDcPrecision() int {
	return int(ctxt.intra_dc_precision)
}

// KeyintMin Return keyint_min
func (ctxt *AvCodecContext) KeyintMin() int {
	return int(ctxt.keyint_min)
}

// LastPredictorCount Return last_predictor_count
func (ctxt *AvCodecContext) LastPredictorCount() int {
	return int(ctxt.last_predictor_count)
}

// Level Return level
func (ctxt *AvCodecContext) Level() int {
	return int(ctxt.level)
}

// LogLevelOffset Return log_level_offset
func (ctxt *AvCodecContext) LogLevelOffset() int {
	return int(ctxt.log_level_offset)
}

// Lowres Return lowres
func (ctxt *AvCodecContext) Lowres() int {
	return int(ctxt.lowres)
}

// LumiMasking Return lumi_masking
func (ctxt *AvCodecContext) LumiMasking() float64 {
	return float64(ctxt.lumi_masking)
}

// MaxBFrames Return max_b_frames
func (ctxt *AvCodecContext) MaxBFrames() int {
	return int(ctxt.max_b_frames)
}

// MaxPredictionOrder Return max_prediction_order
func (ctxt *AvCodecContext) MaxPredictionOrder() int {
	return int(ctxt.max_prediction_order)
}

// MaxQdiff Return max_qdiff
func (ctxt *AvCodecContext) MaxQdiff() int {
	return int(ctxt.max_qdiff)
}

// MbCmp Return mb_cmp
func (ctxt *AvCodecContext) MbCmp() int {
	return int(ctxt.mb_cmp)
}

// MbDecision Return mb_decision
func (ctxt *AvCodecContext) MbDecision() int {
	return int(ctxt.mb_decision)
}

// MbLmax Return mb_lmax
func (ctxt *AvCodecContext) MbLmax() int {
	return int(ctxt.mb_lmax)
}

// MbLmin Return mb_lmin
func (ctxt *AvCodecContext) MbLmin() int {
	return int(ctxt.mb_lmin)
}

// MeCmp Return me_cmp
func (ctxt *AvCodecContext) MeCmp() int {
	return int(ctxt.me_cmp)
}

// MePenaltyCompensation Return me_penalty_compensation
func (ctxt *AvCodecContext) MePenaltyCompensation() int {
	return int(ctxt.me_penalty_compensation)
}

// MePreCmp Return me_pre_cmp
func (ctxt *AvCodecContext) MePreCmp() int {
	return int(ctxt.me_pre_cmp)
}

// MeRange Return me_range
func (ctxt *AvCodecContext) MeRange() int {
	return int(ctxt.me_range)
}

// MeSubCmp Return me_sub_cmp
func (ctxt *AvCodecContext) MeSubCmp() int {
	return int(ctxt.me_sub_cmp)
}

// MeSubpelQuality Return me_subpel_quality
func (ctxt *AvCodecContext) MeSubpelQuality() int {
	return int(ctxt.me_subpel_quality)
}

// MinPredictionOrder Return min_prediction_order
func (ctxt *AvCodecContext) MinPredictionOrder() int {
	return int(ctxt.min_prediction_order)
}

// MiscBits Return misc_bits
func (ctxt *AvCodecContext) MiscBits() int {
	return int(ctxt.misc_bits)
}

// MpegQuant Return mpeg_quant
func (ctxt *AvCodecContext) MpegQuant() int {
	return int(ctxt.mpeg_quant)
}

// Mv0Threshold Return mv0_threshold
func (ctxt *AvCodecContext) Mv0Threshold() int {
	return int(ctxt.mv0_threshold)
}

// MvBits Return mv_bits
func (ctxt *AvCodecContext) MvBits() int {
	return int(ctxt.mv_bits)
}

// NoiseReduction Return noise_reduction
func (ctxt *AvCodecContext) NoiseReduction() int {
	return int(ctxt.noise_reduction)
}

// NsseWeight Return nsse_weight
func (ctxt *AvCodecContext) NsseWeight() int {
	return int(ctxt.nsse_weight)
}

// PCount Return p_count
func (ctxt *AvCodecContext) PCount() int {
	return int(ctxt.p_count)
}

// PMasking Return p_masking
func (ctxt *AvCodecContext) PMasking() float64 {
	return float64(ctxt.p_masking)
}

// PTexBits Return p_tex_bits
func (ctxt *AvCodecContext) PTexBits() int {
	return int(ctxt.p_tex_bits)
}

// PreDiaSize Return pre_dia_size
func (ctxt *AvCodecContext) PreDiaSize() int {
	return int(ctxt.pre_dia_size)
}

// PreMe Return pre_me
func (ctxt *AvCodecContext) PreMe() int {
	return int(ctxt.pre_me)
}

// PredictionMethod Return prediction_method
func (ctxt *AvCodecContext) PredictionMethod() int {
	return int(ctxt.prediction_method)
}

// Profile Return profile
func (ctxt *AvCodecContext) Profile() int {
	return int(ctxt.profile)
}

// Qblur Return qblur
func (ctxt *AvCodecContext) Qblur() float64 {
	return float64(ctxt.qblur)
}

// Qcompress Return qcompress
func (ctxt *AvCodecContext) Qcompress() float64 {
	return float64(ctxt.qcompress)
}

// Qmax Return qmax
func (ctxt *AvCodecContext) Qmax() int {
	return int(ctxt.qmax)
}

// Qmin Return qmin
func (ctxt *AvCodecContext) Qmin() int {
	return int(ctxt.qmin)
}

// RcBufferSize Return rc_buffer_size
func (ctxt *AvCodecContext) RcBufferSize() int {
	return int(ctxt.rc_buffer_size)
}

// RcInitialBufferOccupancy Return rc_initial_buffer_occupancy
func (ctxt *AvCodecContext) RcInitialBufferOccupancy() int {
	return int(ctxt.rc_initial_buffer_occupancy)
}

// RcMaxAvailableVbvUse Return rc_max_available_vbv_use
func (ctxt *AvCodecContext) RcMaxAvailableVbvUse() float64 {
	return float64(ctxt.rc_max_available_vbv_use)
}

// RcMaxRate Return rc_max_rate
func (ctxt *AvCodecContext) RcMaxRate() int {
	return int(ctxt.rc_max_rate)
}

// RcMinRate Return rc_min_rate
func (ctxt *AvCodecContext) RcMinRate() int {
	return int(ctxt.rc_min_rate)
}

// RcMinVbvOverflowUse Return rc_min_vbv_overflow_use
func (ctxt *AvCodecContext) RcMinVbvOverflowUse() float64 {
	return float64(ctxt.rc_min_vbv_overflow_use)
}

// RcOverrideCount Return rc_override_count
func (ctxt *AvCodecContext) RcOverrideCount() int {
	return int(ctxt.rc_override_count)
}

// RefcountedFrames Return refcounted_frames
func (ctxt *AvCodecContext) RefcountedFrames() int {
	return int(ctxt.refcounted_frames)
}

// Refs Return refs
func (ctxt *AvCodecContext) Refs() int {
	return int(ctxt.refs)
}

// RtpPayloadSize Return rtp_payload_size
func (ctxt *AvCodecContext) RtpPayloadSize() int {
	return int(ctxt.rtp_payload_size)
}

// SampleRate Return sample_rate
func (ctxt *AvCodecContext) SampleRate() int {
	return int(ctxt.sample_rate)
}

// ScenechangeThreshold Return scenechange_threshold
func (ctxt *AvCodecContext) ScenechangeThreshold() int {
	return int(ctxt.scenechange_threshold)
}

// SeekPreroll Return seek_preroll
func (ctxt *AvCodecContext) SeekPreroll() int {
	return int(ctxt.seek_preroll)
}

// SideDataOnlyPackets Return side_data_only_packets
func (ctxt *AvCodecContext) SideDataOnlyPackets() int {
	return int(ctxt.side_data_only_packets)
}

// SkipAlpha Return skip_alpha
func (ctxt *AvCodecContext) SkipAlpha() int {
	return int(ctxt.skip_alpha)
}

// SkipBottom Return skip_bottom
func (ctxt *AvCodecContext) SkipBottom() int {
	return int(ctxt.skip_bottom)
}

// SkipCount Return skip_count
func (ctxt *AvCodecContext) SkipCount() int {
	return int(ctxt.skip_count)
}

// SkipTop Return skip_top
func (ctxt *AvCodecContext) SkipTop() int {
	return int(ctxt.skip_top)
}

// SliceCount Return slice_count
func (ctxt *AvCodecContext) SliceCount() int {
	return int(ctxt.slice_count)
}

// SliceFlags Return slice_flags
func (ctxt *AvCodecContext) SliceFlags() int {
	return int(ctxt.slice_flags)
}

// Slices Return slices
func (ctxt *AvCodecContext) Slices() int {
	return int(ctxt.slices)
}

// SpatialCplxMasking Return spatial_cplx_masking
func (ctxt *AvCodecContext) SpatialCplxMasking() float64 {
	return float64(ctxt.spatial_cplx_masking)
}

// StrictStdCompliance Return strict_std_compliance
func (ctxt *AvCodecContext) StrictStdCompliance() int {
	return int(ctxt.strict_std_compliance)
}

// SubCharencMode Return sub_charenc_mode
func (ctxt *AvCodecContext) SubCharencMode() int {
	return int(ctxt.sub_charenc_mode)
}

// SubtitleHeaderSize Return subtitle_header_size
func (ctxt *AvCodecContext) SubtitleHeaderSize() int {
	return int(ctxt.subtitle_header_size)
}

// TemporalCplxMasking Return temporal_cplx_masking
func (ctxt *AvCodecContext) TemporalCplxMasking() float64 {
	return float64(ctxt.temporal_cplx_masking)
}

// ThreadCount Return thread_count
func (ctxt *AvCodecContext) ThreadCount() int {
	return int(ctxt.thread_count)
}

// ThreadSafeCallbacks Return thread_safe_callbacks
func (ctxt *AvCodecContext) ThreadSafeCallbacks() int {
	return int(ctxt.thread_safe_callbacks)
}

// ThreadType Return active_thread_type
func (ctxt *AvCodecContext) ThreadType() int {
	return int(ctxt.thread_type)
}

// TicksPerFrame Return ticks_per_frame
func (ctxt *AvCodecContext) TicksPerFrame() int {
	return int(ctxt.ticks_per_frame)
}

// Trellis Return trellis
func (ctxt *AvCodecContext) Trellis() int {
	return int(ctxt.trellis)
}

// Width Return width
func (ctxt *AvCodecContext) Width() int {
	return int(ctxt.width)
}

// SetWidth Set width
func (ctxt *AvCodecContext) SetWidth(w int) {
	ctxt.width = C.int(w)
}

// WorkaroundBugs Return workaround_bugs
func (ctxt *AvCodecContext) WorkaroundBugs() int {
	return int(ctxt.workaround_bugs)
}

// AudioServiceType Return audio_service_type
func (ctxt *AvCodecContext) AudioServiceType() AvAudioServiceType {
	return (AvAudioServiceType)(ctxt.audio_service_type)
}

// ChromaSampleLocation Return chroma_sample_location
func (ctxt *AvCodecContext) ChromaSampleLocation() AvChromaLocation {
	return (AvChromaLocation)(ctxt.chroma_sample_location)
}

// CodecDescriptor Return codec_descriptor
func (ctxt *AvCodecContext) CodecDescriptor() *AvCodecDescriptor {
	return (*AvCodecDescriptor)(ctxt.codec_descriptor)
}

// AvCodecID Return codec_id
func (ctxt *AvCodecContext) AvCodecID() AvCodecID {
	return (AvCodecID)(ctxt.codec_id)
}

// CodecType Return codec_type
func (ctxt *AvCodecContext) CodecType() AvMediaType {
	return (AvMediaType)(ctxt.codec_type)
}

// ColorPrimaries Return color_primaries
func (ctxt *AvCodecContext) ColorPrimaries() AvColorPrimaries {
	return (AvColorPrimaries)(ctxt.color_primaries)
}

// ColorRange Return color_range
func (ctxt *AvCodecContext) ColorRange() AvColorRange {
	return (AvColorRange)(ctxt.color_range)
}

// ColorTrc Return color_trc
func (ctxt *AvCodecContext) ColorTrc() AvColorTransferCharacteristic {
	return (AvColorTransferCharacteristic)(ctxt.color_trc)
}

// Colorspace Return colorspace
func (ctxt *AvCodecContext) Colorspace() AvColorSpace {
	return (AvColorSpace)(ctxt.colorspace)
}

// FieldOrder Return field_order
func (ctxt *AvCodecContext) FieldOrder() AvFieldOrder {
	return (AvFieldOrder)(ctxt.field_order)
}

// PixFmt Return pix_fmt
func (ctxt *AvCodecContext) PixFmt() AvPixelFormat {
	return (AvPixelFormat)(ctxt.pix_fmt)
}

// RequestSampleFmt Return request_sample_fmt
func (ctxt *AvCodecContext) RequestSampleFmt() AvSampleFormat {
	return (AvSampleFormat)(ctxt.request_sample_fmt)
}

// SampleFmt Return sample_fmt
func (ctxt *AvCodecContext) SampleFmt() AvSampleFormat {
	return (AvSampleFormat)(ctxt.sample_fmt)
}

// SkipFrame Return skip_frame
func (ctxt *AvCodecContext) SkipFrame() AvDiscard {
	return (AvDiscard)(ctxt.skip_frame)
}

// SkipIdct Return skip_idct
func (ctxt *AvCodecContext) SkipIdct() AvDiscard {
	return (AvDiscard)(ctxt.skip_idct)
}

// SkipLoopFilter Return skip_loop_filter
func (ctxt *AvCodecContext) SkipLoopFilter() AvDiscard {
	return (AvDiscard)(ctxt.skip_loop_filter)
}
