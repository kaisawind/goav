// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package swscale

//#cgo pkg-config: libswscale libavutil
//#include <libswscale/swscale.h>
import "C"
import (
	"unsafe"
)

// AllocVec Allocate and return an uninitialized vector with length coefficients.
func AllocVec(l int) *Vector {
	return (*Vector)(C.sws_allocVec(C.int(l)))
}

// GetGaussianVec Return a normalized Gaussian curve used to filter stuff quality = 3 is high quality, lower is lower quality.
func GetGaussianVec(v, q float64) *Vector {
	return (*Vector)(unsafe.Pointer(C.sws_getGaussianVec(C.double(v), C.double(q))))
}

// ScaleVec Scale all the coefficients of a by the scalar value.
func (a *Vector) ScaleVec(s float64) {
	C.sws_scaleVec((*C.struct_SwsVector)(unsafe.Pointer(a)), C.double(s))
}

// NormalizeVec Scale all the coefficients of a so that their sum equals height.
func (a *Vector) NormalizeVec(h float64) {
	C.sws_normalizeVec((*C.struct_SwsVector)(a), C.double(h))
}

func (a *Vector) FreeVec() {
	C.sws_freeVec((*C.struct_SwsVector)(a))
}
