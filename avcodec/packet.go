// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package avcodec

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
func (p *Packet) AvInitPacket() {
	C.av_init_packet((*C.struct_AVPacket)(p))
	p.size = 0
	p.data = nil
}

// AvNewPacket Allocate the payload of a packet and initialize its fields with default values.
func (p *Packet) AvNewPacket(s int) int {
	return int(C.av_new_packet((*C.struct_AVPacket)(p), C.int(s)))
}

// AvShrinkPacket Reduce packet size, correctly zeroing padding.
func (p *Packet) AvShrinkPacket(s int) {
	C.av_shrink_packet((*C.struct_AVPacket)(p), C.int(s))
}

// AvGrowPacket Increase packet size, correctly zeroing padding.
func (p *Packet) AvGrowPacket(s int) int {
	return int(C.av_grow_packet((*C.struct_AVPacket)(p), C.int(s)))
}

// AvPacketFromData Initialize a reference-counted packet from av_malloc()ed data.
func (p *Packet) AvPacketFromData(d *uint8, s int) int {
	return int(C.av_packet_from_data((*C.struct_AVPacket)(p), (*C.uint8_t)(d), C.int(s)))

}

// AvPacketNewSideData Allocate new information of a packet.
func (p *Packet) AvPacketNewSideData(t AvPacketSideDataType, s int) *uint8 {
	return (*uint8)(C.av_packet_new_side_data((*C.struct_AVPacket)(p), (C.enum_AVPacketSideDataType)(t), C.int(s)))
}

// AvPacketShrinkSideData Shrink the already allocated side data buffer.
func (p *Packet) AvPacketShrinkSideData(t AvPacketSideDataType, s int) int {
	return int(C.av_packet_shrink_side_data((*C.struct_AVPacket)(p), (C.enum_AVPacketSideDataType)(t), C.int(s)))
}

// AvPacketGetSideData Get side information from packet.
func (p *Packet) AvPacketGetSideData(t AvPacketSideDataType, s *int) *uint8 {
	return (*uint8)(C.av_packet_get_side_data((*C.struct_AVPacket)(p), (C.enum_AVPacketSideDataType)(t), (*C.int)(unsafe.Pointer(s))))
}

// AvPacketFreeSideData Convenience function to free all the side data stored.
func (p *Packet) AvPacketFreeSideData() {
	C.av_packet_free_side_data((*C.struct_AVPacket)(p))
}

// AvPacketRef Setup a new reference to the data described by a given packet.
func (p *Packet) AvPacketRef(s *Packet) int {
	return int(C.av_packet_ref((*C.struct_AVPacket)(p), (*C.struct_AVPacket)(s)))
}

// AvPacketUnref Wipe the packet.
func (p *Packet) AvPacketUnref() {
	C.av_packet_unref((*C.struct_AVPacket)(p))
}

// AvPacketMoveRef Move every field in src to dst and reset src.
func (p *Packet) AvPacketMoveRef(s *Packet) {
	C.av_packet_move_ref((*C.struct_AVPacket)(p), (*C.struct_AVPacket)(s))
}

// AvPacketCopyProps Copy only "properties" fields from src to dst.
func (p *Packet) AvPacketCopyProps(s *Packet) int {
	return int(C.av_packet_copy_props((*C.struct_AVPacket)(p), (*C.struct_AVPacket)(s)))
}

// AvPacketRescaleTs Convert valid timing fields (timestamps / durations) in a packet from one timebase to another.
func (p *Packet) AvPacketRescaleTs(r, r2 Rational) {
	C.av_packet_rescale_ts((*C.struct_AVPacket)(p), (C.struct_AVRational)(r), (C.struct_AVRational)(r2))
}
