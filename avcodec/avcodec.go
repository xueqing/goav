// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

//Package avcodec contains the codecs (decoders and encoders) provided by the libavcodec library
//Provides some generic global options, which can be set on all the encoders and decoders.
package avcodec

//#cgo pkg-config: libavformat libavcodec libavutil libswresample
//#include <stdio.h>
//#include <stdlib.h>
//#include <inttypes.h>
//#include <stdint.h>
//#include <string.h>
//#include <libavformat/avformat.h>
//#include <libavcodec/avcodec.h>
//#include <libavutil/avutil.h>
import "C"
import (
	"unsafe"
)

type (
	Codec                         C.struct_AVCodec
	Context                       C.struct_AVCodecContext
	Descriptor                    C.struct_AVCodecDescriptor
	Parser                        C.struct_AVCodecParser
	ParserContext                 C.struct_AVCodecParserContext
	Dictionary                    C.struct_AVDictionary
	Frame                         C.struct_AVFrame
	MediaType                     C.enum_AVMediaType
	Packet                        C.struct_AVPacket
	BitStreamFilter               C.struct_AVBitStreamFilter
	BitStreamFilterContext        C.struct_AVBitStreamFilterContext
	Rational                      C.struct_AVRational
	Class                         C.struct_AVClass
	AvCodecParameters             C.struct_AVCodecParameters
	AvHWAccel                     C.struct_AVHWAccel
	AvPacketSideData              C.struct_AVPacketSideData
	AvPanScan                     C.struct_AVPanScan
	Picture                       C.struct_AVPicture
	AvProfile                     C.struct_AVProfile
	AvSubtitle                    C.struct_AVSubtitle
	AvSubtitleRect                C.struct_AVSubtitleRect
	RcOverride                    C.struct_RcOverride
	AvBufferRef                   C.struct_AVBufferRef
	AvAudioServiceType            C.enum_AVAudioServiceType
	AvChromaLocation              C.enum_AVChromaLocation
	CodecID                       C.enum_AVCodecID
	AvColorPrimaries              C.enum_AVColorPrimaries
	AvColorRange                  C.enum_AVColorRange
	AvColorSpace                  C.enum_AVColorSpace
	AvColorTransferCharacteristic C.enum_AVColorTransferCharacteristic
	AvDiscard                     C.enum_AVDiscard
	AvFieldOrder                  C.enum_AVFieldOrder
	AvPacketSideDataType          C.enum_AVPacketSideDataType
	PixelFormat                   C.enum_AVPixelFormat
	AvSampleFormat                C.enum_AVSampleFormat
)

// AvCodecGetID Return codec_id
func (cp *AvCodecParameters) AvCodecGetID() CodecID {
	return *((*CodecID)(unsafe.Pointer(&cp.codec_id)))
}

// AvCodecGetType Return codec_type
func (cp *AvCodecParameters) AvCodecGetType() MediaType {
	return *((*MediaType)(unsafe.Pointer(&cp.codec_type)))
}

// AvCodecGetWidth Return width
func (cp *AvCodecParameters) AvCodecGetWidth() int {
	return (int)(*((*int32)(unsafe.Pointer(&cp.width))))
}

// AvCodecGetHeight Return height
func (cp *AvCodecParameters) AvCodecGetHeight() int {
	return (int)(*((*int32)(unsafe.Pointer(&cp.height))))
}

// AvCodecGetChannels Return channels
func (cp *AvCodecParameters) AvCodecGetChannels() int {
	return *((*int)(unsafe.Pointer(&cp.channels)))
}

// AvCodecGetSampleRate Return sample_rate
func (cp *AvCodecParameters) AvCodecGetSampleRate() int {
	return *((*int)(unsafe.Pointer(&cp.sample_rate)))
}

// AvGetProfileName Return a name for the specified profile, if available.
func (c *Codec) AvGetProfileName(p int) string {
	return C.GoString(C.av_get_profile_name((*C.struct_AVCodec)(c), C.int(p)))
}

// AvcodecAllocContext3 Allocate an Context and set its fields to default values.
func (c *Codec) AvcodecAllocContext3() *Context {
	return (*Context)(C.avcodec_alloc_context3((*C.struct_AVCodec)(c)))
}

// AvCodecIsEncoder Return a non-zero number if codec is an encoder, zero otherwise
func (c *Codec) AvCodecIsEncoder() int {
	return int(C.av_codec_is_encoder((*C.struct_AVCodec)(c)))
}

// AvCodecIsDecoder Return a non-zero number if codec is a decoder, zero otherwise
func (c *Codec) AvCodecIsDecoder() int {
	return int(C.av_codec_is_decoder((*C.struct_AVCodec)(c)))
}

// AvFastPaddedMalloc Same behaviour av_fast_malloc but the buffer has additional FF_INPUT_BUFFER_PADDING_SIZE at the end which will always be 0.
func AvFastPaddedMalloc(p unsafe.Pointer, s *uint, t uintptr) {
	C.av_fast_padded_malloc(p, (*C.uint)(unsafe.Pointer(s)), (C.size_t)(t))
}

// AvcodecVersion Return the LIBAvCODEC_VERSION_INT constant.
func AvcodecVersion() uint {
	return uint(C.avcodec_version())
}

// AvcodecConfiguration Return the libavcodec build-time configuration.
func AvcodecConfiguration() string {
	return C.GoString(C.avcodec_configuration())

}

// AvcodecLicense Return the libavcodec license.
func AvcodecLicense() string {
	return C.GoString(C.avcodec_license())
}

// AvcodecGetClass Get the Class for Context.
func AvcodecGetClass() *Class {
	return (*Class)(C.avcodec_get_class())
}

// AvcodecGetFrameClass Get the Class for Frame.
func AvcodecGetFrameClass() *Class {
	return (*Class)(C.avcodec_get_frame_class())
}

// AvcodecGetSubtitleRectClass Get the Class for AvSubtitleRect.
func AvcodecGetSubtitleRectClass() *Class {
	return (*Class)(C.avcodec_get_subtitle_rect_class())
}

// AvsubtitleFree Free all allocated data in the given subtitle struct.
func AvsubtitleFree(s *AvSubtitle) {
	C.avsubtitle_free((*C.struct_AVSubtitle)(s))
}

// AvPacketAlloc Allocate an AVPacket and set its fields to default values.  The resulting struct must be freed using av_packet_free().
func AvPacketAlloc() *Packet {
	return (*Packet)(C.av_packet_alloc())
}

// AvPacketPackDictionary Pack a dictionary for use in side_data.
func AvPacketPackDictionary(d *Dictionary, s *int) *uint8 {
	return (*uint8)(C.av_packet_pack_dictionary((*C.struct_AVDictionary)(d), (*C.int)(unsafe.Pointer(s))))
}

// AvPacketUnpackDictionary Unpack a dictionary from side_data.
func AvPacketUnpackDictionary(d *uint8, s int, dt **Dictionary) int {
	return int(C.av_packet_unpack_dictionary((*C.uint8_t)(d), C.int(s), (**C.struct_AVDictionary)(unsafe.Pointer(dt))))
}

// AvcodecFindDecoder Find a registered decoder with a matching codec ID.
func AvcodecFindDecoder(id CodecID) *Codec {
	return (*Codec)(C.avcodec_find_decoder((C.enum_AVCodecID)(id)))
}

// AvCodecIterate Iterate over all registered codecs.
func AvCodecIterate(p *unsafe.Pointer) *Codec {
	return (*Codec)(C.av_codec_iterate(p))
}

// AvcodecFindDecoderByName Find a registered decoder with the specified name.
func AvcodecFindDecoderByName(n string) *Codec {
	return (*Codec)(C.avcodec_find_decoder_by_name(C.CString(n)))
}

// AvcodecEnumToChromaPos Converts AvChromaLocation to swscale x/y chroma position.
func AvcodecEnumToChromaPos(x, y *int, l AvChromaLocation) int {
	return int(C.avcodec_enum_to_chroma_pos((*C.int)(unsafe.Pointer(x)), (*C.int)(unsafe.Pointer(y)), (C.enum_AVChromaLocation)(l)))
}

// AvcodecChromaPosToEnum Converts swscale x/y chroma position to AvChromaLocation.
func AvcodecChromaPosToEnum(x, y int) AvChromaLocation {
	return (AvChromaLocation)(C.avcodec_chroma_pos_to_enum(C.int(x), C.int(y)))
}

// AvcodecFindEncoder Find a registered encoder with a matching codec ID.
func AvcodecFindEncoder(id CodecID) *Codec {
	return (*Codec)(C.avcodec_find_encoder((C.enum_AVCodecID)(id)))
}

// AvcodecFindEncoderByName Find a registered encoder with the specified name.
func AvcodecFindEncoderByName(c string) *Codec {
	return (*Codec)(C.avcodec_find_encoder_by_name(C.CString(c)))
}

// AvcodecString Put a string representing the codec tag codec_tag in buf.
func AvcodecString(b string, bs int, ctxt *Context, e int) {
	C.avcodec_string(C.CString(b), C.int(bs), (*C.struct_AVCodecContext)(ctxt), C.int(e))
}

// AvcodecFillAudioFrame Fill Frame audio data and linesize pointers.
func AvcodecFillAudioFrame(f *Frame, c int, s AvSampleFormat, b *uint8, bs, a int) int {
	return int(C.avcodec_fill_audio_frame((*C.struct_AVFrame)(f), C.int(c), (C.enum_AVSampleFormat)(s), (*C.uint8_t)(b), C.int(bs), C.int(a)))
}

// AvGetBitsPerSample Return codec bits per sample.
func AvGetBitsPerSample(c CodecID) int {
	return int(C.av_get_bits_per_sample((C.enum_AVCodecID)(c)))
}

// AvGetPcmCodec Return the PCM codec associated with a sample format.
func AvGetPcmCodec(f AvSampleFormat, b int) CodecID {
	return (CodecID)(C.av_get_pcm_codec((C.enum_AVSampleFormat)(f), C.int(b)))
}

// AvGetExactBitsPerSample Return codec bits per sample.
func AvGetExactBitsPerSample(c CodecID) int {
	return int(C.av_get_exact_bits_per_sample((C.enum_AVCodecID)(c)))
}

// AvFastPaddedMallocz Same behaviour av_fast_padded_malloc except that buffer will always be 0-initialized after call.
func AvFastPaddedMallocz(p unsafe.Pointer, s *uint, t uintptr) {
	C.av_fast_padded_mallocz(p, (*C.uint)(unsafe.Pointer(s)), (C.size_t)(t))
}

// AvXiphlacing Encode extradata length to a buffer.
func AvXiphlacing(s *string, v uint) uint {
	return uint(C.av_xiphlacing((*C.uchar)(unsafe.Pointer(s)), (C.uint)(v)))
}

// AvcodecGetType Get the type of the given codec.
func AvcodecGetType(c CodecID) MediaType {
	return (MediaType)(C.avcodec_get_type((C.enum_AVCodecID)(c)))
}

// AvcodecGetName Get the name of a codec.
func AvcodecGetName(d CodecID) string {
	return C.GoString(C.avcodec_get_name((C.enum_AVCodecID)(d)))
}

// AvcodecDescriptorGet Return descriptor for given codec ID or NULL if no descriptor exists.
func AvcodecDescriptorGet(id CodecID) *Descriptor {
	return (*Descriptor)(C.avcodec_descriptor_get((C.enum_AVCodecID)(id)))
}

// AvcodecDescriptorNext Iterate over all codec descriptors known to libavcodec.
func (d *Descriptor) AvcodecDescriptorNext() *Descriptor {
	return (*Descriptor)(C.avcodec_descriptor_next((*C.struct_AVCodecDescriptor)(d)))
}

// AvcodecDescriptorGetByName Return codec descriptor with the given name or NULL if no such descriptor exists.
func AvcodecDescriptorGetByName(n string) *Descriptor {
	return (*Descriptor)(C.avcodec_descriptor_get_by_name(C.CString(n)))
}

// Pts Return Pts
func (f *Frame) Pts() int64 {
	return int64(f.pts)
}
