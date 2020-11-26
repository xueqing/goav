// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package libswresample

/*
	#cgo pkg-config: libswresample
	#include <libswresample/swresample.h>
*/
import "C"
import (
	"unsafe"
)

// SwrInit Initialize context after user parameters have been set.
func (s *SwrContext) SwrInit() int {
	return int(C.swr_init((*C.struct_SwrContext)(s)))
}

// SwrIsInitialized Check whether an swr context has been initialized or not.
func (s *SwrContext) SwrIsInitialized() int {
	return int(C.swr_is_initialized((*C.struct_SwrContext)(s)))
}

// SwrAllocSetOpts Allocate Context if needed and set/reset common parameters.
func (s *SwrContext) SwrAllocSetOpts(ocl int64, osf AvSampleFormat, osr int, icl int64, isf AvSampleFormat, isr, lo, lc int) *SwrContext {
	return (*SwrContext)(C.swr_alloc_set_opts((*C.struct_SwrContext)(s), C.int64_t(ocl), (C.enum_AVSampleFormat)(osf), C.int(osr), C.int64_t(icl), (C.enum_AVSampleFormat)(isf), C.int(isr), C.int(lo), unsafe.Pointer(&lc)))
}

// SwrFree Context destructor functions. Free the given Context and set the pointer to NULL.
func (s *SwrContext) SwrFree() {
	C.swr_free((**C.struct_SwrContext)(unsafe.Pointer(s)))
}

// SwrClose Closes the context so that swr_is_initialized() returns 0.
func (s *SwrContext) SwrClose() {
	C.swr_close((*C.struct_SwrContext)(s))
}

// SwrConvert Core conversion functions. Convert audio
func (s *SwrContext) SwrConvert(out **uint8, oc int, in **uint8, ic int) int {
	return int(C.swr_convert((*C.struct_SwrContext)(s), (**C.uint8_t)(unsafe.Pointer(out)), C.int(oc), (**C.uint8_t)(unsafe.Pointer(in)), C.int(ic)))
}

// SwrNextPts Convert the next timestamp from input to output timestamps are in 1/(in_sample_rate * out_sample_rate) units.
func (s *SwrContext) SwrNextPts(pts int64) int64 {
	return int64(C.swr_next_pts((*C.struct_SwrContext)(s), C.int64_t(pts)))
}

// SwrSetCompensation Activate resampling compensation ("soft" compensation). This function is
// internally called when needed in swr_next_pts().
func (s *SwrContext) SwrSetCompensation(sd, cd int) int {
	return int(C.swr_set_compensation((*C.struct_SwrContext)(s), C.int(sd), C.int(cd)))
}

// SwrSetChannelMapping Set a customized input channel mapping.
func (s *SwrContext) SwrSetChannelMapping(cm *int) int {
	return int(C.swr_set_channel_mapping((*C.struct_SwrContext)(s), (*C.int)(unsafe.Pointer(cm))))
}

// SwrSetMatrix Set a customized remix matrix.
func (s *SwrContext) SwrSetMatrix(m *int, t int) int {
	return int(C.swr_set_matrix((*C.struct_SwrContext)(s), (*C.double)(unsafe.Pointer(m)), C.int(t)))
}

// SwrDropOutput Sample handling functions. Drops the specified number of output samples.
func (s *SwrContext) SwrDropOutput(c int) int {
	return int(C.swr_drop_output((*C.struct_SwrContext)(s), C.int(c)))
}

// SwrInjectSilence Injects the specified number of silence samples.
func (s *SwrContext) SwrInjectSilence(c int) int {
	return int(C.swr_inject_silence((*C.struct_SwrContext)(s), C.int(c)))
}

// SwrGetDelay Gets the delay the next input sample will experience relative to the next output sample.
func (s *SwrContext) SwrGetDelay(b int64) int64 {
	return int64(C.swr_get_delay((*C.struct_SwrContext)(s), C.int64_t(b)))
}

// SwrConvertFrame AvFrame based API. Convert the samples in the input AvFrame and write them to the output AvFrame.
func (s *SwrContext) SwrConvertFrame(o, i *AvFrame) int {
	return int(C.swr_convert_frame((*C.struct_SwrContext)(s), (*C.struct_AVFrame)(o), (*C.struct_AVFrame)(i)))
}

// SwrConfigFrame Configure or reconfigure the Context using the information provided by the AvFrames.
func (s *SwrContext) SwrConfigFrame(o, i *AvFrame) int {
	return int(C.swr_config_frame((*C.struct_SwrContext)(s), (*C.struct_AVFrame)(o), (*C.struct_AVFrame)(i)))
}
