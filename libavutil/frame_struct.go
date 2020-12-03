package libavutil

//#cgo pkg-config: libavutil
//#include <libavutil/frame.h>
import "C"

import (
	"fmt"
	"unsafe"
)

// Data Return data
func (frame *AvFrame) Data() **uint8 {
	return (**uint8)(unsafe.Pointer(&frame.data[0]))
}

// Linesize Return linesize
func (frame *AvFrame) Linesize() *int32 {
	return (*int32)(unsafe.Pointer(&frame.linesize[0]))
}

// Pts Return Pts
func (frame *AvFrame) Pts() int64 {
	return int64(frame.pts)
}

// SetPts Set Pts
func (frame *AvFrame) SetPts(pts int64) {
	frame.pts = C.int64_t(pts)
}

// BestEffortTimestamp Return best_effort_timestamp
func (frame *AvFrame) BestEffortTimestamp() int64 {
	return int64(frame.best_effort_timestamp)
}

// Metadata Return metadata
func (frame *AvFrame) Metadata() *AvDictionary {
	return (*AvDictionary)(unsafe.Pointer(frame.metadata))
}

// GetInfo Return frame info.
func (frame *AvFrame) GetInfo() (width int, height int, linesize [8]int32, data [8]*uint8) {
	width = int(frame.linesize[0])
	height = int(frame.height)
	for i := range linesize {
		linesize[i] = int32(frame.linesize[i])
	}
	for i := range data {
		data[i] = (*uint8)(frame.data[i])
	}
	// log.Println("Linesize is ", frame.linesize, "Data is", data)
	return
}

// SetInfo Allocate new buffer(s) for audio or video data.
func (frame *AvFrame) SetInfo(width int, height int, pixFmt int) (err error) {
	frame.width = C.int(width)
	frame.height = C.int(height)
	frame.format = C.int(pixFmt)
	if ret := C.av_frame_get_buffer((*C.struct_AVFrame)(unsafe.Pointer(frame)), 32 /*alignment*/); ret < 0 {
		err = fmt.Errorf("Error allocating avframe buffer. Err: %v", ret)
		return
	}
	return
}

// FrameData return data
func FrameData(frame *AvFrame) (data [8]*uint8) {
	for i := range data {
		data[i] = (*uint8)(frame.data[i])
	}
	return
}

// FrameLinesize Return linesize
func FrameLinesize(frame *AvFrame) (linesize [8]int32) {
	for i := range linesize {
		linesize[i] = int32(frame.linesize[i])
	}
	return
}
