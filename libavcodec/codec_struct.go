package libavcodec

//#cgo pkg-config: libavcodec
//#include <libavcodec/avcodec.h>
/*
int count_pix_fmts(const enum AVPixelFormat *pix_fmts)
{
    int cnt = 0;
	if (pix_fmts)
		while (pix_fmts[cnt] != AV_PIX_FMT_NONE)
			cnt++;
    return cnt;
}

int count_sample_fmts(const enum AVSampleFormat *sample_fmts)
{
	  int cnt = 0;
	if (sample_fmts)
		while (sample_fmts[cnt] != AV_SAMPLE_FMT_NONE)
			cnt++;
    return cnt;
}
*/
import "C"
import (
	"reflect"
	"unsafe"
)

// Name Return name
func (codec *AvCodec) Name() string {
	return C.GoString(codec.name)
}

// Capabilities Return capabilities
func (codec *AvCodec) Capabilities() int {
	return int(codec.capabilities)
}

// PixFmts Return pix_fmts
func (codec *AvCodec) PixFmts() []AvPixelFormat {
	cnt := int(C.count_pix_fmts(codec.pix_fmts))
	if cnt == 0 {
		return nil
	}
	return *(*[]AvPixelFormat)(unsafe.Pointer(&reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(codec.pix_fmts)),
		Len:  cnt,
		Cap:  cnt,
	}))
}

// SampleFmts Return sample_fmts
func (codec *AvCodec) SampleFmts() []AvSampleFormat {
	cnt := int(C.count_sample_fmts(codec.sample_fmts))
	if cnt == 0 {
		return nil
	}
	return *(*[]AvSampleFormat)(unsafe.Pointer(&reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(codec.sample_fmts)),
		Len:  cnt,
		Cap:  cnt,
	}))
}
