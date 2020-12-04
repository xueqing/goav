// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package libavcodec

//#cgo pkg-config: libavcodec
//#include <libavcodec/avcodec.h>
import "C"
import (
	"unsafe"
)

// AV_PKT_xxx
const (
	AvPktFlagKey     = int(C.AV_PKT_FLAG_KEY)
	AvPktFlagCorrupt = int(C.AV_PKT_FLAG_CORRUPT)
	AvPktFlagDiscard = int(C.AV_PKT_FLAG_DISCARD)
)

// AvInitPacket Initialize optional fields of a packet with default values.
func (p *AvPacket) AvInitPacket() {
	C.av_init_packet((*C.struct_AVPacket)(p))
	p.size = 0
	p.data = nil
}

// AvNewPacket Allocate the payload of a packet and initialize its fields with default values.
func (p *AvPacket) AvNewPacket(size int) int {
	return int(C.av_new_packet((*C.struct_AVPacket)(p), C.int(size)))
}

// AvShrinkPacket Reduce packet size, correctly zeroing padding.
func (p *AvPacket) AvShrinkPacket(size int) {
	C.av_shrink_packet((*C.struct_AVPacket)(p), C.int(size))
}

// AvGrowPacket Increase packet size, correctly zeroing padding.
func (p *AvPacket) AvGrowPacket(size int) int {
	return int(C.av_grow_packet((*C.struct_AVPacket)(p), C.int(size)))
}

// AvPacketFromData Initialize a reference-counted packet from av_malloc()ed data.
func (p *AvPacket) AvPacketFromData(data *uint8, size int) int {
	return int(C.av_packet_from_data((*C.struct_AVPacket)(p), (*C.uint8_t)(data), C.int(size)))

}

// AvPacketNewSideData Allocate new information of a packet.
func (p *AvPacket) AvPacketNewSideData(typ AvPacketSideDataType, size int) *uint8 {
	return (*uint8)(C.av_packet_new_side_data((*C.struct_AVPacket)(p),
		(C.enum_AVPacketSideDataType)(typ), C.int(size)))
}

// AvPacketShrinkSideData Shrink the already allocated side data buffer.
func (p *AvPacket) AvPacketShrinkSideData(typ AvPacketSideDataType, size int) int {
	return int(C.av_packet_shrink_side_data((*C.struct_AVPacket)(p),
		(C.enum_AVPacketSideDataType)(typ), C.int(size)))
}

// AvPacketGetSideData Get side information from packet.
func (p *AvPacket) AvPacketGetSideData(typ AvPacketSideDataType, size *int) *uint8 {
	return (*uint8)(C.av_packet_get_side_data((*C.struct_AVPacket)(p),
		(C.enum_AVPacketSideDataType)(typ), (*C.int)(unsafe.Pointer(size))))
}

// AvPacketFreeSideData Convenience function to free all the side data stored.
func (p *AvPacket) AvPacketFreeSideData() {
	C.av_packet_free_side_data((*C.struct_AVPacket)(p))
}

// AvPacketRef Setup a new reference to the data described by a given packet.
func (p *AvPacket) AvPacketRef(src *AvPacket) int {
	return int(C.av_packet_ref((*C.struct_AVPacket)(p), (*C.struct_AVPacket)(src)))
}

// AvPacketUnref Wipe the packet.
func (p *AvPacket) AvPacketUnref() {
	C.av_packet_unref((*C.struct_AVPacket)(p))
}

// AvPacketMoveRef Move every field in src to dst and reset src.
func (p *AvPacket) AvPacketMoveRef(s *AvPacket) {
	C.av_packet_move_ref((*C.struct_AVPacket)(p), (*C.struct_AVPacket)(s))
}

// AvPacketCopyProps Copy only "properties" fields from src to dst.
func (p *AvPacket) AvPacketCopyProps(src *AvPacket) int {
	return int(C.av_packet_copy_props((*C.struct_AVPacket)(p), (*C.struct_AVPacket)(src)))
}

// AvPacketRescaleTs Convert valid timing fields (timestamps / durations) in a packet from one timebase to another.
func (p *AvPacket) AvPacketRescaleTs(tbSrc, tbDst AvRational) {
	C.av_packet_rescale_ts((*C.struct_AVPacket)(p), (C.struct_AVRational)(tbSrc),
		(C.struct_AVRational)(tbDst))
}
