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
	"github.com/xueqing/goav/libavutil"
)

type (
	AvProbeData                C.struct_AVProbeData
	AvInputFormat              C.struct_AVInputFormat
	AvOutputFormat             C.struct_AVOutputFormat
	AVFormatContext            C.struct_AVFormatContext
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

// AvGetPacket Allocate and read the payload of a packet and initialize its fields with default values.
func (ctxt *AvIOContext) AvGetPacket(pkt *libavcodec.AvPacket, s int) int {
	return int(C.av_get_packet((*C.struct_AVIOContext)(ctxt), toCPacket(pkt), C.int(s)))
}

// AvAppendPacket Read data and append it to the current content of the Packet.
func (ctxt *AvIOContext) AvAppendPacket(pkt *libavcodec.AvPacket, s int) int {
	return int(C.av_append_packet((*C.struct_AVIOContext)(ctxt), toCPacket(pkt), C.int(s)))
}

// Close ...
func (ctxt *AvIOContext) Close() error {
	return libavutil.ErrorFromCode(int(C.avio_close((*C.AVIOContext)(unsafe.Pointer(ctxt)))))
}

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

// AvformatAllocContext Allocate an Context.
func AvformatAllocContext() *AVFormatContext {
	return (*AVFormatContext)(C.avformat_alloc_context())
}

// AvformatGetClass Get the Class for Context.
func AvformatGetClass() *AvClass {
	return (*AvClass)(C.avformat_get_class())
}

// AvStreamGetSideData Get side information from stream.
func (s *AvStream) AvStreamGetSideData(t AvPacketSideDataType, z int) *uint8 {
	return (*uint8)(C.av_stream_get_side_data((*C.struct_AVStream)(s), (C.enum_AVPacketSideDataType)(t), (*C.int)(unsafe.Pointer(&z))))
}

// AvformatAllocOutputContext2 Allocate an Context for an output format.
func AvformatAllocOutputContext2(ctx **AVFormatContext, o *AvOutputFormat, fo, fi string) int {
	CformatName := C.CString(fo)
	defer C.free(unsafe.Pointer(CformatName))

	Cfilename := C.CString(fi)
	defer C.free(unsafe.Pointer(Cfilename))

	return int(C.avformat_alloc_output_context2((**C.struct_AVFormatContext)(unsafe.Pointer(ctx)), (*C.struct_AVOutputFormat)(o), CformatName, Cfilename))
}

// AvFindInputFormat Find InputFormat based on the short name of the input format.
func AvFindInputFormat(s string) *AvInputFormat {
	CshortName := C.CString(s)
	defer C.free(unsafe.Pointer(CshortName))

	return (*AvInputFormat)(C.av_find_input_format(CshortName))
}

// AvProbeInputFormat Guess the file format.
func AvProbeInputFormat(pd *AvProbeData, i int) *AvInputFormat {
	return (*AvInputFormat)(C.av_probe_input_format((*C.struct_AVProbeData)(pd), C.int(i)))
}

// AvProbeInputFormat2 Guess the file format.
func AvProbeInputFormat2(pd *AvProbeData, o int, sm *int) *AvInputFormat {
	return (*AvInputFormat)(C.av_probe_input_format2((*C.struct_AVProbeData)(pd), C.int(o), (*C.int)(unsafe.Pointer(sm))))
}

// AvProbeInputFormat3 Guess the file format.
func AvProbeInputFormat3(pd *AvProbeData, o int, sl *int) *AvInputFormat {
	return (*AvInputFormat)(C.av_probe_input_format3((*C.struct_AVProbeData)(pd), C.int(o), (*C.int)(unsafe.Pointer(sl))))
}

// AvProbeInputBuffer2 Probe a bytestream to determine the input format.
func AvProbeInputBuffer2(pb *AvIOContext, f **AvInputFormat, fi string, l int, o, m uint) int {
	Curl := C.CString(fi)
	defer C.free(unsafe.Pointer(Curl))

	return int(C.av_probe_input_buffer2((*C.struct_AVIOContext)(pb), (**C.struct_AVInputFormat)(unsafe.Pointer(f)), Curl, unsafe.Pointer(&l), C.uint(o), C.uint(m)))
}

// AvProbeInputBuffer Like av_probe_input_buffer2() but returns 0 on success.
func AvProbeInputBuffer(pb *AvIOContext, f **AvInputFormat, fi string, l int, o, m uint) int {
	Curl := C.CString(fi)
	defer C.free(unsafe.Pointer(Curl))

	return int(C.av_probe_input_buffer((*C.struct_AVIOContext)(pb), (**C.struct_AVInputFormat)(unsafe.Pointer(f)), Curl, unsafe.Pointer(&l), C.uint(o), C.uint(m)))
}

// AvformatOpenInput Open an input stream and read the header.
func AvformatOpenInput(ps **AVFormatContext, fi string, fmt *AvInputFormat, d **libavutil.AvDictionary) int {
	Cfi := C.CString(fi)
	defer C.free(unsafe.Pointer(Cfi))

	return int(C.avformat_open_input((**C.struct_AVFormatContext)(unsafe.Pointer(ps)), Cfi, (*C.struct_AVInputFormat)(fmt), (**C.struct_AVDictionary)(unsafe.Pointer(d))))
}

// AvGuessFormat Return the output format in the list of registered output formats which best matches the provided parameters, or return NULL if there is no match.
func AvGuessFormat(sn, f, mt string) *AvOutputFormat {
	CshortName := C.CString(sn)
	defer C.free(unsafe.Pointer(CshortName))

	Cfilename := C.CString(f)
	defer C.free(unsafe.Pointer(Cfilename))

	CmimeType := C.CString(mt)
	defer C.free(unsafe.Pointer(CmimeType))

	return (*AvOutputFormat)(C.av_guess_format(CshortName, Cfilename, CmimeType))
}

// AvGuessCodec Guess the codec ID based upon muxer and filename.
func AvGuessCodec(fmt *AvOutputFormat, sn, f, mt string, t AvMediaType) AvCodecID {
	CshortName := C.CString(sn)
	defer C.free(unsafe.Pointer(CshortName))

	Cfilename := C.CString(f)
	defer C.free(unsafe.Pointer(Cfilename))

	CmimeType := C.CString(mt)
	defer C.free(unsafe.Pointer(CmimeType))

	return (AvCodecID)(C.av_guess_codec((*C.struct_AVOutputFormat)(fmt), CshortName, Cfilename, CmimeType, (C.enum_AVMediaType)(t)))
}

// AvHexDump Send a nice hexadecimal dump of a buffer to the specified file stream.
func AvHexDump(f *File, b *uint8, s int) {
	C.av_hex_dump((*C.FILE)(f), (*C.uint8_t)(b), C.int(s))
}

// AvHexDumpLog Send a nice hexadecimal dump of a buffer to the log.
func AvHexDumpLog(a, l int, b *uint8, s int) {
	C.av_hex_dump_log(unsafe.Pointer(&a), C.int(l), (*C.uint8_t)(b), C.int(s))
}

// AvPktDump2 Send a nice dump of a packet to the specified file stream.
func AvPktDump2(f *File, pkt *libavcodec.AvPacket, dp int, st *AvStream) {
	C.av_pkt_dump2((*C.FILE)(f), toCPacket(pkt), C.int(dp), (*C.struct_AVStream)(st))
}

// AvPktDumpLog2 Send a nice dump of a packet to the log.
func AvPktDumpLog2(a int, l int, pkt *libavcodec.AvPacket, dp int, st *AvStream) {
	C.av_pkt_dump_log2(unsafe.Pointer(&a), C.int(l), toCPacket(pkt), C.int(dp), (*C.struct_AVStream)(st))
}

// AvCodecGetID Get the AvCodecID for the given codec tag tag.
func AvCodecGetID(t **AvCodecTag, tag uint) AvCodecID {
	return (AvCodecID)(C.av_codec_get_id((**C.struct_AVCodecTag)(unsafe.Pointer(t)), C.uint(tag)))
}

// AvCodecGetTag Get the codec tag for the given codec id id.
func AvCodecGetTag(t **AvCodecTag, id AvCodecID) uint {
	return uint(C.av_codec_get_tag((**C.struct_AVCodecTag)(unsafe.Pointer(t)), (C.enum_AVCodecID)(id)))
}

// AvCodecGetTag2 Get the codec tag for the given codec id.
func AvCodecGetTag2(t **AvCodecTag, id AvCodecID, tag *uint) int {
	return int(C.av_codec_get_tag2((**C.struct_AVCodecTag)(unsafe.Pointer(t)), (C.enum_AVCodecID)(id), (*C.uint)(unsafe.Pointer(tag))))
}

// AvIndexSearchTimestamp Get the index for a specific timestamp.
func AvIndexSearchTimestamp(st *AvStream, t int64, f int) int {
	return int(C.av_index_search_timestamp((*C.struct_AVStream)(st), C.int64_t(t), C.int(f)))
}

// AvAddIndexEntry Add an index entry into a sorted list.
func AvAddIndexEntry(st *AvStream, pos, t, int64, s, d, f int) int {
	return int(C.av_add_index_entry((*C.struct_AVStream)(st), C.int64_t(pos), C.int64_t(t), C.int(s), C.int(d), C.int(f)))
}

// AvURLSplit Split a URL string into components.
func AvURLSplit(protoSize, authorizationSize, hostnameSize int, pp *int, pathSize int, url string) (proto, authorization, hostname, path string) {
	Cproto := (*C.char)(C.malloc(C.sizeof_char * C.ulong(protoSize)))
	defer C.free(unsafe.Pointer(Cproto))

	Cauthorization := (*C.char)(C.malloc(C.sizeof_char * C.ulong(authorizationSize)))
	defer C.free(unsafe.Pointer(Cauthorization))

	Chostname := (*C.char)(C.malloc(C.sizeof_char * C.ulong(hostnameSize)))
	defer C.free(unsafe.Pointer(Chostname))

	Cpath := (*C.char)(C.malloc(C.sizeof_char * C.ulong(pathSize)))
	defer C.free(unsafe.Pointer(Cpath))

	Curl := C.CString(url)
	defer C.free(unsafe.Pointer(Curl))

	C.av_url_split(
		Cproto, C.int(protoSize),
		Cauthorization, C.int(authorizationSize),
		Chostname, C.int(hostnameSize),
		(*C.int)(unsafe.Pointer(pp)),
		Cpath, C.int(pathSize),
		Curl,
	)

	return C.GoString(Cproto), C.GoString(Cauthorization), C.GoString(Chostname), C.GoString(Cpath)
}

// AvGetFrameFilename Return in 'buf' the path with 'd' replaced by a number.
func AvGetFrameFilename(bufSize int, path string, number int) (int, string) {
	Cbuf := (*C.char)(C.malloc(C.sizeof_char * C.ulong(bufSize)))
	defer C.free(unsafe.Pointer(Cbuf))

	Cpath := C.CString(path)
	defer C.free(unsafe.Pointer(Cpath))

	ret := int(C.av_get_frame_filename(Cbuf, C.int(bufSize), Cpath, C.int(number)))

	return ret, C.GoString(Cbuf)
}

// AvFilenameNumberTest Check whether filename actually is a numbered sequence generator.
func AvFilenameNumberTest(filename string) int {
	Cfilename := C.CString(filename)
	defer C.free(unsafe.Pointer(Cfilename))

	return int(C.av_filename_number_test(Cfilename))
}

// AvSdpCreate Generate an SDP for an RTP session.
func AvSdpCreate(ac **AVFormatContext, nFiles int, bufSize int) (int, string) {
	Cbuf := (*C.char)(C.malloc(C.sizeof_char * C.ulong(bufSize)))
	defer C.free(unsafe.Pointer(Cbuf))

	ret := int(C.av_sdp_create((**C.struct_AVFormatContext)(unsafe.Pointer(ac)), C.int(nFiles), Cbuf, C.int(bufSize)))

	return ret, C.GoString(Cbuf)
}

// AvMatchExt Return a positive value if the given filename has one of the given extensions, 0 otherwise.
func AvMatchExt(filename, extensions string) int {
	Cfilename := C.CString(filename)
	defer C.free(unsafe.Pointer(Cfilename))

	Cextensions := C.CString(extensions)
	defer C.free(unsafe.Pointer(Cextensions))

	return int(C.av_match_ext(Cfilename, Cextensions))
}

// AvformatQueryCodec Test if the given container can store a codec.
func AvformatQueryCodec(o *AvOutputFormat, cd AvCodecID, sc int) int {
	return int(C.avformat_query_codec((*C.struct_AVOutputFormat)(o), (C.enum_AVCodecID)(cd), C.int(sc)))
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

// AvIOOpen Create and initialize a AVIOContext for accessing the resource indicated by url.
func AvIOOpen(url string, flags int) (res *AvIOContext, err error) {
	urlStr := C.CString(url)
	defer C.free(unsafe.Pointer(urlStr))
	err = libavutil.ErrorFromCode(int(C.avio_open((**C.AVIOContext)(unsafe.Pointer(&res)), urlStr, C.int(flags))))
	return
}
