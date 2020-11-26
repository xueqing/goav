// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package libswscale

//#cgo pkg-config: libswscale libavutil
//#include <libswscale/swscale.h>
import "C"
import (
	"unsafe"
)

// SwsAllocvec Allocate and return an uninitialized vector with length coefficients.
func SwsAllocvec(length int) *SwsVector {
	return (*SwsVector)(C.sws_allocVec(C.int(length)))
}

// SwsGetgaussianvec Return a normalized Gaussian curve used to filter stuff quality = 3 is high quality, lower is lower quality.
func SwsGetgaussianvec(variance, quality float64) *SwsVector {
	return (*SwsVector)(unsafe.Pointer(C.sws_getGaussianVec(C.double(variance), C.double(quality))))
}

// SwsScalevec Scale all the coefficients of a by the scalar value.
func (vec *SwsVector) SwsScalevec(scalar float64) {
	C.sws_scaleVec((*C.struct_SwsVector)(unsafe.Pointer(vec)), C.double(scalar))
}

// SwsNormalizevec Scale all the coefficients of a so that their sum equals height.
func (vec *SwsVector) SwsNormalizevec(height float64) {
	C.sws_normalizeVec((*C.struct_SwsVector)(vec), C.double(height))
}

// SwsFreevec Free struct
func (vec *SwsVector) SwsFreevec() {
	C.sws_freeVec((*C.struct_SwsVector)(vec))
}
