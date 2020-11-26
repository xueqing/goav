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
func (fctx *AVFormatContext) Chapters() **AvChapter {
	return (**AvChapter)(unsafe.Pointer(fctx.chapters))
}

// AudioCodec Return audio_codec
func (fctx *AVFormatContext) AudioCodec() *AvCodec {
	return (*AvCodec)(unsafe.Pointer(fctx.audio_codec))
}

// SubtitleCodec Return subtitle_codec
func (fctx *AVFormatContext) SubtitleCodec() *AvCodec {
	return (*AvCodec)(unsafe.Pointer(fctx.subtitle_codec))
}

// VideoCodec Return video_codec
func (fctx *AVFormatContext) VideoCodec() *AvCodec {
	return (*AvCodec)(unsafe.Pointer(fctx.video_codec))
}

// Metadata Return metadata
func (fctx *AVFormatContext) Metadata() *libavutil.AvDictionary {
	return (*libavutil.AvDictionary)(unsafe.Pointer(fctx.metadata))
}

// Internal Return internal
func (fctx *AVFormatContext) Internal() *AvFormatInternal {
	return (*AvFormatInternal)(unsafe.Pointer(fctx.internal))
}

// Pb Return pb
func (fctx *AVFormatContext) Pb() *AvIOContext {
	return (*AvIOContext)(unsafe.Pointer(fctx.pb))
}

// InterruptCallback Return interrupt_callback
func (fctx *AVFormatContext) InterruptCallback() AvIOInterruptCB {
	return AvIOInterruptCB(fctx.interrupt_callback)
}

// Programs Return programs
func (fctx *AVFormatContext) Programs() []*AvProgram {
	header := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(fctx.programs)),
		Len:  int(fctx.NbPrograms()),
		Cap:  int(fctx.NbPrograms()),
	}

	return *((*[]*AvProgram)(unsafe.Pointer(&header)))
}

// Streams Return streams
func (fctx *AVFormatContext) Streams() []*AvStream {
	header := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(fctx.streams)),
		Len:  int(fctx.NbStreams()),
		Cap:  int(fctx.NbStreams()),
	}

	return *((*[]*AvStream)(unsafe.Pointer(&header)))
}

// Filename Return filename
func (fctx *AVFormatContext) Filename() string {
	return C.GoString((*C.char)(unsafe.Pointer(&fctx.filename[0])))
}

// func (fctx *AVFormatContext) CodecWhitelist() string {
// 	return C.GoString(fctx.codec_whitelist)
// }

// func (fctx *AVFormatContext) FormatWhitelist() string {
// 	return C.GoString(fctx.format_whitelist)
// }

// AudioCodecID Return audio_codec_id
func (fctx *AVFormatContext) AudioCodecID() AvCodecID {
	return AvCodecID(fctx.audio_codec_id)
}

// SubtitleCodecID Return subtitle_codec_id
func (fctx *AVFormatContext) SubtitleCodecID() AvCodecID {
	return AvCodecID(fctx.subtitle_codec_id)
}

// VideoCodecID Return video_codec_id
func (fctx *AVFormatContext) VideoCodecID() AvCodecID {
	return AvCodecID(fctx.video_codec_id)
}

// DurationEstimationMethod Return duration_estimation_method
func (fctx *AVFormatContext) DurationEstimationMethod() AvDurationEstimationMethod {
	return AvDurationEstimationMethod(fctx.duration_estimation_method)
}

// AudioPreload Return audio_preload
func (fctx *AVFormatContext) AudioPreload() int {
	return int(fctx.audio_preload)
}

// AvioFlags Return avio_flags
func (fctx *AVFormatContext) AvioFlags() int {
	return int(fctx.avio_flags)
}

// AvoidNegativeTs Return avoid_negative_ts
func (fctx *AVFormatContext) AvoidNegativeTs() int {
	return int(fctx.avoid_negative_ts)
}

// BitRate Return bit_rate
func (fctx *AVFormatContext) BitRate() int {
	return int(fctx.bit_rate)
}

// CtxFlags Return ctx_flags
func (fctx *AVFormatContext) CtxFlags() int {
	return int(fctx.ctx_flags)
}

// Debug Return debug
func (fctx *AVFormatContext) Debug() int {
	return int(fctx.debug)
}

// ErrorRecognition Return error_recognition
func (fctx *AVFormatContext) ErrorRecognition() int {
	return int(fctx.error_recognition)
}

// EventFlags Return event_flags
func (fctx *AVFormatContext) EventFlags() int {
	return int(fctx.event_flags)
}

// Flags Return flags
func (fctx *AVFormatContext) Flags() int {
	return int(fctx.flags)
}

// FlushPackets Return flush_packets
func (fctx *AVFormatContext) FlushPackets() int {
	return int(fctx.flush_packets)
}

// FormatProbesize Return format_probesize
func (fctx *AVFormatContext) FormatProbesize() int {
	return int(fctx.format_probesize)
}

// FpsProbeSize Return fps_probe_size
func (fctx *AVFormatContext) FpsProbeSize() int {
	return int(fctx.fps_probe_size)
}

// IoRepositioned Return io_repositioned
func (fctx *AVFormatContext) IoRepositioned() int {
	return int(fctx.io_repositioned)
}

// Keylen Return keylen
func (fctx *AVFormatContext) Keylen() int {
	return int(fctx.keylen)
}

// MaxChunkDuration Return max_chunk_duration
func (fctx *AVFormatContext) MaxChunkDuration() int {
	return int(fctx.max_chunk_duration)
}

// MaxChunkSize Return max_chunk_size
func (fctx *AVFormatContext) MaxChunkSize() int {
	return int(fctx.max_chunk_size)
}

// MaxDelay Return max_delay
func (fctx *AVFormatContext) MaxDelay() int {
	return int(fctx.max_delay)
}

// MaxTsProbe Return max_ts_probe
func (fctx *AVFormatContext) MaxTsProbe() int {
	return int(fctx.max_ts_probe)
}

// MetadataHeaderPadding Return metadata_header_padding
func (fctx *AVFormatContext) MetadataHeaderPadding() int {
	return int(fctx.metadata_header_padding)
}

// ProbeScore Return probe_score
func (fctx *AVFormatContext) ProbeScore() int {
	return int(fctx.probe_score)
}

// Seek2any Return seek2any
func (fctx *AVFormatContext) Seek2any() int {
	return int(fctx.seek2any)
}

// StrictStdCompliance Return strict_std_compliance
func (fctx *AVFormatContext) StrictStdCompliance() int {
	return int(fctx.strict_std_compliance)
}

// TsID Return ts_id
func (fctx *AVFormatContext) TsID() int {
	return int(fctx.ts_id)
}

// UseWallclockAsTimestamps Return use_wallclock_as_timestamps
func (fctx *AVFormatContext) UseWallclockAsTimestamps() int {
	return int(fctx.use_wallclock_as_timestamps)
}

// Duration Return duration
func (fctx *AVFormatContext) Duration() int64 {
	return int64(fctx.duration)
}

// MaxAnalyzeDuration2 Return max_analyze_duration
func (fctx *AVFormatContext) MaxAnalyzeDuration2() int64 {
	return int64(fctx.max_analyze_duration)
}

// MaxInterleaveDelta Return max_interleave_delta
func (fctx *AVFormatContext) MaxInterleaveDelta() int64 {
	return int64(fctx.max_interleave_delta)
}

// OutputTsOffset Return output_ts_offset
func (fctx *AVFormatContext) OutputTsOffset() int64 {
	return int64(fctx.output_ts_offset)
}

// Probesize2 Return probesize
func (fctx *AVFormatContext) Probesize2() int64 {
	return int64(fctx.probesize)
}

// SkipInitialBytes Return skip_initial_bytes
func (fctx *AVFormatContext) SkipInitialBytes() int64 {
	return int64(fctx.skip_initial_bytes)
}

// StartTime Return start_time
func (fctx *AVFormatContext) StartTime() int64 {
	return int64(fctx.start_time)
}

// StartTimeRealtime Return start_time_realtime
func (fctx *AVFormatContext) StartTimeRealtime() int64 {
	return int64(fctx.start_time_realtime)
}

// Iformat Return iformat
func (fctx *AVFormatContext) Iformat() *AvInputFormat {
	return (*AvInputFormat)(unsafe.Pointer(fctx.iformat))
}

// Oformat Return oformat
func (fctx *AVFormatContext) Oformat() *AvOutputFormat {
	return (*AvOutputFormat)(unsafe.Pointer(fctx.oformat))
}

// func (fctx *AVFormatContext) DumpSeparator() uint8 {
// 	return uint8(fctx.dump_separator)
// }

// CorrectTsOverflow Return correct_ts_overflow
func (fctx *AVFormatContext) CorrectTsOverflow() int {
	return int(fctx.correct_ts_overflow)
}

// MaxIndexSize Return max_index_size
func (fctx *AVFormatContext) MaxIndexSize() uint {
	return uint(fctx.max_index_size)
}

// MaxPictureBuffer Return max_picture_buffer
func (fctx *AVFormatContext) MaxPictureBuffer() uint {
	return uint(fctx.max_picture_buffer)
}

// NbChapters Return nb_chapters
func (fctx *AVFormatContext) NbChapters() uint {
	return uint(fctx.nb_chapters)
}

// NbPrograms Return nb_programs
func (fctx *AVFormatContext) NbPrograms() uint {
	return uint(fctx.nb_programs)
}

// NbStreams Return nb_streams
func (fctx *AVFormatContext) NbStreams() uint {
	return uint(fctx.nb_streams)
}

// PacketSize Return packet_size
func (fctx *AVFormatContext) PacketSize() uint {
	return uint(fctx.packet_size)
}

// Probesize Return probesize
func (fctx *AVFormatContext) Probesize() uint {
	return uint(fctx.probesize)
}

// SetPb Set pb
func (fctx *AVFormatContext) SetPb(pb *AvIOContext) {
	fctx.pb = (*C.struct_AVIOContext)(unsafe.Pointer(pb))
}

// Pb2 Return pb
func (fctx *AVFormatContext) Pb2() **AvIOContext {
	return (**AvIOContext)(unsafe.Pointer(&fctx.pb))
}
