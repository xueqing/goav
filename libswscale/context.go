// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package libswscale

//#cgo pkg-config: libswscale libavutil
//#include <libswscale/swscale.h>
import "C"
import (
	"unsafe"
)

// SwsAllocContext Allocate an empty Context.
func SwsAllocContext() *SwsContext {
	return (*SwsContext)(C.sws_alloc_context())
}

// SwsInitContext Initialize the swscaler context sws_context.
func SwsInitContext(ctx *SwsContext, srcFilter, dstFilter *SwsFilter) int {
	return int(C.sws_init_context((*C.struct_SwsContext)(ctx),
		(*C.struct_SwsFilter)(srcFilter), (*C.struct_SwsFilter)(dstFilter)))
}

// SwsFreecontext Free the swscaler context swsContext.
func SwsFreecontext(ctx *SwsContext) {
	C.sws_freeContext((*C.struct_SwsContext)(ctx))
}

// SwsGetcontext Allocate and return an Context.
func SwsGetcontext(srcW, srcH int, srcFormat AvPixelFormat, dstW, dstH int, dstFormat AvPixelFormat,
	flags int, srcFilter, dstFilter *SwsFilter, param *int) *SwsContext {
	return (*SwsContext)(C.sws_getContext(
		C.int(srcW), C.int(srcH), (C.enum_AVPixelFormat)(srcFormat),
		C.int(dstW), C.int(dstH), (C.enum_AVPixelFormat)(dstFormat),
		C.int(flags), (*C.struct_SwsFilter)(srcFilter),
		(*C.struct_SwsFilter)(dstFilter), (*C.double)(unsafe.Pointer(param))))
}

// SwsGetcachedcontext Check if context can be reused, otherwise reallocate a new one.
func SwsGetcachedcontext(ctx *SwsContext, srcW, srcH int, srcFormat AvPixelFormat, dstW, dstH int, dstFormat AvPixelFormat,
	flags int, srcFilter, dstFilter *SwsFilter, param *float64) *SwsContext {
	return (*SwsContext)(C.sws_getCachedContext((*C.struct_SwsContext)(ctx),
		C.int(srcW), C.int(srcH), (C.enum_AVPixelFormat)(srcFormat),
		C.int(dstW), C.int(dstH), (C.enum_AVPixelFormat)(dstFormat),
		C.int(flags), (*C.struct_SwsFilter)(srcFilter),
		(*C.struct_SwsFilter)(dstFilter), (*C.double)(param)))
}
