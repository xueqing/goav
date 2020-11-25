package avutil

//#cgo pkg-config: libavutil
//#include <libavutil/imgutils.h>
import "C"
import "unsafe"

// AvImageFillArrays Setup the data pointers and linesizes based on the specified image parameters and the provided array.
func AvImageFillArrays(dstData [8]*uint8, dstLine [8]int32, src *uint8, pixFmt PixelFormat, width, height, align int) int {
	return int(C.av_image_fill_arrays((**C.uint8_t)(unsafe.Pointer(&dstData[0])), (*C.int)(unsafe.Pointer(&dstLine[0])), (*C.uint8_t)(src),
		C.enum_AVPixelFormat(pixFmt), C.int(width), C.int(height), C.int(align)))
}

// AvImageGetBufferSize Return the size in bytes of the amount of data required to store an image with the given parameters.
func AvImageGetBufferSize(pixFmt PixelFormat, width, height, align int) int {
	return int(C.av_image_get_buffer_size(C.enum_AVPixelFormat(pixFmt), C.int(width), C.int(height), C.int(align)))
}
