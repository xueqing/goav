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

// Iformat Return iformat
func (fctx *AvFormatContext) Iformat() *AvInputFormat {
	return (*AvInputFormat)(unsafe.Pointer(fctx.iformat))
}

// Oformat Return oformat
func (fctx *AvFormatContext) Oformat() *AvOutputFormat {
	return (*AvOutputFormat)(unsafe.Pointer(fctx.oformat))
}

// Pb Return pb
func (fctx *AvFormatContext) Pb() *AvIOContext {
	return (*AvIOContext)(unsafe.Pointer(fctx.pb))
}

// SetPb Set pb
func (fctx *AvFormatContext) SetPb(pb *AvIOContext) {
	fctx.pb = (*C.struct_AVIOContext)(unsafe.Pointer(pb))
}

// CtxFlags Return ctx_flags
func (fctx *AvFormatContext) CtxFlags() int {
	return int(fctx.ctx_flags)
}

// NbStreams Return nb_streams
func (fctx *AvFormatContext) NbStreams() uint {
	return uint(fctx.nb_streams)
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

// StartTime Return start_time
func (fctx *AvFormatContext) StartTime() int64 {
	return int64(fctx.start_time)
}

// Duration Return duration
func (fctx *AvFormatContext) Duration() int64 {
	return int64(fctx.duration)
}

// BitRate Return bit_rate
func (fctx *AvFormatContext) BitRate() int {
	return int(fctx.bit_rate)
}

// PacketSize Return packet_size
func (fctx *AvFormatContext) PacketSize() uint {
	return uint(fctx.packet_size)
}

// MaxDelay Return max_delay
func (fctx *AvFormatContext) MaxDelay() int {
	return int(fctx.max_delay)
}

// Flags Return flags
func (fctx *AvFormatContext) Flags() int {
	return int(fctx.flags)
}

// Probesize Return probesize
func (fctx *AvFormatContext) Probesize() int64 {
	return int64(fctx.probesize)
}

// MaxAnalyzeDuration2 Return max_analyze_duration
func (fctx *AvFormatContext) MaxAnalyzeDuration2() int64 {
	return int64(fctx.max_analyze_duration)
}

// Keylen Return keylen
func (fctx *AvFormatContext) Keylen() int {
	return int(fctx.keylen)
}

// NbPrograms Return nb_programs
func (fctx *AvFormatContext) NbPrograms() uint {
	return uint(fctx.nb_programs)
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

// VideoCodecID Return video_codec_id
func (fctx *AvFormatContext) VideoCodecID() AvCodecID {
	return AvCodecID(fctx.video_codec_id)
}

// AudioCodecID Return audio_codec_id
func (fctx *AvFormatContext) AudioCodecID() AvCodecID {
	return AvCodecID(fctx.audio_codec_id)
}

// SubtitleCodecID Return subtitle_codec_id
func (fctx *AvFormatContext) SubtitleCodecID() AvCodecID {
	return AvCodecID(fctx.subtitle_codec_id)
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

// Chapters Return chapters
func (fctx *AvFormatContext) Chapters() **AvChapter {
	return (**AvChapter)(unsafe.Pointer(fctx.chapters))
}

// Metadata Return metadata
func (fctx *AvFormatContext) Metadata() *libavutil.AvDictionary {
	return (*libavutil.AvDictionary)(unsafe.Pointer(fctx.metadata))
}

// StartTimeRealtime Return start_time_realtime
func (fctx *AvFormatContext) StartTimeRealtime() int64 {
	return int64(fctx.start_time_realtime)
}

// FpsProbeSize Return fps_probe_size
func (fctx *AvFormatContext) FpsProbeSize() int {
	return int(fctx.fps_probe_size)
}

// ErrorRecognition Return error_recognition
func (fctx *AvFormatContext) ErrorRecognition() int {
	return int(fctx.error_recognition)
}

// InterruptCallback Return interrupt_callback
func (fctx *AvFormatContext) InterruptCallback() AvIOInterruptCB {
	return AvIOInterruptCB(fctx.interrupt_callback)
}

// Debug Return debug
func (fctx *AvFormatContext) Debug() int {
	return int(fctx.debug)
}

// MaxInterleaveDelta Return max_interleave_delta
func (fctx *AvFormatContext) MaxInterleaveDelta() int64 {
	return int64(fctx.max_interleave_delta)
}

// StrictStdCompliance Return strict_std_compliance
func (fctx *AvFormatContext) StrictStdCompliance() int {
	return int(fctx.strict_std_compliance)
}

// EventFlags Return event_flags
func (fctx *AvFormatContext) EventFlags() int {
	return int(fctx.event_flags)
}

// MaxTsProbe Return max_ts_probe
func (fctx *AvFormatContext) MaxTsProbe() int {
	return int(fctx.max_ts_probe)
}

// AvoidNegativeTs Return avoid_negative_ts
func (fctx *AvFormatContext) AvoidNegativeTs() int {
	return int(fctx.avoid_negative_ts)
}

// TsID Return ts_id
func (fctx *AvFormatContext) TsID() int {
	return int(fctx.ts_id)
}

// AudioPreload Return audio_preload
func (fctx *AvFormatContext) AudioPreload() int {
	return int(fctx.audio_preload)
}

// MaxChunkDuration Return max_chunk_duration
func (fctx *AvFormatContext) MaxChunkDuration() int {
	return int(fctx.max_chunk_duration)
}

// MaxChunkSize Return max_chunk_size
func (fctx *AvFormatContext) MaxChunkSize() int {
	return int(fctx.max_chunk_size)
}

// UseWallclockAsTimestamps Return use_wallclock_as_timestamps
func (fctx *AvFormatContext) UseWallclockAsTimestamps() int {
	return int(fctx.use_wallclock_as_timestamps)
}

// AvioFlags Return avio_flags
func (fctx *AvFormatContext) AvioFlags() int {
	return int(fctx.avio_flags)
}

// DurationEstimationMethod Return duration_estimation_method
func (fctx *AvFormatContext) DurationEstimationMethod() AvDurationEstimationMethod {
	return AvDurationEstimationMethod(fctx.duration_estimation_method)
}

// SkipInitialBytes Return skip_initial_bytes
func (fctx *AvFormatContext) SkipInitialBytes() int64 {
	return int64(fctx.skip_initial_bytes)
}

// CorrectTsOverflow Return correct_ts_overflow
func (fctx *AvFormatContext) CorrectTsOverflow() int {
	return int(fctx.correct_ts_overflow)
}

// Seek2any Return seek2any
func (fctx *AvFormatContext) Seek2any() int {
	return int(fctx.seek2any)
}

// FlushPackets Return flush_packets
func (fctx *AvFormatContext) FlushPackets() int {
	return int(fctx.flush_packets)
}

// ProbeScore Return probe_score
func (fctx *AvFormatContext) ProbeScore() int {
	return int(fctx.probe_score)
}

// FormatProbesize Return format_probesize
func (fctx *AvFormatContext) FormatProbesize() int {
	return int(fctx.format_probesize)
}

// CodecWhitelist Return codec_whitelist
func (fctx *AvFormatContext) CodecWhitelist() string {
	return C.GoString(fctx.codec_whitelist)
}

// FormatWhitelist Return format_whitelist
func (fctx *AvFormatContext) FormatWhitelist() string {
	return C.GoString(fctx.format_whitelist)
}

// Internal Return internal
func (fctx *AvFormatContext) Internal() *AvFormatInternal {
	return (*AvFormatInternal)(unsafe.Pointer(fctx.internal))
}

// IoRepositioned Return io_repositioned
func (fctx *AvFormatContext) IoRepositioned() int {
	return int(fctx.io_repositioned)
}

// VideoCodec Return video_codec
func (fctx *AvFormatContext) VideoCodec() *AvCodec {
	return (*AvCodec)(unsafe.Pointer(fctx.video_codec))
}

// AudioCodec Return audio_codec
func (fctx *AvFormatContext) AudioCodec() *AvCodec {
	return (*AvCodec)(unsafe.Pointer(fctx.audio_codec))
}

// SubtitleCodec Return subtitle_codec
func (fctx *AvFormatContext) SubtitleCodec() *AvCodec {
	return (*AvCodec)(unsafe.Pointer(fctx.subtitle_codec))
}

// MetadataHeaderPadding Return metadata_header_padding
func (fctx *AvFormatContext) MetadataHeaderPadding() int {
	return int(fctx.metadata_header_padding)
}

// OutputTsOffset Return output_ts_offset
func (fctx *AvFormatContext) OutputTsOffset() int64 {
	return int64(fctx.output_ts_offset)
}
