// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

//Package libavcodec contains the codecs (decoders and encoders) provided by the libavcodec library
//Provides some generic global options, which can be set on all the encoders and decoders.
package libavcodec

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
	AvCodec                       C.struct_AVCodec
	AvCodecContext                C.struct_AVCodecContext
	AvCodecDescriptor             C.struct_AVCodecDescriptor
	AvCodecParser                 C.struct_AVCodecParser
	AvCodecParserContext          C.struct_AVCodecParserContext
	AvDictionary                  C.struct_AVDictionary
	AvFrame                       C.struct_AVFrame
	AvMediaType                   C.enum_AVMediaType
	AvPacket                      C.struct_AVPacket
	AvRational                    C.struct_AVRational
	AvClass                       C.struct_AVClass
	AvCodecParameters             C.struct_AVCodecParameters
	AvHWAccel                     C.struct_AVHWAccel
	AvPacketSideData              C.struct_AVPacketSideData
	AvPanScan                     C.struct_AVPanScan
	AvPicture                     C.struct_AVPicture
	AvProfile                     C.struct_AVProfile
	AvSubtitle                    C.struct_AVSubtitle
	AvSubtitleRect                C.struct_AVSubtitleRect
	RcOverride                    C.struct_RcOverride
	AvBufferRef                   C.struct_AVBufferRef
	AvAudioServiceType            C.enum_AVAudioServiceType
	AvChromaLocation              C.enum_AVChromaLocation
	AvCodecID                     C.enum_AVCodecID
	AvColorPrimaries              C.enum_AVColorPrimaries
	AvColorRange                  C.enum_AVColorRange
	AvColorSpace                  C.enum_AVColorSpace
	AvColorTransferCharacteristic C.enum_AVColorTransferCharacteristic
	AvDiscard                     C.enum_AVDiscard
	AvFieldOrder                  C.enum_AVFieldOrder
	AvPacketSideDataType          C.enum_AVPacketSideDataType
	AvPixelFormat                 C.enum_AVPixelFormat
	AvSampleFormat                C.enum_AVSampleFormat
	AvRounding                    C.enum_AVRounding
)

// AvGetProfileName Return a name for the specified profile, if available.
func (codec *AvCodec) AvGetProfileName(profile int) string {
	return C.GoString(C.av_get_profile_name((*C.struct_AVCodec)(codec), C.int(profile)))
}

// AvcodecAllocContext3 Allocate an Context and set its fields to default values.
func (codec *AvCodec) AvcodecAllocContext3() *AvCodecContext {
	return (*AvCodecContext)(C.avcodec_alloc_context3((*C.struct_AVCodec)(codec)))
}

// AvCodecIsEncoder Return a non-zero number if codec is an encoder, zero otherwise
func (codec *AvCodec) AvCodecIsEncoder() int {
	return int(C.av_codec_is_encoder((*C.struct_AVCodec)(codec)))
}

// AvCodecIsDecoder Return a non-zero number if codec is a decoder, zero otherwise
func (codec *AvCodec) AvCodecIsDecoder() int {
	return int(C.av_codec_is_decoder((*C.struct_AVCodec)(codec)))
}

// AvFastPaddedMalloc Same behaviour av_fast_malloc but the buffer has additional
// FF_INPUT_BUFFER_PADDING_SIZE at the end which will always be 0.
func AvFastPaddedMalloc(ptr unsafe.Pointer, size *uint, minSize uintptr) {
	C.av_fast_padded_malloc(ptr, (*C.uint)(unsafe.Pointer(size)), (C.size_t)(minSize))
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
func AvcodecGetClass() *AvClass {
	return (*AvClass)(C.avcodec_get_class())
}

// AvcodecGetFrameClass Get the Class for Frame.
func AvcodecGetFrameClass() *AvClass {
	return (*AvClass)(C.avcodec_get_frame_class())
}

// AvcodecGetSubtitleRectClass Get the Class for AvSubtitleRect.
func AvcodecGetSubtitleRectClass() *AvClass {
	return (*AvClass)(C.avcodec_get_subtitle_rect_class())
}

// AvsubtitleFree Free all allocated data in the given subtitle struct.
func AvsubtitleFree(sub *AvSubtitle) {
	C.avsubtitle_free((*C.struct_AVSubtitle)(sub))
}

// AvPacketAlloc Allocate an AvPacket and set its fields to default values.
// The resulting struct must be freed using av_packet_free().
func AvPacketAlloc() *AvPacket {
	return (*AvPacket)(C.av_packet_alloc())
}

// AvPacketPackDictionary Pack a dictionary for use in side_data.
func AvPacketPackDictionary(dict *AvDictionary, size *int) *uint8 {
	return (*uint8)(C.av_packet_pack_dictionary((*C.struct_AVDictionary)(dict),
		(*C.int)(unsafe.Pointer(size))))
}

// AvPacketUnpackDictionary Unpack a dictionary from side_data.
func AvPacketUnpackDictionary(data *uint8, size int, dict **AvDictionary) int {
	return int(C.av_packet_unpack_dictionary((*C.uint8_t)(data), C.int(size),
		(**C.struct_AVDictionary)(unsafe.Pointer(dict))))
}

// AvcodecFindDecoder Find a registered decoder with a matching codec ID.
func AvcodecFindDecoder(id AvCodecID) *AvCodec {
	return (*AvCodec)(C.avcodec_find_decoder((C.enum_AVCodecID)(id)))
}

// AvCodecIterate Iterate over all registered codecs.
func AvCodecIterate(opaque *unsafe.Pointer) *AvCodec {
	return (*AvCodec)(C.av_codec_iterate(opaque))
}

// AvcodecFindDecoderByName Find a registered decoder with the specified name.
func AvcodecFindDecoderByName(name string) *AvCodec {
	return (*AvCodec)(C.avcodec_find_decoder_by_name(C.CString(name)))
}

// AvcodecEnumToChromaPos Converts AvChromaLocation to swscale x/y chroma position.
func AvcodecEnumToChromaPos(xpos, ypos *int, pos AvChromaLocation) int {
	return int(C.avcodec_enum_to_chroma_pos((*C.int)(unsafe.Pointer(xpos)),
		(*C.int)(unsafe.Pointer(ypos)), (C.enum_AVChromaLocation)(pos)))
}

// AvcodecChromaPosToEnum Converts swscale x/y chroma position to AvChromaLocation.
func AvcodecChromaPosToEnum(xpos, ypos int) AvChromaLocation {
	return (AvChromaLocation)(C.avcodec_chroma_pos_to_enum(C.int(xpos), C.int(ypos)))
}

// AvcodecFindEncoder Find a registered encoder with a matching codec ID.
func AvcodecFindEncoder(id AvCodecID) *AvCodec {
	return (*AvCodec)(C.avcodec_find_encoder((C.enum_AVCodecID)(id)))
}

// AvcodecFindEncoderByName Find a registered encoder with the specified name.
func AvcodecFindEncoderByName(name string) *AvCodec {
	return (*AvCodec)(C.avcodec_find_encoder_by_name(C.CString(name)))
}

// AvcodecString Put a string representing the codec tag codec_tag in buf.
func AvcodecString(buf string, bufSize int, enc *AvCodecContext, encode int) {
	C.avcodec_string(C.CString(buf), C.int(bufSize), (*C.struct_AVCodecContext)(enc), C.int(encode))
}

// AvcodecFillAudioFrame Fill Frame audio data and linesize pointers.
func AvcodecFillAudioFrame(frame *AvFrame, nbChannels int, sampleFmt AvSampleFormat,
	buf *uint8, bufSize, align int) int {
	return int(C.avcodec_fill_audio_frame((*C.struct_AVFrame)(frame), C.int(nbChannels),
		(C.enum_AVSampleFormat)(sampleFmt), (*C.uint8_t)(buf), C.int(bufSize), C.int(align)))
}

// AvGetBitsPerSample Return codec bits per sample.
func AvGetBitsPerSample(id AvCodecID) int {
	return int(C.av_get_bits_per_sample((C.enum_AVCodecID)(id)))
}

// AvGetPcmCodec Return the PCM codec associated with a sample format.
func AvGetPcmCodec(fmt AvSampleFormat, be int) AvCodecID {
	return (AvCodecID)(C.av_get_pcm_codec((C.enum_AVSampleFormat)(fmt), C.int(be)))
}

// AvGetExactBitsPerSample Return codec bits per sample.
func AvGetExactBitsPerSample(id AvCodecID) int {
	return int(C.av_get_exact_bits_per_sample((C.enum_AVCodecID)(id)))
}

// AvFastPaddedMallocz Same behaviour av_fast_padded_malloc except that buffer
// will always be 0-initialized after call.
func AvFastPaddedMallocz(ptr unsafe.Pointer, size *uint, minSize uintptr) {
	C.av_fast_padded_mallocz(ptr, (*C.uint)(unsafe.Pointer(size)), (C.size_t)(minSize))
}

// AvXiphlacing Encode extradata length to a buffer.
func AvXiphlacing(s *string, v uint) uint {
	return uint(C.av_xiphlacing((*C.uchar)(unsafe.Pointer(s)), (C.uint)(v)))
}

// AvcodecGetType Get the type of the given codec.
func AvcodecGetType(id AvCodecID) AvMediaType {
	return (AvMediaType)(C.avcodec_get_type((C.enum_AVCodecID)(id)))
}

// AvcodecGetName Get the name of a codec.
func AvcodecGetName(id AvCodecID) string {
	return C.GoString(C.avcodec_get_name((C.enum_AVCodecID)(id)))
}

// AvcodecDescriptorGet Return descriptor for given codec ID or NULL if no descriptor exists.
func AvcodecDescriptorGet(id AvCodecID) *AvCodecDescriptor {
	return (*AvCodecDescriptor)(C.avcodec_descriptor_get((C.enum_AVCodecID)(id)))
}

// AvcodecDescriptorGetByName Return codec descriptor with the given name or NULL if no such descriptor exists.
func AvcodecDescriptorGetByName(name string) *AvCodecDescriptor {
	return (*AvCodecDescriptor)(C.avcodec_descriptor_get_by_name(C.CString(name)))
}
