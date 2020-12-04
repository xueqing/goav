package libavcodec

//#cgo pkg-config: libavcodec
//#include <libavcodec/avcodec.h>
import "C"

// AvcodecParametersCopy Copy the contents of src to dst. Any allocated fields in dst are freed and
// replaced with newly allocated duplicates of the corresponding fields in src.
func (cp *AvCodecParameters) AvcodecParametersCopy(src *AvCodecParameters) int {
	return int(C.avcodec_parameters_copy((*C.struct_AVCodecParameters)(cp), (*C.struct_AVCodecParameters)(src)))
}
