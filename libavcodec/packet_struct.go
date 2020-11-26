// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package libavcodec

//#cgo pkg-config: libavcodec
//#include <libavcodec/avcodec.h>
import "C"

// Buf Return buf
func (p *Packet) Buf() *AvBufferRef {
	return (*AvBufferRef)(p.buf)
}

// Duration Return duration
func (p *Packet) Duration() int {
	return int(p.duration)
}

// SetDuration Set duration
func (p *Packet) SetDuration(dur int64) {
	p.duration = C.int64_t(dur)
}

// Flags Return flags
func (p *Packet) Flags() int {
	return int(p.flags)
}

// SetFlags Set flags
func (p *Packet) SetFlags(flags int) {
	p.flags = C.int(flags)
}

// SideDataElems Return side_data_elems
func (p *Packet) SideDataElems() int {
	return int(p.side_data_elems)
}

// Size Return size
func (p *Packet) Size() int {
	return int(p.size)
}

// StreamIndex Return stream_index
func (p *Packet) StreamIndex() int {
	return int(p.stream_index)
}

// SetStreamIndex Set idx
func (p *Packet) SetStreamIndex(idx int) {
	p.stream_index = C.int(idx)
}

// ConvergenceDuration Return convergence_duration
func (p *Packet) ConvergenceDuration() int64 {
	return int64(p.convergence_duration)
}

// Dts Return dts
func (p *Packet) Dts() int64 {
	return int64(p.dts)
}

// SetDts Set dts
func (p *Packet) SetDts(dts int64) {
	p.dts = C.int64_t(dts)
}

// Pos Return pos
func (p *Packet) Pos() int64 {
	return int64(p.pos)
}

// SetPos Set pos
func (p *Packet) SetPos(pos int64) {
	p.pos = C.int64_t(pos)
}

// Pts Return pts
func (p *Packet) Pts() int64 {
	return int64(p.pts)
}

// SetPts Set pts
func (p *Packet) SetPts(pts int64) {
	p.pts = C.int64_t(pts)
}

// Data Return data
func (p *Packet) Data() *uint8 {
	return (*uint8)(p.data)
}
