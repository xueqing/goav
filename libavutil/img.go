package libavutil

//#cgo pkg-config: libavutil
//#include <libavutil/imgutils.h>
import "C"
import "unsafe"

// AvImageFillArrays Setup the data pointers and linesizes based on the specified image parameters and the provided array.
func AvImageFillArrays(dstData **uint8, dstLine *int32, src *uint8, pixFmt AvPixelFormat, width, height, align int) int {
	return int(C.av_image_fill_arrays((**C.uint8_t)(unsafe.Pointer(dstData)), (*C.int)(dstLine), (*C.uint8_t)(src),
		C.enum_AVPixelFormat(pixFmt), C.int(width), C.int(height), C.int(align)))
}

// AvImageGetBufferSize Return the size in bytes of the amount of data required to store an image with the given parameters.
func AvImageGetBufferSize(pixFmt AvPixelFormat, width, height, align int) int {
	return int(C.av_image_get_buffer_size(C.enum_AVPixelFormat(pixFmt), C.int(width), C.int(height), C.int(align)))
}
