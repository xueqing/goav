package libavformat

//#cgo pkg-config: libavformat
//#include <libavformat/avformat.h>
import "C"
import (
	"unsafe"

	"github.com/xueqing/goav/libavcodec"
	"github.com/xueqing/goav/libavutil"
)

// AvIOOpen Create and initialize a AVIOContext for accessing the resource indicated by url.
func AvIOOpen(url string, flags int) (ioctx *AvIOContext, err error) {
	cURL := C.CString(url)
	defer C.free(unsafe.Pointer(cURL))
	err = libavutil.ErrorFromCode(int(C.avio_open((**C.AVIOContext)(unsafe.Pointer(&ioctx)), cURL, C.int(flags))))
	return
}

// AvioClose Close the resource accessed by the AVIOContext s and free it.
// This function can only be used if s was opened by avio_open().
// The internal buffer is automatically flushed before closing the resource.
func (ctxt *AvIOContext) AvioClose() error {
	return libavutil.ErrorFromCode(int(C.avio_close((*C.AVIOContext)(unsafe.Pointer(ctxt)))))
}

// AvioClosep the resource accessed by the AVIOContext *s, free it and set the pointer pointing to it to NULL.
// This function can only be used if s was opened by avio_open().
func AvioClosep(ctxt *AvIOContext) error {
	return libavutil.ErrorFromCode(int(C.avio_closep((**C.AVIOContext)(unsafe.Pointer(&ctxt)))))
}

// AvGetPacket Allocate and read the payload of a packet and initialize its fields with default values.
func (ctxt *AvIOContext) AvGetPacket(pkt *libavcodec.AvPacket, size int) int {
	return int(C.av_get_packet((*C.struct_AVIOContext)(ctxt),
		toCPacket(pkt), C.int(size)))
}

// AvAppendPacket Read data and append it to the current content of the Packet.
func (ctxt *AvIOContext) AvAppendPacket(pkt *libavcodec.AvPacket, size int) int {
	return int(C.av_append_packet((*C.struct_AVIOContext)(ctxt),
		toCPacket(pkt), C.int(size)))
}
