package libavcodec

//#cgo pkg-config: libavcodec
//#include <libavcodec/avcodec.h>
import "C"

import "unsafe"

// CodecType Return codec_type
func (cp *AvCodecParameters) CodecType() AvMediaType {
	return *((*AvMediaType)(unsafe.Pointer(&cp.codec_type)))
}

// CodecID Return codec_id
func (cp *AvCodecParameters) CodecID() AvCodecID {
	return *((*AvCodecID)(unsafe.Pointer(&cp.codec_id)))
}

// Format Return format
func (cp *AvCodecParameters) Format() int {
	return *((*int)(unsafe.Pointer(&cp.format)))
}

// SetFormat Set format
func (cp *AvCodecParameters) SetFormat(f int) {
	cp.format = C.int(f)
}

// Width Return width
func (cp *AvCodecParameters) Width() int {
	return (int)(*((*int32)(unsafe.Pointer(&cp.width))))
}

// SetWidth Set width
func (cp *AvCodecParameters) SetWidth(w int) {
	cp.width = C.int(w)
}

// Height Return height
func (cp *AvCodecParameters) Height() int {
	return (int)(*((*int32)(unsafe.Pointer(&cp.height))))
}

// SetHeight Set height
func (cp *AvCodecParameters) SetHeight(h int) {
	cp.height = C.int(h)
}

// SampleAspectRatio Return sample_aspect_ratio
func (cp *AvCodecParameters) SampleAspectRatio() AvRational {
	return NewAvRational(int(cp.sample_aspect_ratio.num), int(cp.sample_aspect_ratio.den))
}

// SetSampleAspectRatio Set sample_aspect_ratio
func (cp *AvCodecParameters) SetSampleAspectRatio(sampleAspectRatio AvRational) {
	cp.sample_aspect_ratio.num = C.int(sampleAspectRatio.Num())
	cp.sample_aspect_ratio.den = C.int(sampleAspectRatio.Den())
}

// ChannelLayout Return channel_layout
func (cp *AvCodecParameters) ChannelLayout() uint64 {
	return *((*uint64)(unsafe.Pointer(&cp.channel_layout)))
}

// SetChannelLayout Set channel_layout
func (cp *AvCodecParameters) SetChannelLayout(cl uint64) {
	cp.channel_layout = C.uint64_t(cl)
}

// Channels Return channels
func (cp *AvCodecParameters) Channels() int {
	return *((*int)(unsafe.Pointer(&cp.channels)))
}

// SetChannels Set channels
func (cp *AvCodecParameters) SetChannels(nc int) {
	cp.channels = C.int(nc)
}

// SampleRate Return sample_rate
func (cp *AvCodecParameters) SampleRate() int {
	return *((*int)(unsafe.Pointer(&cp.sample_rate)))
}

// SetSampleRate Set sample_rate
func (cp *AvCodecParameters) SetSampleRate(sr int) {
	cp.sample_rate = C.int(sr)
}
