// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

// Package libswresample provides a high-level interface to the libswresample library audio resampling utilities
// The process of changing the sampling rate of a discrete signal to obtain a new discrete representation of the underlying continuous signal.
package libswresample

/*
	#cgo pkg-config: libswresample
	#include <libswresample/swresample.h>
*/
import "C"

type (
	SwrContext     C.struct_SwrContext
	AvFrame        C.struct_AVFrame
	AvClass        C.struct_AVClass
	AvSampleFormat C.enum_AVSampleFormat
)

// SwrGetClass Get the Class for Context.
func SwrGetClass() *AvClass {
	return (*AvClass)(C.swr_get_class())
}

// SwrAlloc Context constructor functions.Allocate Context.
func SwrAlloc() *SwrContext {
	return (*SwrContext)(C.swr_alloc())
}

// SwresampleVersion Configuration accessors
func SwresampleVersion() uint {
	return uint(C.swresample_version())
}

// SwresampleConfiguration Return the swr build-time configuration.
func SwresampleConfiguration() string {
	return C.GoString(C.swresample_configuration())
}

// SwresampleLicense Return the swr license.
func SwresampleLicense() string {
	return C.GoString(C.swresample_license())
}
