// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package libavcodec

//#cgo pkg-config: libavcodec
//#include <libavcodec/avcodec.h>
import "C"
import (
	"unsafe"
)

// AvcodecPixFmtToCodecTag Return a value representing the fourCC code associated to the pixel format pix_fmt, or 0 if no associated fourCC code can be found.
func (p AvPixelFormat) AvcodecPixFmtToCodecTag() uint {
	return uint(C.avcodec_pix_fmt_to_codec_tag((C.enum_AVPixelFormat)(p)))
}

// AvcodecGetPixFmtLoss see av_get_pix_fmt_loss()
func (p AvPixelFormat) AvcodecGetPixFmtLoss(f AvPixelFormat, a int) int {
	return int(C.avcodec_get_pix_fmt_loss((C.enum_AVPixelFormat)(p), (C.enum_AVPixelFormat)(f), C.int(a)))
}

// AvcodecFindBestPixFmtOfList Find the best pixel format to convert to given a certain source pixel format.
func (p *AvPixelFormat) AvcodecFindBestPixFmtOfList(s AvPixelFormat, a int, l *int) AvPixelFormat {
	return (AvPixelFormat)(C.avcodec_find_best_pix_fmt_of_list((*C.enum_AVPixelFormat)(p), (C.enum_AVPixelFormat)(s), C.int(a), (*C.int)(unsafe.Pointer(l))))
}

// AvcodecFindBestPixFmtOf2 see av_find_best_pix_fmt_of_2()
func (p AvPixelFormat) AvcodecFindBestPixFmtOf2(f2, s AvPixelFormat, a int, l *int) AvPixelFormat {
	return (AvPixelFormat)(C.avcodec_find_best_pix_fmt_of_2((C.enum_AVPixelFormat)(p), (C.enum_AVPixelFormat)(f2), (C.enum_AVPixelFormat)(s), C.int(a), (*C.int)(unsafe.Pointer(l))))
}
