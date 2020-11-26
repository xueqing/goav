// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package libavformat

//#cgo pkg-config: libavformat
//#include <libavformat/avformat.h>
import "C"
import (
	"time"
	"unsafe"

	"github.com/xueqing/goav/libavcodec"
	"github.com/xueqing/goav/libavutil"
)

// AvSeekFlagxxx
const (
	AvseekFlagBackward = 1 ///< seek backward
	AvseekFlagByte     = 2 ///< seeking based on position in bytes
	AvseekFlagAny      = 4 ///< seek to any frame, even non-keyframes
	AvseekFlagFrame    = 8 ///< seeking based on frame number
)

// AvFormatInjectGlobalSideData This function will cause global side data to be injected in the next packet of each stream as well as after any subsequent seek.
func (fctx *AVFormatContext) AvFormatInjectGlobalSideData() {
	C.av_format_inject_global_side_data((*C.struct_AVFormatContext)(fctx))
}

// AvFmtCtxGetDurationEstimationMethod Returns the method used to set ctx->duration.
func (fctx *AVFormatContext) AvFmtCtxGetDurationEstimationMethod() AvDurationEstimationMethod {
	return (AvDurationEstimationMethod)(C.av_fmt_ctx_get_duration_estimation_method((*C.struct_AVFormatContext)(fctx)))
}

// AvformatFreeContext Free an Context and all its streams.
func (fctx *AVFormatContext) AvformatFreeContext() {
	C.avformat_free_context((*C.struct_AVFormatContext)(fctx))
}

// AvformatNewStream Add a new stream to a media file.
func (fctx *AVFormatContext) AvformatNewStream(codec *AvCodec) *AvStream {
	return (*AvStream)(C.avformat_new_stream((*C.struct_AVFormatContext)(fctx),
		(*C.struct_AVCodec)(codec)))
}

// AvNewProgram ...
func (fctx *AVFormatContext) AvNewProgram(id int) *AvProgram {
	return (*AvProgram)(C.av_new_program((*C.struct_AVFormatContext)(fctx), C.int(id)))
}

// AvformatFindStreamInfo Read packets of a media file to get stream information.
func (fctx *AVFormatContext) AvformatFindStreamInfo(options **libavutil.AvDictionary) int {
	return int(C.avformat_find_stream_info((*C.struct_AVFormatContext)(fctx),
		(**C.struct_AVDictionary)(unsafe.Pointer(options))))
}

// AvFindProgramFromStream Find the programs which belong to a given stream.
func (fctx *AVFormatContext) AvFindProgramFromStream(last *AvProgram, streamIdx int) *AvProgram {
	return (*AvProgram)(C.av_find_program_from_stream((*C.struct_AVFormatContext)(fctx),
		(*C.struct_AVProgram)(last), C.int(streamIdx)))
}

// AvFindBestStream Find the "best" stream in the file.
func AvFindBestStream(fctx *AVFormatContext, typ AvMediaType, wantedStreamNb, relatedStream int, decoderRet **AvCodec, flags int) int {
	return int(C.av_find_best_stream((*C.struct_AVFormatContext)(fctx),
		(C.enum_AVMediaType)(typ), C.int(wantedStreamNb), C.int(relatedStream),
		(**C.struct_AVCodec)(unsafe.Pointer(decoderRet)), C.int(flags)))
}

// AvReadFrame Return the next frame of a stream.
func (fctx *AVFormatContext) AvReadFrame(pkt *libavcodec.AvPacket) int {
	return int(C.av_read_frame((*C.struct_AVFormatContext)(unsafe.Pointer(fctx)), toCPacket(pkt)))
}

// AvSeekFrame Seek to the keyframe at timestamp.
func (fctx *AVFormatContext) AvSeekFrame(streamIdx int, timestamp int64, flags int) int {
	return int(C.av_seek_frame((*C.struct_AVFormatContext)(fctx),
		C.int(streamIdx), C.int64_t(timestamp), C.int(flags)))
}

// AvSeekFrameTime seeks to a specified time location.
// |timebase| is codec specific and can be obtained by calling AvCodecGetPktTimebase2
func (fctx *AVFormatContext) AvSeekFrameTime(streamIdx int, dura time.Duration, timebase libavcodec.AvRational) int {
	timestamp := C.double(C.double(dura.Seconds())*C.double(timebase.Den())) / (C.double(timebase.Num()))
	// log.Printf("Seeking to time :%v TimebaseTime:%v ActualTimebase:%v", dura, timestamp, timebase)
	return int(C.av_seek_frame((*C.struct_AVFormatContext)(fctx),
		C.int(streamIdx), C.int64_t(timestamp), AvseekFlagBackward))
}

// AvformatSeekFile Seek to timestamp ts.
func (fctx *AVFormatContext) AvformatSeekFile(streamIdx int, minTs, ts, maxts int64, flags int) int {
	return int(C.avformat_seek_file((*C.struct_AVFormatContext)(fctx),
		C.int(streamIdx), C.int64_t(minTs), C.int64_t(ts), C.int64_t(maxts), C.int(flags)))
}

// AvReadPlay Start playing a network-based stream (e.g.
func (fctx *AVFormatContext) AvReadPlay() int {
	return int(C.av_read_play((*C.struct_AVFormatContext)(fctx)))
}

// AvReadPause Pause a network-based stream (e.g.
func (fctx *AVFormatContext) AvReadPause() int {
	return int(C.av_read_pause((*C.struct_AVFormatContext)(fctx)))
}

// AvformatCloseInput Close an opened input Context.
func (fctx *AVFormatContext) AvformatCloseInput() {
	C.avformat_close_input((**C.struct_AVFormatContext)(unsafe.Pointer(&fctx)))
}

// AvformatWriteHeader Allocate the stream private data and write the stream header to an output media file.
func (fctx *AVFormatContext) AvformatWriteHeader(options **libavutil.AvDictionary) int {
	return int(C.avformat_write_header((*C.struct_AVFormatContext)(fctx),
		(**C.struct_AVDictionary)(unsafe.Pointer(options))))
}

// AvWriteFrame Write a packet to an output media file.
func (fctx *AVFormatContext) AvWriteFrame(pkt *libavcodec.AvPacket) int {
	return int(C.av_write_frame((*C.struct_AVFormatContext)(fctx), toCPacket(pkt)))
}

// AvInterleavedWriteFrame Write a packet to an output media file ensuring correct interleaving.
func (fctx *AVFormatContext) AvInterleavedWriteFrame(pkt *libavcodec.AvPacket) int {
	return int(C.av_interleaved_write_frame((*C.struct_AVFormatContext)(fctx), toCPacket(pkt)))
}

// AvWriteUncodedFrame Write a uncoded frame to an output media file.
func (fctx *AVFormatContext) AvWriteUncodedFrame(streamIdx int, frame *AvFrame) int {
	return int(C.av_write_uncoded_frame((*C.struct_AVFormatContext)(fctx),
		C.int(streamIdx), (*C.struct_AVFrame)(frame)))
}

// AvInterleavedWriteUncodedFrame Write a uncoded frame to an output media file.
func (fctx *AVFormatContext) AvInterleavedWriteUncodedFrame(streamIdx int, frame *AvFrame) int {
	return int(C.av_interleaved_write_uncoded_frame((*C.struct_AVFormatContext)(fctx),
		C.int(streamIdx), (*C.struct_AVFrame)(frame)))
}

// AvWriteUncodedFrameQuery Test whether a muxer supports uncoded frame.
func (fctx *AVFormatContext) AvWriteUncodedFrameQuery(streamIdx int) int {
	return int(C.av_write_uncoded_frame_query((*C.struct_AVFormatContext)(fctx), C.int(streamIdx)))
}

// AvWriteTrailer Write the stream trailer to an output media file and free the file private data.
func (fctx *AVFormatContext) AvWriteTrailer() int {
	return int(C.av_write_trailer((*C.struct_AVFormatContext)(fctx)))
}

// AvGetOutputTimestamp Get timing information for the data currently output.
func (fctx *AVFormatContext) AvGetOutputTimestamp(stream int, dts, wall *int) int {
	return int(C.av_get_output_timestamp((*C.struct_AVFormatContext)(fctx),
		C.int(stream), (*C.int64_t)(unsafe.Pointer(&dts)), (*C.int64_t)(unsafe.Pointer(&wall))))
}

// AvFindDefaultStreamIndex ...
func (fctx *AVFormatContext) AvFindDefaultStreamIndex() int {
	return int(C.av_find_default_stream_index((*C.struct_AVFormatContext)(fctx)))
}

// AvDumpFormat Print detailed information about the input or output format, such as duration, bitrate, streams, container, programs, metadata, side data, codec and time base.
func (fctx *AVFormatContext) AvDumpFormat(idx int, url string, isOutput int) {
	cURL := C.CString(url)
	defer C.free(unsafe.Pointer(cURL))

	C.av_dump_format((*C.struct_AVFormatContext)(unsafe.Pointer(fctx)),
		C.int(idx), cURL, C.int(isOutput))
}

// AvGuessSampleAspectRatio Guess the sample aspect ratio of a frame, based on both the stream and the frame aspect ratio.
func (fctx *AVFormatContext) AvGuessSampleAspectRatio(st *AvStream, frame *AvFrame) libavcodec.AvRational {
	return newRational(C.av_guess_sample_aspect_ratio((*C.struct_AVFormatContext)(fctx),
		(*C.struct_AVStream)(st), (*C.struct_AVFrame)(frame)))
}

// AvGuessFrameRate Guess the frame rate, based on both the container and codec information.
func (fctx *AVFormatContext) AvGuessFrameRate(st *AvStream, frame *AvFrame) libavcodec.AvRational {
	return newRational(C.av_guess_frame_rate((*C.struct_AVFormatContext)(fctx),
		(*C.struct_AVStream)(st), (*C.struct_AVFrame)(frame)))
}

// AvformatMatchStreamSpecifier Check if the stream st contained in fctx is matched by the stream specifier spec.
func (fctx *AVFormatContext) AvformatMatchStreamSpecifier(st *AvStream, spec string) int {
	cSpec := C.CString(spec)
	defer C.free(unsafe.Pointer(cSpec))

	return int(C.avformat_match_stream_specifier((*C.struct_AVFormatContext)(fctx),
		(*C.struct_AVStream)(st), cSpec))
}

// AvformatQueueAttachedPictures ...
func (fctx *AVFormatContext) AvformatQueueAttachedPictures() int {
	return int(C.avformat_queue_attached_pictures((*C.struct_AVFormatContext)(fctx)))
}

// AvformatNewStream2 Add a new stream to a media file.
func (fctx *AVFormatContext) AvformatNewStream2(codec *AvCodec) *AvStream {
	stream := (*AvStream)(C.avformat_new_stream((*C.struct_AVFormatContext)(fctx), (*C.struct_AVCodec)(codec)))
	stream.codec.pix_fmt = int32(libavcodec.AvPixFmtYuv)
	stream.codec.width = 640
	stream.codec.height = 480
	stream.time_base.num = 1
	stream.time_base.num = 25
	return stream
}

// //av_format_control_message av_format_get_control_message_cb (const AVFormatContext *fctx)
// func (fctx *AVFormatContext) AvFormatControlMessage() C.av_format_get_control_message_cb {
// 	return C.av_format_get_control_message_cb((*C.struct_AVFormatContext)(fctx))
// }

// //void av_format_set_control_message_cb (AVFormatContext *fctx, av_format_control_message callback)
// func (fctx *AVFormatContext) AvFormatSetControlMessageCb(c AvFormatControlMessage) C.av_format_get_control_message_cb {
// 	C.av_format_set_control_message_cb((*C.struct_AVFormatContext)(fctx), (C.struct_AVFormatControlMessage)(c))
// }

// //AvCodec * av_format_get_data_codec (const AVFormatContext *fctx)
// func (fctx *AVFormatContext)AvFormatGetDataCodec() *AvCodec {
// 	return (*AvCodec)(C.av_format_get_data_codec((*C.struct_AVFormatContext)(fctx)))
// }

// //void av_format_set_data_codec (AVFormatContext *fctx, AvCodec *c)
// func (fctx *AVFormatContext)AvFormatSetDataCodec( c *AvCodec) {
// 	C.av_format_set_data_codec((*C.struct_AVFormatContext)(fctx), (*C.struct_AVCodec)(c))
// }
