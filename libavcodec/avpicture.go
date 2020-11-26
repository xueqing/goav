package libavcodec

//#cgo pkg-config: libavcodec
//#include <libavcodec/avcodec.h>
import "C"

// AvpictureFill Setup the picture fields based on the specified image
// parameters and the provided image data buffer.
func (pict *AvPicture) AvpictureFill(ptr *uint8, pixFmt AvPixelFormat, width, height int) int {
	return int(C.avpicture_fill((*C.struct_AVPicture)(pict), (*C.uint8_t)(ptr),
		(C.enum_AVPixelFormat)(pixFmt), C.int(width), C.int(height)))
}
