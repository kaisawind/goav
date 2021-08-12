// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package avfilter

/*
	#cgo pkg-config: libavfilter
	#include <libavfilter/avfilter.h>
	#include <libavfilter/buffersink.h>
	#include <libavfilter/buffersrc.h>
*/
import "C"
import (
	"github.com/asticode/goav/avutil"
	"unsafe"
)

const (
	AvBuffersrcFlagKeepRef = C.AV_BUFFERSRC_FLAG_KEEP_REF
)

// GraphAlloc Allocate a filter graph.
func GraphAlloc() *Graph {
	return (*Graph)(C.avfilter_graph_alloc())
}

// GraphAllocFilter Create a new filter instance in a filter graph.
func (g *Graph) GraphAllocFilter(f *Filter, n string) *Context {
	cn := C.CString(n)
	defer C.free(unsafe.Pointer(cn))
	return (*Context)(C.avfilter_graph_alloc_filter((*C.struct_AVFilterGraph)(g), (*C.struct_AVFilter)(f), cn))
}

// GraphGetFilter Get a filter instance identified by instance name from graph.
func (g *Graph) GraphGetFilter(n string) *Context {
	cn := C.CString(n)
	defer C.free(unsafe.Pointer(cn))
	return (*Context)(C.avfilter_graph_get_filter((*C.struct_AVFilterGraph)(g), cn))
}

// GraphSetAutoConvert Enable or disable automatic format conversion inside the graph.
func (g *Graph) GraphSetAutoConvert(f uint) {
	C.avfilter_graph_set_auto_convert((*C.struct_AVFilterGraph)(g), C.uint(f))
}

// GraphConfig Check validity and configure all the links and formats in the graph.
func (g *Graph) GraphConfig(l *int) int {
	return int(C.avfilter_graph_config((*C.struct_AVFilterGraph)(g), unsafe.Pointer(l)))
}

// GraphFree Free a graph, destroy its links, and set *graph to NULL.
func (g *Graph) GraphFree() {
	C.avfilter_graph_free((**C.struct_AVFilterGraph)(unsafe.Pointer(&g)))
}

// GraphParse Add a graph described by a string to a graph.
func (g *Graph) GraphParse(f string, i, o *Input, l int) int {
	cf := C.CString(f)
	defer C.free(unsafe.Pointer(cf))
	return int(C.avfilter_graph_parse((*C.struct_AVFilterGraph)(g), cf, (*C.struct_AVFilterInOut)(i), (*C.struct_AVFilterInOut)(o), unsafe.Pointer(&l)))
}

// GraphParsePtr Add a graph described by a string to a graph.
func (g *Graph) GraphParsePtr(f string, i, o **Input, l *int) int {
	cf := C.CString(f)
	defer C.free(unsafe.Pointer(cf))
	return int(C.avfilter_graph_parse_ptr((*C.struct_AVFilterGraph)(g), cf, (**C.struct_AVFilterInOut)(unsafe.Pointer(i)), (**C.struct_AVFilterInOut)(unsafe.Pointer(o)), unsafe.Pointer(l)))
}

// GraphParse2 Add a graph described by a string to a graph.
func (g *Graph) GraphParse2(f string, i, o **Input) int {
	cf := C.CString(f)
	defer C.free(unsafe.Pointer(cf))
	return int(C.avfilter_graph_parse2((*C.struct_AVFilterGraph)(g), cf, (**C.struct_AVFilterInOut)(unsafe.Pointer(i)), (**C.struct_AVFilterInOut)(unsafe.Pointer(o))))
}

// GraphSendCommand Send a command to one or more filter instances.
func (g *Graph) GraphSendCommand(t, cmd, arg, res string, resl, f int) int {
	ct := C.CString(t)
	defer C.free(unsafe.Pointer(ct))
	cc := C.CString(cmd)
	defer C.free(unsafe.Pointer(cc))
	ca := C.CString(arg)
	defer C.free(unsafe.Pointer(ca))
	cr := C.CString(res)
	defer C.free(unsafe.Pointer(cr))
	return int(C.avfilter_graph_send_command((*C.struct_AVFilterGraph)(g), ct, cc, ca, cr, C.int(resl), C.int(f)))
}

// GraphQueueCommand Queue a command for one or more filter instances.
func (g *Graph) GraphQueueCommand(t, cmd, arg string, f int, ts C.double) int {
	ct := C.CString(t)
	defer C.free(unsafe.Pointer(ct))
	cc := C.CString(cmd)
	defer C.free(unsafe.Pointer(cc))
	ca := C.CString(arg)
	defer C.free(unsafe.Pointer(ca))
	return int(C.avfilter_graph_queue_command((*C.struct_AVFilterGraph)(g), ct, cc, ca, C.int(f), ts))
}

// GraphDump Dump a graph into a human-readable string representation.
func (g *Graph) GraphDump(o string) string {
	co := C.CString(o)
	defer C.free(unsafe.Pointer(co))
	return C.GoString(C.avfilter_graph_dump((*C.struct_AVFilterGraph)(g), co))
}

// GraphRequestOldest Request a frame on the oldest sink
func (g *Graph) GraphRequestOldest() int {
	return int(C.avfilter_graph_request_oldest((*C.struct_AVFilterGraph)(g)))
}

// GraphCreateFilter Create and add a filter instance into an existing graph.
func GraphCreateFilter(cx **Context, f *Filter, n, a string, o *int, g *Graph) int {
	ca := (*C.char)(nil)
	if len(a) > 0 {
		ca = C.CString(a)
		defer C.free(unsafe.Pointer(ca))
	}
	cn := C.CString(n)
	defer C.free(unsafe.Pointer(cn))
	return int(C.avfilter_graph_create_filter((**C.struct_AVFilterContext)(unsafe.Pointer(cx)), (*C.struct_AVFilter)(f), cn, ca, unsafe.Pointer(o), (*C.struct_AVFilterGraph)(g)))
}

func (g *Graph) AvBuffersrcAddFrameFlags(cx *Context, frame *avutil.Frame, flags int) int {
	return int(C.av_buffersrc_add_frame_flags((*C.struct_AVFilterContext)(cx), (*C.struct_AVFrame)(unsafe.Pointer(frame)), C.int(flags)))
}

func (g *Graph) AvBuffersinkGetFrame(cx *Context, frame *avutil.Frame) int {
	return int(C.av_buffersink_get_frame((*C.struct_AVFilterContext)(cx), (*C.struct_AVFrame)(unsafe.Pointer(frame))))
}
