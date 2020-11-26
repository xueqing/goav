// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package libavcodec

//#cgo pkg-config: libavcodec
//#include <libavcodec/avcodec.h>
import "C"

// Buf Return buf
func (p *AvPacket) Buf() *AvBufferRef {
	return (*AvBufferRef)(p.buf)
}

// Duration Return duration
func (p *AvPacket) Duration() int {
	return int(p.duration)
}

// SetDuration Set duration
func (p *AvPacket) SetDuration(dur int64) {
	p.duration = C.int64_t(dur)
}

// Flags Return flags
func (p *AvPacket) Flags() int {
	return int(p.flags)
}

// SetFlags Set flags
func (p *AvPacket) SetFlags(flags int) {
	p.flags = C.int(flags)
}

// SideDataElems Return side_data_elems
func (p *AvPacket) SideDataElems() int {
	return int(p.side_data_elems)
}

// Size Return size
func (p *AvPacket) Size() int {
	return int(p.size)
}

// StreamIndex Return stream_index
func (p *AvPacket) StreamIndex() int {
	return int(p.stream_index)
}

// SetStreamIndex Set idx
func (p *AvPacket) SetStreamIndex(idx int) {
	p.stream_index = C.int(idx)
}

// ConvergenceDuration Return convergence_duration
func (p *AvPacket) ConvergenceDuration() int64 {
	return int64(p.convergence_duration)
}

// Dts Return dts
func (p *AvPacket) Dts() int64 {
	return int64(p.dts)
}

// SetDts Set dts
func (p *AvPacket) SetDts(dts int64) {
	p.dts = C.int64_t(dts)
}

// Pos Return pos
func (p *AvPacket) Pos() int64 {
	return int64(p.pos)
}

// SetPos Set pos
func (p *AvPacket) SetPos(pos int64) {
	p.pos = C.int64_t(pos)
}

// Pts Return pts
func (p *AvPacket) Pts() int64 {
	return int64(p.pts)
}

// SetPts Set pts
func (p *AvPacket) SetPts(pts int64) {
	p.pts = C.int64_t(pts)
}

// Data Return data
func (p *AvPacket) Data() *uint8 {
	return (*uint8)(p.data)
}
