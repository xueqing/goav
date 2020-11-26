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
func (s *AVFormatContext) AvFormatInjectGlobalSideData() {
	C.av_format_inject_global_side_data((*C.struct_AVFormatContext)(s))
}

// AvFmtCtxGetDurationEstimationMethod Returns the method used to set ctx->duration.
func (s *AVFormatContext) AvFmtCtxGetDurationEstimationMethod() AvDurationEstimationMethod {
	return (AvDurationEstimationMethod)(C.av_fmt_ctx_get_duration_estimation_method((*C.struct_AVFormatContext)(s)))
}

// AvformatFreeContext Free an Context and all its streams.
func (s *AVFormatContext) AvformatFreeContext() {
	C.avformat_free_context((*C.struct_AVFormatContext)(s))
}

// AvformatNewStream Add a new stream to a media file.
func (s *AVFormatContext) AvformatNewStream(c *AvCodec) *AvStream {
	return (*AvStream)(C.avformat_new_stream((*C.struct_AVFormatContext)(s), (*C.struct_AVCodec)(c)))
}

// AvNewProgram ...
func (s *AVFormatContext) AvNewProgram(id int) *AvProgram {
	return (*AvProgram)(C.av_new_program((*C.struct_AVFormatContext)(s), C.int(id)))
}

// AvformatFindStreamInfo Read packets of a media file to get stream information.
func (s *AVFormatContext) AvformatFindStreamInfo(d **libavutil.AvDictionary) int {
	return int(C.avformat_find_stream_info((*C.struct_AVFormatContext)(s), (**C.struct_AVDictionary)(unsafe.Pointer(d))))
}

// AvFindProgramFromStream Find the programs which belong to a given stream.
func (s *AVFormatContext) AvFindProgramFromStream(l *AvProgram, su int) *AvProgram {
	return (*AvProgram)(C.av_find_program_from_stream((*C.struct_AVFormatContext)(s), (*C.struct_AVProgram)(l), C.int(su)))
}

// AvFindBestStream Find the "best" stream in the file.
func AvFindBestStream(ic *AVFormatContext, t AvMediaType, ws, rs int, c **AvCodec, f int) int {
	return int(C.av_find_best_stream((*C.struct_AVFormatContext)(ic), (C.enum_AVMediaType)(t), C.int(ws), C.int(rs), (**C.struct_AVCodec)(unsafe.Pointer(c)), C.int(f)))
}

// AvReadFrame Return the next frame of a stream.
func (s *AVFormatContext) AvReadFrame(pkt *libavcodec.AvPacket) int {
	return int(C.av_read_frame((*C.struct_AVFormatContext)(unsafe.Pointer(s)), toCPacket(pkt)))
}

// AvSeekFrame Seek to the keyframe at timestamp.
func (s *AVFormatContext) AvSeekFrame(st int, t int64, f int) int {
	return int(C.av_seek_frame((*C.struct_AVFormatContext)(s), C.int(st), C.int64_t(t), C.int(f)))
}

// AvSeekFrameTime seeks to a specified time location.
// |timebase| is codec specific and can be obtained by calling AvCodecGetPktTimebase2
func (s *AVFormatContext) AvSeekFrameTime(st int, at time.Duration, timebase libavcodec.AvRational) int {
	t2 := C.double(C.double(at.Seconds())*C.double(timebase.Den())) / (C.double(timebase.Num()))
	// log.Printf("Seeking to time :%v TimebaseTime:%v ActualTimebase:%v", at, t2, timebase)
	return int(C.av_seek_frame((*C.struct_AVFormatContext)(s), C.int(st), C.int64_t(t2), AvseekFlagBackward))
}

// AvformatSeekFile Seek to timestamp ts.
func (s *AVFormatContext) AvformatSeekFile(si int, mit, ts, mat int64, f int) int {
	return int(C.avformat_seek_file((*C.struct_AVFormatContext)(s), C.int(si), C.int64_t(mit), C.int64_t(ts), C.int64_t(mat), C.int(f)))
}

// AvReadPlay Start playing a network-based stream (e.g.
func (s *AVFormatContext) AvReadPlay() int {
	return int(C.av_read_play((*C.struct_AVFormatContext)(s)))
}

// AvReadPause Pause a network-based stream (e.g.
func (s *AVFormatContext) AvReadPause() int {
	return int(C.av_read_pause((*C.struct_AVFormatContext)(s)))
}

// AvformatCloseInput Close an opened input Context.
func (s *AVFormatContext) AvformatCloseInput() {
	C.avformat_close_input((**C.struct_AVFormatContext)(unsafe.Pointer(&s)))
}

// AvformatWriteHeader Allocate the stream private data and write the stream header to an output media file.
func (s *AVFormatContext) AvformatWriteHeader(o **libavutil.AvDictionary) int {
	return int(C.avformat_write_header((*C.struct_AVFormatContext)(s), (**C.struct_AVDictionary)(unsafe.Pointer(o))))
}

// AvWriteFrame Write a packet to an output media file.
func (s *AVFormatContext) AvWriteFrame(pkt *libavcodec.AvPacket) int {
	return int(C.av_write_frame((*C.struct_AVFormatContext)(s), toCPacket(pkt)))
}

// AvInterleavedWriteFrame Write a packet to an output media file ensuring correct interleaving.
func (s *AVFormatContext) AvInterleavedWriteFrame(pkt *libavcodec.AvPacket) int {
	return int(C.av_interleaved_write_frame((*C.struct_AVFormatContext)(s), toCPacket(pkt)))
}

// AvWriteUncodedFrame Write a uncoded frame to an output media file.
func (s *AVFormatContext) AvWriteUncodedFrame(si int, f *AvFrame) int {
	return int(C.av_write_uncoded_frame((*C.struct_AVFormatContext)(s), C.int(si), (*C.struct_AVFrame)(f)))
}

// AvInterleavedWriteUncodedFrame Write a uncoded frame to an output media file.
func (s *AVFormatContext) AvInterleavedWriteUncodedFrame(si int, f *AvFrame) int {
	return int(C.av_interleaved_write_uncoded_frame((*C.struct_AVFormatContext)(s), C.int(si), (*C.struct_AVFrame)(f)))
}

// AvWriteUncodedFrameQuery Test whether a muxer supports uncoded frame.
func (s *AVFormatContext) AvWriteUncodedFrameQuery(si int) int {
	return int(C.av_write_uncoded_frame_query((*C.struct_AVFormatContext)(s), C.int(si)))
}

// AvWriteTrailer Write the stream trailer to an output media file and free the file private data.
func (s *AVFormatContext) AvWriteTrailer() int {
	return int(C.av_write_trailer((*C.struct_AVFormatContext)(s)))
}

// AvGetOutputTimestamp Get timing information for the data currently output.
func (s *AVFormatContext) AvGetOutputTimestamp(st int, dts, wall *int) int {
	return int(C.av_get_output_timestamp((*C.struct_AVFormatContext)(s), C.int(st), (*C.int64_t)(unsafe.Pointer(&dts)), (*C.int64_t)(unsafe.Pointer(&wall))))
}

// AvFindDefaultStreamIndex ...
func (s *AVFormatContext) AvFindDefaultStreamIndex() int {
	return int(C.av_find_default_stream_index((*C.struct_AVFormatContext)(s)))
}

// AvDumpFormat Print detailed information about the input or output format, such as duration, bitrate, streams, container, programs, metadata, side data, codec and time base.
func (s *AVFormatContext) AvDumpFormat(i int, url string, io int) {
	Curl := C.CString(url)
	defer C.free(unsafe.Pointer(Curl))

	C.av_dump_format((*C.struct_AVFormatContext)(unsafe.Pointer(s)), C.int(i), Curl, C.int(io))
}

// AvGuessSampleAspectRatio Guess the sample aspect ratio of a frame, based on both the stream and the frame aspect ratio.
func (s *AVFormatContext) AvGuessSampleAspectRatio(st *AvStream, fr *AvFrame) libavcodec.AvRational {
	return newRational(C.av_guess_sample_aspect_ratio((*C.struct_AVFormatContext)(s), (*C.struct_AVStream)(st), (*C.struct_AVFrame)(fr)))
}

// AvGuessFrameRate Guess the frame rate, based on both the container and codec information.
func (s *AVFormatContext) AvGuessFrameRate(st *AvStream, fr *AvFrame) libavcodec.AvRational {
	return newRational(C.av_guess_frame_rate((*C.struct_AVFormatContext)(s), (*C.struct_AVStream)(st), (*C.struct_AVFrame)(fr)))
}

// AvformatMatchStreamSpecifier Check if the stream st contained in s is matched by the stream specifier spec.
func (s *AVFormatContext) AvformatMatchStreamSpecifier(st *AvStream, spec string) int {
	Cspec := C.CString(spec)
	defer C.free(unsafe.Pointer(Cspec))

	return int(C.avformat_match_stream_specifier((*C.struct_AVFormatContext)(s), (*C.struct_AVStream)(st), Cspec))
}

// AvformatQueueAttachedPictures ...
func (s *AVFormatContext) AvformatQueueAttachedPictures() int {
	return int(C.avformat_queue_attached_pictures((*C.struct_AVFormatContext)(s)))
}

// AvformatNewStream2 Add a new stream to a media file.
func (s *AVFormatContext) AvformatNewStream2(c *AvCodec) *AvStream {
	stream := (*AvStream)(C.avformat_new_stream((*C.struct_AVFormatContext)(s), (*C.struct_AVCodec)(c)))
	stream.codec.pix_fmt = int32(libavcodec.AvPixFmtYuv)
	stream.codec.width = 640
	stream.codec.height = 480
	stream.time_base.num = 1
	stream.time_base.num = 25
	return stream
}

// //av_format_control_message av_format_get_control_message_cb (const AVFormatContext *s)
// func (s *AVFormatContext) AvFormatControlMessage() C.av_format_get_control_message_cb {
// 	return C.av_format_get_control_message_cb((*C.struct_AVFormatContext)(s))
// }

// //void av_format_set_control_message_cb (AVFormatContext *s, av_format_control_message callback)
// func (s *AVFormatContext) AvFormatSetControlMessageCb(c AvFormatControlMessage) C.av_format_get_control_message_cb {
// 	C.av_format_set_control_message_cb((*C.struct_AVFormatContext)(s), (C.struct_AVFormatControlMessage)(c))
// }

// //AvCodec * av_format_get_data_codec (const AVFormatContext *s)
// func (s *AVFormatContext)AvFormatGetDataCodec() *AvCodec {
// 	return (*AvCodec)(C.av_format_get_data_codec((*C.struct_AVFormatContext)(s)))
// }

// //void av_format_set_data_codec (AVFormatContext *s, AvCodec *c)
// func (s *AVFormatContext)AvFormatSetDataCodec( c *AvCodec) {
// 	C.av_format_set_data_codec((*C.struct_AVFormatContext)(s), (*C.struct_AVCodec)(c))
// }
