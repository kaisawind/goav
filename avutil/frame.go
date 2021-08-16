// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package avutil

//#cgo pkg-config: libavutil
//#include <libavutil/frame.h>
//#include <libavutil/imgutils.h>
//#include <stdlib.h>
import "C"
import "unsafe"

type (
	AvBuffer            C.struct_AVBuffer
	AvBufferRef         C.struct_AVBufferRef
	AvBufferPool        C.struct_AVBufferPool
	Frame               C.struct_AVFrame
	AvFrameSideData     C.struct_AVFrameSideData
	AvFrameSideDataType C.enum_AVFrameSideDataType
)

const (
	AvNumDataPointers = C.AV_NUM_DATA_POINTERS
)

// AvFrameAlloc Allocate an Frame and set its fields to default values.
func AvFrameAlloc() *Frame {
	return (*Frame)(unsafe.Pointer(C.av_frame_alloc()))
}

// AvFrameFree Free the frame and any dynamically allocated objects in it, e.g.
func AvFrameFree(f *Frame) {
	var ptr *C.struct_AVFrame = (*C.struct_AVFrame)(unsafe.Pointer(f))
	C.av_frame_free(&ptr)
}

// Free Free the frame and any dynamically allocated objects in it, e.g.
func (f *Frame) Free() {
	var ptr *C.struct_AVFrame = (*C.struct_AVFrame)(unsafe.Pointer(f))
	C.av_frame_free(&ptr)
}

// GetBuffer Allocate new buffer(s) for audio or video data.
func (f *Frame) GetBuffer(a int) int {
	return int(C.av_frame_get_buffer((*C.struct_AVFrame)(unsafe.Pointer(f)), C.int(a)))
}

// Ref Setup a new reference to the data described by an given frame.
func (f *Frame) Ref(d *Frame) int {
	return int(C.av_frame_ref((*C.struct_AVFrame)(unsafe.Pointer(d)), (*C.struct_AVFrame)(unsafe.Pointer(f))))
}

// Clone Create a new frame that references the same data as src.
func (f *Frame) Clone() *Frame {
	return (*Frame)(unsafe.Pointer(C.av_frame_clone((*C.struct_AVFrame)(unsafe.Pointer(f)))))
}

// Unref Unreference all the buffers referenced by frame and reset the frame fields.
func (f *Frame) Unref() {
	cf := (*C.struct_AVFrame)(unsafe.Pointer(f))
	C.av_frame_unref(cf)
}

// MoveRef Move everythnig contained in src to dst and reset src.
func (f *Frame) MoveRef(d *Frame) {
	C.av_frame_move_ref((*C.struct_AVFrame)(unsafe.Pointer(d)), (*C.struct_AVFrame)(unsafe.Pointer(f)))
}

// IsWritable Check if the frame data is writable.
func (f *Frame) IsWritable() int {
	return int(C.av_frame_is_writable((*C.struct_AVFrame)(unsafe.Pointer(f))))
}

// MakeWritable Ensure that the frame data is writable, avoiding data copy if possible.
func (f *Frame) MakeWritable() int {
	return int(C.av_frame_make_writable((*C.struct_AVFrame)(unsafe.Pointer(f))))
}

// CopyProps Copy only "metadata" fields from src to dst.
func (f *Frame) CopyProps(d *Frame) int {
	return int(C.av_frame_copy_props((*C.struct_AVFrame)(unsafe.Pointer(d)), (*C.struct_AVFrame)(unsafe.Pointer(f))))
}

// GetPlaneBuffer Get the buffer reference a given data plane is stored in.
func (f *Frame) GetPlaneBuffer(p int) *AvBufferRef {
	return (*AvBufferRef)(unsafe.Pointer(C.av_frame_get_plane_buffer((*C.struct_AVFrame)(unsafe.Pointer(f)), C.int(p))))
}

// NewSideData Add a new side data to a frame.
func (f *Frame) NewSideData(d AvFrameSideDataType, s int) *AvFrameSideData {
	return (*AvFrameSideData)(unsafe.Pointer(C.av_frame_new_side_data((*C.struct_AVFrame)(unsafe.Pointer(f)), (C.enum_AVFrameSideDataType)(d), C.int(s))))
}

func (f *Frame) GetSideData(t AvFrameSideDataType) *AvFrameSideData {
	return (*AvFrameSideData)(unsafe.Pointer(C.av_frame_get_side_data((*C.struct_AVFrame)(unsafe.Pointer(f)), (C.enum_AVFrameSideDataType)(t))))
}

func (f *Frame) AvImageAlloc() int {
	ret := C.av_image_alloc(
		(**C.uint8_t)(unsafe.Pointer(&f.data)),
		(*C.int)(unsafe.Pointer(&f.linesize)),
		C.int(f.Width()),
		C.int(f.Height()),
		int32(f.Format()),
		32)
	return int(ret)
}
