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
func SwsAllocvec(l int) *SwsVector {
	return (*SwsVector)(C.sws_allocVec(C.int(l)))
}

// SwsGetgaussianvec Return a normalized Gaussian curve used to filter stuff quality = 3 is high quality, lower is lower quality.
func SwsGetgaussianvec(v, q float64) *SwsVector {
	return (*SwsVector)(unsafe.Pointer(C.sws_getGaussianVec(C.double(v), C.double(q))))
}

// SwsScalevec Scale all the coefficients of a by the scalar value.
func (a *SwsVector) SwsScalevec(s float64) {
	C.sws_scaleVec((*C.struct_SwsVector)(unsafe.Pointer(a)), C.double(s))
}

// SwsNormalizevec Scale all the coefficients of a so that their sum equals height.
func (a *SwsVector) SwsNormalizevec(h float64) {
	C.sws_normalizeVec((*C.struct_SwsVector)(a), C.double(h))
}

// SwsFreevec Free struct
func (a *SwsVector) SwsFreevec() {
	C.sws_freeVec((*C.struct_SwsVector)(a))
}
