// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package avcodec

// ActiveThreadType Return active_thread_type
func (ctxt *Context) ActiveThreadType() int {
	return int(ctxt.active_thread_type)
}

// BFrameStrategy Return b_frame_strategy
func (ctxt *Context) BFrameStrategy() int {
	return int(ctxt.b_frame_strategy)
}

// BQuantFactor Return b_quant_factor
func (ctxt *Context) BQuantFactor() float64 {
	return float64(ctxt.b_quant_factor)
}

// BQuantOffset Return b_quant_offset
func (ctxt *Context) BQuantOffset() float64 {
	return float64(ctxt.b_quant_offset)
}

// BSensitivity Return b_sensitivity
func (ctxt *Context) BSensitivity() int {
	return int(ctxt.b_sensitivity)
}

// BidirRefine Return bidir_refine
func (ctxt *Context) BidirRefine() int {
	return int(ctxt.bidir_refine)
}

// BitRate Return bit_rate
func (ctxt *Context) BitRate() int {
	return int(ctxt.bit_rate)
}

// BitRateTolerance Return bit_rate_tolerance
func (ctxt *Context) BitRateTolerance() int {
	return int(ctxt.bit_rate_tolerance)
}

// BitsPerCodedSample Return bits_per_coded_sample
func (ctxt *Context) BitsPerCodedSample() int {
	return int(ctxt.bits_per_coded_sample)
}

// BitsPerRawSample Return bits_per_raw_sample
func (ctxt *Context) BitsPerRawSample() int {
	return int(ctxt.bits_per_raw_sample)
}

// BlockAlign Return block_align
func (ctxt *Context) BlockAlign() int {
	return int(ctxt.block_align)
}

// BrdScale Return brd_scale
func (ctxt *Context) BrdScale() int {
	return int(ctxt.brd_scale)
}

// Channels Return channels
func (ctxt *Context) Channels() int {
	return int(ctxt.channels)
}

// Chromaoffset Return chromaoffset
func (ctxt *Context) Chromaoffset() int {
	return int(ctxt.chromaoffset)
}

// CodedHeight Return coded_height
func (ctxt *Context) CodedHeight() int {
	return int(ctxt.coded_height)
}

// CodedWidth Return coded_width
func (ctxt *Context) CodedWidth() int {
	return int(ctxt.coded_width)
}

// CoderType Return coder_type
func (ctxt *Context) CoderType() int {
	return int(ctxt.coder_type)
}

// CompressionLevel Return compression_level
func (ctxt *Context) CompressionLevel() int {
	return int(ctxt.compression_level)
}

// ContextModel Return context_model
func (ctxt *Context) ContextModel() int {
	return int(ctxt.context_model)
}

// Cutoff Return cutoff
func (ctxt *Context) Cutoff() int {
	return int(ctxt.cutoff)
}

// DarkMasking Return dark_masking
func (ctxt *Context) DarkMasking() float64 {
	return float64(ctxt.dark_masking)
}

// DctAlgo Return dct_algo
func (ctxt *Context) DctAlgo() int {
	return int(ctxt.dct_algo)
}

// Debug Return debug
func (ctxt *Context) Debug() int {
	return int(ctxt.debug)
}

// DebugMv Return debug_mv
func (ctxt *Context) DebugMv() int {
	return int(ctxt.debug_mv)
}

// Delay Return delay
func (ctxt *Context) Delay() int {
	return int(ctxt.delay)
}

// DiaSize Return dia_size
func (ctxt *Context) DiaSize() int {
	return int(ctxt.dia_size)
}

// ErrRecognition Return err_recognition
func (ctxt *Context) ErrRecognition() int {
	return int(ctxt.err_recognition)
}

// ErrorConcealment Return error_concealment
func (ctxt *Context) ErrorConcealment() int {
	return int(ctxt.error_concealment)
}

// ExtradataSize Return extradata_size
func (ctxt *Context) ExtradataSize() int {
	return int(ctxt.extradata_size)
}

// Flags Return flags
func (ctxt *Context) Flags() int {
	return int(ctxt.flags)
}

// Flags2 Return flags2
func (ctxt *Context) Flags2() int {
	return int(ctxt.flags2)
}

// FrameBits Return frame_bits
func (ctxt *Context) FrameBits() int {
	return int(ctxt.frame_bits)
}

// FrameNumber Return frame_number
func (ctxt *Context) FrameNumber() int {
	return int(ctxt.frame_number)
}

// FrameSize Return frame_size
func (ctxt *Context) FrameSize() int {
	return int(ctxt.frame_size)
}

// FrameSkipCmp Return frame_skip_cmp
func (ctxt *Context) FrameSkipCmp() int {
	return int(ctxt.frame_skip_cmp)
}

// FrameSkipExp Return frame_skip_exp
func (ctxt *Context) FrameSkipExp() int {
	return int(ctxt.frame_skip_exp)
}

// FrameSkipFactor Return frame_skip_factor
func (ctxt *Context) FrameSkipFactor() int {
	return int(ctxt.frame_skip_factor)
}

// FrameSkipThreshold Return frame_skip_threshold
func (ctxt *Context) FrameSkipThreshold() int {
	return int(ctxt.frame_skip_threshold)
}

// GlobalQuality Return global_quality
func (ctxt *Context) GlobalQuality() int {
	return int(ctxt.global_quality)
}

// GopSize Return gop_size
func (ctxt *Context) GopSize() int {
	return int(ctxt.gop_size)
}

// HasBFrames Return has_b_frames
func (ctxt *Context) HasBFrames() int {
	return int(ctxt.has_b_frames)
}

// HeaderBits Return header_bits
func (ctxt *Context) HeaderBits() int {
	return int(ctxt.header_bits)
}

// Height Return height
func (ctxt *Context) Height() int {
	return int(ctxt.height)
}

// ICount Return i_count
func (ctxt *Context) ICount() int {
	return int(ctxt.i_count)
}

// IQuantFactor Return i_quant_factor
func (ctxt *Context) IQuantFactor() float64 {
	return float64(ctxt.i_quant_factor)
}

// IQuantOffset Return i_quant_offset
func (ctxt *Context) IQuantOffset() float64 {
	return float64(ctxt.i_quant_offset)
}

// ITexBits Return i_tex_bits
func (ctxt *Context) ITexBits() int {
	return int(ctxt.i_tex_bits)
}

// IdctAlgo Return idct_algo
func (ctxt *Context) IdctAlgo() int {
	return int(ctxt.idct_algo)
}

// IldctCmp Return ildct_cmp
func (ctxt *Context) IldctCmp() int {
	return int(ctxt.ildct_cmp)
}

// IntraDcPrecision Return intra_dc_precision
func (ctxt *Context) IntraDcPrecision() int {
	return int(ctxt.intra_dc_precision)
}

// KeyintMin Return keyint_min
func (ctxt *Context) KeyintMin() int {
	return int(ctxt.keyint_min)
}

// LastPredictorCount Return last_predictor_count
func (ctxt *Context) LastPredictorCount() int {
	return int(ctxt.last_predictor_count)
}

// Level Return level
func (ctxt *Context) Level() int {
	return int(ctxt.level)
}

// LogLevelOffset Return log_level_offset
func (ctxt *Context) LogLevelOffset() int {
	return int(ctxt.log_level_offset)
}

// Lowres Return lowres
func (ctxt *Context) Lowres() int {
	return int(ctxt.lowres)
}

// LumiMasking Return lumi_masking
func (ctxt *Context) LumiMasking() float64 {
	return float64(ctxt.lumi_masking)
}

// MaxBFrames Return max_b_frames
func (ctxt *Context) MaxBFrames() int {
	return int(ctxt.max_b_frames)
}

// MaxPredictionOrder Return max_prediction_order
func (ctxt *Context) MaxPredictionOrder() int {
	return int(ctxt.max_prediction_order)
}

// MaxQdiff Return max_qdiff
func (ctxt *Context) MaxQdiff() int {
	return int(ctxt.max_qdiff)
}

// MbCmp Return mb_cmp
func (ctxt *Context) MbCmp() int {
	return int(ctxt.mb_cmp)
}

// MbDecision Return mb_decision
func (ctxt *Context) MbDecision() int {
	return int(ctxt.mb_decision)
}

// MbLmax Return mb_lmax
func (ctxt *Context) MbLmax() int {
	return int(ctxt.mb_lmax)
}

// MbLmin Return mb_lmin
func (ctxt *Context) MbLmin() int {
	return int(ctxt.mb_lmin)
}

// MeCmp Return me_cmp
func (ctxt *Context) MeCmp() int {
	return int(ctxt.me_cmp)
}

// MePenaltyCompensation Return me_penalty_compensation
func (ctxt *Context) MePenaltyCompensation() int {
	return int(ctxt.me_penalty_compensation)
}

// MePreCmp Return me_pre_cmp
func (ctxt *Context) MePreCmp() int {
	return int(ctxt.me_pre_cmp)
}

// MeRange Return me_range
func (ctxt *Context) MeRange() int {
	return int(ctxt.me_range)
}

// MeSubCmp Return me_sub_cmp
func (ctxt *Context) MeSubCmp() int {
	return int(ctxt.me_sub_cmp)
}

// MeSubpelQuality Return me_subpel_quality
func (ctxt *Context) MeSubpelQuality() int {
	return int(ctxt.me_subpel_quality)
}

// MinPredictionOrder Return min_prediction_order
func (ctxt *Context) MinPredictionOrder() int {
	return int(ctxt.min_prediction_order)
}

// MiscBits Return misc_bits
func (ctxt *Context) MiscBits() int {
	return int(ctxt.misc_bits)
}

// MpegQuant Return mpeg_quant
func (ctxt *Context) MpegQuant() int {
	return int(ctxt.mpeg_quant)
}

// Mv0Threshold Return mv0_threshold
func (ctxt *Context) Mv0Threshold() int {
	return int(ctxt.mv0_threshold)
}

// MvBits Return mv_bits
func (ctxt *Context) MvBits() int {
	return int(ctxt.mv_bits)
}

// NoiseReduction Return noise_reduction
func (ctxt *Context) NoiseReduction() int {
	return int(ctxt.noise_reduction)
}

// NsseWeight Return nsse_weight
func (ctxt *Context) NsseWeight() int {
	return int(ctxt.nsse_weight)
}

// PCount Return p_count
func (ctxt *Context) PCount() int {
	return int(ctxt.p_count)
}

// PMasking Return p_masking
func (ctxt *Context) PMasking() float64 {
	return float64(ctxt.p_masking)
}

// PTexBits Return p_tex_bits
func (ctxt *Context) PTexBits() int {
	return int(ctxt.p_tex_bits)
}

// PreDiaSize Return pre_dia_size
func (ctxt *Context) PreDiaSize() int {
	return int(ctxt.pre_dia_size)
}

// PreMe Return pre_me
func (ctxt *Context) PreMe() int {
	return int(ctxt.pre_me)
}

// PredictionMethod Return prediction_method
func (ctxt *Context) PredictionMethod() int {
	return int(ctxt.prediction_method)
}

// Profile Return profile
func (ctxt *Context) Profile() int {
	return int(ctxt.profile)
}

// Qblur Return qblur
func (ctxt *Context) Qblur() float64 {
	return float64(ctxt.qblur)
}

// Qcompress Return qcompress
func (ctxt *Context) Qcompress() float64 {
	return float64(ctxt.qcompress)
}

// Qmax Return qmax
func (ctxt *Context) Qmax() int {
	return int(ctxt.qmax)
}

// Qmin Return qmin
func (ctxt *Context) Qmin() int {
	return int(ctxt.qmin)
}

// RcBufferSize Return rc_buffer_size
func (ctxt *Context) RcBufferSize() int {
	return int(ctxt.rc_buffer_size)
}

// RcInitialBufferOccupancy Return rc_initial_buffer_occupancy
func (ctxt *Context) RcInitialBufferOccupancy() int {
	return int(ctxt.rc_initial_buffer_occupancy)
}

// RcMaxAvailableVbvUse Return rc_max_available_vbv_use
func (ctxt *Context) RcMaxAvailableVbvUse() float64 {
	return float64(ctxt.rc_max_available_vbv_use)
}

// RcMaxRate Return rc_max_rate
func (ctxt *Context) RcMaxRate() int {
	return int(ctxt.rc_max_rate)
}

// RcMinRate Return rc_min_rate
func (ctxt *Context) RcMinRate() int {
	return int(ctxt.rc_min_rate)
}

// RcMinVbvOverflowUse Return rc_min_vbv_overflow_use
func (ctxt *Context) RcMinVbvOverflowUse() float64 {
	return float64(ctxt.rc_min_vbv_overflow_use)
}

// RcOverrideCount Return rc_override_count
func (ctxt *Context) RcOverrideCount() int {
	return int(ctxt.rc_override_count)
}

// RefcountedFrames Return refcounted_frames
func (ctxt *Context) RefcountedFrames() int {
	return int(ctxt.refcounted_frames)
}

// Refs Return refs
func (ctxt *Context) Refs() int {
	return int(ctxt.refs)
}

// RtpPayloadSize Return rtp_payload_size
func (ctxt *Context) RtpPayloadSize() int {
	return int(ctxt.rtp_payload_size)
}

// SampleRate Return sample_rate
func (ctxt *Context) SampleRate() int {
	return int(ctxt.sample_rate)
}

// ScenechangeThreshold Return scenechange_threshold
func (ctxt *Context) ScenechangeThreshold() int {
	return int(ctxt.scenechange_threshold)
}

// SeekPreroll Return seek_preroll
func (ctxt *Context) SeekPreroll() int {
	return int(ctxt.seek_preroll)
}

// SideDataOnlyPackets Return side_data_only_packets
func (ctxt *Context) SideDataOnlyPackets() int {
	return int(ctxt.side_data_only_packets)
}

// SkipAlpha Return skip_alpha
func (ctxt *Context) SkipAlpha() int {
	return int(ctxt.skip_alpha)
}

// SkipBottom Return skip_bottom
func (ctxt *Context) SkipBottom() int {
	return int(ctxt.skip_bottom)
}

// SkipCount Return skip_count
func (ctxt *Context) SkipCount() int {
	return int(ctxt.skip_count)
}

// SkipTop Return skip_top
func (ctxt *Context) SkipTop() int {
	return int(ctxt.skip_top)
}

// SliceCount Return slice_count
func (ctxt *Context) SliceCount() int {
	return int(ctxt.slice_count)
}

// SliceFlags Return slice_flags
func (ctxt *Context) SliceFlags() int {
	return int(ctxt.slice_flags)
}

// Slices Return slices
func (ctxt *Context) Slices() int {
	return int(ctxt.slices)
}

// SpatialCplxMasking Return spatial_cplx_masking
func (ctxt *Context) SpatialCplxMasking() float64 {
	return float64(ctxt.spatial_cplx_masking)
}

// StrictStdCompliance Return strict_std_compliance
func (ctxt *Context) StrictStdCompliance() int {
	return int(ctxt.strict_std_compliance)
}

// SubCharencMode Return sub_charenc_mode
func (ctxt *Context) SubCharencMode() int {
	return int(ctxt.sub_charenc_mode)
}

// SubtitleHeaderSize Return subtitle_header_size
func (ctxt *Context) SubtitleHeaderSize() int {
	return int(ctxt.subtitle_header_size)
}

// TemporalCplxMasking Return temporal_cplx_masking
func (ctxt *Context) TemporalCplxMasking() float64 {
	return float64(ctxt.temporal_cplx_masking)
}

// ThreadCount Return thread_count
func (ctxt *Context) ThreadCount() int {
	return int(ctxt.thread_count)
}

// ThreadSafeCallbacks Return thread_safe_callbacks
func (ctxt *Context) ThreadSafeCallbacks() int {
	return int(ctxt.thread_safe_callbacks)
}

// ThreadType Return active_thread_type
func (ctxt *Context) ThreadType() int {
	return int(ctxt.thread_type)
}

// TicksPerFrame Return ticks_per_frame
func (ctxt *Context) TicksPerFrame() int {
	return int(ctxt.ticks_per_frame)
}

// Trellis Return trellis
func (ctxt *Context) Trellis() int {
	return int(ctxt.trellis)
}

// Width Return width
func (ctxt *Context) Width() int {
	return int(ctxt.width)
}

// WorkaroundBugs Return workaround_bugs
func (ctxt *Context) WorkaroundBugs() int {
	return int(ctxt.workaround_bugs)
}

// AudioServiceType Return audio_service_type
func (ctxt *Context) AudioServiceType() AvAudioServiceType {
	return (AvAudioServiceType)(ctxt.audio_service_type)
}

// ChromaSampleLocation Return chroma_sample_location
func (ctxt *Context) ChromaSampleLocation() AvChromaLocation {
	return (AvChromaLocation)(ctxt.chroma_sample_location)
}

// CodecDescriptor Return codec_descriptor
func (ctxt *Context) CodecDescriptor() *Descriptor {
	return (*Descriptor)(ctxt.codec_descriptor)
}

// CodecID Return codec_id
func (ctxt *Context) CodecID() CodecID {
	return (CodecID)(ctxt.codec_id)
}

// CodecType Return codec_type
func (ctxt *Context) CodecType() MediaType {
	return (MediaType)(ctxt.codec_type)
}

// ColorPrimaries Return color_primaries
func (ctxt *Context) ColorPrimaries() AvColorPrimaries {
	return (AvColorPrimaries)(ctxt.color_primaries)
}

// ColorRange Return color_range
func (ctxt *Context) ColorRange() AvColorRange {
	return (AvColorRange)(ctxt.color_range)
}

// ColorTrc Return color_trc
func (ctxt *Context) ColorTrc() AvColorTransferCharacteristic {
	return (AvColorTransferCharacteristic)(ctxt.color_trc)
}

// Colorspace Return colorspace
func (ctxt *Context) Colorspace() AvColorSpace {
	return (AvColorSpace)(ctxt.colorspace)
}

// FieldOrder Return field_order
func (ctxt *Context) FieldOrder() AvFieldOrder {
	return (AvFieldOrder)(ctxt.field_order)
}

// PixFmt Return pix_fmt
func (ctxt *Context) PixFmt() PixelFormat {
	return (PixelFormat)(ctxt.pix_fmt)
}

// RequestSampleFmt Return request_sample_fmt
func (ctxt *Context) RequestSampleFmt() AvSampleFormat {
	return (AvSampleFormat)(ctxt.request_sample_fmt)
}

// SampleFmt Return sample_fmt
func (ctxt *Context) SampleFmt() AvSampleFormat {
	return (AvSampleFormat)(ctxt.sample_fmt)
}

// SkipFrame Return skip_frame
func (ctxt *Context) SkipFrame() AvDiscard {
	return (AvDiscard)(ctxt.skip_frame)
}

// SkipIdct Return skip_idct
func (ctxt *Context) SkipIdct() AvDiscard {
	return (AvDiscard)(ctxt.skip_idct)
}

// SkipLoopFilter Return skip_loop_filter
func (ctxt *Context) SkipLoopFilter() AvDiscard {
	return (AvDiscard)(ctxt.skip_loop_filter)
}
