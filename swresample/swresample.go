// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

// Package swresample provides a high-level interface to the libswresample library audio resampling utilities
// The process of changing the sampling rate of a discrete signal to obtain a new discrete representation of the underlying continuous signal.
package swresample

/*
	#cgo pkg-config: libswresample
	#include <libswresample/swresample.h>
*/
import "C"

type (
	Context        C.struct_SwrContext
	Frame          C.struct_AVFrame
	Class          C.struct_AVClass
	AvSampleFormat C.enum_AVSampleFormat
)

// GetClass Get the Class for Context.
func GetClass() *Class {
	return (*Class)(C.swr_get_class())
}

// Alloc Context constructor functions.Allocate Context.
func Alloc() *Context {
	return (*Context)(C.swr_alloc())
}

// Version Configuration accessors
func Version() uint {
	return uint(C.swresample_version())
}

func Configuration() string {
	return C.GoString(C.swresample_configuration())
}

func License() string {
	return C.GoString(C.swresample_license())
}
