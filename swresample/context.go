// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package swresample

/*
	#cgo pkg-config: libswresample
	#include <libswresample/swresample.h>
*/
import "C"
import (
	"unsafe"
)

// Init Initialize context after user parameters have been set.
func (ctxt *Context) Init() int {
	return int(C.swr_init((*C.struct_SwrContext)(ctxt)))
}

// IsInitialized Check whether an swr context has been initialized or not.
func (ctxt *Context) IsInitialized() int {
	return int(C.swr_is_initialized((*C.struct_SwrContext)(ctxt)))
}

// AllocSetOpts Allocate Context if needed and set/reset common parameters.
func (ctxt *Context) AllocSetOpts(ocl int64, osf AvSampleFormat, osr int, icl int64, isf AvSampleFormat, isr, lo, lc int) *Context {
	return (*Context)(C.swr_alloc_set_opts((*C.struct_SwrContext)(ctxt), C.int64_t(ocl), (C.enum_AVSampleFormat)(osf), C.int(osr), C.int64_t(icl), (C.enum_AVSampleFormat)(isf), C.int(isr), C.int(lo), unsafe.Pointer(&lc)))
}

// Free Context destructor functions. Free the given Context and set the pointer to NULL.
func (ctxt *Context) Free() {
	C.swr_free((**C.struct_SwrContext)(unsafe.Pointer(&ctxt)))
}

// Close Closes the context so that swr_is_initialized() returns 0.
func (ctxt *Context) Close() {
	C.swr_close((*C.struct_SwrContext)(ctxt))
}

// Convert Core conversion functions. Convert audio
func (ctxt *Context) Convert(out **uint8, oc int, in **uint8, ic int) int {
	return int(C.swr_convert((*C.struct_SwrContext)(ctxt), (**C.uint8_t)(unsafe.Pointer(out)), C.int(oc), (**C.uint8_t)(unsafe.Pointer(in)), C.int(ic)))
}

// NextPts Convert the next timestamp from input to output timestamps are in 1/(in_sample_rate * out_sample_rate) units.
func (ctxt *Context) NextPts(pts int64) int64 {
	return int64(C.swr_next_pts((*C.struct_SwrContext)(ctxt), C.int64_t(pts)))
}

// SetCompensation Low-level option setting functions
//These functons provide a means to set low-level options that is not possible with the AvOption API.
//Activate resampling compensation ("soft" compensation).
func (ctxt *Context) SetCompensation(sd, cd int) int {
	return int(C.swr_set_compensation((*C.struct_SwrContext)(ctxt), C.int(sd), C.int(cd)))
}

// SetChannelMapping Set a customized input channel mapping.
func (ctxt *Context) SetChannelMapping(cm *int) int {
	return int(C.swr_set_channel_mapping((*C.struct_SwrContext)(ctxt), (*C.int)(unsafe.Pointer(cm))))
}

// SetMatrix Set a customized remix matrix.
func (ctxt *Context) SetMatrix(m *int, t int) int {
	return int(C.swr_set_matrix((*C.struct_SwrContext)(ctxt), (*C.double)(unsafe.Pointer(m)), C.int(t)))
}

// DropOutput Sample handling functions. Drops the specified number of output samples.
func (ctxt *Context) DropOutput(c int) int {
	return int(C.swr_drop_output((*C.struct_SwrContext)(ctxt), C.int(c)))
}

// InjectSilence Injects the specified number of silence samples.
func (ctxt *Context) InjectSilence(c int) int {
	return int(C.swr_inject_silence((*C.struct_SwrContext)(ctxt), C.int(c)))
}

// GetDelay Gets the delay the next input sample will experience relative to the next output sample.
func (ctxt *Context) GetDelay(b int64) int64 {
	return int64(C.swr_get_delay((*C.struct_SwrContext)(ctxt), C.int64_t(b)))
}

// ConvertFrame Frame based API. Convert the samples in the input Frame and write them to the output Frame.
func (ctxt *Context) ConvertFrame(o, i *Frame) int {
	return int(C.swr_convert_frame((*C.struct_SwrContext)(ctxt), (*C.struct_AVFrame)(o), (*C.struct_AVFrame)(i)))
}

// ConfigFrame Configure or reconfigure the Context using the information provided by the AvFrames.
func (ctxt *Context) ConfigFrame(o, i *Frame) int {
	return int(C.swr_config_frame((*C.struct_SwrContext)(ctxt), (*C.struct_AVFrame)(o), (*C.struct_AVFrame)(i)))
}
