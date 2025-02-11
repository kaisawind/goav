// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

//Package avfilter contains methods that deal with ffmpeg filters
//filters in the same linear chain are separated by commas, and distinct linear chains of filters are separated by semicolons.
//FFmpeg is enabled through the "C" libavfilter library
package avfilter

/*
	#cgo pkg-config: libavfilter
	#include <libavfilter/avfilter.h>
*/
import "C"
import (
	"unsafe"

	"github.com/asticode/goav/avutil"
)

type (
	Filter    C.struct_AVFilter
	Context   C.struct_AVFilterContext
	Link      C.struct_AVFilterLink
	Graph     C.struct_AVFilterGraph
	Input     C.struct_AVFilterInOut
	Pad       C.struct_AVFilterPad
	Class     C.struct_AVClass
	MediaType C.enum_AVMediaType
)

const (
	MaxArraySize = 1<<29 - 1
)

// Version Return the LIBAvFILTER_VERSION_INT constant.
func Version() uint {
	return uint(C.avfilter_version())
}

// Configuration Return the libavfilter build-time configuration.
func Configuration() string {
	return C.GoString(C.avfilter_configuration())
}

// License Return the libavfilter license.
func License() string {
	return C.GoString(C.avfilter_license())
}

// PadCount Get the number of elements in a NULL-terminated array of Pads (e.g.
func PadCount(p *Pad) int {
	return int(C.avfilter_pad_count((*C.struct_AVFilterPad)(p)))
}

// PadGetName Get the name of an Pad.
func PadGetName(p *Pad, pi int) string {
	return C.GoString(C.avfilter_pad_get_name((*C.struct_AVFilterPad)(p), C.int(pi)))
}

// PadGetType Get the type of an Pad.
func PadGetType(p *Pad, pi int) MediaType {
	return (MediaType)(C.avfilter_pad_get_type((*C.struct_AVFilterPad)(p), C.int(pi)))
}

// LinkFilter Link two filters together.
func LinkFilter(s *Context, sp uint, d *Context, dp uint) int {
	return int(C.avfilter_link((*C.struct_AVFilterContext)(s), C.uint(sp), (*C.struct_AVFilterContext)(d), C.uint(dp)))
}

// LinkFree Free the link in *link, and set its pointer to NULL.
func LinkFree(l **Link) {
	C.avfilter_link_free((**C.struct_AVFilterLink)(unsafe.Pointer(l)))
}

// ConfigLinks Negotiate the media format, dimensions, etc of all inputs to a filter.
func ConfigLinks(f *Context) int {
	return int(C.avfilter_config_links((*C.struct_AVFilterContext)(f)))
}

// ProcessCommand Make the filter instance process a command.
func ProcessCommand(f *Context, cmd, arg, res string, l, fl int) int {
	cc := C.CString(cmd)
	defer C.free(unsafe.Pointer(cc))
	ca := C.CString(arg)
	defer C.free(unsafe.Pointer(ca))
	cr := C.CString(res)
	defer C.free(unsafe.Pointer(cr))
	return int(C.avfilter_process_command((*C.struct_AVFilterContext)(f), cc, ca, cr, C.int(l), C.int(fl)))
}

// InitStr Initialize a filter with the supplied parameters.
func (ctx *Context) InitStr(args string) int {
	ca := C.CString(args)
	defer C.free(unsafe.Pointer(ca))
	return int(C.avfilter_init_str((*C.struct_AVFilterContext)(ctx), ca))
}

// InitDict Initialize a filter with the supplied dictionary of options.
func (ctx *Context) InitDict(o **avutil.Dictionary) int {
	return int(C.avfilter_init_dict((*C.struct_AVFilterContext)(ctx), (**C.struct_AVDictionary)(unsafe.Pointer(o))))
}

// Free free a filter context.
func (ctx *Context) Free() {
	C.avfilter_free((*C.struct_AVFilterContext)(ctx))
}

func (ctx *Context) NbInputs() uint {
	return uint(ctx.nb_inputs)
}

func (ctx *Context) NbOutputs() uint {
	return uint(ctx.nb_outputs)
}

func (ctx *Context) Inputs() []*Link {
	if ctx.NbInputs() == 0 {
		return nil
	}

	arr := (*[MaxArraySize]*Link)(unsafe.Pointer(ctx.inputs))

	if arr == nil {
		return nil
	}

	return arr[:ctx.NbInputs()]
}

func (ctx *Context) Outputs() []*Link {
	if ctx.NbOutputs() == 0 {
		return nil
	}

	arr := (*[MaxArraySize]*Link)(unsafe.Pointer(ctx.outputs))

	return arr[:ctx.NbOutputs()]
}

// InsertFilter Insert a filter in the middle of an existing link.
func InsertFilter(l *Link, f *Context, fsi, fdi uint) int {
	return int(C.avfilter_insert_filter((*C.struct_AVFilterLink)(l), (*C.struct_AVFilterContext)(f), C.uint(fsi), C.uint(fdi)))
}

// GetClass get class
func GetClass() *Class {
	return (*Class)(C.avfilter_get_class())
}

// InoutAlloc Allocate a single Input entry.
func InoutAlloc() *Input {
	return (*Input)(C.avfilter_inout_alloc())
}

// InoutFree Free the supplied list of Input and set *inout to NULL.
func InoutFree(i **Input) {
	C.avfilter_inout_free((**C.struct_AVFilterInOut)(unsafe.Pointer(i)))
}

func (i *Input) SetName(n string) {
	i.name = C.CString(n)
}

func (i *Input) SetFilterCtx(ctx *Context) {
	i.filter_ctx = (*C.struct_AVFilterContext)(ctx)
}

func (i *Input) SetPadIdx(idx int) {
	i.pad_idx = C.int(idx)
}

func (i *Input) SetNext(n *Input) {
	i.next = (*C.struct_AVFilterInOut)(n)
}

func (l *Link) TimeBase() avutil.Rational {
	return *(*avutil.Rational)(unsafe.Pointer(&l.time_base))
}
