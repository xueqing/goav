package libavcodec

//#cgo pkg-config: libavcodec
//#include <libavcodec/avcodec.h>
import "C"

// AvcodecDescriptorNext Iterate over all codec descriptors known to libavcodec.
func (desc *AvCodecDescriptor) AvcodecDescriptorNext() *AvCodecDescriptor {
	return (*AvCodecDescriptor)(C.avcodec_descriptor_next((*C.struct_AVCodecDescriptor)(desc)))
}
