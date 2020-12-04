// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package libavformat

//#cgo pkg-config: libavformat
//#include <libavformat/avformat.h>
import "C"
import "unsafe"

//AvStreamGetParser Return parser
func (st *AvStream) AvStreamGetParser() *AvCodecParserContext {
	return (*AvCodecParserContext)(C.av_stream_get_parser((*C.struct_AVStream)(st)))
}

// AvStreamGetEndPts Returns the pts of the last muxed packet + its duration.
func (st *AvStream) AvStreamGetEndPts() int64 {
	return int64(C.av_stream_get_end_pts((*C.struct_AVStream)(st)))
}

// AvStreamGetSideData Get side information from stream.
func (st *AvStream) AvStreamGetSideData(typ AvPacketSideDataType, size int) *uint8 {
	return (*uint8)(C.av_stream_get_side_data((*C.struct_AVStream)(st),
		(C.enum_AVPacketSideDataType)(typ), (*C.int)(unsafe.Pointer(&size))))
}
