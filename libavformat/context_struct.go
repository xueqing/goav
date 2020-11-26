// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package libavformat

//#cgo pkg-config: libavformat
//#include <libavformat/avformat.h>
import "C"
import (
	"reflect"
	"unsafe"

	"github.com/xueqing/goav/libavutil"
)

// Chapters Return chapters
func (ctxt *Context) Chapters() **AvChapter {
	return (**AvChapter)(unsafe.Pointer(ctxt.chapters))
}

// AudioCodec Return audio_codec
func (ctxt *Context) AudioCodec() *AvCodec {
	return (*AvCodec)(unsafe.Pointer(ctxt.audio_codec))
}

// SubtitleCodec Return subtitle_codec
func (ctxt *Context) SubtitleCodec() *AvCodec {
	return (*AvCodec)(unsafe.Pointer(ctxt.subtitle_codec))
}

// VideoCodec Return video_codec
func (ctxt *Context) VideoCodec() *AvCodec {
	return (*AvCodec)(unsafe.Pointer(ctxt.video_codec))
}

// Metadata Return metadata
func (ctxt *Context) Metadata() *libavutil.Dictionary {
	return (*libavutil.Dictionary)(unsafe.Pointer(ctxt.metadata))
}

// Internal Return internal
func (ctxt *Context) Internal() *AvFormatInternal {
	return (*AvFormatInternal)(unsafe.Pointer(ctxt.internal))
}

// Pb Return pb
func (ctxt *Context) Pb() *AvIOContext {
	return (*AvIOContext)(unsafe.Pointer(ctxt.pb))
}

// InterruptCallback Return interrupt_callback
func (ctxt *Context) InterruptCallback() AvIOInterruptCB {
	return AvIOInterruptCB(ctxt.interrupt_callback)
}

// Programs Return programs
func (ctxt *Context) Programs() []*AvProgram {
	header := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(ctxt.programs)),
		Len:  int(ctxt.NbPrograms()),
		Cap:  int(ctxt.NbPrograms()),
	}

	return *((*[]*AvProgram)(unsafe.Pointer(&header)))
}

// Streams Return streams
func (ctxt *Context) Streams() []*Stream {
	header := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(ctxt.streams)),
		Len:  int(ctxt.NbStreams()),
		Cap:  int(ctxt.NbStreams()),
	}

	return *((*[]*Stream)(unsafe.Pointer(&header)))
}

// Filename Return filename
func (ctxt *Context) Filename() string {
	return C.GoString((*C.char)(unsafe.Pointer(&ctxt.filename[0])))
}

// func (ctxt *Context) CodecWhitelist() string {
// 	return C.GoString(ctxt.codec_whitelist)
// }

// func (ctxt *Context) FormatWhitelist() string {
// 	return C.GoString(ctxt.format_whitelist)
// }

// AudioCodecID Return audio_codec_id
func (ctxt *Context) AudioCodecID() CodecID {
	return CodecID(ctxt.audio_codec_id)
}

// SubtitleCodecID Return subtitle_codec_id
func (ctxt *Context) SubtitleCodecID() CodecID {
	return CodecID(ctxt.subtitle_codec_id)
}

// VideoCodecID Return video_codec_id
func (ctxt *Context) VideoCodecID() CodecID {
	return CodecID(ctxt.video_codec_id)
}

// DurationEstimationMethod Return duration_estimation_method
func (ctxt *Context) DurationEstimationMethod() AvDurationEstimationMethod {
	return AvDurationEstimationMethod(ctxt.duration_estimation_method)
}

// AudioPreload Return audio_preload
func (ctxt *Context) AudioPreload() int {
	return int(ctxt.audio_preload)
}

// AvioFlags Return avio_flags
func (ctxt *Context) AvioFlags() int {
	return int(ctxt.avio_flags)
}

// AvoidNegativeTs Return avoid_negative_ts
func (ctxt *Context) AvoidNegativeTs() int {
	return int(ctxt.avoid_negative_ts)
}

// BitRate Return bit_rate
func (ctxt *Context) BitRate() int {
	return int(ctxt.bit_rate)
}

// CtxFlags Return ctx_flags
func (ctxt *Context) CtxFlags() int {
	return int(ctxt.ctx_flags)
}

// Debug Return debug
func (ctxt *Context) Debug() int {
	return int(ctxt.debug)
}

// ErrorRecognition Return error_recognition
func (ctxt *Context) ErrorRecognition() int {
	return int(ctxt.error_recognition)
}

// EventFlags Return event_flags
func (ctxt *Context) EventFlags() int {
	return int(ctxt.event_flags)
}

// Flags Return flags
func (ctxt *Context) Flags() int {
	return int(ctxt.flags)
}

// FlushPackets Return flush_packets
func (ctxt *Context) FlushPackets() int {
	return int(ctxt.flush_packets)
}

// FormatProbesize Return format_probesize
func (ctxt *Context) FormatProbesize() int {
	return int(ctxt.format_probesize)
}

// FpsProbeSize Return fps_probe_size
func (ctxt *Context) FpsProbeSize() int {
	return int(ctxt.fps_probe_size)
}

// IoRepositioned Return io_repositioned
func (ctxt *Context) IoRepositioned() int {
	return int(ctxt.io_repositioned)
}

// Keylen Return keylen
func (ctxt *Context) Keylen() int {
	return int(ctxt.keylen)
}

// MaxChunkDuration Return max_chunk_duration
func (ctxt *Context) MaxChunkDuration() int {
	return int(ctxt.max_chunk_duration)
}

// MaxChunkSize Return max_chunk_size
func (ctxt *Context) MaxChunkSize() int {
	return int(ctxt.max_chunk_size)
}

// MaxDelay Return max_delay
func (ctxt *Context) MaxDelay() int {
	return int(ctxt.max_delay)
}

// MaxTsProbe Return max_ts_probe
func (ctxt *Context) MaxTsProbe() int {
	return int(ctxt.max_ts_probe)
}

// MetadataHeaderPadding Return metadata_header_padding
func (ctxt *Context) MetadataHeaderPadding() int {
	return int(ctxt.metadata_header_padding)
}

// ProbeScore Return probe_score
func (ctxt *Context) ProbeScore() int {
	return int(ctxt.probe_score)
}

// Seek2any Return seek2any
func (ctxt *Context) Seek2any() int {
	return int(ctxt.seek2any)
}

// StrictStdCompliance Return strict_std_compliance
func (ctxt *Context) StrictStdCompliance() int {
	return int(ctxt.strict_std_compliance)
}

// TsID Return ts_id
func (ctxt *Context) TsID() int {
	return int(ctxt.ts_id)
}

// UseWallclockAsTimestamps Return use_wallclock_as_timestamps
func (ctxt *Context) UseWallclockAsTimestamps() int {
	return int(ctxt.use_wallclock_as_timestamps)
}

// Duration Return duration
func (ctxt *Context) Duration() int64 {
	return int64(ctxt.duration)
}

// MaxAnalyzeDuration2 Return max_analyze_duration
func (ctxt *Context) MaxAnalyzeDuration2() int64 {
	return int64(ctxt.max_analyze_duration)
}

// MaxInterleaveDelta Return max_interleave_delta
func (ctxt *Context) MaxInterleaveDelta() int64 {
	return int64(ctxt.max_interleave_delta)
}

// OutputTsOffset Return output_ts_offset
func (ctxt *Context) OutputTsOffset() int64 {
	return int64(ctxt.output_ts_offset)
}

// Probesize2 Return probesize
func (ctxt *Context) Probesize2() int64 {
	return int64(ctxt.probesize)
}

// SkipInitialBytes Return skip_initial_bytes
func (ctxt *Context) SkipInitialBytes() int64 {
	return int64(ctxt.skip_initial_bytes)
}

// StartTime Return start_time
func (ctxt *Context) StartTime() int64 {
	return int64(ctxt.start_time)
}

// StartTimeRealtime Return start_time_realtime
func (ctxt *Context) StartTimeRealtime() int64 {
	return int64(ctxt.start_time_realtime)
}

// Iformat Return iformat
func (ctxt *Context) Iformat() *InputFormat {
	return (*InputFormat)(unsafe.Pointer(ctxt.iformat))
}

// Oformat Return oformat
func (ctxt *Context) Oformat() *OutputFormat {
	return (*OutputFormat)(unsafe.Pointer(ctxt.oformat))
}

// func (ctxt *Context) DumpSeparator() uint8 {
// 	return uint8(ctxt.dump_separator)
// }

// CorrectTsOverflow Return correct_ts_overflow
func (ctxt *Context) CorrectTsOverflow() int {
	return int(ctxt.correct_ts_overflow)
}

// MaxIndexSize Return max_index_size
func (ctxt *Context) MaxIndexSize() uint {
	return uint(ctxt.max_index_size)
}

// MaxPictureBuffer Return max_picture_buffer
func (ctxt *Context) MaxPictureBuffer() uint {
	return uint(ctxt.max_picture_buffer)
}

// NbChapters Return nb_chapters
func (ctxt *Context) NbChapters() uint {
	return uint(ctxt.nb_chapters)
}

// NbPrograms Return nb_programs
func (ctxt *Context) NbPrograms() uint {
	return uint(ctxt.nb_programs)
}

// NbStreams Return nb_streams
func (ctxt *Context) NbStreams() uint {
	return uint(ctxt.nb_streams)
}

// PacketSize Return packet_size
func (ctxt *Context) PacketSize() uint {
	return uint(ctxt.packet_size)
}

// Probesize Return probesize
func (ctxt *Context) Probesize() uint {
	return uint(ctxt.probesize)
}

// SetPb Set pb
func (ctxt *Context) SetPb(pb *AvIOContext) {
	ctxt.pb = (*C.struct_AVIOContext)(unsafe.Pointer(pb))
}

// Pb2 Return pb
func (ctxt *Context) Pb2() **AvIOContext {
	return (**AvIOContext)(unsafe.Pointer(&ctxt.pb))
}
