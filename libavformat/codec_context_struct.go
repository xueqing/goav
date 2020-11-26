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

// Type Return codec_type as AvMediaType
func (cctx *AvCodecContext) Type() AvMediaType {
	return AvMediaType(cctx.codec_type)
}

// SetBitRate Set bit_rate
func (cctx *AvCodecContext) SetBitRate(br int64) {
	cctx.bit_rate = C.int64_t(br)
}

// GetCodecID Return codec_id
func (cctx *AvCodecContext) GetCodecID() AvCodecID {
	return AvCodecID(cctx.codec_id)
}

// SetCodecID Set codec_id
func (cctx *AvCodecContext) SetCodecID(codecID AvCodecID) {
	cctx.codec_id = C.enum_AVCodecID(codecID)
}

// GetCodecType Return codec_type
func (cctx *AvCodecContext) GetCodecType() AvMediaType {
	return AvMediaType(cctx.codec_type)
}

// SetCodecType Set codec_type
func (cctx *AvCodecContext) SetCodecType(ctype AvMediaType) {
	cctx.codec_type = C.enum_AVMediaType(ctype)
}

// GetTimeBase Return time_base
func (cctx *AvCodecContext) GetTimeBase() libavcodec.AvRational {
	return newRational(cctx.time_base)
}

// SetTimeBase Set codec_type
func (cctx *AvCodecContext) SetTimeBase(timeBase libavcodec.AvRational) {
	cctx.time_base.num = C.int(timeBase.Num())
	cctx.time_base.den = C.int(timeBase.Den())
}

// GetWidth Return width
func (cctx *AvCodecContext) GetWidth() int {
	return int(cctx.width)
}

// SetWidth Set width
func (cctx *AvCodecContext) SetWidth(w int) {
	cctx.width = C.int(w)
}

// GetHeight Return height
func (cctx *AvCodecContext) GetHeight() int {
	return int(cctx.height)
}

// SetHeight Set height
func (cctx *AvCodecContext) SetHeight(h int) {
	cctx.height = C.int(h)
}

// GetPixelFormat Return pix_fmt
func (cctx *AvCodecContext) GetPixelFormat() libavcodec.AvPixelFormat {
	return libavcodec.AvPixelFormat(C.int(cctx.pix_fmt))
}

// SetPixelFormat Set pix_fmt
func (cctx *AvCodecContext) SetPixelFormat(fmt libavcodec.AvPixelFormat) {
	cctx.pix_fmt = C.enum_AVPixelFormat(C.int(fmt))
}

// GetFlags Return flags
func (cctx *AvCodecContext) GetFlags() int {
	return int(cctx.flags)
}

// SetFlags Set flags
func (cctx *AvCodecContext) SetFlags(flags int) {
	cctx.flags = C.int(flags)
}

// GetMeRange Return me_range
func (cctx *AvCodecContext) GetMeRange() int {
	return int(cctx.me_range)
}

// SetMeRange Set me_range
func (cctx *AvCodecContext) SetMeRange(r int) {
	cctx.me_range = C.int(r)
}

// GetMaxQDiff Return max_qdiff
func (cctx *AvCodecContext) GetMaxQDiff() int {
	return int(cctx.max_qdiff)
}

// SetMaxQDiff Set max_qdiff
func (cctx *AvCodecContext) SetMaxQDiff(v int) {
	cctx.max_qdiff = C.int(v)
}

// GetQMin Return qmin
func (cctx *AvCodecContext) GetQMin() int {
	return int(cctx.qmin)
}

// SetQMin Set qmin
func (cctx *AvCodecContext) SetQMin(v int) {
	cctx.qmin = C.int(v)
}

// GetQMax Return qmax
func (cctx *AvCodecContext) GetQMax() int {
	return int(cctx.qmax)
}

// SetQMax Set qmax
func (cctx *AvCodecContext) SetQMax(v int) {
	cctx.qmax = C.int(v)
}

// GetQCompress Return qcompress
func (cctx *AvCodecContext) GetQCompress() float32 {
	return float32(cctx.qcompress)
}

// SetQCompress Set qcompress
func (cctx *AvCodecContext) SetQCompress(v float32) {
	cctx.qcompress = C.float(v)
}

// GetExtraData Return extradata
func (cctx *AvCodecContext) GetExtraData() []byte {
	header := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(cctx.extradata)),
		Len:  int(cctx.extradata_size),
		Cap:  int(cctx.extradata_size),
	}

	return *((*[]byte)(unsafe.Pointer(&header)))
}

// SetExtraData Return data
func (cctx *AvCodecContext) SetExtraData(data []byte) {
	header := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	cctx.extradata = (*C.uint8_t)(unsafe.Pointer(header.Data))
	cctx.extradata_size = C.int(header.Len)
}

// Release ...
func (cctx *AvCodecContext) Release() {
	C.avcodec_close((*C.struct_AVCodecContext)(unsafe.Pointer(cctx)))
	C.av_freep(unsafe.Pointer(cctx))
}
