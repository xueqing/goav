// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package libavcodec

//#cgo pkg-config: libavcodec
//#include <libavcodec/avcodec.h>
import "C"
import (
	"unsafe"
)

// AvcodecFreeContext Free the codec context and everything associated with it
// and write NULL to the provided pointer.
func (cctx *AvCodecContext) AvcodecFreeContext() {
	C.avcodec_free_context((**C.struct_AVCodecContext)(unsafe.Pointer(&cctx)))
}

// AvcodecGetContextDefaults3 Set the fields of the given Context to default values
// corresponding to the given codec (defaults may be codec-dependent).
func (cctx *AvCodecContext) AvcodecGetContextDefaults3(codec *AvCodec) int {
	return int(C.avcodec_get_context_defaults3((*C.struct_AVCodecContext)(cctx),
		(*C.struct_AVCodec)(codec)))
}

// AvcodecParametersFromContext Fill the parameters struct based on the values
// from the supplied codec context.
func (cctx *AvCodecContext) AvcodecParametersFromContext(par *AvCodecParameters) int {
	return int(C.avcodec_parameters_from_context((*C.struct_AVCodecParameters)(par),
		(*C.struct_AVCodecContext)(cctx)))
}

// AvcodecParametersToContext Fill the codec context based on the values
// from the supplied codec parameters.
func (cctx *AvCodecContext) AvcodecParametersToContext(par *AvCodecParameters) int {
	return int(C.avcodec_parameters_to_context((*C.struct_AVCodecContext)(cctx),
		(*C.struct_AVCodecParameters)(par)))
}

// AvcodecOpen2 Initialize the Context to use the given Codec
func (cctx *AvCodecContext) AvcodecOpen2(codec *AvCodec, options **AvDictionary) int {
	return int(C.avcodec_open2((*C.struct_AVCodecContext)(cctx),
		(*C.struct_AVCodec)(codec), (**C.struct_AVDictionary)(unsafe.Pointer(options))))
}

// AvcodecClose Close a given Context and free all the data associated with it
// (but not the Context itself).
func (cctx *AvCodecContext) AvcodecClose() int {
	return int(C.avcodec_close((*C.struct_AVCodecContext)(cctx)))
}

// AvcodecDefaultGetBuffer2 The default callback for Context.get_buffer2().
func (cctx *AvCodecContext) AvcodecDefaultGetBuffer2(frame *AvFrame, flags int) int {
	return int(C.avcodec_default_get_buffer2((*C.struct_AVCodecContext)(cctx),
		(*C.struct_AVFrame)(frame), C.int(flags)))
}

// AvcodecAlignDimensions Modify width and height values so that they will result in
// a memory buffer that is acceptable for the codec if you do not use any horizontal padding.
func (cctx *AvCodecContext) AvcodecAlignDimensions(width, height *int) {
	C.avcodec_align_dimensions((*C.struct_AVCodecContext)(cctx),
		(*C.int)(unsafe.Pointer(width)), (*C.int)(unsafe.Pointer(height)))
}

// AvcodecAlignDimensions2 Modify width and height values so that they will
// result in a memory buffer that is acceptable for the codec if you also
// ensure that all line sizes are a multiple of the respective linesize_align[i].
func (cctx *AvCodecContext) AvcodecAlignDimensions2(width, height *int, linesizeAlign int) {
	C.avcodec_align_dimensions2((*C.struct_AVCodecContext)(cctx),
		(*C.int)(unsafe.Pointer(width)), (*C.int)(unsafe.Pointer(height)),
		(*C.int)(unsafe.Pointer(&linesizeAlign)))
}

// AvcodecDecodeSubtitle2 Decode a subtitle message.
func (cctx *AvCodecContext) AvcodecDecodeSubtitle2(sub *AvSubtitle, gotSubPtr *int, pkt *AvPacket) int {
	return int(C.avcodec_decode_subtitle2((*C.struct_AVCodecContext)(cctx),
		(*C.struct_AVSubtitle)(sub), (*C.int)(unsafe.Pointer(gotSubPtr)),
		(*C.struct_AVPacket)(pkt)))
}

// AvcodecSendPacket Supply raw packet data as input to a decoder.
func (cctx *AvCodecContext) AvcodecSendPacket(pkt *AvPacket) int {
	return (int)(C.avcodec_send_packet((*C.struct_AVCodecContext)(cctx), (*C.struct_AVPacket)(pkt)))
}

// AvcodecReceiveFrame Return decoded output data from a decoder.
func (cctx *AvCodecContext) AvcodecReceiveFrame(frame *AvFrame) int {
	return (int)(C.avcodec_receive_frame((*C.struct_AVCodecContext)(cctx), (*C.struct_AVFrame)(frame)))
}

// AvcodecSendFrame Supply a raw video or audio frame to the encoder.
func (cctx *AvCodecContext) AvcodecSendFrame(frame *AvFrame) int {
	return (int)(C.avcodec_send_frame((*C.struct_AVCodecContext)(cctx), (*C.struct_AVFrame)(frame)))
}

// AvcodecReceivePacket Return decoded output data from a decoder.
func (cctx *AvCodecContext) AvcodecReceivePacket(pkt *AvPacket) int {
	return (int)(C.avcodec_receive_packet((*C.struct_AVCodecContext)(cctx), (*C.struct_AVPacket)(pkt)))
}

// AvParserInit ...
func AvParserInit(codecID int) *AvCodecParserContext {
	return (*AvCodecParserContext)(C.av_parser_init(C.int(codecID)))
}

// AvParserParse2 Parse a packet.
func (cctx *AvCodecContext) AvParserParse2(cpctx *AvCodecParserContext, poutbuf **uint8, poutbufSize *int, buf *uint8, bufSize int, pts, dts, pos int64) int {
	return int(C.av_parser_parse2((*C.struct_AVCodecParserContext)(cpctx),
		(*C.struct_AVCodecContext)(cctx), (**C.uint8_t)(unsafe.Pointer(poutbuf)),
		(*C.int)(unsafe.Pointer(poutbufSize)), (*C.uint8_t)(buf), C.int(bufSize),
		(C.int64_t)(pts), (C.int64_t)(dts), (C.int64_t)(pos)))
}

// AvParserChange Return 0 if the output buffer is a subset of the input, 1 if it is allocated and must be freed
func (cctx *AvCodecContext) AvParserChange(cpctx *AvCodecParserContext, poutbuf **uint8, poutbufSize *int, buf *uint8, bufSize, keyFrame int) int {
	return int(C.av_parser_change((*C.struct_AVCodecParserContext)(cpctx),
		(*C.struct_AVCodecContext)(cctx), (**C.uint8_t)(unsafe.Pointer(poutbuf)),
		(*C.int)(unsafe.Pointer(poutbufSize)), (*C.uint8_t)(buf), C.int(bufSize), C.int(keyFrame)))
}

// AvParserClose ...
func AvParserClose(cpctx *AvCodecParserContext) {
	C.av_parser_close((*C.struct_AVCodecParserContext)(cpctx))
}

// AvcodecEncodeSubtitle ...
func (cctx *AvCodecContext) AvcodecEncodeSubtitle(buf *uint8, bufSize int, sub *AvSubtitle) int {
	return int(C.avcodec_encode_subtitle((*C.struct_AVCodecContext)(cctx),
		(*C.uint8_t)(buf), C.int(bufSize), (*C.struct_AVSubtitle)(sub)))
}

// AvcodecDefaultGetFormat ...
func (cctx *AvCodecContext) AvcodecDefaultGetFormat(pixFmt *AvPixelFormat) AvPixelFormat {
	return (AvPixelFormat)(C.avcodec_default_get_format((*C.struct_AVCodecContext)(cctx),
		(*C.enum_AVPixelFormat)(pixFmt)))
}

// AvcodecFlushBuffers Reset the internal decoder state / flush internal buffers.
func (cctx *AvCodecContext) AvcodecFlushBuffers() {
	C.avcodec_flush_buffers((*C.struct_AVCodecContext)(cctx))
}

// AvGetAudioFrameDuration Return audio frame duration.
func (cctx *AvCodecContext) AvGetAudioFrameDuration(frameBytes int) int {
	return int(C.av_get_audio_frame_duration((*C.struct_AVCodecContext)(cctx), C.int(frameBytes)))
}

// AvcodecIsOpen Return a positive value if s is open (i.e. avcodec_open2() was called on it with no corresponding avcodec_close()), 0 otherwise.
func (cctx *AvCodecContext) AvcodecIsOpen() int {
	return int(C.avcodec_is_open((*C.struct_AVCodecContext)(cctx)))
}

// AvcodecEncodeAudio2 deprecated
func (cctx *AvCodecContext) AvcodecEncodeAudio2(pkt *AvPacket, frame *AvFrame, gotPacketPtr *int) int {
	var gotPacket C.int
	ret := int(C.avcodec_encode_audio2((*C.struct_AVCodecContext)(cctx),
		(*C.struct_AVPacket)(pkt), (*C.struct_AVFrame)(frame), &gotPacket))
	*gotPacketPtr = int(gotPacket)
	return ret
}

// AvcodecEncodeVideo2 deprecated
func (cctx *AvCodecContext) AvcodecEncodeVideo2(pkt *AvPacket, frame *AvFrame, gotPacketPtr *int) int {
	var gotPacket C.int
	ret := int(C.avcodec_encode_video2((*C.struct_AVCodecContext)(cctx),
		(*C.struct_AVPacket)(pkt), (*C.struct_AVFrame)(frame), &gotPacket))
	*gotPacketPtr = int(gotPacket)
	return ret
}

// AvcodecDecodeAudio4 deprecated
func (cctx *AvCodecContext) AvcodecDecodeAudio4(frame *AvFrame, gotFramePtr *int, pkt *AvPacket) int {
	var gotFrame C.int
	ret := int(C.avcodec_decode_audio4((*C.struct_AVCodecContext)(cctx),
		(*C.struct_AVFrame)(frame), &gotFrame, (*C.struct_AVPacket)(pkt)))
	*gotFramePtr = int(gotFrame)
	return ret
}

// AvcodecDecodeVideo2 deprecated
func (cctx *AvCodecContext) AvcodecDecodeVideo2(frame *AvFrame, gotFramePtr *int, pkt *AvPacket) int {
	var gotFrame C.int
	ret := int(C.avcodec_decode_video2((*C.struct_AVCodecContext)(cctx),
		(*C.struct_AVFrame)(frame), &gotFrame, (*C.struct_AVPacket)(pkt)))
	*gotFramePtr = int(gotFrame)
	return ret
}
