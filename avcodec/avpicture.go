// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)
package avcodec

//#cgo pkg-config: libavcodec
//#include <libavcodec/avcodec.h>
import "C"

// AvpictureFill - Setup the picture fields based on the specified image parameters and the provided image data buffer.
func (p *Picture) AvpictureFill(pt *uint8, pf PixelFormat, w, h int) int {
	return int(C.avpicture_fill((*C.struct_AVPicture)(p), (*C.uint8_t)(pt), (C.enum_AVPixelFormat)(pf), C.int(w), C.int(h)))
}
