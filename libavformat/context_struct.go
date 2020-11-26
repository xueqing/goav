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
func (fctx *AvFormatContext) Chapters() **AvChapter {
	return (**AvChapter)(unsafe.Pointer(fctx.chapters))
}

// AudioCodec Return audio_codec
func (fctx *AvFormatContext) AudioCodec() *AvCodec {
	return (*AvCodec)(unsafe.Pointer(fctx.audio_codec))
}

// SubtitleCodec Return subtitle_codec
func (fctx *AvFormatContext) SubtitleCodec() *AvCodec {
	return (*AvCodec)(unsafe.Pointer(fctx.subtitle_codec))
}

// VideoCodec Return video_codec
func (fctx *AvFormatContext) VideoCodec() *AvCodec {
	return (*AvCodec)(unsafe.Pointer(fctx.video_codec))
}

// Metadata Return metadata
func (fctx *AvFormatContext) Metadata() *libavutil.AvDictionary {
	return (*libavutil.AvDictionary)(unsafe.Pointer(fctx.metadata))
}

// Internal Return internal
func (fctx *AvFormatContext) Internal() *AvFormatInternal {
	return (*AvFormatInternal)(unsafe.Pointer(fctx.internal))
}

// Pb Return pb
func (fctx *AvFormatContext) Pb() *AvIOContext {
	return (*AvIOContext)(unsafe.Pointer(fctx.pb))
}

// InterruptCallback Return interrupt_callback
func (fctx *AvFormatContext) InterruptCallback() AvIOInterruptCB {
	return AvIOInterruptCB(fctx.interrupt_callback)
}

// Programs Return programs
func (fctx *AvFormatContext) Programs() []*AvProgram {
	header := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(fctx.programs)),
		Len:  int(fctx.NbPrograms()),
		Cap:  int(fctx.NbPrograms()),
	}

	return *((*[]*AvProgram)(unsafe.Pointer(&header)))
}

// Streams Return streams
func (fctx *AvFormatContext) Streams() []*AvStream {
	header := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(fctx.streams)),
		Len:  int(fctx.NbStreams()),
		Cap:  int(fctx.NbStreams()),
	}

	return *((*[]*AvStream)(unsafe.Pointer(&header)))
}

// Filename Return filename
func (fctx *AvFormatContext) Filename() string {
	return C.GoString((*C.char)(unsafe.Pointer(&fctx.filename[0])))
}

// func (fctx *AvFormatContext) CodecWhitelist() string {
// 	return C.GoString(fctx.codec_whitelist)
// }

// func (fctx *AvFormatContext) FormatWhitelist() string {
// 	return C.GoString(fctx.format_whitelist)
// }

// AudioCodecID Return audio_codec_id
func (fctx *AvFormatContext) AudioCodecID() AvCodecID {
	return AvCodecID(fctx.audio_codec_id)
}

// SubtitleCodecID Return subtitle_codec_id
func (fctx *AvFormatContext) SubtitleCodecID() AvCodecID {
	return AvCodecID(fctx.subtitle_codec_id)
}

// VideoCodecID Return video_codec_id
func (fctx *AvFormatContext) VideoCodecID() AvCodecID {
	return AvCodecID(fctx.video_codec_id)
}

// DurationEstimationMethod Return duration_estimation_method
func (fctx *AvFormatContext) DurationEstimationMethod() AvDurationEstimationMethod {
	return AvDurationEstimationMethod(fctx.duration_estimation_method)
}

// AudioPreload Return audio_preload
func (fctx *AvFormatContext) AudioPreload() int {
	return int(fctx.audio_preload)
}

// AvioFlags Return avio_flags
func (fctx *AvFormatContext) AvioFlags() int {
	return int(fctx.avio_flags)
}

// AvoidNegativeTs Return avoid_negative_ts
func (fctx *AvFormatContext) AvoidNegativeTs() int {
	return int(fctx.avoid_negative_ts)
}

// BitRate Return bit_rate
func (fctx *AvFormatContext) BitRate() int {
	return int(fctx.bit_rate)
}

// CtxFlags Return ctx_flags
func (fctx *AvFormatContext) CtxFlags() int {
	return int(fctx.ctx_flags)
}

// Debug Return debug
func (fctx *AvFormatContext) Debug() int {
	return int(fctx.debug)
}

// ErrorRecognition Return error_recognition
func (fctx *AvFormatContext) ErrorRecognition() int {
	return int(fctx.error_recognition)
}

// EventFlags Return event_flags
func (fctx *AvFormatContext) EventFlags() int {
	return int(fctx.event_flags)
}

// Flags Return flags
func (fctx *AvFormatContext) Flags() int {
	return int(fctx.flags)
}

// FlushPackets Return flush_packets
func (fctx *AvFormatContext) FlushPackets() int {
	return int(fctx.flush_packets)
}

// FormatProbesize Return format_probesize
func (fctx *AvFormatContext) FormatProbesize() int {
	return int(fctx.format_probesize)
}

// FpsProbeSize Return fps_probe_size
func (fctx *AvFormatContext) FpsProbeSize() int {
	return int(fctx.fps_probe_size)
}

// IoRepositioned Return io_repositioned
func (fctx *AvFormatContext) IoRepositioned() int {
	return int(fctx.io_repositioned)
}

// Keylen Return keylen
func (fctx *AvFormatContext) Keylen() int {
	return int(fctx.keylen)
}

// MaxChunkDuration Return max_chunk_duration
func (fctx *AvFormatContext) MaxChunkDuration() int {
	return int(fctx.max_chunk_duration)
}

// MaxChunkSize Return max_chunk_size
func (fctx *AvFormatContext) MaxChunkSize() int {
	return int(fctx.max_chunk_size)
}

// MaxDelay Return max_delay
func (fctx *AvFormatContext) MaxDelay() int {
	return int(fctx.max_delay)
}

// MaxTsProbe Return max_ts_probe
func (fctx *AvFormatContext) MaxTsProbe() int {
	return int(fctx.max_ts_probe)
}

// MetadataHeaderPadding Return metadata_header_padding
func (fctx *AvFormatContext) MetadataHeaderPadding() int {
	return int(fctx.metadata_header_padding)
}

// ProbeScore Return probe_score
func (fctx *AvFormatContext) ProbeScore() int {
	return int(fctx.probe_score)
}

// Seek2any Return seek2any
func (fctx *AvFormatContext) Seek2any() int {
	return int(fctx.seek2any)
}

// StrictStdCompliance Return strict_std_compliance
func (fctx *AvFormatContext) StrictStdCompliance() int {
	return int(fctx.strict_std_compliance)
}

// TsID Return ts_id
func (fctx *AvFormatContext) TsID() int {
	return int(fctx.ts_id)
}

// UseWallclockAsTimestamps Return use_wallclock_as_timestamps
func (fctx *AvFormatContext) UseWallclockAsTimestamps() int {
	return int(fctx.use_wallclock_as_timestamps)
}

// Duration Return duration
func (fctx *AvFormatContext) Duration() int64 {
	return int64(fctx.duration)
}

// MaxAnalyzeDuration2 Return max_analyze_duration
func (fctx *AvFormatContext) MaxAnalyzeDuration2() int64 {
	return int64(fctx.max_analyze_duration)
}

// MaxInterleaveDelta Return max_interleave_delta
func (fctx *AvFormatContext) MaxInterleaveDelta() int64 {
	return int64(fctx.max_interleave_delta)
}

// OutputTsOffset Return output_ts_offset
func (fctx *AvFormatContext) OutputTsOffset() int64 {
	return int64(fctx.output_ts_offset)
}

// Probesize2 Return probesize
func (fctx *AvFormatContext) Probesize2() int64 {
	return int64(fctx.probesize)
}

// SkipInitialBytes Return skip_initial_bytes
func (fctx *AvFormatContext) SkipInitialBytes() int64 {
	return int64(fctx.skip_initial_bytes)
}

// StartTime Return start_time
func (fctx *AvFormatContext) StartTime() int64 {
	return int64(fctx.start_time)
}

// StartTimeRealtime Return start_time_realtime
func (fctx *AvFormatContext) StartTimeRealtime() int64 {
	return int64(fctx.start_time_realtime)
}

// Iformat Return iformat
func (fctx *AvFormatContext) Iformat() *AvInputFormat {
	return (*AvInputFormat)(unsafe.Pointer(fctx.iformat))
}

// Oformat Return oformat
func (fctx *AvFormatContext) Oformat() *AvOutputFormat {
	return (*AvOutputFormat)(unsafe.Pointer(fctx.oformat))
}

// func (fctx *AvFormatContext) DumpSeparator() uint8 {
// 	return uint8(fctx.dump_separator)
// }

// CorrectTsOverflow Return correct_ts_overflow
func (fctx *AvFormatContext) CorrectTsOverflow() int {
	return int(fctx.correct_ts_overflow)
}

// MaxIndexSize Return max_index_size
func (fctx *AvFormatContext) MaxIndexSize() uint {
	return uint(fctx.max_index_size)
}

// MaxPictureBuffer Return max_picture_buffer
func (fctx *AvFormatContext) MaxPictureBuffer() uint {
	return uint(fctx.max_picture_buffer)
}

// NbChapters Return nb_chapters
func (fctx *AvFormatContext) NbChapters() uint {
	return uint(fctx.nb_chapters)
}

// NbPrograms Return nb_programs
func (fctx *AvFormatContext) NbPrograms() uint {
	return uint(fctx.nb_programs)
}

// NbStreams Return nb_streams
func (fctx *AvFormatContext) NbStreams() uint {
	return uint(fctx.nb_streams)
}

// PacketSize Return packet_size
func (fctx *AvFormatContext) PacketSize() uint {
	return uint(fctx.packet_size)
}

// Probesize Return probesize
func (fctx *AvFormatContext) Probesize() uint {
	return uint(fctx.probesize)
}

// SetPb Set pb
func (fctx *AvFormatContext) SetPb(pb *AvIOContext) {
	fctx.pb = (*C.struct_AVIOContext)(unsafe.Pointer(pb))
}

// Pb2 Return pb
func (fctx *AvFormatContext) Pb2() **AvIOContext {
	return (**AvIOContext)(unsafe.Pointer(&fctx.pb))
}
