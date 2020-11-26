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
func (ctx *SwrContext) SwrInit() int {
	return int(C.swr_init((*C.struct_SwrContext)(ctx)))
}

// SwrIsInitialized Check whether an swr context has been initialized or not.
func (ctx *SwrContext) SwrIsInitialized() int {
	return int(C.swr_is_initialized((*C.struct_SwrContext)(ctx)))
}

// SwrAllocSetOpts Allocate Context if needed and set/reset common parameters.
func (ctx *SwrContext) SwrAllocSetOpts(outChLayout int64, outSampleFmt AvSampleFormat, outSampleRate int,
	inChLayout int64, inSampleFmt AvSampleFormat, inSampleRate, logOffset, logCtx int) *SwrContext {
	return (*SwrContext)(C.swr_alloc_set_opts((*C.struct_SwrContext)(ctx),
		C.int64_t(outChLayout), (C.enum_AVSampleFormat)(outSampleFmt), C.int(outSampleRate),
		C.int64_t(inChLayout), (C.enum_AVSampleFormat)(inSampleFmt), C.int(inSampleRate),
		C.int(logOffset), unsafe.Pointer(&logCtx)))
}

// SwrFree Context destructor functions. Free the given Context and set the pointer to NULL.
func (ctx *SwrContext) SwrFree() {
	C.swr_free((**C.struct_SwrContext)(unsafe.Pointer(ctx)))
}

// SwrClose Closes the context so that swr_is_initialized() returns 0.
func (ctx *SwrContext) SwrClose() {
	C.swr_close((*C.struct_SwrContext)(ctx))
}

// SwrConvert Core conversion functions. Convert audio
func (ctx *SwrContext) SwrConvert(out **uint8, outCount int, in **uint8, inCount int) int {
	return int(C.swr_convert((*C.struct_SwrContext)(ctx),
		(**C.uint8_t)(unsafe.Pointer(out)), C.int(outCount),
		(**C.uint8_t)(unsafe.Pointer(in)), C.int(inCount)))
}

// SwrNextPts Convert the next timestamp from input to output timestamps are in 1/(in_sample_rate * out_sample_rate) units.
func (ctx *SwrContext) SwrNextPts(pts int64) int64 {
	return int64(C.swr_next_pts((*C.struct_SwrContext)(ctx), C.int64_t(pts)))
}

// SwrSetCompensation Activate resampling compensation ("soft" compensation). This function is
// internally called when needed in swr_next_pts().
func (ctx *SwrContext) SwrSetCompensation(sampleDelta, compensationDistance int) int {
	return int(C.swr_set_compensation((*C.struct_SwrContext)(ctx),
		C.int(sampleDelta), C.int(compensationDistance)))
}

// SwrSetChannelMapping Set a customized input channel mapping.
func (ctx *SwrContext) SwrSetChannelMapping(channelMap *int) int {
	return int(C.swr_set_channel_mapping((*C.struct_SwrContext)(ctx),
		(*C.int)(unsafe.Pointer(channelMap))))
}

// SwrSetMatrix Set a customized remix matrix.
func (ctx *SwrContext) SwrSetMatrix(matrix *int, stride int) int {
	return int(C.swr_set_matrix((*C.struct_SwrContext)(ctx),
		(*C.double)(unsafe.Pointer(matrix)), C.int(stride)))
}

// SwrDropOutput Sample handling functions. Drops the specified number of output samples.
func (ctx *SwrContext) SwrDropOutput(count int) int {
	return int(C.swr_drop_output((*C.struct_SwrContext)(ctx), C.int(count)))
}

// SwrInjectSilence Injects the specified number of silence samples.
func (ctx *SwrContext) SwrInjectSilence(count int) int {
	return int(C.swr_inject_silence((*C.struct_SwrContext)(ctx), C.int(count)))
}

// SwrGetDelay Gets the delay the next input sample will experience relative to the next output sample.
func (ctx *SwrContext) SwrGetDelay(base int64) int64 {
	return int64(C.swr_get_delay((*C.struct_SwrContext)(ctx), C.int64_t(base)))
}

// SwrConvertFrame AvFrame based API. Convert the samples in the input AvFrame and write them to the output AvFrame.
func (ctx *SwrContext) SwrConvertFrame(output, input *AvFrame) int {
	return int(C.swr_convert_frame((*C.struct_SwrContext)(ctx),
		(*C.struct_AVFrame)(output), (*C.struct_AVFrame)(input)))
}

// SwrConfigFrame Configure or reconfigure the Context using the information provided by the AvFrames.
func (ctx *SwrContext) SwrConfigFrame(output, input *AvFrame) int {
	return int(C.swr_config_frame((*C.struct_SwrContext)(ctx),
		(*C.struct_AVFrame)(output), (*C.struct_AVFrame)(input)))
}
