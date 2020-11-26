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
	"fmt"
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

// AvprivFrameGetMetadatap Return metadata
func AvprivFrameGetMetadatap(f *AvFrame) *AvDictionary {
	return (*AvDictionary)(unsafe.Pointer(f.metadata))
}

// AvFrameAlloc Allocate an Frame and set its fields to default values.
func AvFrameAlloc() *AvFrame {
	return (*AvFrame)(unsafe.Pointer(C.av_frame_alloc()))
}

// AvFrameFree Free the frame and any dynamically allocated objects in it, e.g.
func AvFrameFree(f *AvFrame) {
	C.av_frame_free((**C.struct_AVFrame)(unsafe.Pointer(&f)))
}

// AvFrameGetBuffer Allocate new buffer(s) for audio or video data.
func AvFrameGetBuffer(f *AvFrame, a int) int {
	return int(C.av_frame_get_buffer((*C.struct_AVFrame)(unsafe.Pointer(f)), C.int(a)))
}

// AvFrameRef Setup a new reference to the data described by an given frame.
func AvFrameRef(d, s *AvFrame) int {
	return int(C.av_frame_ref((*C.struct_AVFrame)(unsafe.Pointer(d)), (*C.struct_AVFrame)(unsafe.Pointer(s))))
}

// AvFrameClone Create a new frame that references the same data as src.
func AvFrameClone(f *AvFrame) *AvFrame {
	return (*AvFrame)(C.av_frame_clone((*C.struct_AVFrame)(unsafe.Pointer(f))))
}

// AvFrameUnref Unreference all the buffers referenced by frame and reset the frame fields.
func AvFrameUnref(f *AvFrame) {
	cf := (*C.struct_AVFrame)(unsafe.Pointer(f))
	C.av_frame_unref(cf)
}

// AvFrameMoveRef Move everythnig contained in src to dst and reset src.
func AvFrameMoveRef(d, s *AvFrame) {
	C.av_frame_move_ref((*C.struct_AVFrame)(unsafe.Pointer(d)), (*C.struct_AVFrame)(unsafe.Pointer(s)))
}

// AvFrameIsWritable Check if the frame data is writable.
func AvFrameIsWritable(f *AvFrame) int {
	return int(C.av_frame_is_writable((*C.struct_AVFrame)(unsafe.Pointer(f))))
}

// AvFrameMakeWritable Ensure that the frame data is writable, avoiding data copy if possible.
func AvFrameMakeWritable(f *AvFrame) int {
	return int(C.av_frame_make_writable((*C.struct_AVFrame)(unsafe.Pointer(f))))
}

// AvFrameCopyProps Copy only "metadata" fields from src to dst.
func AvFrameCopyProps(d, s *AvFrame) int {
	return int(C.av_frame_copy_props((*C.struct_AVFrame)(unsafe.Pointer(d)), (*C.struct_AVFrame)(unsafe.Pointer(s))))
}

// AvFrameGetPlaneBuffer Get the buffer reference a given data plane is stored in.
func AvFrameGetPlaneBuffer(f *AvFrame, p int) *AvBufferRef {
	return (*AvBufferRef)(C.av_frame_get_plane_buffer((*C.struct_AVFrame)(unsafe.Pointer(f)), C.int(p)))
}

// AvFrameNewSideData Add a new side data to a frame.
func AvFrameNewSideData(f *AvFrame, d AvFrameSideDataType, s int) *AvFrameSideData {
	return (*AvFrameSideData)(C.av_frame_new_side_data((*C.struct_AVFrame)(unsafe.Pointer(f)), (C.enum_AVFrameSideDataType)(d), C.int(s)))
}

// AvFrameGetSideData Return a pointer to the side data of a given type on success, NULL if there is no side data with such type in this frame.
func AvFrameGetSideData(f *AvFrame, t AvFrameSideDataType) *AvFrameSideData {
	return (*AvFrameSideData)(C.av_frame_get_side_data((*C.struct_AVFrame)(unsafe.Pointer(f)), (C.enum_AVFrameSideDataType)(t)))
}

// Data return data
func Data(f *AvFrame) (data [8]*uint8) {
	for i := range data {
		data[i] = (*uint8)(f.data[i])
	}
	return
}

// Linesize Return linesize
func Linesize(f *AvFrame) (linesize [8]int32) {
	for i := range linesize {
		linesize[i] = int32(f.linesize[i])
	}
	return
}

// GetPicture creates a YCbCr image from the frame
func GetPicture(f *AvFrame) (img *image.YCbCr, err error) {
	// For 4:4:4, CStride == YStride/1 && len(Cb) == len(Cr) == len(Y)/1.
	// For 4:2:2, CStride == YStride/2 && len(Cb) == len(Cr) == len(Y)/2.
	// For 4:2:0, CStride == YStride/2 && len(Cb) == len(Cr) == len(Y)/4.
	// For 4:4:0, CStride == YStride/1 && len(Cb) == len(Cr) == len(Y)/2.
	// For 4:1:1, CStride == YStride/4 && len(Cb) == len(Cr) == len(Y)/4.
	// For 4:1:0, CStride == YStride/4 && len(Cb) == len(Cr) == len(Y)/8.

	w := int(f.linesize[0])
	h := int(f.height)
	r := image.Rectangle{image.Point{0, 0}, image.Point{w, h}}
	// TODO: Use the sub sample ratio from the input image 'f.format'
	img = image.NewYCbCr(r, image.YCbCrSubsampleRatio420)
	// convert the frame data data to a Go byte array
	img.Y = C.GoBytes(unsafe.Pointer(f.data[0]), C.int(w*h))

	wCb := int(f.linesize[1])
	if unsafe.Pointer(f.data[1]) != nil {
		img.Cb = C.GoBytes(unsafe.Pointer(f.data[1]), C.int(wCb*h/2))
	}

	wCr := int(f.linesize[2])
	if unsafe.Pointer(f.data[2]) != nil {
		img.Cr = C.GoBytes(unsafe.Pointer(f.data[2]), C.int(wCr*h/2))
	}
	return
}

// SetPicture sets the image pointer of |f| to the image pointers of |img|
func SetPicture(f *AvFrame, img *image.YCbCr) {
	d := Data(f)
	// l := Linesize(f)
	// FIXME: Save the original pointers somewhere, this is a memory leak
	d[0] = (*uint8)(unsafe.Pointer(&img.Y[0]))
	// d[1] = (*uint8)(unsafe.Pointer(&img.Cb[0]))
}

// GetPictureRGB ...
func GetPictureRGB(f *AvFrame) (img *image.RGBA, err error) {
	w := int(f.linesize[0])
	h := int(f.height)
	r := image.Rectangle{image.Point{0, 0}, image.Point{w, h}}
	// TODO: Use the sub sample ratio from the input image 'f.format'
	img = image.NewRGBA(r)
	// convert the frame data data to a Go byte array
	img.Pix = C.GoBytes(unsafe.Pointer(f.data[0]), C.int(w*h))
	img.Stride = w
	log.Println("w", w, "h", h)
	return
}

// AvSetFrame Allocate new buffer(s) for audio or video data.
func AvSetFrame(f *AvFrame, w int, h int, pixFmt int) (err error) {
	f.width = C.int(w)
	f.height = C.int(h)
	f.format = C.int(pixFmt)
	if ret := C.av_frame_get_buffer((*C.struct_AVFrame)(unsafe.Pointer(f)), 32 /*alignment*/); ret < 0 {
		err = fmt.Errorf("Error allocating avframe buffer. Err: %v", ret)
		return
	}
	return
}

// AvFrameGetInfo Return frame info.
func AvFrameGetInfo(f *AvFrame) (width int, height int, linesize [8]int32, data [8]*uint8) {
	width = int(f.linesize[0])
	height = int(f.height)
	for i := range linesize {
		linesize[i] = int32(f.linesize[i])
	}
	for i := range data {
		data[i] = (*uint8)(f.data[i])
	}
	// log.Println("Linesize is ", f.linesize, "Data is", data)
	return
}

// GetBestEffortTimestamp Return best_effort_timestamp
func GetBestEffortTimestamp(f *AvFrame) int64 {
	return int64(f.best_effort_timestamp)
}

// //static int get_video_buffer (AvFrame *frame, int align)
// func GetVideoBuffer(f *AvFrame, a int) int {
// 	return int(C.get_video_buffer(f, C.int(a)))
// }

// //static int get_audio_buffer (AvFrame *frame, int align)
// func GetAudioBuffer(f *AvFrame, a int) int {
// 	return C.get_audio_buffer(f, C.int(a))
// }

// //static void get_frame_defaults (AvFrame *frame)
// func GetFrameDefaults(f *AvFrame) {
// 	C.get_frame_defaults(*C.struct_AVFrame(f))
// }
