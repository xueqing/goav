// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

//Package libavformat provides some generic global options, which can be set on all the muxers and demuxers.
//In addition each muxer or demuxer may support so-called private options, which are specific for that component.
//Supported formats (muxers and demuxers) provided by the libavformat library
package libavformat

//#cgo pkg-config: libavformat libavcodec libavutil libavdevice libavfilter libswresample libswscale
//#include <stdio.h>
//#include <stdlib.h>
//#include <inttypes.h>
//#include <stdint.h>
//#include <string.h>
//#include <libavformat/avformat.h>
//#include <libavcodec/avcodec.h>
//#include <libavutil/avutil.h>
//#include <libavutil/opt.h>
//#include <libavdevice/avdevice.h>
import "C"
import (
	"unsafe"

	"github.com/xueqing/goav/libavcodec"
)

type (
	AvProbeData                C.struct_AVProbeData
	AvInputFormat              C.struct_AVInputFormat
	AvOutputFormat             C.struct_AVOutputFormat
	AvFormatContext            C.struct_AVFormatContext
	AvFrame                    C.struct_AVFrame
	AvCodecContext             C.struct_AVCodecContext
	AvIndexEntry               C.struct_AVIndexEntry
	AvStream                   C.struct_AVStream
	AvProgram                  C.struct_AVProgram
	AvChapter                  C.struct_AVChapter
	AvPacketList               C.struct_AVPacketList
	AvCodecParserContext       C.struct_AVCodecParserContext
	AvIOContext                C.struct_AVIOContext
	AvCodec                    C.struct_AVCodec
	AvCodecTag                 C.struct_AVCodecTag
	AvClass                    C.struct_AVClass
	AvFormatInternal           C.struct_AVFormatInternal
	AvIOInterruptCB            C.struct_AVIOInterruptCB
	AvPacketSideData           C.struct_AVPacketSideData
	FFFrac                     C.struct_FFFrac
	AvStreamParseType          C.enum_AVStreamParseType
	AvDiscard                  C.enum_AVDiscard
	AvMediaType                C.enum_AVMediaType
	AvDurationEstimationMethod C.enum_AVDurationEstimationMethod
	AvPacketSideDataType       C.enum_AVPacketSideDataType
	AvCodecID                  C.enum_AVCodecID

	File C.FILE
)

// AvformatVersion Return the LIBAvFORMAT_VERSION_INT constant.
func AvformatVersion() uint {
	return uint(C.avformat_version())
}

// AvformatConfiguration Return the libavformat build-time configuration.
func AvformatConfiguration() string {
	return C.GoString(C.avformat_configuration())
}

// AvformatLicense Return the libavformat license.
func AvformatLicense() string {
	return C.GoString(C.avformat_license())
}

// AvformatNetworkInit Do global initialization of network components.
func AvformatNetworkInit() int {
	return int(C.avformat_network_init())
}

// AvformatNetworkDeinit Undo the initialization done by avformat_network_init.
func AvformatNetworkDeinit() int {
	return int(C.avformat_network_deinit())
}

// AvformatGetClass Get the Class for Context.
func AvformatGetClass() *AvClass {
	return (*AvClass)(C.avformat_get_class())
}

// AvFindInputFormat Find InputFormat based on the short name of the input format.
func AvFindInputFormat(name string) *AvInputFormat {
	cShortName := C.CString(name)
	defer C.free(unsafe.Pointer(cShortName))

	return (*AvInputFormat)(C.av_find_input_format(cShortName))
}

// AvProbeInputFormat Guess the file format.
func AvProbeInputFormat(pd *AvProbeData, isOpened int) *AvInputFormat {
	return (*AvInputFormat)(C.av_probe_input_format((*C.struct_AVProbeData)(pd),
		C.int(isOpened)))
}

// AvProbeInputFormat2 Guess the file format.
func AvProbeInputFormat2(pd *AvProbeData, isOpened int, scoreMax *int) *AvInputFormat {
	return (*AvInputFormat)(C.av_probe_input_format2((*C.struct_AVProbeData)(pd),
		C.int(isOpened), (*C.int)(unsafe.Pointer(scoreMax))))
}

// AvProbeInputFormat3 Guess the file format.
func AvProbeInputFormat3(pd *AvProbeData, isOpened int, scoreRet *int) *AvInputFormat {
	return (*AvInputFormat)(C.av_probe_input_format3((*C.struct_AVProbeData)(pd),
		C.int(isOpened), (*C.int)(unsafe.Pointer(scoreRet))))
}

// AvProbeInputBuffer Like av_probe_input_buffer2() but returns 0 on success.
func AvProbeInputBuffer(pb *AvIOContext, format **AvInputFormat, url string, logCtx int, offset, maxProbeSize uint) int {
	cURL := C.CString(url)
	defer C.free(unsafe.Pointer(cURL))

	return int(C.av_probe_input_buffer((*C.struct_AVIOContext)(pb),
		(**C.struct_AVInputFormat)(unsafe.Pointer(format)), cURL,
		unsafe.Pointer(&logCtx), C.uint(offset), C.uint(maxProbeSize)))
}

// AvProbeInputBuffer2 Probe a bytestream to determine the input format.
func AvProbeInputBuffer2(pb *AvIOContext, format **AvInputFormat, url string, logCtx int, offset, maxProbeSize uint) int {
	cURL := C.CString(url)
	defer C.free(unsafe.Pointer(cURL))

	return int(C.av_probe_input_buffer2((*C.struct_AVIOContext)(pb),
		(**C.struct_AVInputFormat)(unsafe.Pointer(format)), cURL,
		unsafe.Pointer(&logCtx), C.uint(offset), C.uint(maxProbeSize)))
}

// AvGuessFormat Return the output format in the list of registered output formats which best matches the provided parameters, or return NULL if there is no match.
func AvGuessFormat(shortName, fileName, mimeType string) *AvOutputFormat {
	cShortName := C.CString(shortName)
	defer C.free(unsafe.Pointer(cShortName))

	cFileName := C.CString(fileName)
	defer C.free(unsafe.Pointer(cFileName))

	cMimeType := C.CString(mimeType)
	defer C.free(unsafe.Pointer(cMimeType))

	return (*AvOutputFormat)(C.av_guess_format(cShortName, cFileName, cMimeType))
}

// AvGuessCodec Guess the codec ID based upon muxer and filename.
func AvGuessCodec(format *AvOutputFormat, shortName, fileName, mimeType string, typ AvMediaType) AvCodecID {
	cShortName := C.CString(shortName)
	defer C.free(unsafe.Pointer(cShortName))

	cFileName := C.CString(fileName)
	defer C.free(unsafe.Pointer(cFileName))

	cMimeType := C.CString(mimeType)
	defer C.free(unsafe.Pointer(cMimeType))

	return (AvCodecID)(C.av_guess_codec((*C.struct_AVOutputFormat)(format),
		cShortName, cFileName, cMimeType, (C.enum_AVMediaType)(typ)))
}

// AvHexDump Send a nice hexadecimal dump of a buffer to the specified file stream.
func AvHexDump(f *File, buf *uint8, size int) {
	C.av_hex_dump((*C.FILE)(f), (*C.uint8_t)(buf), C.int(size))
}

// AvHexDumpLog Send a nice hexadecimal dump of a buffer to the log.
func AvHexDumpLog(avcl, level int, buf *uint8, size int) {
	C.av_hex_dump_log(unsafe.Pointer(&avcl), C.int(level), (*C.uint8_t)(buf), C.int(size))
}

// AvPktDump2 Send a nice dump of a packet to the specified file stream.
func AvPktDump2(f *File, pkt *libavcodec.AvPacket, dumpPayload int, st *AvStream) {
	C.av_pkt_dump2((*C.FILE)(f), toCPacket(pkt), C.int(dumpPayload), (*C.struct_AVStream)(st))
}

// AvPktDumpLog2 Send a nice dump of a packet to the log.
func AvPktDumpLog2(avcl, level int, pkt *libavcodec.AvPacket, dumpPayload int, st *AvStream) {
	C.av_pkt_dump_log2(unsafe.Pointer(&avcl), C.int(level),
		toCPacket(pkt), C.int(dumpPayload), (*C.struct_AVStream)(st))
}

// AvCodecGetID Get the AvCodecID for the given codec tag tag.
func AvCodecGetID(tags **AvCodecTag, tag uint) AvCodecID {
	return (AvCodecID)(C.av_codec_get_id((**C.struct_AVCodecTag)(unsafe.Pointer(tags)), C.uint(tag)))
}

// AvCodecGetTag Get the codec tag for the given codec id id.
func AvCodecGetTag(tags **AvCodecTag, id AvCodecID) uint {
	return uint(C.av_codec_get_tag((**C.struct_AVCodecTag)(unsafe.Pointer(tags)), (C.enum_AVCodecID)(id)))
}

// AvCodecGetTag2 Get the codec tag for the given codec id.
func AvCodecGetTag2(tags **AvCodecTag, id AvCodecID, tag *uint) int {
	return int(C.av_codec_get_tag2((**C.struct_AVCodecTag)(unsafe.Pointer(tags)),
		(C.enum_AVCodecID)(id), (*C.uint)(unsafe.Pointer(tag))))
}

// AvIndexSearchTimestamp Get the index for a specific timestamp.
func AvIndexSearchTimestamp(st *AvStream, timestamp int64, flags int) int {
	return int(C.av_index_search_timestamp((*C.struct_AVStream)(st),
		C.int64_t(timestamp), C.int(flags)))
}

// AvAddIndexEntry Add an index entry into a sorted list.
func AvAddIndexEntry(st *AvStream, pos, timestamp, int64, size, distance, flags int) int {
	return int(C.av_add_index_entry((*C.struct_AVStream)(st),
		C.int64_t(pos), C.int64_t(timestamp), C.int(size), C.int(distance), C.int(flags)))
}

// AvURLSplit Split a URL string into components.
func AvURLSplit(protoSize, authorizationSize, hostnameSize int, portPtr *int, pathSize int, url string) (proto, authorization, hostname, path string) {
	cProto := (*C.char)(C.malloc(C.sizeof_char * C.ulong(protoSize)))
	defer C.free(unsafe.Pointer(cProto))

	cAuthorization := (*C.char)(C.malloc(C.sizeof_char * C.ulong(authorizationSize)))
	defer C.free(unsafe.Pointer(cAuthorization))

	cHostname := (*C.char)(C.malloc(C.sizeof_char * C.ulong(hostnameSize)))
	defer C.free(unsafe.Pointer(cHostname))

	cPath := (*C.char)(C.malloc(C.sizeof_char * C.ulong(pathSize)))
	defer C.free(unsafe.Pointer(cPath))

	cURL := C.CString(url)
	defer C.free(unsafe.Pointer(cURL))

	C.av_url_split(
		cProto, C.int(protoSize),
		cAuthorization, C.int(authorizationSize),
		cHostname, C.int(hostnameSize),
		(*C.int)(unsafe.Pointer(portPtr)),
		cPath, C.int(pathSize),
		cURL,
	)

	return C.GoString(cProto), C.GoString(cAuthorization), C.GoString(cHostname), C.GoString(cPath)
}

// AvGetFrameFilename Return in 'buf' the path with 'd' replaced by a number.
func AvGetFrameFilename(bufSize int, path string, number int) (ret int, buf string) {
	cBuf := (*C.char)(C.malloc(C.sizeof_char * C.ulong(bufSize)))
	defer C.free(unsafe.Pointer(cBuf))

	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))

	ret = int(C.av_get_frame_filename(cBuf, C.int(bufSize), cPath, C.int(number)))
	buf = C.GoString(cBuf)
	return
}

// AvFilenameNumberTest Check whether filename actually is a numbered sequence generator.
func AvFilenameNumberTest(fileName string) int {
	cFileName := C.CString(fileName)
	defer C.free(unsafe.Pointer(cFileName))

	return int(C.av_filename_number_test(cFileName))
}

// AvSdpCreate Generate an SDP for an RTP session.
func AvSdpCreate(ac **AvFormatContext, nFiles int, bufSize int) (ret int, buf string) {
	cBuf := (*C.char)(C.malloc(C.sizeof_char * C.ulong(bufSize)))
	defer C.free(unsafe.Pointer(cBuf))

	ret = int(C.av_sdp_create((**C.struct_AVFormatContext)(unsafe.Pointer(ac)),
		C.int(nFiles), cBuf, C.int(bufSize)))
	buf = C.GoString(cBuf)
	return
}

// AvMatchExt Return a positive value if the given filename has one of the given extensions, 0 otherwise.
func AvMatchExt(fileName, extensions string) int {
	cFileName := C.CString(fileName)
	defer C.free(unsafe.Pointer(cFileName))

	cExtensions := C.CString(extensions)
	defer C.free(unsafe.Pointer(cExtensions))

	return int(C.av_match_ext(cFileName, cExtensions))
}

// AvformatQueryCodec Test if the given container can store a codec.
func AvformatQueryCodec(oFmt *AvOutputFormat, codecID AvCodecID, stdCompliance int) int {
	return int(C.avformat_query_codec((*C.struct_AVOutputFormat)(oFmt),
		(C.enum_AVCodecID)(codecID), C.int(stdCompliance)))
}

// AvformatGetRiffVideoTags Return the table mapping RIFF FourCCs for video to libavcodec AvCodecID.
func AvformatGetRiffVideoTags() *AvCodecTag {
	return (*AvCodecTag)(C.avformat_get_riff_video_tags())
}

// AvformatGetRiffAudioTags Return the table mapping RIFF FourCCs for audio to AvCodecID.
func AvformatGetRiffAudioTags() *AvCodecTag {
	return (*AvCodecTag)(C.avformat_get_riff_audio_tags())
}

// AvformatGetMovVideoTags Return the table mapping MOV FourCCs for video to libavcodec AvCodecID.
func AvformatGetMovVideoTags() *AvCodecTag {
	return (*AvCodecTag)(C.avformat_get_mov_video_tags())
}

// AvformatGetMovAudioTags Return the table mapping MOV FourCCs for audio to AvCodecID.
func AvformatGetMovAudioTags() *AvCodecTag {
	return (*AvCodecTag)(C.avformat_get_mov_audio_tags())
}
