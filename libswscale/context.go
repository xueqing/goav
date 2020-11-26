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
func SwsInitContext(ctxt *SwsContext, sf, df *SwsFilter) int {
	return int(C.sws_init_context((*C.struct_SwsContext)(ctxt), (*C.struct_SwsFilter)(sf), (*C.struct_SwsFilter)(df)))
}

// SwsFreecontext Free the swscaler context swsContext.
func SwsFreecontext(ctxt *SwsContext) {
	C.sws_freeContext((*C.struct_SwsContext)(ctxt))
}

// SwsGetcontext Allocate and return an Context.
func SwsGetcontext(sw, sh int, sf AvPixelFormat, dw, dh int, df AvPixelFormat, f int, sfl, dfl *SwsFilter, p *int) *SwsContext {
	return (*SwsContext)(C.sws_getContext(C.int(sw), C.int(sh), (C.enum_AVPixelFormat)(sf), C.int(dw), C.int(dh), (C.enum_AVPixelFormat)(df), C.int(f), (*C.struct_SwsFilter)(sfl), (*C.struct_SwsFilter)(dfl), (*C.double)(unsafe.Pointer(p))))
}

// SwsGetcachedcontext Check if context can be reused, otherwise reallocate a new one.
func SwsGetcachedcontext(ctxt *SwsContext, sw, sh int, sf AvPixelFormat, dw, dh int, df AvPixelFormat, f int, sfl, dfl *SwsFilter, p *float64) *SwsContext {
	return (*SwsContext)(C.sws_getCachedContext((*C.struct_SwsContext)(ctxt), C.int(sw), C.int(sh), (C.enum_AVPixelFormat)(sf), C.int(dw), C.int(dh), (C.enum_AVPixelFormat)(df), C.int(f), (*C.struct_SwsFilter)(sfl), (*C.struct_SwsFilter)(dfl), (*C.double)(p)))
}
