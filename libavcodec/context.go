// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package libavcodec

//#cgo pkg-config: libavcodec
//#include <libavcodec/avcodec.h>
import "C"
import (
	"unsafe"
)

// AvcodecFreeContext Free the codec context and everything associated with it and write NULL to the provided pointer.
func (ctxt *Context) AvcodecFreeContext() {
	C.avcodec_free_context((**C.struct_AVCodecContext)(unsafe.Pointer(ctxt)))
}

// AvcodecGetContextDefaults3 Set the fields of the given Context to default values corresponding to the given codec (defaults may be codec-dependent).
func (ctxt *Context) AvcodecGetContextDefaults3(c *Codec) int {
	return int(C.avcodec_get_context_defaults3((*C.struct_AVCodecContext)(ctxt), (*C.struct_AVCodec)(c)))
}

// AvcodecParametersFromContext Fill the parameters struct based on the values from the supplied codec context.
func (ctxt *Context) AvcodecParametersFromContext(cp *AvCodecParameters) int {
	return int(C.avcodec_parameters_from_context((*C.struct_AVCodecParameters)(cp), (*C.struct_AVCodecContext)(ctxt)))
}

// AvcodecParametersToContext Fill the codec context based on the values from the supplied codec parameters.
func (ctxt *Context) AvcodecParametersToContext(cp *AvCodecParameters) int {
	return int(C.avcodec_parameters_to_context((*C.struct_AVCodecContext)(ctxt), (*C.struct_AVCodecParameters)(cp)))
}

// AvcodecOpen2 Initialize the Context to use the given Codec
func (ctxt *Context) AvcodecOpen2(c *Codec, d **Dictionary) int {
	return int(C.avcodec_open2((*C.struct_AVCodecContext)(ctxt), (*C.struct_AVCodec)(c), (**C.struct_AVDictionary)(unsafe.Pointer(d))))
}

// AvcodecClose Close a given Context and free all the data associated with it (but not the Context itself).
func (ctxt *Context) AvcodecClose() int {
	return int(C.avcodec_close((*C.struct_AVCodecContext)(ctxt)))
}

// AvcodecDefaultGetBuffer2 The default callback for Context.get_buffer2().
func (ctxt *Context) AvcodecDefaultGetBuffer2(f *Frame, l int) int {
	return int(C.avcodec_default_get_buffer2((*C.struct_AVCodecContext)(ctxt), (*C.struct_AVFrame)(f), C.int(l)))
}

// AvcodecAlignDimensions Modify width and height values so that they will result in a memory buffer that is acceptable for the codec if you do not use any horizontal padding.
func (ctxt *Context) AvcodecAlignDimensions(w, h *int) {
	C.avcodec_align_dimensions((*C.struct_AVCodecContext)(ctxt), (*C.int)(unsafe.Pointer(w)), (*C.int)(unsafe.Pointer(h)))
}

// AvcodecAlignDimensions2 Modify width and height values so that they will result in a memory buffer that is acceptable for the codec if you also ensure that all line sizes are a multiple of the respective linesize_align[i].
func (ctxt *Context) AvcodecAlignDimensions2(w, h *int, l int) {
	C.avcodec_align_dimensions2((*C.struct_AVCodecContext)(ctxt), (*C.int)(unsafe.Pointer(w)), (*C.int)(unsafe.Pointer(h)), (*C.int)(unsafe.Pointer(&l)))
}

// AvcodecDecodeSubtitle2 Decode a subtitle message.
func (ctxt *Context) AvcodecDecodeSubtitle2(s *AvSubtitle, g *int, a *Packet) int {
	return int(C.avcodec_decode_subtitle2((*C.struct_AVCodecContext)(ctxt), (*C.struct_AVSubtitle)(s), (*C.int)(unsafe.Pointer(g)), (*C.struct_AVPacket)(a)))
}

// AvcodecEncodeSubtitle ...
func (ctxt *Context) AvcodecEncodeSubtitle(b *uint8, bs int, s *AvSubtitle) int {
	return int(C.avcodec_encode_subtitle((*C.struct_AVCodecContext)(ctxt), (*C.uint8_t)(b), C.int(bs), (*C.struct_AVSubtitle)(s)))
}

// AvcodecDefaultGetFormat ...
func (ctxt *Context) AvcodecDefaultGetFormat(f *PixelFormat) PixelFormat {
	return (PixelFormat)(C.avcodec_default_get_format((*C.struct_AVCodecContext)(ctxt), (*C.enum_AVPixelFormat)(f)))
}

// AvcodecFlushBuffers Reset the internal decoder state / flush internal buffers.
func (ctxt *Context) AvcodecFlushBuffers() {
	C.avcodec_flush_buffers((*C.struct_AVCodecContext)(ctxt))
}

// AvGetAudioFrameDuration Return audio frame duration.
func (ctxt *Context) AvGetAudioFrameDuration(f int) int {
	return int(C.av_get_audio_frame_duration((*C.struct_AVCodecContext)(ctxt), C.int(f)))
}

// AvcodecIsOpen Return a positive value if s is open (i.e. avcodec_open2() was called on it with no corresponding avcodec_close()), 0 otherwise.
func (ctxt *Context) AvcodecIsOpen() int {
	return int(C.avcodec_is_open((*C.struct_AVCodecContext)(ctxt)))
}

// AvParserParse2 Parse a packet.
func (ctxt *Context) AvParserParse2(ctxtp *ParserContext, p **uint8, ps *int, b *uint8, bs int, pt, dt, po int64) int {
	return int(C.av_parser_parse2((*C.struct_AVCodecParserContext)(ctxtp), (*C.struct_AVCodecContext)(ctxt), (**C.uint8_t)(unsafe.Pointer(p)), (*C.int)(unsafe.Pointer(ps)), (*C.uint8_t)(b), C.int(bs), (C.int64_t)(pt), (C.int64_t)(dt), (C.int64_t)(po)))
}

// AvParserChange Return 0 if the output buffer is a subset of the input, 1 if it is allocated and must be freed
func (ctxt *Context) AvParserChange(ctxtp *ParserContext, pb **uint8, pbs *int, b *uint8, bs, k int) int {
	return int(C.av_parser_change((*C.struct_AVCodecParserContext)(ctxtp), (*C.struct_AVCodecContext)(ctxt), (**C.uint8_t)(unsafe.Pointer(pb)), (*C.int)(unsafe.Pointer(pbs)), (*C.uint8_t)(b), C.int(bs), C.int(k)))
}

// AvParserInit ...
func AvParserInit(c int) *ParserContext {
	return (*ParserContext)(C.av_parser_init(C.int(c)))
}

// AvParserClose ...
func AvParserClose(ctxtp *ParserContext) {
	C.av_parser_close((*C.struct_AVCodecParserContext)(ctxtp))
}

// SetTimebase Set time_base
func (ctxt *Context) SetTimebase(num1 int, den1 int) {
	ctxt.time_base.num = C.int(num1)
	ctxt.time_base.den = C.int(den1)
}

// SetEncodeParams2 ...
func (ctxt *Context) SetEncodeParams2(width int, height int, pxlFmt PixelFormat, hasBframes bool, gopSize int) {
	ctxt.width = C.int(width)
	ctxt.height = C.int(height)
	// ctxt.bit_rate = 1000000
	ctxt.gop_size = C.int(gopSize)
	// ctxt.max_b_frames = 2
	if hasBframes {
		ctxt.has_b_frames = 1
	} else {
		ctxt.has_b_frames = 0
	}
	// ctxt.extradata = nil
	// ctxt.extradata_size = 0
	// ctxt.channels = 0
	ctxt.pix_fmt = int32(pxlFmt)
	// C.av_opt_set(ctxt.priv_data, "preset", "ultrafast", 0)
}

// SetEncodeParams ...
func (ctxt *Context) SetEncodeParams(width int, height int, pxlFmt PixelFormat) {
	ctxt.SetEncodeParams2(width, height, pxlFmt, false /*no b frames*/, 10)
}

// AvcodecSendPacket Supply raw packet data as input to a decoder.
func (ctxt *Context) AvcodecSendPacket(packet *Packet) int {
	return (int)(C.avcodec_send_packet((*C.struct_AVCodecContext)(ctxt), (*C.struct_AVPacket)(packet)))
}

// AvcodecReceiveFrame Return decoded output data from a decoder.
func (ctxt *Context) AvcodecReceiveFrame(frame *Frame) int {
	return (int)(C.avcodec_receive_frame((*C.struct_AVCodecContext)(ctxt), (*C.struct_AVFrame)(frame)))
}
