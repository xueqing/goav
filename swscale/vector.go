// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package swscale

//#cgo pkg-config: libswscale libavutil
//#include <libswscale/swscale.h>
import "C"
import (
	"unsafe"
)

//Allocate and return an uninitialized vector with length coefficients.
func SwsAllocvec(l int) *Vector {
	return (*Vector)(C.sws_allocVec(C.int(l)))
}

//Return a normalized Gaussian curve used to filter stuff quality = 3 is high quality, lower is lower quality.
func SwsGetgaussianvec(v, q float64) *Vector {
	return (*Vector)(unsafe.Pointer(C.sws_getGaussianVec(C.double(v), C.double(q))))
}

//Scale all the coefficients of a by the scalar value.
func (a *Vector) SwsScalevec(s float64) {
	C.sws_scaleVec((*C.struct_SwsVector)(unsafe.Pointer(a)), C.double(s))
}

//Scale all the coefficients of a so that their sum equals height.
func (a *Vector) SwsNormalizevec(h float64) {
	C.sws_normalizeVec((*C.struct_SwsVector)(a), C.double(h))
}

func (a *Vector) SwsFreevec() {
	C.sws_freeVec((*C.struct_SwsVector)(a))
}
