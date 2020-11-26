// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package libavformat

//#cgo pkg-config: libavformat
//#include <libavformat/avformat.h>
import "C"

//AvStreamGetParser Return parser
func (s *AvStream) AvStreamGetParser() *AvCodecParserContext {
	return (*AvCodecParserContext)(C.av_stream_get_parser((*C.struct_AVStream)(s)))
}

// AvStreamGetEndPts Returns the pts of the last muxed packet + its duration.
func (s *AvStream) AvStreamGetEndPts() int64 {
	return int64(C.av_stream_get_end_pts((*C.struct_AVStream)(s)))
}
