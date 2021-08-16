// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package avcodec

//#cgo pkg-config: libavcodec
//#include <libavcodec/avcodec.h>
import "C"
import (
	"unsafe"

	"github.com/asticode/goav/avutil"
)

// FreeContext Free the codec context and everything associated with it and write NULL to the provided pointer.
func FreeContext(ctxt *Context) {
	var ptr *C.struct_AVCodecContext = (*C.struct_AVCodecContext)(unsafe.Pointer(ctxt))
	C.avcodec_free_context(&ptr)
}

// Free Free the codec context and everything associated with it and write NULL to the provided pointer.
func (ctxt *Context) Free() {
	var ptr *C.struct_AVCodecContext = (*C.struct_AVCodecContext)(unsafe.Pointer(ctxt))
	C.avcodec_free_context(&ptr)
}

// GetContextDefaults3 Set the fields of the given Context to default values corresponding to the given codec (defaults may be codec-dependent).
func (ctxt *Context) GetContextDefaults3(c *Codec) int {
	return int(C.avcodec_get_context_defaults3((*C.struct_AVCodecContext)(unsafe.Pointer(ctxt)), (*C.struct_AVCodec)(c)))
}

// Open2 Initialize the Context to use the given Codec
func (ctxt *Context) Open2(c *Codec, d **avutil.Dictionary) int {
	return int(C.avcodec_open2((*C.struct_AVCodecContext)(unsafe.Pointer(ctxt)), (*C.struct_AVCodec)(c), (**C.struct_AVDictionary)(unsafe.Pointer(d))))
}

// Close Close a given Context and free all the data associated with it (but not the Context itself).
func (ctxt *Context) Close() int {
	return int(C.avcodec_close((*C.struct_AVCodecContext)(unsafe.Pointer(ctxt))))
}

// DefaultGetBuffer2 The default callback for Context.get_buffer2().
func (ctxt *Context) DefaultGetBuffer2(f *Frame, l int) int {
	return int(C.avcodec_default_get_buffer2((*C.struct_AVCodecContext)(unsafe.Pointer(ctxt)), (*C.struct_AVFrame)(f), C.int(l)))
}

// AlignDimensions Modify width and height values so that they will result in a memory buffer that is acceptable for the codec if you do not use any horizontal padding.
func (ctxt *Context) AlignDimensions(w, h *int) {
	C.avcodec_align_dimensions((*C.struct_AVCodecContext)(unsafe.Pointer(ctxt)), (*C.int)(unsafe.Pointer(w)), (*C.int)(unsafe.Pointer(h)))
}

// AlignDimensions2 Modify width and height values so that they will result in a memory buffer that is acceptable for the codec if you also ensure that all line sizes are a multiple of the respective linesize_align[i].
func (ctxt *Context) AlignDimensions2(w, h *int, l int) {
	C.avcodec_align_dimensions2((*C.struct_AVCodecContext)(unsafe.Pointer(ctxt)), (*C.int)(unsafe.Pointer(w)), (*C.int)(unsafe.Pointer(h)), (*C.int)(unsafe.Pointer(&l)))
}

func (ctxt *Context) ReceivePacket(a *Packet) int {
	return ReceivePacket(ctxt, a)
}

// SendPacket sends a packet to the context for decoding
// OO form of AvcodecSendPacket
func (ctxt *Context) SendPacket(a *Packet) int {
	return SendPacket(ctxt, a)
}

// ReceiveFrame receives a decoded from from a context
// OO form of AvcodecReceiveFrame
func (ctxt *Context) ReceiveFrame(f *avutil.Frame) int {
	return ReceiveFrame(ctxt, f)
}

func (ctxt *Context) SendFrame(f *avutil.Frame) int {
	return SendFrame(ctxt, f)
}

// DecodeSubtitle2 Decode a subtitle message.
func (ctxt *Context) DecodeSubtitle2(s *AvSubtitle, g *int, a *Packet) int {
	return int(C.avcodec_decode_subtitle2((*C.struct_AVCodecContext)(unsafe.Pointer(ctxt)), (*C.struct_AVSubtitle)(s), (*C.int)(unsafe.Pointer(g)), (*C.struct_AVPacket)(a)))
}

func (ctxt *Context) EncodeSubtitle(b *uint8, bs int, s *AvSubtitle) int {
	return int(C.avcodec_encode_subtitle((*C.struct_AVCodecContext)(unsafe.Pointer(ctxt)), (*C.uint8_t)(b), C.int(bs), (*C.struct_AVSubtitle)(s)))
}

func (ctxt *Context) DefaultGetFormat(f *avutil.PixelFormat) avutil.PixelFormat {
	return (avutil.PixelFormat)(C.avcodec_default_get_format((*C.struct_AVCodecContext)(unsafe.Pointer(ctxt)), (*C.enum_AVPixelFormat)(f)))
}

// FlushBuffers Reset the internal decoder state / flush internal buffers.
func (ctxt *Context) FlushBuffers() {
	C.avcodec_flush_buffers((*C.struct_AVCodecContext)(unsafe.Pointer(ctxt)))
}

// AvGetAudioFrameDuration Return audio frame duration.
func (ctxt *Context) AvGetAudioFrameDuration(f int) int {
	return int(C.av_get_audio_frame_duration((*C.struct_AVCodecContext)(unsafe.Pointer(ctxt)), C.int(f)))
}

func (ctxt *Context) IsOpen() int {
	return int(C.avcodec_is_open((*C.struct_AVCodecContext)(unsafe.Pointer(ctxt))))
}

// AvParserParse2 Parse a packet.
func (ctxt *Context) AvParserParse2(ctxtp *ParserContext, p **uint8, ps *int, b *uint8, bs int, pt, dt, po int64) int {
	return int(C.av_parser_parse2((*C.struct_AVCodecParserContext)(ctxtp), (*C.struct_AVCodecContext)(unsafe.Pointer(ctxt)), (**C.uint8_t)(unsafe.Pointer(p)), (*C.int)(unsafe.Pointer(ps)), (*C.uint8_t)(b), C.int(bs), (C.int64_t)(pt), (C.int64_t)(dt), (C.int64_t)(po)))
}

func AvParserInit(c int) *ParserContext {
	return (*ParserContext)(C.av_parser_init(C.int(c)))
}

func AvParserClose(ctxtp *ParserContext) {
	C.av_parser_close((*C.struct_AVCodecParserContext)(ctxtp))
}
