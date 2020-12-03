// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package libavutil

/*
	#cgo pkg-config: libavutil
	#include <libavutil/frame.h>
	#include <stdlib.h>
*/
import "C"
import (
	"image"
	"log"
	"unsafe"
)

type (
	AvBuffer            C.struct_AVBuffer
	AvBufferRef         C.struct_AVBufferRef
	AvBufferPool        C.struct_AVBufferPool
	AvFrame             C.struct_AVFrame
	AvFrameSideData     C.struct_AVFrameSideData
	AvFrameSideDataType C.enum_AVFrameSideDataType
)

// AvFrameAlloc Allocate an Frame and set its fields to default values.
func AvFrameAlloc() *AvFrame {
	return (*AvFrame)(unsafe.Pointer(C.av_frame_alloc()))
}

// AvFrameFree Free the frame and any dynamically allocated objects in it, e.g.
func AvFrameFree(frame *AvFrame) {
	C.av_frame_free((**C.struct_AVFrame)(unsafe.Pointer(&frame)))
}

// AvFrameGetBuffer Allocate new buffer(s) for audio or video data.
func AvFrameGetBuffer(frame *AvFrame, align int) int {
	return int(C.av_frame_get_buffer((*C.struct_AVFrame)(unsafe.Pointer(frame)), C.int(align)))
}

// AvFrameRef Setup a new reference to the data described by an given frame.
func AvFrameRef(dst, src *AvFrame) int {
	return int(C.av_frame_ref((*C.struct_AVFrame)(unsafe.Pointer(dst)),
		(*C.struct_AVFrame)(unsafe.Pointer(src))))
}

// AvFrameClone Create a new frame that references the same data as src.
func AvFrameClone(src *AvFrame) *AvFrame {
	return (*AvFrame)(C.av_frame_clone((*C.struct_AVFrame)(unsafe.Pointer(src))))
}

// AvFrameUnref Unreference all the buffers referenced by frame and reset the frame fields.
func AvFrameUnref(frame *AvFrame) {
	C.av_frame_unref((*C.struct_AVFrame)(unsafe.Pointer(frame)))
}

// AvFrameMoveRef Move everythnig contained in src to dst and reset src.
func AvFrameMoveRef(dst, src *AvFrame) {
	C.av_frame_move_ref((*C.struct_AVFrame)(unsafe.Pointer(dst)),
		(*C.struct_AVFrame)(unsafe.Pointer(src)))
}

// AvFrameIsWritable Check if the frame data is writable.
func AvFrameIsWritable(frame *AvFrame) int {
	return int(C.av_frame_is_writable((*C.struct_AVFrame)(unsafe.Pointer(frame))))
}

// AvFrameMakeWritable Ensure that the frame data is writable, avoiding data copy if possible.
func AvFrameMakeWritable(frame *AvFrame) int {
	return int(C.av_frame_make_writable((*C.struct_AVFrame)(unsafe.Pointer(frame))))
}

// AvFrameCopyProps Copy only "metadata" fields from src to dst.
func AvFrameCopyProps(dst, src *AvFrame) int {
	return int(C.av_frame_copy_props((*C.struct_AVFrame)(unsafe.Pointer(dst)),
		(*C.struct_AVFrame)(unsafe.Pointer(src))))
}

// AvFrameGetPlaneBuffer Get the buffer reference a given data plane is stored in.
func AvFrameGetPlaneBuffer(frame *AvFrame, plane int) *AvBufferRef {
	return (*AvBufferRef)(C.av_frame_get_plane_buffer((*C.struct_AVFrame)(unsafe.Pointer(frame)),
		C.int(plane)))
}

// AvFrameNewSideData Add a new side data to a frame.
func AvFrameNewSideData(frame *AvFrame, typ AvFrameSideDataType, size int) *AvFrameSideData {
	return (*AvFrameSideData)(C.av_frame_new_side_data((*C.struct_AVFrame)(unsafe.Pointer(frame)),
		(C.enum_AVFrameSideDataType)(typ), C.int(size)))
}

// AvFrameGetSideData Return a pointer to the side data of a given type on success, NULL if there is no side data with such type in this frame.
func AvFrameGetSideData(frame *AvFrame, typ AvFrameSideDataType) *AvFrameSideData {
	return (*AvFrameSideData)(C.av_frame_get_side_data((*C.struct_AVFrame)(unsafe.Pointer(frame)),
		(C.enum_AVFrameSideDataType)(typ)))
}

// GetPicture creates a YCbCr image from the frame
func GetPicture(frame *AvFrame) (img *image.YCbCr, err error) {
	// For 4:4:4, CStride == YStride/1 && len(Cb) == len(Cr) == len(Y)/1.
	// For 4:2:2, CStride == YStride/2 && len(Cb) == len(Cr) == len(Y)/2.
	// For 4:2:0, CStride == YStride/2 && len(Cb) == len(Cr) == len(Y)/4.
	// For 4:4:0, CStride == YStride/1 && len(Cb) == len(Cr) == len(Y)/2.
	// For 4:1:1, CStride == YStride/4 && len(Cb) == len(Cr) == len(Y)/4.
	// For 4:1:0, CStride == YStride/4 && len(Cb) == len(Cr) == len(Y)/8.

	w := int(frame.linesize[0])
	h := int(frame.height)
	r := image.Rectangle{image.Point{0, 0}, image.Point{w, h}}
	// TODO: Use the sub sample ratio from the input image 'frame.format'
	img = image.NewYCbCr(r, image.YCbCrSubsampleRatio420)
	// convert the frame data data to a Go byte array
	img.Y = C.GoBytes(unsafe.Pointer(frame.data[0]), C.int(w*h))

	wCb := int(frame.linesize[1])
	if unsafe.Pointer(frame.data[1]) != nil {
		img.Cb = C.GoBytes(unsafe.Pointer(frame.data[1]), C.int(wCb*h/2))
	}

	wCr := int(frame.linesize[2])
	if unsafe.Pointer(frame.data[2]) != nil {
		img.Cr = C.GoBytes(unsafe.Pointer(frame.data[2]), C.int(wCr*h/2))
	}
	return
}

// GetPictureRGB ...
func GetPictureRGB(frame *AvFrame) (img *image.RGBA, err error) {
	w := int(frame.linesize[0])
	h := int(frame.height)
	r := image.Rectangle{image.Point{0, 0}, image.Point{w, h}}
	// TODO: Use the sub sample ratio from the input image 'frame.format'
	img = image.NewRGBA(r)
	// convert the frame data data to a Go byte array
	img.Pix = C.GoBytes(unsafe.Pointer(frame.data[0]), C.int(w*h))
	img.Stride = w
	log.Println("w", w, "h", h)
	return
}

// //static int get_video_buffer (AvFrame *frame, int align)
// func GetVideoBuffer(frame *AvFrame, a int) int {
// 	return int(C.get_video_buffer(frame, C.int(a)))
// }

// //static int get_audio_buffer (AvFrame *frame, int align)
// func GetAudioBuffer(frame *AvFrame, a int) int {
// 	return C.get_audio_buffer(frame, C.int(a))
// }

// //static void get_frame_defaults (AvFrame *frame)
// func GetFrameDefaults(frame *AvFrame) {
// 	C.get_frame_defaults(*C.struct_AVFrame(frame))
// }
