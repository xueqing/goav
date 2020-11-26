package libavcodec

//#cgo pkg-config: libavcodec
//#include <libavcodec/avcodec.h>
import "C"

// AvpictureFill Setup the picture fields based on the specified image parameters and the provided image data buffer.
func (p *AvPicture) AvpictureFill(pt *uint8, pf AvPixelFormat, w, h int) int {
	return int(C.avpicture_fill((*C.struct_AVPicture)(p), (*C.uint8_t)(pt), (C.enum_AVPixelFormat)(pf), C.int(w), C.int(h)))
}
