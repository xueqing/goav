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
import "C"
import (
	"reflect"
	"unsafe"

	"github.com/xueqing/goav/libavcodec"
)

// Type Return codec_type as MediaType
func (cctx *CodecContext) Type() MediaType {
	return MediaType(cctx.codec_type)
}

// SetBitRate Set bit_rate
func (cctx *CodecContext) SetBitRate(br int64) {
	cctx.bit_rate = C.int64_t(br)
}

// GetCodecID Return codec_id
func (cctx *CodecContext) GetCodecID() CodecID {
	return CodecID(cctx.codec_id)
}

// SetCodecID Set codec_id
func (cctx *CodecContext) SetCodecID(codecID CodecID) {
	cctx.codec_id = C.enum_AVCodecID(codecID)
}

// GetCodecType Return codec_type
func (cctx *CodecContext) GetCodecType() MediaType {
	return MediaType(cctx.codec_type)
}

// SetCodecType Set codec_type
func (cctx *CodecContext) SetCodecType(ctype MediaType) {
	cctx.codec_type = C.enum_AVMediaType(ctype)
}

// GetTimeBase Return time_base
func (cctx *CodecContext) GetTimeBase() libavcodec.Rational {
	return newRational(cctx.time_base)
}

// SetTimeBase Set codec_type
func (cctx *CodecContext) SetTimeBase(timeBase libavcodec.Rational) {
	cctx.time_base.num = C.int(timeBase.Num())
	cctx.time_base.den = C.int(timeBase.Den())
}

// GetWidth Return width
func (cctx *CodecContext) GetWidth() int {
	return int(cctx.width)
}

// SetWidth Set width
func (cctx *CodecContext) SetWidth(w int) {
	cctx.width = C.int(w)
}

// GetHeight Return height
func (cctx *CodecContext) GetHeight() int {
	return int(cctx.height)
}

// SetHeight Set height
func (cctx *CodecContext) SetHeight(h int) {
	cctx.height = C.int(h)
}

// GetPixelFormat Return pix_fmt
func (cctx *CodecContext) GetPixelFormat() libavcodec.PixelFormat {
	return libavcodec.PixelFormat(C.int(cctx.pix_fmt))
}

// SetPixelFormat Set pix_fmt
func (cctx *CodecContext) SetPixelFormat(fmt libavcodec.PixelFormat) {
	cctx.pix_fmt = C.enum_AVPixelFormat(C.int(fmt))
}

// GetFlags Return flags
func (cctx *CodecContext) GetFlags() int {
	return int(cctx.flags)
}

// SetFlags Set flags
func (cctx *CodecContext) SetFlags(flags int) {
	cctx.flags = C.int(flags)
}

// GetMeRange Return me_range
func (cctx *CodecContext) GetMeRange() int {
	return int(cctx.me_range)
}

// SetMeRange Set me_range
func (cctx *CodecContext) SetMeRange(r int) {
	cctx.me_range = C.int(r)
}

// GetMaxQDiff Return max_qdiff
func (cctx *CodecContext) GetMaxQDiff() int {
	return int(cctx.max_qdiff)
}

// SetMaxQDiff Set max_qdiff
func (cctx *CodecContext) SetMaxQDiff(v int) {
	cctx.max_qdiff = C.int(v)
}

// GetQMin Return qmin
func (cctx *CodecContext) GetQMin() int {
	return int(cctx.qmin)
}

// SetQMin Set qmin
func (cctx *CodecContext) SetQMin(v int) {
	cctx.qmin = C.int(v)
}

// GetQMax Return qmax
func (cctx *CodecContext) GetQMax() int {
	return int(cctx.qmax)
}

// SetQMax Set qmax
func (cctx *CodecContext) SetQMax(v int) {
	cctx.qmax = C.int(v)
}

// GetQCompress Return qcompress
func (cctx *CodecContext) GetQCompress() float32 {
	return float32(cctx.qcompress)
}

// SetQCompress Set qcompress
func (cctx *CodecContext) SetQCompress(v float32) {
	cctx.qcompress = C.float(v)
}

// GetExtraData Return extradata
func (cctx *CodecContext) GetExtraData() []byte {
	header := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(cctx.extradata)),
		Len:  int(cctx.extradata_size),
		Cap:  int(cctx.extradata_size),
	}

	return *((*[]byte)(unsafe.Pointer(&header)))
}

// SetExtraData Return data
func (cctx *CodecContext) SetExtraData(data []byte) {
	header := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	cctx.extradata = (*C.uint8_t)(unsafe.Pointer(header.Data))
	cctx.extradata_size = C.int(header.Len)
}

// Release ...
func (cctx *CodecContext) Release() {
	C.avcodec_close((*C.struct_AVCodecContext)(unsafe.Pointer(cctx)))
	C.av_freep(unsafe.Pointer(cctx))
}
